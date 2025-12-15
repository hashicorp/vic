// Copyright IBM Corp. 2016, 2025

package log

import "google.golang.org/grpc/grpclog"

func init() {
	// completely replace the grpc logger with the logrus logger.
	grpclog.SetLogger(L)
}
