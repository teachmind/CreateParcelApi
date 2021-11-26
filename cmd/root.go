package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: 		"create-parcel-api",
	Short: 		"create-parcel-api",
	Long: 		`create-parcel-api`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init()  {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
