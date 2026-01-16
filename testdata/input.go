package testdata

import "time"

type Simple struct {
	Name *string
	Age  *int
}

type Complex struct {
	Tags    *[]string
	Timeout *time.Duration
	// Should not generate for these
	Normal string
	Slice  []int
}
