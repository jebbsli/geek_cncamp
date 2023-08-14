package main

import "fmt"

/*
 * 给定一个字符串数组
 * ["I", "am", "stupid", "and", "weak"]
 * 用for循环遍历该数组并修改为
 * ["I", "am", "smart", "and", "strong"]
 */
func main() {
	sourceArray := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("修改前：")
	for _, v := range sourceArray {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

	for i, v := range sourceArray {
		if v == "stupid" {
			sourceArray[i] = "smart"
		} else if v == "weak" {
			sourceArray[i] = "strong"
		}
	}

	fmt.Println("修改后：")
	for _, v := range sourceArray {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}
