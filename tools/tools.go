//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/Khan/genqlient"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/swaggo/swag/cmd/swag"       // @v1.16.4
	_ "github.com/tilt-dev/ctlptl/cmd/ctlptl" // @v0.8.34
	_ "go.k6.io/k6"                           // @v0.54.0
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "honnef.co/go/tools/cmd/staticcheck" // @2024.1.1
	_ "sigs.k8s.io/kind"                   // @v0.24.0
	//  "github.com/goreleaser/goreleaser"
	//  "github.com/fullstorydev/grpcui/cmd/grpcui"
	//  "github.com/vektra/mockery/v2"
)
