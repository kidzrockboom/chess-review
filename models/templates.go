package models

import "github.com/google/uuid"

type Player struct {
	Username string
	Rating   int
	Result   string
	Id       string
	UUID     uuid.UUID
}

type Game struct {
	Url          string
	Pgn          string
	TimeControl  string
	EndTime      int32
	Rated        bool
	Tcn          string
	UUID         uuid.UUID
	InitialSetup string
	Fen          string
	TimeClass    string
	Rules        string
	White        Player
	Black        Player
	Eco          string
}
