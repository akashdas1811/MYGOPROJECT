package main

import "fmt"

func main() {
	tokenBucketObj := NewBucket(3, 0)

	fmt.Printf("Request processed: %v\n", tokenBucketObj.AllowRequest(1))
	fmt.Printf("Request processed: %v\n", tokenBucketObj.AllowRequest(1))

}
