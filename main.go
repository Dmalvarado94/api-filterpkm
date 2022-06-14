package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonPokemon struct {
	Pokemon []Pokemon `json:"pokemon"`
}
type pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Pokemon struct {
	Pokemon pokemon `json:"pokemon"`
	Slot    int     `json:"slot"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/type/1")
	if err != nil {
		fmt.Printf("404 not Found %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	JsonPokemon := json.RawMessage(`{"Name":"pidgey", "url":"https://pokeapi.co/api/v2/pokemon/16/"}`)
	var pkm2 pokemon
	err2 := json.Unmarshal(JsonPokemon, &pkm2)
	log.Print(err2)
	log.Printf("%+v", pkm2)

	// Agregar unmarshall de datos y en el GET hacer un ciclo for de todos los tipos

}
