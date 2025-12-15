// Copyright IBM Corp. 2016, 2025

//go:generate protoc -I.:../../../../vendor --gogoswarm_out=plugins=deepcopy,import_path=github.com/docker/swarmkit/protobuf/plugin/deepcopy/test:. deepcopy.proto

package test
