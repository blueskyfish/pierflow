package api

import "embed"

type ServerConfig struct {
	Port     int
	Host     string
	DbPath   string
	BasePath string
	Log      string
	Web      *embed.FS
}
