package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type PokemonDetail struct {
	Name  string       `json:"name"`
	ID    string       `json:"id"`
	Types PokemonTypes `json:"types"`
	Stats PokemonStats `json:"stats"`
}

type PokemonTypes struct {
	Type1 string `json:"type1"`
	Type2 string `json:"type2"`
}

type PokemonStats struct {
	HP     string `json:"HP"`
	Attack string `json:"Attack"`
	Defense string `json:"Defense"`
	Speed  string `json:"Speed"`
	SpAtk  string `json:"Sp Atk"`
	SpDef  string `json:"Sp Def"`
}

func main() {
	pokemonMap := make(map[int]PokemonDetail)

	// Loop through Pokémon IDs from 1 to 151
	for id := 1; id <= 2; id++ {
		// Construct the URL with the current ID
		url := fmt.Sprintf("https://pokedex.org/#/pokemon/%d", id)

		// Fetch Pokémon details
		pokemonDetail, err := fetchPokemonDetail(url)
		if err != nil {
			log.Printf("Error fetching details for Pokémon with ID %d: %v\n", id, err)
			continue
		}

		// Store the Pokémon details in the map
		pokemonMap[id] = pokemonDetail
	}

	// Marshal the Pokémon map into JSON format
	jsonData, err := json.MarshalIndent(pokemonMap, "", "    ")
	if err != nil {
		log.Fatal("Error marshalling Pokémon details into JSON: ", err)
	}
	// Write the JSON data to a file
	file, err := os.Create("pokemon_details.json")
	if err != nil {
		log.Fatal("Error creating file for Pokémon details: ", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing Pokémon details to file: ", err)
	}

	log.Println("Pokémon details have been saved to pokemon_details.json")
}

func fetchPokemonDetail(url string) (PokemonDetail, error) {
	var pokemonDetail PokemonDetail

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return pokemonDetail, err
	}

	// Extracting Pokémon name
	pokemonDetail.Name = doc.Find(".detail-panel-header").Text()

	doc.Find(".detail-types-and-num").Each(func(_ int, s *goquery.Selection) {
		id := s.Find(".detail-national-id span").Text()
		types := PokemonTypes{}
		s.Find(".detail-types span").Each(func(i int, t *goquery.Selection) {
			if i == 0 {
				types.Type1 = t.Text()
			} else if i == 1 {
				types.Type2 = t.Text()
			}
		})
		pokemonDetail.ID = id
		pokemonDetail.Types = types
	})

	doc.Find(".detail-stats .detail-stats-row").Each(func(_ int, s *goquery.Selection) {
		statName := s.Find("span").First().Text()
		statValue := s.Find(".stat-bar-fg").Text()

		switch statName {
		case "HP":
			pokemonDetail.Stats.HP = statValue
		case "Attack":
			pokemonDetail.Stats.Attack = statValue
		case "Defense":
			pokemonDetail.Stats.Defense = statValue
		case "Speed":
			pokemonDetail.Stats.Speed = statValue
		case "Sp Atk":
			pokemonDetail.Stats.SpAtk = statValue
		case "Sp Def":
			pokemonDetail.Stats.SpDef = statValue
		}
	})

	return pokemonDetail, nil
}
