package projects

type CreatePayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Path        string `json:"path" validate:"required"`
	GitUrl      string `json:"giturl" validate:"required,url"`
	User        string `json:"user" validate:"required"`
	Token       string `json:"token" validate:"required"`
}

type CommandPayload struct {
	Message string `json:"message" validate:"required"`
}

type CheckoutPayload struct {
	Branch  string `json:"branch" validate:"required"`
	Place   string `json:"place" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// deprecated: It is removed
type TaskFileProjectPayload struct {
	TaskFile string `json:"taskfile" validate:"required"`
	Message  string `json:"message" validate:"required"`
}

type ChangeTaskfilePayload struct {
	Taskfile string `json:"taskfile" validate:"required"`
}
