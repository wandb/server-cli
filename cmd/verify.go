package cmd

import "github.com/spf13/cobra"

var verify = &cobra.Command{
	Use:   "verify",
	Short: "Runs a few tests to make sure the instance is configured correctly.",
	/*
	- should test if the bucket has the correct CORS policy
	- should test if the bucket has the right access policy
	- should test if EKS cluster has access to the bucket
	- should test if SQS queue is setup, if it is setup - check if it's subscribed to s3 notifications
	- should test if MySQL DB has automatic backups enabled
	- should test if DB and bucket have deletion_protection enabled
	*/

}

func init() {
	RootCmd.AddCommand(verify)
}
