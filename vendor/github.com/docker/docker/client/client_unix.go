// Copyright IBM Corp. 2016, 2025

// +build linux freebsd solaris openbsd darwin

package client

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/docker.sock"
