package main

import (
	"fmt"
	"github.com/drael/GOnetstat"
	"log"
	"net/http"
	"strconv"
)

const (
	localPort = 8001
)

var allowedUsers = []string{"reda", "pi"}

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(localPort), nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	tcp_data := GOnetstat.Tcp()
	searchStr := r.RemoteAddr
	//fmt.Fprint(w, tcp_data)
	for _, p := range tcp_data {
		if p.ForeignPort == localPort && searchStr == fmt.Sprintf("%s:%d", p.Ip, p.Port) {
			if contains(allowedUsers, p.User) {
				fmt.Fprintf(w, "Remote = %s\n", r.RemoteAddr)
				fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
				fmt.Fprintf(w, "User %s passed the test\n", p.User)
			} else {
				fmt.Fprintf(w, "User %s didn't pass the test\n", p.User)
			}
			return
		}
	}
	fmt.Fprintf(w, "You are not running on localhost! go away!\n")
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
