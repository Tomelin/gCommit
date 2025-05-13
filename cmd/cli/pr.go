/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v57/github"
	"github.com/spf13/cobra"
)

type PRStructure struct {
	branch string
	title  string
	head   string
	body   string
}

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Create a pull request",
	Long:  `Create a pull request`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pr called")
		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'branch':", err.Error())
			os.Exit(1)
		}
		if branch == "" {
			fmt.Println("A flag 'branch' é obrigatória")
			os.Exit(1)
		}

		title, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'title':", err.Error())
			os.Exit(1)
		}

		if title == "" {
			fmt.Println("A flag 'title' é obrigatória")
			os.Exit(1)
		}

		body, err := cmd.Flags().GetString("body")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'body':", err.Error())
			os.Exit(1)
		}
		if body == "" {
			fmt.Println("A flag 'body' é obrigatória")
			os.Exit(1)
		}

		head, err := cmd.Flags().GetString("head")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'head':", err.Error())
			os.Exit(1)
		}

		log.Println("Title:", title)
		log.Println("Branch:", branch)
		log.Println("Body:", body)
		log.Println("Head:", head)

		pr := PRStructure{
			branch: branch,
			title:  title,
			head:   head,
			body:   body,
		}
		createPR(&pr)
	},
}

func init() {

	rootCmd.AddCommand(prCmd)
	prCmd.Flags().StringP("branch", "b", "", "Destination branch")
	prCmd.Flags().StringP("title", "t", "", "PR title")
	prCmd.Flags().StringP("head", "x", "", "Feature branch")
	prCmd.Flags().StringP("body", "d", "", "Create body")

}

func createPR(prData *PRStructure) {

	client := github.NewClient(nil)
	newPR := &github.NewPullRequest{
		Title:               github.String(prData.title),
		Head:                github.String(prData.head),
		Base:                github.String(prData.branch),
		Body:                github.String(prData.body),
		MaintainerCanModify: github.Bool(true), // Permite que o mantenedor modifique o PR
	}
	ctx := context.Background()
	pr, response, err := client.PullRequests.Create(ctx, "Tomelin", "gCommit", newPR)
	if err != nil {
		log.Fatalf("Erro ao criar o Pull Request: %v", err)
	}

	log.Println(response)
	fmt.Printf("Pull Request criado com sucesso! URL: %s\n", pr.GetHTMLURL())

}
