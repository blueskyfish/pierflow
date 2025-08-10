package tasker

type TaskItem struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Summary string `json:"summary,omitempty"`
}
