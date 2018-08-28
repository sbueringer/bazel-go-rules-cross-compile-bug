package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  ``,
}

func main() {

	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("test", "t", "", "descr")

}
