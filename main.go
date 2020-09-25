package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/monaco-io/request"
)

var letters = "qazwsxedcrfvtgbyhnujmikolp"
var digits = "1234567890"
var discoveries = []string{}
var domainExtensions = []string{".com", ".net", ".org"}
var searches = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()
	router.HandleFunc("/spread", func(w http.ResponseWriter, r *http.Request) {

		site := findRandomSite(rand.Intn(8-5) + 5)
		io.WriteString(w, "<a href=' http://"+site+"'>site</a>")

	})
	router.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<a href=' http://"+discoveries[rand.Intn(len(discoveries)-1)]+"'>site</a>")

	})
	log.Println("-> http://0.0.0.0:8079")
	panic(http.ListenAndServe(os.Getenv("PORT"), router))
}

func findRandomSite(length int) string {
	retVal := ""
	lettersSlice := strings.Split(letters, "")
	digitsSlice := strings.Split(digits, "")
	//Create
	attempt := func() string {
		searches++
		build := ""
		for iteration := 0; iteration < length; iteration++ {
			typeDecision := rand.Intn(6)
			if typeDecision < 5 {
				build += lettersSlice[rand.Intn(25)]
			} else {
				build += digitsSlice[rand.Intn(9)]

			}
		}
		build += domainExtensions[rand.Intn(3)]
		log.Println("Searches:", searches)
		log.Println("Discoveries:", len(discoveries))
		if checkSite(build) {
			return build
		}
		return ""
	}
	for retVal == "" {
		retVal = attempt()

	}
	log.Println("Success!")
	discoveries = append(discoveries, retVal)
	return retVal
}
func checkSite(s string) bool {
	client := request.Client{
		URL:     "http://" + s,
		Method:  "GET",
		Timeout: 2, //seconds
	}
	r, err := client.Do()
	if err != nil {
		return false
	}

	if len(r.Data) < 60 {
		return false
	}

	return true
}
