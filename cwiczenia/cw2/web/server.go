package web

import (
	"fmt"
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
	// {name: string, priority: string, description: string, id: int64}
}


// TODO: Define url: /api/add_issue -> {name: string, priority: string, description: string}, POST
// TODO: Define url: /api/remove_issue/id -> POST
