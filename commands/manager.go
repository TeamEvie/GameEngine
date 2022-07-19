package commands

import "eviecoin/database"

type Manager struct {
	Commands map[string]Command
	Database database.Database
}

func NewManager(db database.Database) *Manager {
	return &Manager{
		Commands: make(map[string]Command),
		Database: db,
	}
}
