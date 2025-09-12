package api

import (
	"embed"

	"github.com/moby/moby/api/types/events"
)

type ServerConfig struct {
	Port          int
	Host          string
	DbPath        string
	BasePath      string
	Log           string
	DockerActions []events.Action
	Web           *embed.FS
}
