package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/google/go-github/v23/github"
)

// GitHub ...
type GitHub struct {
	Commits json.RawMessage `json:"commits"`
	After   string          `json:"after"`
}

// Commits :)
type Commits struct {
	Modified []string `json:"modified"`
}

func main() {
	fmt.Println("Starting server in port :8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	client := github.NewClient()

}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.github.com/repos/dealmeidafernando/testGithub/commits")
	if err != nil {
		panic(err)
	}

	bla := &GitHub{}
	jErr := json.NewDecoder(resp.Body).Decode(bla)
	if jErr != nil {
		panic(jErr)
	}
	fmt.Println(bla)

	// bodyBytes, _ := ioutil.ReadAll(r.Body)
	// bodyString := string(bodyBytes)
	//
	// fmt.Fprintf(w, "R E S P ===>", resp)
}
