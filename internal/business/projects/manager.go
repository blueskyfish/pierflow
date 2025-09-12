package projects

import (
	"path"
	"pierflow/internal/docker"
	"pierflow/internal/eventer"
	"pierflow/internal/gitter"
	"pierflow/internal/logger"
	"pierflow/internal/tasker"

	"github.com/glebarez/sqlite"
	"github.com/moby/moby/api/types/events"
	"gorm.io/gorm"
)

const DbMaxIdleConnections = 10  // Maximum number of idle connections in the pool
const DbMaxOpenConnections = 100 // Maximum number of open connections to the database
const DbConnMaxLifetime = 0      // Maximum amount of time a connection may be reused

type ProjectManager struct {
	db            *gorm.DB             // Database connection
	basePath      string               // Base path for all projects
	gitClient     gitter.GitClient     // Git client for repository operations
	taskClient    tasker.TaskClient    // Task client for task operations
	eventServe    eventer.EventServe   // Event serve for event operations with server-sent events
	composeClient docker.ComposeClient // Docker Compose client for docker events
}

// NewProjectManager creates a new instance of ProjectManager.
//
// The database is under the specified `dbPath` within the `basePath`.
// The `basePath` is the root directory where all project repositories will be managed.
func NewProjectManager(basePath, dbPath string) (*ProjectManager, error) {
	db, err := gorm.Open(sqlite.Open(path.Join(basePath, dbPath)), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(DbMaxIdleConnections) // Maximum number of idle connections in the pool
	sqlDb.SetMaxOpenConns(DbMaxOpenConnections) // Maximum number of open connections to the database
	sqlDb.SetConnMaxLifetime(DbConnMaxLifetime) // No limit on connection lifetime

	// Automatically migrate the DbProject model to create the table if it doesn't exist
	err = db.AutoMigrate(&DbProject{})
	if err != nil {
		logger.Errorf("Failed to auto-migrate models: %s", err.Error())
		return nil, err
	}

	return &ProjectManager{
		db:         db,
		basePath:   basePath,
		gitClient:  gitter.NewGitClient(basePath),
		taskClient: tasker.NewTaskClient(basePath),
		eventServe: eventer.NewEventServe(),
		composeClient: docker.NewComposeClient([]events.Action{
			events.ActionStart,
			events.ActionRestart,
			events.ActionStop,
			events.ActionDie,
			events.ActionKill,
		}),
	}, nil
}
