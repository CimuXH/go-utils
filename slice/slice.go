package slice

import (
	"errors"
)

// Ordered 自定义约束(整型，浮点型，字符串)
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 |
	~string
}

// MaxSlice 返回切片中的最大值(整型，浮点型，字符串切片)
func MaxSlice[T Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var t T
		return t, errors.New("slice is empty")
	}

	// 遍历切片，查找最大值
	maxp := slice[0]
	for _, v := range slice[1:] {
		if v > maxp {
			maxp = v
		}
	}
	return maxp, nil
}

// MinSlice 返回切片中的最小值(整型，浮点型，字符串切片)
func MinSlice[T Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var t T
		return t, errors.New("slice is empty")
	}

	// 遍历切片，查找最小值
	minp := slice[0]
	for _, v := range slice[1:] {
		if v < minp {
			minp = v
		}
	}
	return minp, nil
}

// SliceIndex 在切片中查找 k 的下标并返回(k可以是整型，浮点型，字符串，结构体(字段必须都是comparable字段))
func SliceIndex[T comparable](slice []T, k T) int {
	for i, v := range slice {
		if v == k {
			return i
		}
	}

	// 未找到返回 -1
	return -1
}
