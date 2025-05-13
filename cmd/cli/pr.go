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
	"golang.org/x/oauth2"
)

type PRStructure struct {
	branch          string
	title           string
	head            string
	body            string
	gitToken        string
	ownerRepository string
	repositoryName  string
	typeFeature     string
}

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Create a pull request",
	Long:  `Create a pull request`,

	Run: func(cmd *cobra.Command, args []string) {

		branch, err := cmd.Flags().GetString("destination")
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

		head, err := cmd.Flags().GetString("source")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'head':", err.Error())
			os.Exit(1)
		}

		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'token':", err.Error())
			os.Exit(1)
		}

		owner, err := cmd.Flags().GetString("owner")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'owner':", err.Error())
			os.Exit(1)
		}

		repository, err := cmd.Flags().GetString("repository")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'repository':", err.Error())
			os.Exit(1)
		}

		typeFeature, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Erro ao obter a flag 'repository':", err.Error())
			os.Exit(1)
		}

		pr := PRStructure{
			branch:          branch,
			title:           title,
			head:            head,
			body:            body,
			gitToken:        token,
			ownerRepository: owner,
			repositoryName:  repository,
			typeFeature:     typeFeature,
		}
		createPR(&pr)
	},
}

func init() {

	rootCmd.AddCommand(prCmd)
	prCmd.Flags().StringP("destination", "d", "", "Destination branch")
	prCmd.Flags().StringP("title", "t", "", "PR title")
	prCmd.Flags().StringP("source", "s", "", "Source branch")
	prCmd.Flags().StringP("body", "b", "", "Create body")
	prCmd.Flags().String("type", "f", "Commit type (feat, fix, chore, docs, style, refactor, perf, test, ci, build, revert)")
	prCmd.Flags().StringP("token", "e", "", "Git Token Environment variable default (GIT_TOKEN)")
	prCmd.Flags().StringP("owner", "o", "", "Git owner repository  Environment variable default (GIT_OWNER)")
	prCmd.Flags().String("repository", "r", "Git repository name")

	prCmd.MarkFlagRequired("repository")
	prCmd.MarkFlagRequired("type")
}

func createPR(prData *PRStructure) {

	token := os.Getenv("GIT_TOKEN")
	if prData.gitToken != "" {
		token = prData.gitToken
	}
	if token == "" {
		log.Fatal("A variável de ambiente GITHUB_TOKEN não está definida.")
	}

	repoOwner := os.Getenv("GIT_OWNER")
	if prData.ownerRepository != "" {
		repoOwner = prData.ownerRepository
	}
	if repoOwner == "" {
		log.Fatal("A variável de ambiente GIT_OWNER não está definida.")
	}

	repoName := os.Getenv("GIT_REPO")
	if prData.repositoryName == "" && repoName == "" {
		log.Fatal("repository name is required")
	}

	if prData.repositoryName != "" {
		repoName = prData.repositoryName
	}

	ctx := context.Background()

	// Autenticação
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	title := fmt.Sprintf("%s: %s", prData.typeFeature, prData.title)
	newPR := &github.NewPullRequest{
		Title:               github.String(title),
		Head:                github.String(prData.head),
		Base:                github.String(prData.branch),
		Body:                github.String(prData.body),
		MaintainerCanModify: github.Bool(true), // Permite que o mantenedor modifique o PR
	}

	pr, _, err := client.PullRequests.Create(ctx, repoOwner, prData.repositoryName, newPR)
	if err != nil {
		log.Fatalf("Erro ao criar o Pull Request: %v", err)
	}

	fmt.Printf("Pull Request criado com sucesso! URL: %s\n", pr.GetHTMLURL())
}
