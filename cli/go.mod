module github.com/ericbutera/amalgam/cli

go 1.23.2

require (
	github.com/ericbutera/amalgam/pkg v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

// https://go.dev/doc/tutorial/call-module-code
replace github.com/ericbutera/amalgam/pkg => ../pkg
