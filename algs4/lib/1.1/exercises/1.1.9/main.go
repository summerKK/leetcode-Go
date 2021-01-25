package main

import (
	"fmt"
	"math/rand"
)

// 1.1.9 编写一段代码，将一个正整数 N 用二进制表示并转换为一个 String 类型的值 s。
func main() {
	N := rand.Intn(1<<31 - 1)
	fmt.Printf("N:%d\n", N)
	s := ""
	// 100转换为二进制:
	// 100/2=50....(余数为0);
	// 50/2=25.....(余数为0);
	// 25/2=12.....(余数为1);
	// 12/2=6......(余数为0);
	// 6/2=3.......(余数为0);
	// 3/2=1.......(余数为1);
	// 1/2=0.......(余数为1);
	// 所以100的二进制表示形式为1100100;
	for i := N; i > 0; i /= 2 {
		s = fmt.Sprintf("%d", i%2) + s
	}

	fmt.Print(s)
}
