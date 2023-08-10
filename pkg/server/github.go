package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
)

var orgList = []string{"Crossbell-Box", "NaturalSelectionLabs"}

func handleListGithubOrg(c *gin.Context) {
	ctx := context.Background()
	var orgs []*github.Organization

	for _, org := range orgList {
		o, _, _ := client.Organizations.Get(ctx, org)
		orgs = append(orgs, o)
	}

	c.JSON(http.StatusOK, orgs)
}

func handleListGithubRepo(c *gin.Context) {
	ctx := context.Background()

	org := c.Param("org")

	repos, _, err := client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
		Type:        "all",
		ListOptions: getListOption(c),
	})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, repos)
}

func getListOption(c *gin.Context) github.ListOptions {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("per_page"))
	return github.ListOptions{
		Page:    page,
		PerPage: perPage,
	}
}

func InitGithubClient(token string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}
