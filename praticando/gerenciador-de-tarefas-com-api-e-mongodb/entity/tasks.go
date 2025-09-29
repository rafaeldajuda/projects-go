package entity

type Task struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Completed   *bool  `json:"completed,omitempty"`
}
