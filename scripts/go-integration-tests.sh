#!/bin/bash
INTEGRATION_TEST_PATTERN="*_integration_test.go"

packages=$(find . -name $INTEGRATION_TEST_PATTERN -exec dirname {} \; | sort -u | xargs go list)

packages="$packages ./integration"

if [ -z "$packages" ]; then
    echo "No integration tests found"
else
    echo "Running integration tests in packages:"
    echo "$packages"
    go test -v $packages
fi
