package main

import (
	"encoding/json"
	// "fmt"
	// "io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Song struct {
	IP               net.IP
	Time             time.Time
	ID, Name, Artist string
}

type Vote struct {
	IP      net.IP
	Time    time.Time
	ID      string
	Against bool
}

var playlist []Song
var votes []Vote

type player struct{}

func (m *player) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method + " " + r.URL.EscapedPath() {
	case "GET /shift":
		if len(playlist) > 0 {
			song := playlist[0]
			playlist = playlist[1:]
			w.Write([]byte(song.ID))
		} else {
			http.NotFound(w, r)
		}
	default:
		http.ServeFile(w, r, "player.html")
	}
}

func getbody(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

// grooveshark.im
// groovesharks.org

func client(w http.ResponseWriter, r *http.Request) {
	ip := net.ParseIP(strings.Trim(r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")], "[]"))

	switch r.Method + " " + r.URL.EscapedPath() {
	case "POST /playlist":
		name := r.FormValue("name")
		artist := r.FormValue("artist")
		youtube := struct{ ID string }{}
		err := json.Unmarshal(getbody("http://grooveshark.im/music/getYoutube?track="+url.QueryEscape(name)+"&artist="+url.QueryEscape(artist)), &youtube)
		if err != nil {
			panic(err)
		}
		song := Song{IP: ip, ID: youtube.ID, Time: time.Now(), Name: name, Artist: artist}
		playlist = append(playlist, song)
		votes = append(votes, Vote{IP: ip, ID: youtube.ID, Time: time.Now()})
	case "GET /playlist":
		body, err := json.Marshal(playlist)
		if err != nil {
			panic(err)
		}
		w.Write(body)
	case "POST /votes":
		votes = append(votes, Vote{IP: ip, ID: r.FormValue("id"), Time: time.Now()})
	case "GET /search":
		w.Write(getbody("http://grooveshark.im/music/typeahead?query=" + url.QueryEscape(r.URL.Query().Get("q"))))
	default:
		http.ServeFile(w, r, "client.html")
	}
}

func main() {
	go http.ListenAndServe("localhost:"+os.Getenv("PORT"), &player{})

	http.HandleFunc("/", client)
	http.ListenAndServe(":80", nil)
}
