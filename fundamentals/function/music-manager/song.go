package main

import "errors"

type Song struct {
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Artist   string `json:"artist"`
}

func New(title, artist string, duration int) (*Song, error) {
	if title == "" || artist == "" || duration == 0 {
		return nil, errors.New("title, artist and duration is required")
	}

	return &Song{
		Title:    title,
		Duration: duration,
		Artist:   artist,
	}, nil
}
