#!/bin/bash
echo "mode: atomic" > coverage.out

go test ./... -v -coverprofile=profile.out -covermode=atomic; exit_code=$?

if [[ $exit_code -eq 1 ]]; then
    echo "Test Error"
    exit 1
fi

if [ -f profile.out ]; then
    tail -n +2 profile.out >> coverage.out; rm profile.out
    go tool cover -func=coverage.out
fi
