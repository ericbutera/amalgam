package main

import "github.com/ericbutera/amalgam/services/api/cmd"

// TODO: @host should be set to the actual host, not "api:8080 or localhost:8080"

// @title           Feed API
// @version         1.0
// @scheme          http
// @host			localhost:8080
// @basepath        /v1

func main() {
	cmd.Execute()
}
