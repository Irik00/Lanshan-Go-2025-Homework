package main

import "fmt"

func average(sum int, count int) float64 {
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
func main() {
	var count, sum, i int
	for {
		fmt.Println("请输入一个整数（输入0结束）：")
		fmt.Scanln(&i)
		if i == 0 {
			break
		}
		sum += i
		count++
	}
	k := average(sum, count)
	if k >= 60 {
		fmt.Printf("平均成绩为 %.2f,成绩合格", k)
	} else {
		fmt.Printf("平均成绩为 %.2f,成绩不合格", k)
	}

}
