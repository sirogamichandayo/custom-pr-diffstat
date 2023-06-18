package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// The following parameters should be set:
	owner := "sirogamichandayo"
	repo := "custom-pr-diffstat"
	splinted := strings.Split(os.Getenv("GITHUB_REF"), "/")
	prNumber, err := strconv.Atoi(splinted[2])
	if err != nil {
		fmt.Printf("Error parsing PR number: %v\n", err)
		return
	}

	comment := &github.IssueComment{
		Body: github.String("This is a comment from Go code."),
	}

	_, _, err = client.Issues.CreateComment(ctx, owner, repo, prNumber, comment)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Comment successfully created.")
}
