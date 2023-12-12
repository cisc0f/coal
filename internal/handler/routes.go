package handler

import (
	"context"
	"io"
	"net/url"
	"strings"

	coal "github.com/cisc0f/coal/internal/core"
	fire "github.com/cisc0f/coal/internal/db"
)

const (
	firebaseRef = "songs"
)

type CoalServerResponse struct {
	Id           int
	Title        string
	Author       string
	MatchingRate float64
}

func GetAllSongs(database *fire.DB) ([]fire.Song, error) {
	var songs []fire.Song

	ref := database.Client.NewRef(firebaseRef)

	err := ref.Get(context.TODO(), &songs)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func PostCompare(database *fire.DB, storage *fire.Storage, payload *io.Reader) (*CoalServerResponse, error) {
	songs, err := GetAllSongs(database)
	if err != nil {
		return nil, err
	}

	// Get Notes On from the input song
	inputSong := coal.ReadNotesOn(payload)

	// fmt.Print(inputSong)

	res := &CoalServerResponse{}
	var maxMatchingRate float64 = 0.0

	// Loop through every song in the fireDB
	for idx, song := range songs {

		parsedUrl, err := url.Parse(song.SongUrl)
		if err != nil {
			return nil, err
		}

		cleanPath := strings.TrimLeft(parsedUrl.Path, "/")
		reader, err := storage.Client.Bucket(parsedUrl.Host).Object(cleanPath).NewReader(context.Background())
		if err != nil {
			return nil, err
		}

		var ioReader io.Reader = reader

		matchingRate := coal.Start(inputSong, &ioReader)

		if matchingRate > maxMatchingRate {
			maxMatchingRate = matchingRate

			res.Id = idx
			res.Author = song.Author
			res.Title = song.Title
			res.MatchingRate = matchingRate
		}

	}

	return res, nil
}
