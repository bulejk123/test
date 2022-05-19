package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("")

	// 前景 背景 颜色
	// ---------------------------------------
	// 30  40  黑色
	// 31  41  红色
	// 32  42  绿色
	// 33  43  黄色
	// 34  44  蓝色
	// 35  45  紫红色
	// 36  46  青蓝色
	// 37  47  白色
	//
	// 代码 意义
	// -------------------------
	//  0  终端默认设置
	//  1  高亮显示
	//  4  使用下划线
	//  5  闪烁
	//  7  反白显示
	//  8  不可见

	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}

	var byAllCards [108]byte
	for i := 0; i < 54; i++ {
		for j := 0; j < 2; j++ {
			byAllCards[i+j*54] = byte(i) + 1
		}
	}
	//demo(byAllCards)
	//demo1(byAllCards)
}

func demo(byAllCards [108]byte){
	rand.Shuffle(len(byAllCards), func(i, j int) {
		byAllCards[i],byAllCards[j]= byAllCards[j],byAllCards[i]
	})
	fmt.Println(byAllCards)
}


func demo1(byAllCards [108]byte){
	for i := 0; i < 1000; i++ {
		rand_num := rand.Intn(1000)
		m := rand_num % (54 * 2)
		rand_num_2 := rand.Intn(1000)
		n := rand_num_2 % (54 * 2)
		zz := byAllCards[m]
		byAllCards[m] = byAllCards[n]
		byAllCards[n] = zz
	}
	fmt.Println(byAllCards)
}