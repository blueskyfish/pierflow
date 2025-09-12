package docker

import "github.com/moby/moby/api/types/events"

// ComposeEvent represents a Docker Compose related-event with relevant details.
//
// It transforms from a event.Message from Docker SDK to a more convenient structure for Compose events
// with toComposeEvent function.
type ComposeEvent struct {
	ID         string        `json:"id"`          //The actor id
	Action     events.Action `json:"action"`      // action type, e.g., "start", "stop", "create", etc.
	Project    string        `json:"project"`     // project name
	WorkingDir string        `json:"working_dir"` // absolute path of the project working directory
	Image      string        `json:"image"`       // image name
	Name       string        `json:"name"`        // container name
	Service    string        `json:"service"`     // service name
	Container  string        `json:"container"`   // container number as string
	Time       int64         `json:"time"`        // event timestamp as unix time
}

/*
  {
    "status": "create",
    "id": "b6278a486f6a82b96fbf8a883a9d48e1acee8e486068f04e40a2736e0b7a3c89",
    "from": "mockary-web-mockary",
    "Type": "container",
*   "Action": "create",
    "Actor": {
*     "ID": "b6278a486f6a82b96fbf8a883a9d48e1acee8e486068f04e40a2736e0b7a3c89",
      "Attributes": {
        "com.docker.compose.config-hash": "bacc4deedfe7537843f0592b9eb08a9a6e655bb10df455a8316d7536b6337389",
*       "com.docker.compose.container-number": "1",
        "com.docker.compose.depends_on": "",
        "com.docker.compose.image": "sha256:c7d7d5178e0b74869e052ed196fbfccbf938260c1220c3e793d2dc5641ce876c",
        "com.docker.compose.oneoff": "False",
*       "com.docker.compose.project": "mockary-web",
        "com.docker.compose.project.config_files": "/Users/sarah/Projects/Kirchnerei/github.com/blueskyfish/pierflow/var/projects/mockary/prod.compose.yml",
*       "com.docker.compose.project.working_dir": "/Users/sarah/Projects/Kirchnerei/github.com/blueskyfish/pierflow/var/projects/mockary",
*       "com.docker.compose.service": "mockary",
        "com.docker.compose.version": "2.39.2",
        "desktop.docker.io/ports.scheme": "v2",
        "desktop.docker.io/ports/8080/tcp": ":30080",
*       "image": "mockary-web-mockary",
        "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>",
*       "name": "blueskyfish_mockary"
      }
    },
    "scope": "local",
*   "time": 1757623133,
    "timeNano": 1757623133084229927
  },

*/
