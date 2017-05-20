package main

import (
	"fmt"
	"github.com/takama/router"
	"net/http"
)

// Run server: go run main.go
// Try few requests:
// - curl http://localhost:8000
// - curl http://localhost:8000/test-some-path
func main() {
	r := router.New()

	a := 5

	r.POST("/api/v1/users", mw(a, abc))

	r.Listen(":8000")
}

func mw(a int, h router.Handle) router.Handle {
	// ? как передать a в h?
	return h
}

func pampam(c *router.Control) {
	fmt.Fprintf(
		c.Writer, "URL.Path = %q\n", c.Request.URL.Path,
	)

}

func abc(c *router.Control) {
	fmt.Fprintf(
		c.Writer, "a = ?",
	)

}

func home(c *router.Control) {
	fmt.Fprintf(
		c.Writer, "URL.Path = %q\n", c.Request.URL.Path,
	)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
