package projects

// ProjectStatus represents the status of a project in the system.
// It is an integer type that can be used to represent various states of a project,
// such as created, cloning, built, running, etc.
// The values are defined as constants for easy reference and comparison.
type ProjectStatus int

// StatusUnknown is used to represent an unknown status.
const StatusUnknown = -1

const (
	StatusCreated = iota + 1000
	StatusCloning
	StatusCloned
	StatusCheckingOut
	StatusCheckedOut
	StatusPulling // Currently pulling the latest changes from the repository
	StatusPulled  // Successfully pulled the latest changes from the repository
	StatusBuilding
	StatusBuilt
	StatusRunning
	StatusRun
	StatusStopping
	StatusStopped
	StatusDeleting
	StatusDeleted
	StatusTagging
	StatusTagged
	StatusFailed
	StatusError
)

func (s ProjectStatus) String() string {
	switch s {
	case StatusCreated:
		return "created"
	case StatusCloning:
		return "cloning"
	case StatusCloned:
		return "cloned"
	case StatusCheckingOut:
		return "checking-out"
	case StatusCheckedOut:
		return "checked-out"
	case StatusPulling:
		return "pulling"
	case StatusPulled:
		return "pulled"
	case StatusBuilding:
		return "building"
	case StatusBuilt:
		return "built"
	case StatusRunning:
		return "running"
	case StatusRun:
		return "run"
	case StatusStopping:
		return "stopping"
	case StatusStopped:
		return "stopped"
	case StatusDeleting:
		return "deleting"
	case StatusDeleted:
		return "deleted"
	case StatusTagging:
		return "tagging"
	case StatusTagged:
		return "tagged"
	case StatusError:
		return "error"
	case StatusFailed:
		return "failed"
	case StatusUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

func ProjectStatusFrom(s string) ProjectStatus {
	switch s {
	case "created":
		return StatusCreated
	case "cloning":
		return StatusCloning
	case "cloned":
		return StatusCloned
	case "checking-out":
		return StatusCheckingOut
	case "checked-out":
		return StatusCheckedOut
	case "building":
		return StatusBuilding
	case "built":
		return StatusBuilt
	case "running":
		return StatusRunning
	case "run":
		return StatusRun
	case "stopping":
		return StatusStopping
	case "stopped":
		return StatusStopped
	case "deleting":
		return StatusDeleting
	case "deleted":
		return StatusDeleted
	case "tagging":
		return StatusTagging
	case "tagged":
		return StatusTagged
	case "error":
		return StatusError
	case "failed":
		return StatusFailed
	case "unknown":
		return StatusUnknown
	default:
		return StatusUnknown
	}
}
