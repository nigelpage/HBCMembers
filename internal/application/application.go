package application

import "log"

type App struct {
	config *AppConfig
}

func NewApp() *App {

	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	return &App{config: config}
}

func (a *App) Run() error {
	log.Printf("Running application in %s environment", a.config.Environment)
	log.Printf("Service Name: %s, Ports: %v", a.config.Service.Name, a.config.Service.Ports)
	// Add more application logic here
	return nil
}
