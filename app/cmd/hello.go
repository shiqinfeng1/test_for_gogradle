package cmd

import (
	"fmt"
	"errors"
	"github.com/spf13/cobra"
)

func init() {
  	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Short: "hello",
	Args: func(cmd *cobra.Command, args []string) error {
	  if len(args) < 1 {
		return errors.New("requires at least one arg")
	  }
	  /*
	  if myapp.IsValidColor(args[0]) {
		return nil
	  }
	  */
	  return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Hello, World!")
	},
  }