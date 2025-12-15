// Copyright IBM Corp. 2016, 2025

package executor

import "github.com/smartystreets/goconvey/web/server/contract"

type Parser interface {
	Parse([]*contract.Package)
}

type Tester interface {
	SetBatchSize(batchSize int)
	TestAll(folders []*contract.Package)
}
