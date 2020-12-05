package web

import (
	"cw1/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Run () error {
	fmt.Println("Run server on port: ", "10000")
	http.HandleFunc("/api/get_issues", getIssues)
	http.HandleFunc("/api/add_issue", addIssue)
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(":10000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func getIssues(w http.ResponseWriter, r *http.Request){
	items, err := database.GetAllItems()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputElements := make([]issueOutput, len(items))

	for i, el := range items {
		priorityStr, err := database.PriorityAsString(el.Priority)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		outputElements[i] = issueOutput{
			createIssueInput{
				el.Name,
				priorityStr,
				el.Description,
			},
			0,
		}
	}

	jsonOutput, err := json.Marshal(outputElements)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(jsonOutput)
}

func addIssue(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out := createIssueInput{}
	err := json.NewDecoder(r.Body).Decode(&out)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	priority, isPriorityOk := database.IsValidPriority(out.Priority)
	switch {
	case len(out.Name) == 0:
		sendErrorMessage(w, "missing name or empty")
		return
	case len(out.Description) == 0:
		sendErrorMessage(w, "missing description or empty")
		return
	case !isPriorityOk:
		sendErrorMessage(w, "priority is missing or wrong value")
		return
	}

	_, err = database.AddItem(out.Name, out.Description, priority)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func sendErrorMessage (w http.ResponseWriter, message string) {
	w.WriteHeader(400)
	w.Write([]byte(message))
}

// TODO: Define url: /api/remove_issue/id -> POST
