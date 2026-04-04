package main 

import "fmt"

func main() {
	// if文
	x := 10
	if x > 5 {
		fmt.Println("xは5より大きいです")
	}
	
	// if-else
	age := 17
	if age >= 18 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}
	
	// if文の初期化文
	if err := doSomething(); err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("成功")
	}
}

func doSomething() error {
	return nil;
}