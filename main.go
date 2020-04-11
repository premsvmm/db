package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/premsvmm/db/cmd"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"log"
)

const (
	version = "1.0.7"
)

func main() {
	doSelfUpdate()
	cmd.Execute()
}

func doSelfUpdate() {
	selfupdate.EnableLog()
	v := semver.MustParse(version)
	fmt.Println(v)
	latest, err := selfupdate.UpdateSelf(v, "premsvmm/db")
	fmt.Println(latest)
	if err != nil {
		log.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}
