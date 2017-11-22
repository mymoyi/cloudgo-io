package server

import (
	"net/http"
	"os"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx)
	n.UseHandler(mx)
	return n
}
func initRoutes(mx *mux.Router) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	mx.HandleFunc("/unknown", unknownImp).Methods("GET")
	mx.HandleFunc("/api", apihandle).Methods("GET")
	mx.HandleFunc("/", showTable).Methods("POST")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}
func unknownImp(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "501 Not Implement", http.StatusNotImplemented)
}
func apihandle(w http.ResponseWriter, req *http.Request) {
	rdr := render.New()
	rdr.JSON(w, http.StatusOK, struct {
		ID      string
		Content string
	}{ID: "15331244", Content: "MoYi"})
}
func showTable(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rdr := render.New()
	err := rdr.HTML(rw, http.StatusOK, "table", struct {
		Username string
		Password string
	}{Username: r.Form["username"][0], Password: r.Form["password"][0]})
	if err != nil {
		panic("render fail!")
	}
}
