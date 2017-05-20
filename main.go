package main

import (
	"fmt"
	"log"
	"os"

	"github.com/takama/router"
)

// Run server: go run main.go
// Try few requests:
// - curl http://localhost:8000
// - curl http://localhost:8000/test-some-path
func main() {
	logger := log.New(os.Stdout, "[step-by-step] ", log.LstdFlags)

	logger.Print("Запускаем приложение...")

	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		logger.Fatal("Порт не задан")
	}

	r := router.New()

	r.GET("/", home)

	logger.Printf("Готовимся слушать порт %s...", port)
	r.Listen(":" + port)
}

func home(c *router.Control) {
	log.Print("Хэндлер home поймал запрос")
	fmt.Fprintf(
		c.Writer, "URL.Path = %q\n", c.Request.URL.Path,
	)

}
