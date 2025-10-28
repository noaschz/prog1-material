package listprops

import "fmt"

func ExampleContains() {
	l := []string{"Hallo", "Welt", "gute", "Nacht"}
	fmt.Println(Contains(l, "Welt"))
	fmt.Println(Contains(l, "Haus"))

	// Output:
<<<<<<< HEAD
	//true
	//false
=======
	// true
	// false
>>>>>>> 9a22e09e105f7d8388ac0794ab0dcb95e29cca77
}

func ExampleContainsOnly() {
	l1 := []string{"Hallo", "Welt", "gute", "Nacht"}
	fmt.Println(ContainsOnly(l1, "Welt"))
<<<<<<< HEAD
=======
	l2 := []string{"Welt", "Welt", "Welt", "Welt"}
	fmt.Println(ContainsOnly(l2, "Welt"))
	fmt.Println(ContainsOnly(l2, "Haus"))
	l3 := []string{}
	fmt.Println(ContainsOnly(l3, "Welt"))
>>>>>>> 9a22e09e105f7d8388ac0794ab0dcb95e29cca77

	l2 := []string{"Welt", "Welt", "Welt", "Welt"}
	fmt.Println(ContainsOnly(l2, "Welt"))

	l3 := []string{}
	fmt.Println(ContainsOnly(l3, "Welt"))
	// Output:
<<<<<<< HEAD
	//false
	//true
	//true
=======
	// false
	// true
	// false
	// true
>>>>>>> 9a22e09e105f7d8388ac0794ab0dcb95e29cca77
}

func ExampleContainsN() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt"}
	fmt.Println(ContainsN(l1, "Welt", 2))
	fmt.Println(ContainsN(l1, "Haus", 2))

	// Output:
<<<<<<< HEAD
	//true
	//false
}
func ExampleContainsNRow() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt"}
	fmt.Println(ContainsNRow(l1, "Welt", 2))
	fmt.Println(ContainsNRow(l1, "Welt", 3))

	// Output:
	//true
	//false
=======
	// true
	// false
}
func ExampleContainsNRow() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt", "Welt"}
	fmt.Println(ContainsNRow(l1, "Welt", 2))
	fmt.Println(ContainsNRow(l1, "Welt", 3))
	fmt.Println(ContainsNRow(l1, "Welt", 4))

	// Output:
	// true
	// true
	// false
>>>>>>> 9a22e09e105f7d8388ac0794ab0dcb95e29cca77
}
