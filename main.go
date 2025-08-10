package main

import "pierflow/cmd"

// main is the entry point of the application.
// It executes the root command defined in cmd package.
func main() {
	err := cmd.CommandRoot.Execute()
	if err != nil {
		panic(err)
	}
}
