#!/bin/bash

changed_files=$(git diff --name-only origin/main...HEAD -- '*.go')

if [ -n "$changed_files" ]; then
  changed_dirs=$(echo "$changed_files" | xargs -n1 dirname | sort -u)
  for dir in $changed_dirs; do
    mise exec -- golangci-lint run --fix --config .golangci.yaml "$dir"
  done
else
  echo "No changed .go files to lint."
fi
