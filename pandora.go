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

func (s *stationOutput) AddFeedback(feedback []responses.FeedbackResponse) {
	for _, song := range feedback {
		s.Songs = append(s.Songs, songOutput{
			Artist: song.ArtistName,
			Name:   song.SongName,
		})
	}
}

func main() {

	if len(os.Args) < 3 {
		printHelp()
		return
	}

	username := os.Args[1]
	password := os.Args[2]

	client, err := gopiano.NewClient(gopiano.AndroidClient)
	if err != nil {
		panic(err)
	}

	_, _ = client.AuthPartnerLogin()
	_, err = client.AuthUserLogin(username, password)
	if err != nil {
		panic(err)
	}

	stationsResult, err := client.UserGetStationList(false)
	if err != nil {
		panic(err)
	}

	output := make([]stationOutput, 0)
	stations := stationsResult.Result.Stations

	for i, station := range stations {
		fmt.Println(i, "/", len(stations), station.StationName)

		stationResult, err := client.StationGetStation(station.StationToken, true)
		if err != nil {
			panic(err)
		}

		station = stationResult.Result
		sOutput := stationOutput{
			Name:  station.StationName,
			Songs: make([]songOutput, 0),
		}

		sOutput.AddFeedback(station.Feedback.ThumbsUp)

		output = append(output, sOutput)
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
