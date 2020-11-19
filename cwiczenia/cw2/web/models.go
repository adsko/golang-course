package web

type issueOutput struct {
	ID	int64	`json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Priority string `json:"priority"`
}

