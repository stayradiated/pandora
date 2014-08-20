package main

import (
	"reflect"
	"testing"

	"github.com/cellofellow/gopiano/responses"
)

var expectedStationOutput = stationOutput{
	Name: "station",
	Songs: []songOutput{
		songOutput{
			Name:   "song",
			Artist: "artist",
		},
	},
}

var expectedStationListOutput = []stationOutput{
	expectedStationOutput,
}

func mockStationGetStation(token string, includeFullAttrs bool) (*responses.StationGetStation, error) {
	res := &responses.StationGetStation{}
	res.Result.Feedback.ThumbsUp = make([]responses.FeedbackResponse, 1)

	res.Result.StationName = "station"

	res.Result.Feedback.ThumbsUp[0] = responses.FeedbackResponse{
		ArtistName: "artist",
		SongName:   "song",
	}

	return res, nil
}

func TestFetchStation(t *testing.T) {
	result := make(chan *stationOutput)
	go fetchStation(mockStationGetStation, "id", result)

	output := <-result

	if !reflect.DeepEqual(*output, expectedStationOutput) {
		t.Errorf("does not match expected output")
	}
}

func TestProcessStations(t *testing.T) {
	stationList := make(responses.StationList, 1)
	stationList[0] = responses.Station{
		IsQuickMix:   false,
		IsShared:     false,
		StationToken: "id",
	}

	output := processStations(mockStationGetStation, stationList)

	if !reflect.DeepEqual(output, expectedStationListOutput) {
		t.Errorf("does not match expected output")
	}
}
