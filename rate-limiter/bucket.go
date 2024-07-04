package main

import (
	"time"
)

const (
	MAX_BUCKET_SIZE = 4
	REFILL_RATE     = 1
)

type TockenBucket struct {
	curentBucketSize int
	lastRefillTime   int
}

func NewBucket(currSize int, rate int) *TockenBucket {
	return &TockenBucket{
		curentBucketSize: currSize,
		lastRefillTime:   rate,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t *TockenBucket) refillRate() {
	nowTime := time.Now().Nanosecond()
	tokesToAdd := (nowTime - t.lastRefillTime) * REFILL_RATE
	t.curentBucketSize = min(t.curentBucketSize+tokesToAdd, MAX_BUCKET_SIZE)
	t.lastRefillTime = nowTime
}

func (t *TockenBucket) AllowRequest(tokenReq int) bool {
	t.refillRate()
	if t.curentBucketSize >= tokenReq {
		t.curentBucketSize -= tokenReq
		return true
	}
	return false
}
