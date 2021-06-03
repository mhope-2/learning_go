package main
import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
	"Content-Type",
	"text/html",
)
	io.WriteString(
		res,
		`<DOCTYPE html> <html>
              <head>
                  <title>Hello, World</title>
              </head>
              <body>
                  Hello, World!
              </body>
</html>`,
)
}
func main() {
	// handle route
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9000", nil)

	// handle static files
	http.Handle( "/assets/",
		http.StripPrefix( "/assets/",
			http.FileServer(http.Dir("assets")), ),
	)
}