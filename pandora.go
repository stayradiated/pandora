package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/stayradiated/pandora/pandora-lib"
)

// main takes a username and password to fetch the station and feedback info which it will prints out as JSON
func main() {

	username := flag.String("u", "", "Pandora username")
	password := flag.String("p", "", "Pandora password")
	asJson := flag.Bool("json", false, "Print as json")
	flag.Parse()

	if len(*username) < 1 || len(*password) < 1 {
		printHelp()
		return
	}

	output, err := pandora.FetchStations(*username, *password)
	if err != nil {
		panic(err)
	}

	if *asJson {
		json, err := json.Marshal(output)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	} else {
		for _, station := range output {
			for _, song := range station.Songs {
				fmt.Println(song.Artist, "--", song.Name)
			}
		}
	}

}

// printHelp will print help
func printHelp() {
	fmt.Println(`Pandora prints out your stations and favorites.

Usage:
	pandora -u [username] -p [password] [-json]
	`)
}
