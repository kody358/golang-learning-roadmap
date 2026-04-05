package main

import "fmt"

func main() {
	// forループ
	for i :=0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	
	// for-range
	fruits := []string{"りんご", "みかん", "ぶどう"}
	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}
	// indexを捨てる場合
	for _, fruit := range fruits {
		fmt.Println("果物:", fruit)
	}
	
	// mampのfor-range
	colors := map[string]string {
		"red": "赤",
		"blue": "青",
		"green": "緑",
	}
	for k, v := range colors {
		fmt.Printf("%s = %s\n", k, v)
	}
}
