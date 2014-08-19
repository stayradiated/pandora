package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cellofellow/gopiano"
	"github.com/cellofellow/gopiano/responses"
)

type stationOutput struct {
	Name  string       `json:name`
	Songs []songOutput `json:songs`
}

type songOutput struct {
	Name   string `json:name`
	Artist string `json:artist`
}

func fetchStation(client *gopiano.Client, station responses.Station, result chan<- stationOutput) {
	details, err := client.StationGetStation(station.StationToken, true)
	if err != nil {
		panic(err)
	}

	feedback := details.Result.Feedback.ThumbsUp

	fmt.Println(station.StationName, len(feedback))

	sOutput := stationOutput{
		Name:  details.Result.StationName,
		Songs: make([]songOutput, len(feedback)),
	}

	for i, song := range feedback {
		sOutput.Songs[i] = songOutput{
			Artist: song.ArtistName,
			Name:   song.SongName,
		}
	}

	result <- sOutput
}

func FetchStations(username, password string) ([]stationOutput, error) {
	client, err := gopiano.NewClient(gopiano.AndroidClient)
	if err != nil {
		return nil, err
	}

	_, err = client.AuthPartnerLogin()
	if err != nil {
		return nil, err
	}

	_, err = client.AuthUserLogin(username, password)
	if err != nil {
		return nil, err
	}

	stationsList, err := client.UserGetStationList(false)
	if err != nil {
		return nil, err
	}

	output := make([]stationOutput, 0)
	stations := stationsList.Result.Stations

	result := make(chan stationOutput, 2)
	expectedLength := len(stations)

	for _, station := range stations {
		if station.IsQuickMix || station.IsShared {
			expectedLength -= 1
			continue
		}

		go fetchStation(client, station, result)
	}

	for {
		sOutput := <-result
		output = append(output, sOutput)
		if len(output) == expectedLength {
			break
		}
	}

	return output, nil
}

func main() {

	if len(os.Args) < 3 {
		printHelp()
		return
	}

	username := os.Args[1]
	password := os.Args[2]

	output, err := FetchStations(username, password)
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}

func printHelp() {
	fmt.Println(`Pandora prints out your stations and favorites.

Usage:
	pandora [email [password]
	`)
}
