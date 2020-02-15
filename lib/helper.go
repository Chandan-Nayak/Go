package lib

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"time"
)

func GetFuncName(fc interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fc).Pointer()).Name()
}

func Unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GenerateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func GenerateRandomNumber(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size)
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s > %s\n\n", name, elapsed)
}

func GetTimeComparission(m map[time.Duration]string) {
	var keys []time.Duration
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, key := range keys {
		fmt.Printf("\nFunction: %s Time: %v Times: %.2v Times",m[key], keys[i], keys[i].Seconds()/keys[0].Seconds())
	}
}
