package version

import "fmt"

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/opts"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/main"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "74f62c5fa8f5efc30e79449339b0d626e6b61c37"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-04-14T11:11:16+02:00"
	// Tag lists the Tag on the podbuild, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.0.13+"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/p9c/opts/"
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"Package Info:\n"+
		"	git repository: "+URL+"\n",
		"	branch: "+GitRef+"\n"+
		"	commit: "+GitCommit+"\n"+
		"	built: "+BuildTime+"\n"+
		"	Tag: "+Tag+"\n",
	)
}
