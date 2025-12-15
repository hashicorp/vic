// Copyright IBM Corp. 2016, 2025

package plugin

//go:generate protoc -I.:/usr/local --gogoswarm_out=import_path=github.com/docker/swarmkit/protobuf/plugin,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. plugin.proto
