#!/usr/bin/env bash
# Copyright IBM Corp. 2016, 2025


GOOSARCH="${GOOS}_${GOARCH}"
case "$GOOSARCH" in
_* | *_ | _)
	echo 'undefined $GOOS_$GOARCH:' "$GOOSARCH" 1>&2
	exit 1
	;;
esac

GODEFS="go tool cgo -godefs"

$GODEFS types.go |gofmt > ztypes_$GOARCH.go

case $GOOS in
freebsd)
	$GODEFS types_$GOOS.go |gofmt > ztypes_$GOOSARCH.go
	;;
esac
