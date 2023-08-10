package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v28/github"
)

var client *github.Client

func registerAPI(g *gin.RouterGroup) {
	g.GET("/github/org", handleListGithubOrg)
	g.GET("/github/repo/:org", handleListGithubRepo)
	g.GET("/convert/:owner/:repo", handleConvert)
}

func handleConvert(c *gin.Context) {
	ctx := context.Background()

	owner := c.Param("owner")
	repo := c.Param("repo")

	fmt.Println(owner, repo)

	fileContent, _, _, err := client.Repositories.GetContents(
		ctx,
		owner,
		repo,
		".github/workflows/docker-build-push.yml",
		&github.RepositoryContentGetOptions{},
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(*fileContent.Type)

	//// The API response contains the file content encoded in base64
	//content, err := fileContent.GetContent()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//// Print the file contents
	//c.String(200, content)
}
