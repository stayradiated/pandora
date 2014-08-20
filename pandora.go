package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cellofellow/gopiano"
	"github.com/cellofellow/gopiano/responses"
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

type stationGetter func(token string, includeFullAttrs bool) (*responses.StationGetStation, error)

// fetchStation fetches full information about a station, specifically the thumbed up songs. It then processes this info and creates a stationOutput instance. This is then passed into the result channel.
func fetchStation(getStation stationGetter, stationToken string, result chan<- *stationOutput) {
	// get full station details (including feedback)
	details, err := getStation(stationToken, true)
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

	result <- &sOutput
}

// processStations takes a list of stations and fetches the feedback information for each one
func processStations(getStation stationGetter, stations responses.StationList) []stationOutput {
	output := make([]stationOutput, 0)
	result := make(chan *stationOutput, 2)
	expectedLength := len(stations)

	for _, station := range stations {
		// ignore quick mix and shared stations
		if station.IsQuickMix || station.IsShared {
			expectedLength -= 1
			continue
		}

		// run in parallel as a goroutine
		go fetchStation(getStation, station.StationToken, result)
	}

	// collect responses and assemble into a single slice
	for {
		sOutput := <-result
		output = append(output, *sOutput)
		if len(output) == expectedLength {
			break
		}
	}

	return output
}

// FetchStations connects to pandora using a username/password combo and downloads their entire station list. It then goes through each station and fetches their feedback info.
func FetchStations(username, password string) ([]stationOutput, error) {
	// setup a new gopiano client using the android settings
	client, err := gopiano.NewClient(gopiano.AndroidClient)
	if err != nil {
		return nil, err
	}

	// apparently required before doing a login
	_, err = client.AuthPartnerLogin()
	if err != nil {
		return nil, err
	}

	// auth the user
	_, err = client.AuthUserLogin(username, password)
	if err != nil {
		return nil, err
	}

	// get full station list (doesn't include station feedback)
	stations, err := client.UserGetStationList(false)
	if err != nil {
		return nil, err
	}

	return processStations(client.StationGetStation, stations.Result.Stations), nil
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
