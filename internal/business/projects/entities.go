package projects

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbProject struct {
	ID          string        `json:"id" gorm:"primaryKey;index:uni_project_id;not null"`
	Name        string        `json:"name" gorm:"not null"`
	Description string        `json:"description" gorm:"default:''"`
	Path        string        `json:"path" gorm:"not null"`
	GitUrl      string        `json:"gitUrl" gorm:"not null"`
	Branch      string        `json:"branch" gorm:"default:''"`   // Branch: default: main
	Taskfile    string        `json:"taskfile" gorm:"default:''"` // Taskfile: default: Taskfile.yml
	User        string        `json:"user" gorm:"not null"`
	Token       string        `json:"token" gorm:"not null"`
	Creation    int64         `json:"creation" gorm:"autoCreateTime"`
	Modified    int64         `json:"updated" gorm:"autoUpdateTime"`
	Status      ProjectStatus `json:"status" gorm:"not null"`
}

func (p *DbProject) TableName() string {
	return "pierflow_projects"
}

func (p *DbProject) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

// DbEvent represents an event in the system, such as project changed, docker started, etc.
type DbEvent struct {
	ID        string `json:"id" gorm:"primaryKey;index:uni_event_id;not null"`       // Unique event ID
	Group     string `json:"group" gorm:"not null"`                                  // required: Group to identify related events e.g "project", "docker", "system"
	Event     string `json:"event" gorm:"not null"`                                  // required: The event name, e.g., "checkout-repository", "build-project"
	ValueID   string `json:"value_id" gorm:"primaryKey;index:idx_value_id;not null"` // required: ID of the related entity, e.g., project ID, docker container ID
	Value     string `json:"value" gorm:"not null"`                                  // JSON serialized value of the event
	Timestamp int64  `json:"timestamp" gorm:"not null"`                              // The unix epoch timestamp when the event occurred
}

func (e *DbEvent) TableName() string {
	return "pierflow_events"
}

func (e *DbEvent) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	if e.ValueID == "" {
		return errors.New("value id is required")
	}
	if e.Event == "" {
		return errors.New("event is required")
	}
	if e.Group == "" {
		return errors.New("group is required")
	}
	if e.Timestamp == 0 {
		e.Timestamp = tx.NowFunc().Unix()
	}
	return nil
}
