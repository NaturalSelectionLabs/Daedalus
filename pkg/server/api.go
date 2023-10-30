package server

import (
	"github.com/gin-gonic/gin"
)

//var client *github.Client

func registerAPI(g *gin.RouterGroup) {
	//g.GET("/github/org", handleListGithubOrg)
	//g.GET("/github/repo/:org", handleListGithubRepo)
	//g.GET("/github/content", handleContent)
	//g.GET("/convert/:owner/:repo", handleConvert)
}

//func handleConvert(c *gin.Context) {
//	ctx := context.Background()
//
//	owner := c.Param("owner")
//	repo := c.Param("repo")
//
//	fmt.Println(owner, repo)
//
//	_, directContent, _, err := client.Repositories.GetContents(
//		ctx,
//		owner,
//		repo,
//		".github/workflows",
//		&github.RepositoryContentGetOptions{},
//	)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	//fmt.Println(directContent)
//	c.JSON(http.StatusOK, directContent)
//
//	//// The API response contains the file content encoded in base64
//	//content, err := fileContent.GetContent()
//	//if err != nil {
//	//	fmt.Println("Error:", err)
//	//	return
//	//}
//	//
//	//// Print the file contents
//	//c.String(200, content)
//}
