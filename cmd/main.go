package main

import (
	"fmt"
	"os"

	"github.com/Tomelin/go-convetional-commit/internal/entity"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "commit",
		Short: "Commit a message",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide a commit message")
				return
			}
			commitMessage := args[0]
			fmt.Printf("Commit message received: %s\n", commitMessage)

			options := entity.Option(0)

			prompt := promptui.Select{
				Label: "Select Day",
				Items: options.OptionsList(),
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %q\n", result)
		},
	}

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
