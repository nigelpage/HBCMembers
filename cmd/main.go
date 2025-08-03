package main
import (
    "log"
    
)
func main() {
    app := application.NewApp()
    if err := app.Run(); err != nil {
        log.Fatalf("Application failed: %v", err)
    }
}
