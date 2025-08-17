package serve

import (
	"embed"
	"pierflow/internal/api"

	"github.com/spf13/cobra"
)

//go:embed web
var web embed.FS

var CommandServe = &cobra.Command{
	Use:     "serve",
	Short:   "Start the Pierflow server",
	Long:    `Start the Pierflow server to manage your docker workflow and provide API access.`,
	Example: `pierflow server --port=8080 --host=localhost --dbPath=pierflow.db --log=info --base-path=/path/to/repositories`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")
		dbPath, _ := cmd.Flags().GetString("db-path")
		log, _ := cmd.Flags().GetString("log")
		basePath, _ := cmd.Flags().GetString("base-path")

		options := api.ServerConfig{
			Port:     port,
			Host:     host,
			DbPath:   dbPath,
			Log:      log,
			BasePath: basePath,
			Web:      &web,
		}

		return api.StartApiServer(&options)
	},
}

func init() {
	CommandServe.PersistentFlags().Int("port", 8080, "Port to run the server on")
	CommandServe.PersistentFlags().String("log", "info", "Log level for the server (debug, info, warn, error, fatal, panic)")
	CommandServe.PersistentFlags().String("host", "localhost", "Host to bind the server to")
	CommandServe.PersistentFlags().String("db-path", "pierflow.db", "Path to the database file")
	CommandServe.PersistentFlags().String("base-path", ".", "Base path for the git repositories and the database file")
}
