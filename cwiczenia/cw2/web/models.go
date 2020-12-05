package web

type createIssueInput struct {
	Name string	`json:"name"`
	Priority string `json:"priority"`
	Description string `json:"description"`
}

type issueOutput struct {
	createIssueInput
	ID int64 `json:"id"`
}

