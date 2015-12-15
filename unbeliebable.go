package main // grooveshark.im, groovesharks.org

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if strings.Contains(r.URL.Path, ".") {
			mux.ServeHTTP(w, r)
		} else {
			ip := net.ParseIP(strings.Trim(r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")], "[]"))

			switch r.Method + " " + r.URL.EscapedPath() {
			case "POST /playlist":
				name := r.FormValue("name")
				artist := r.FormValue("artist")
				youtube := struct{ ID string }{}

				if err := json.Unmarshal(getbody("http://groovesharks.org/music/getYoutube?track="+url.QueryEscape(name)+"&artist="+url.QueryEscape(artist)), &youtube); err != nil {
					panic(err)
				}

				playlist.Add(unbeliebable.Song{IP: ip, ID: youtube.ID, Time: time.Now(), Name: name, Artist: artist})

			case "GET /playlist":
				if body, err := json.Marshal(playlist.Songs); err != nil {
					panic(err)
				} else {
					w.Write(body)
				}

			case "GET /nowplaying":
				if body, err := json.Marshal(playlist.NowPlaying); err != nil {
					panic(err)
				} else {
					w.Write(body)
				}

			case "POST /elapsedtime":
				if body, err := ioutil.ReadAll(r.Body); err != nil {
					panic(err)
				} else if i, err := strconv.Atoi(string(body)); err != nil {
					panic(err)
				} else {
					playlist.ElapsedTime = i
				}

			case "GET /elapsedtime":
				w.Write([]byte(strconv.Itoa(playlist.ElapsedTime)))

			case "POST /votes":
				if against, err := strconv.ParseBool(r.FormValue("against")); err != nil {
					panic(err)
				} else {
					playlist.Vote(ip, r.FormValue("id"), against)
				}

			case "GET /search":
				w.Write(getbody("http://groovesharks.org/music/typeahead?query=" + url.QueryEscape(r.URL.Query().Get("q"))))

			default:
				http.ServeFile(w, r, "public/index.html")
			}
		}
	})
	http.ListenAndServe(":80", nil)
}
