package main

import "fmt"

var pl = fmt.Println

func main() {
	sl1 := make([]string, 6) // 為sl1 array宣告所需要的記憶體容量以及型態

	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	pl("Slice size :", len(sl1))                   // 查看sl1的長度
	pl("----------------------------------------") // 分段

	// 印出sl1 array裡面所有元素
	for i := 0; i < len(sl1); i++ {
		fmt.Print(sl1[i] + " ")
	}
	pl("\n----------------------------------------") // 分段

	// 不需要的值可以用底線( _ )省略
	for _, x := range sl1 {
		fmt.Print(x + " ")
	}
	pl("\n----------------------------------------") // 分段

	sArr := []int{1, 2, 3, 4, 5} // 宣告型態為int的sArr slice, 並設定初始值

	sl3 := sArr[0:2]                               // 宣告一個sl3 (pointer) 指向sArr[0:2]
	pl("sl3 :", sl3)                               // 印出sl3指向的值 (1, 2)
	pl("1st 3 :", sArr[:3])                        // 印出sArr slice (1, 2, 3)
	pl("Last 3 :", sArr[2:])                       // 印出sArr slice (3, 4, 5)
	pl("----------------------------------------") // 分段

	fmt.Print("Change sArr[0] to 10\n")
	sArr[0] = 10                                   // 如果sArr被修改了, 那理論上sl3也會被修改 (因為是pointer)
	pl("sl3 :", sl3)                               // 印出sl3檢查是否真的也被修改了 (10, 2)
	pl("----------------------------------------") // 分段

	fmt.Print("Change sl3[0] to 1\n")
	sl3[0] = 1                                     // 如果sl3的值被修改了, 那sArr的值也會被修改
	pl("sArr[:2] :", sArr[:2])                     // 1, 2
	pl("----------------------------------------") // 分段

	sl3 = append(sl3, 12)                          // sl3加入新的元素, 則sArr也會被更動到!!
	pl("sl3 :", sl3)                               // 1, 2, 12
	pl("sArr :", sArr)                             // 1, 2, 12, 4, 5
	pl("----------------------------------------") // 分段

	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])                         // nil
	pl("----------------------------------------") // 分段

	sl5 := make([]int, len(sArr)) // 如果不想要原本的sArr被修改到, 就要用make的方式來取得新的記憶體空間
	copy(sl5, sArr)               // 複製一份sArr到sl5
	pl("sl5 :", sl5)
	pl("sArr :", sArr)
	pl("----------------------------------------") // 分段

	sl5 = append(sl5, 99)
	pl("sl5 :", sl5)
	pl("sArr :", sArr)
}
