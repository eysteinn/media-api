/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/eysteinn/media-api/src/server"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve("8080")
		/*err := filescan.Scan(scandir)
		if err != nil {
			log.Fatal(err)
		}*/
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	//scanCmd.Flags().StringVarP(&scandir, "directory", "d", scandir, "Directory to scan")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
