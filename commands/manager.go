package commands

type Manager struct {
	Commands map[string]Command
}

func NewManager() *Manager {
	return &Manager{
		Commands: make(map[string]Command),
	}
}
