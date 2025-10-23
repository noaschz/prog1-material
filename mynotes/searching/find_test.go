package searching

import "fmt"

func ExampleBinFind1() {
	l1 := []int{1, 3, 6, 38, 92, 190, 246, 457, 748, 904, 999}
	fmt.Println(BinFind1(l1, 3))
	fmt.Println(BinFind1(l1, 4))
	fmt.Println(BinFind1(l1, 67))
	fmt.Println(BinFind1(l1, 92))
	fmt.Println(BinFind1(l1, 457))
	fmt.Println(BinFind1(l1, 1))
	fmt.Println(BinFind1(l1, 44))

	//Output:
}

func ExampleBinFind2() {
	l1 := []int{1, 3, 6, 38, 92, 190, 246, 457, 748, 904, 999}
	fmt.Println(BinFind2(l1, 3))
	fmt.Println(BinFind2(l1, 4))
	fmt.Println(BinFind2(l1, 67))
	fmt.Println(BinFind2(l1, 92))
	fmt.Println(BinFind2(l1, 457))
	fmt.Println(BinFind2(l1, 1))
	fmt.Println(BinFind2(l1, 44))

	//Output:
}
