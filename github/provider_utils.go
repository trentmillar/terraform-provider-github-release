package github

import (
	"os"
)

var testOrganization string = testOrganizationFunc()
var testToken string = os.Getenv("GITHUB_TOKEN")

func testOrganizationFunc() string {
	organization := os.Getenv("GITHUB_ORGANIZATION")
	if organization == "" {
		organization = os.Getenv("GITHUB_TEST_ORGANIZATION")
	}
	return organization
}
