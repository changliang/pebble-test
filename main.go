package main

import (
	"awesomeProject3/pebble"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	number = flag.Int("n", 10000000, "total number")
	size   = flag.Int("s", 32000, "feature size")
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	flag.Parse()

	featureDB, err := pebble.NewFeatureDB("/data/pebble/test")
	if err != nil {
		println("init err")
		return
	}

	for i := 0; i < *number; i++ {
		randVal := String(*size)
		randKey := String(32)
		begin := time.Now()
		err := featureDB.AddFeature(randKey, []byte(randVal))
		if err != nil {
			fmt.Printf("add feature err: %s\n", err)
		}

		cost := time.Since(begin).Nanoseconds() / 1000000
		fmt.Printf("AddFeature cost:%dms\n", cost)

		if cost > 5000 {
			fmt.Printf("slow add cost:%dms\n", cost)
		}

		if i%1000 == 0 {
			fmt.Printf("processed %d, %f\n", i, float32(i)/float32(*number))
		}
	}

}
