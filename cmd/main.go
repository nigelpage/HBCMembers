package main
import (
    "log"
    "/internal/application"
)
func main() {
    app := application.NewApp()
    if err := app.Run(); err != nil {
        log.Fatalf("Application failed: %v", err)
    }
}
