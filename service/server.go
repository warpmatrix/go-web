package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/api/unknown", notImplementedHandler()).Methods("GET")
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/userInfo", postUserInfoHandler(formatter)).Methods("POST")
	mx.HandleFunc("/userInfo", getUserInfoHandler(formatter)).Methods("GET")

	// setup file server
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}
