#!/bin/sh

export PLUGIN_DEBUG=true
export PLUGIN_CLUSTER=sandbox-1
export DRONE_REPO_NAME=demo
export PLUGIN_ZONE=us-east1-d
export PLUGIN_PROJECT=yebo-project
export PLUGIN_CREDENTIALS=$(cat ./secrets/yebo.json)
export CLOUDSDK_CONTAINER_USE_CLIENT_CERTIFICATE=True

go run main.go
