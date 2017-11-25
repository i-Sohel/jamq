package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

type Artist struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Furl     string `json:"facebook_page_url"`
	Event    []Event_info
}

type Event_info struct {
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

/*type Combine struct {
	Artist
	Event
}*/

const (
	DEFAULT_PORT = "8080"
	size         = "5"
	DEF_URL      = "https://rest.bandsintown.com/artists/" // API URL
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func httpserve(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "indexPage", nil)
}

// Get Twitter Serach results in JSON Format using GET request method
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

	URLArtist := DEF_URL + keyword + "?app_id=go"
	URLEvent := DEF_URL + keyword + "/events?app_id=test"
	finalURLArtist, _ := url.Parse(URLArtist)
	finalURLEvent, _ := url.Parse(URLEvent)
	fmt.Println(finalURLEvent)
	res := &Artist{}
	getJSONartist(finalURLArtist.String(), res)
	getJSONevent(finalURLEvent.String(), &res.Event) // Get the data
	// Return the data to the template
	templates.ExecuteTemplate(w, "indexPage", res)
}

func about(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "aboutPage", nil)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", httpserve)
	http.HandleFunc("/results", results)
	http.HandleFunc("/about", about)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
