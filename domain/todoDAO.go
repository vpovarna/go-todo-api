package domain

type TodoDAO struct {
	Id          int    `json:"id,omitempty", `
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed,omitempty"`
}
