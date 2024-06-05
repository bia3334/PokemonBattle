package main

// import (
// 	"bytes"
	
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"golang.org/x/net/html"
// )

type DamegeWhenAttacked struct {
	Element     string  
	Coefficient float64 
}

type Stats struct {
	HP                 int                  
	Defense            int                  
	Speed              int                  
	Sp_Attack          int                  
	Sp_Defense         int                  
}

type GenderRatio struct {
	MaleRatio  int 
	FemaleRatio int 
}

type Profile struct {
	Name 			string               
	Weight          float64              
	CatchRate 	 int                  
	GenderRatio     GenderRatio 
	EggGroup        []string 
	HatchSteps	  int
	Abilities	   []string 
}

type NaturalMoves struct {
	Name string
	Element string
	Power int
	Acc 	int
	PP 		int
	Description string
}

type MachineMoves struct {
	Name string
	Element string
	Power int
	Acc 	int
	PP 		int
	Description string
}

type TutorMoves struct {
	Name string
	Element string
	Power int
	Acc 	int
	PP 		int
	Description string
}

type EggMoves struct {
	Name string
	Element string
	Power int
	Acc 	int
	PP 		int
	Description string
}

type Moves struct {
	NaturalMoves []NaturalMoves
	MachineMoves []MachineMoves
	TutorMoves   []TutorMoves
	EggMoves     []EggMoves
}

type Pokemon struct {
	Elements           []string             
	EV                 int                                
	Profile			Profile              
	DamegeWhenAttacked []DamegeWhenAttacked 
	EvolutionLevel    int
	Moves			  Moves                  
}

// func fetchPokemons(index int) []Pokemon {
// 	url := fmt.Sprintf("https://pokedex.org/#/pokemon/%d", index);
// 		resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error fetching genre page:", err)
// 		return nil
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return nil
// 	}

// 	doc, err := html.Parse(bytes.NewReader(body))
// 	if err != nil {
// 		fmt.Println("Error parsing HTML:", err)
// 		return nil
// 	}

// 	// return findPokemons(doc)
// }

// func findPokemons(n* html.Node) {
// 	var pokemonList []Pokemon
// 	var walk func(*html.Node)
// 	count := 0

// 	walk = func(n *html.Node) {
// 		if count >= 10 {
// 			return
// 		}
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			var link, title string
// 			for _, attr := range n.Attr {
// 				if attr.Key == "href" {
// 					link = attr.Val
// 				}
// 			}
// 			for c := n.FirstChild; c != nil; c = c.NextSibling {
// 				if c.Type == html.ElementNode && c.Data == "div" {
// 					for gc := c.FirstChild; gc != nil; gc = gc.NextSibling {
// 						if gc.Type == html.ElementNode && gc.Data == "p" && hasClass(gc, "subj") {
// 							if gc.FirstChild != nil && gc.FirstChild.Type == html.TextNode {
// 								title = gc.FirstChild.Data
// 							}
// 						}
// 					}
// 				}
// 			}
// 			if link != "" && title != "" {
// 				pokemonList = append(pokemonList, Pokemon{})
// 				count++
// 			}
// 		}
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			walk(c)
// 		}
// 	}

// 	walk(n)
// 	return pokemonList
// }

// func hasClass(n *html.Node, class string) bool {
// 	for _, attr := range n.Attr {
// 		if attr.Key == "class" && attr.Val == class {
// 			return true
// 		}
// 	}
// 	return false
// }

