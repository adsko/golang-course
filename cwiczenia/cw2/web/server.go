package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Run () error {
	fmt.Println("Run server on port: ", "10000")
	http.HandleFunc("/api/get_issues", getIssues)
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(":10000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func getIssues(w http.ResponseWriter, r *http.Request){
	err := outputJSON(w, [1]issueOutput{{
		344,
		"Test",
		"Description",
		"high",
	}})

	if err != nil {
		log.Panic(err)
	}
}

func outputJSON(w http.ResponseWriter, v interface{}) error {
	out, err := json.Marshal(v)

	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(out)
	return err
}

// TODO: Define url: /api/add_issue -> {name: string, priority: string, description: string}, POST
// TODO: Define url: /api/remove_issue/id -> POST
