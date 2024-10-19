//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/swaggo/swag/cmd/swag@latest"
	_ "golang.org/x/lint/golint"
	//  "github.com/goreleaser/goreleaser"
	//  "github.com/fullstorydev/grpcui/cmd/grpcui"
	//  "github.com/golang/protobuf/protoc-gen-go"
	//  "github.com/vektra/mockery/v2"
)
