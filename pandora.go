package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cellofellow/gopiano"
)

// stationOutput contains a station name and a slice of songs
type stationOutput struct {
	Name  string       `json:name`
	Songs []songOutput `json:songs`
}

// songOutput contains a song name and artist
type songOutput struct {
	Name   string `json:name`
	Artist string `json:artist`
}

// fetchStation fetches full information about a station, specifically the thumbed up songs. It then processes this info and creates a stationOutput instance. This is then passed into the result channel.
func fetchStation(client *gopiano.Client, stationToken string, result chan<- stationOutput) {
	details, err := client.StationGetStation(stationToken, true)
	if err != nil {
		panic(err)
	}

	// alias
	feedback := details.Result.Feedback.ThumbsUp

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

// FetchStations connects to pandora using a username/password combo and downloads their entire station list. It then goes through each station and fetches their feedback info.
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

		go fetchStation(client, station.StationToken, result)
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

// main takes a username and password to fetch the station and feedback info which it will prints out as JSON
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

// printHelp will print help
func printHelp() {
	fmt.Println(`Pandora prints out your stations and favorites.

Usage:
	pandora [email [password]
	`)
}
