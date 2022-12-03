package cmd

import "github.com/spf13/cobra"

var logs = &cobra.Command{
	Use:   "logs",
	Short: "Downloads instance debug bundle.",
	/* 
	Need to check if the user requesting debug bundle 
	is authenticated as admin. If they are we download
	the debug bundle. If not, throw an error and to 
	contact admin user for debug bundle.
	*/
	/*
	- We might need to check the role of the user against the database
    - This needs to connect to the database in whichever env it's running
	- [nice] Check
	*/
}

func init() {
	RootCmd.AddCommand(logs)
}
