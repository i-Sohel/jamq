package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
)

type Artist struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Furl     string `json:"facebook_page_url"`
	Event    []EventInfo
}

type EventInfo struct {
	ID             string `json:"id"`
	ArtistID       string `json:"artist_id"`
	URL            string `json:"url"`
	OnSaleDatetime string `json:"on_sale_datetime"`
	Datetime       string `json:"datetime"`
	Venue          struct {
		Name      string `json:"name"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
		City      string `json:"city"`
		Region    string `json:"region"`
		Country   string `json:"country"`
	} `json:"venue"`
}

const (
	DEFAULT_PORT = "8080"
	DEF_URL      = "https://rest.bandsintown.com/artists/" // API URL
)

// var templates = template.Must(template.ParseGlob("templates/*.html")) // Parse templates locally

var _, filename, _, ok = runtime.Caller(0)
var tmpl = template.Must(template.ParseGlob(path.Dir(filename) + "/templates/*.html"))

func httpserve(w http.ResponseWriter, req *http.Request) {
	tmpl.ExecuteTemplate(w, "indexPage", nil)
}

// Get Artist results in JSON Format using GET request method
func getJSONartist(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	res := json.NewDecoder(r.Body).Decode(target)
	fmt.Println(res)
	return res
}

// Get Event results in JSON Format using GET request method
func getJSONevent(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(r.Body)
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func results(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	keyword := req.FormValue("query") // Get the keywords from the template

	// Construct the URL
	URLArtist := DEF_URL + keyword + "?app_id=go"
	URLEvent := DEF_URL + keyword + "/events?app_id=test"
	finalURLArtist, _ := url.Parse(URLArtist)
	finalURLEvent, _ := url.Parse(URLEvent)

	res := &Artist{}
	getJSONartist(finalURLArtist.String(), res)      // Get the artist data
	getJSONevent(finalURLEvent.String(), &res.Event) // Get the event data

	tmpl.ExecuteTemplate(w, "indexPage", res) // Return the data to the template
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", httpserve)
	http.HandleFunc("/results", results)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
