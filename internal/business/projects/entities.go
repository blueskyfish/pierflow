package projects

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbProject struct {
	ID          string        `json:"id" gorm:"primaryKey;index:uni_project_id;not null"`
	Name        string        `json:"name" gorm:"not null"`
	Description string        `json:"description"`
	Path        string        `json:"path" gorm:"not null"`
	GitUrl      string        `json:"gitUrl" gorm:"not null"`
	Branch      string        `json:"branch"`
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
