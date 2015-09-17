package main

import (
	"dj"
	"encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	// "net/url"
	"os"
	"strings"
	"time"
)

var votes []dj.Vote
var requests []dj.Request

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

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := r.Method + " " + r.URL.EscapedPath()
	ip := net.ParseIP(strings.Trim(r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")], "[]"))

	switch cmd {
	case "POST /requests":
		request := dj.Request{IP: ip, Song: dj.Song{Name: "xyz"}}
		requests = append(requests, request)
		votes = append(votes, dj.Vote{IP: ip, Request: request, Time: time.Now()})
	case "POST /votes":
		votes = append(votes, dj.Vote{IP: ip, Time: time.Now()})
	case "GET /requests":
		body, err := json.Marshal(requests)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(body))
	case "GET /search":
		io.WriteString(w, string(getbody("http://grooveshark.im/music/typeahead?query="+r.URL.Query().Get("q"))))
	case "GET /id":
		q := r.URL.Query()
		youtube := struct{ ID string }{}
		err := json.Unmarshal(getbody("http://grooveshark.im/music/getYoutube?track="+q.Get("name")+"&artist="+q.Get("artist")), &youtube)
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
	requests = append(requests, dj.Request{IP: net.ParseIP("::1"), Song: dj.Song{ID: "YgSPaXgAdzE", Name: "Loser", Artist: "Beck"}})

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
