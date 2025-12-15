// Copyright IBM Corp. 2016, 2025

package ratelimiter

type Storage interface {
	GetBucketFor(string) (*LeakyBucket, error)
	SetBucketFor(string, LeakyBucket) error
}
