package taskmodels

type Task struct {
	Title             string `json:"title"`
	Assignee_UserName string `json:"assignee_username"`
	Created_UserName  string `json:"created_username"`
	Description       string `json:"description"`
	Status            string `json:"status"`
}
