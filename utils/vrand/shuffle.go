package vrand

import (
	"math/rand"
	"reflect"
	"time"
)

// Fisher–Yates shuffle

func ShuffleSlice(arr interface{}) {
	contentType := reflect.TypeOf(arr)
	if contentType.Kind() != reflect.Slice {
		panic("expects a slice type")
	}
	contentValue := reflect.ValueOf(arr)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	length := contentValue.Len()
	for i := length - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		x, y := contentValue.Index(i).Interface(), contentValue.Index(j).Interface()
		contentValue.Index(i).Set(reflect.ValueOf(y))
		contentValue.Index(j).Set(reflect.ValueOf(x))
	}
}

func ShuffleStringSlice(x []string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(x) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
}
