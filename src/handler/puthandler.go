package handler

import (
	"codes"
)

func put(bucket string, key string, value string) int {

	lock.RLock()
	l := bucketsLock[bucket]
	defer lock.Unlock()

	l.Lock()
	if bucketMap, ok := buckets[bucket]; ok {
		bucketMap[key] = value
		return codes.OK
	} else {
		return codes.BUCKET_NOT_EXIST
	}
	defer l.Unlock()
	return codes.OK
}
