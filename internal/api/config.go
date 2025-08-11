package api

import "embed"

type ServerConfig struct {
	Port     int       `json:"port"`
	Host     string    `json:"host"`
	DbPath   string    `json:"dbPath"`
	BasePath string    `json:"basePath"`
	Log      string    `json:"log"`
	Web      *embed.FS `json:"-"`
}
