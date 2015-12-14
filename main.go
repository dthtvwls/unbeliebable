package main

// grooveshark.im
// groovesharks.org

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unbeliebable"
)

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

func main() {
	playlist := unbeliebable.Playlist{}
	go http.ListenAndServe("localhost:"+os.Getenv("PORT"), &unbeliebable.Player{Playlist: &playlist})

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".") {
			mux.ServeHTTP(w, r)
		} else {
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
				song := unbeliebable.Song{IP: ip, ID: youtube.ID, Time: time.Now(), Name: name, Artist: artist}
				playlist.Add(song)
				// votes = append(votes, Vote{IP: ip, ID: youtube.ID, Time: time.Now()})
			case "GET /playlist":
				body, err := json.Marshal(playlist.Songs)
				if err != nil {
					panic(err)
				}
				w.Write(body)
			case "POST /votes":
				// votes = append(votes, Vote{IP: ip, ID: r.FormValue("id"), Time: time.Now()})
			case "GET /search":
				w.Write(getbody("http://grooveshark.im/music/typeahead?query=" + url.QueryEscape(r.URL.Query().Get("q"))))
			default:
				http.ServeFile(w, r, "public/index.html")
			}
		}
	})
	http.ListenAndServe(":80", nil)
}
