package main

import (
	"encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	"junmusic"
	"net"
	"net/http"
	// "net/url"
	"os"
	"strings"
	"time"
)

var votes []junmusic.Vote
var requests []junmusic.Request

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

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := r.Method + " " + r.URL.EscapedPath()
	ip := net.ParseIP(strings.Trim(r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")], "[]"))

	switch cmd {
	case "POST /requests":
		request := junmusic.Request{IP: ip, Song: junmusic.Song{Name: "xyz"}}
		requests = append(requests, request)
		votes = append(votes, junmusic.Vote{IP: ip, Request: request, Time: time.Now()})
	case "POST /votes":
		votes = append(votes, junmusic.Vote{IP: ip, Time: time.Now()})
	case "GET /requests":
		body, err := json.Marshal(requests)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(body))
	case "GET /search":
		io.WriteString(w, string(getbody("http://groovesharks.org/music/typeahead?query="+r.URL.Query().Get("q"))))
	case "GET /id":
		q := r.URL.Query()
		youtube := struct{ ID string }{}
		err := json.Unmarshal(getbody("http://groovesharks.org/music/getYoutube?track="+q.Get("name")+"&artist="+q.Get("artist")), &youtube)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, youtube.ID)
	case "GET /":
		http.ServeFile(w, r, "public/index.html")
	default:
		http.NotFound(w, r)
	}
}

func main() {
	requests = append(requests, junmusic.Request{IP: net.ParseIP("::1"), Song: junmusic.Song{ID: "YgSPaXgAdzE", Name: "Loser", Artist: "Beck"}})

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
