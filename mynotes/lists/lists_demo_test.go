package lists

import "fmt"

//Sucht x in l und liefert die Position. Falls x nicht in l vorkommt, wird l geliefert
func Find(l []int, x int) int {
	//TODO -> done
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}

//sucht x in l und liefert alle vorkommnisse von x in l als liste von positionen
func FindAll(l []int, x int) []int {
	result := []int{}
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			result = append(result, i)
		}
	}
	return result
}

func ExampleLists() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := FindAll(l1, 5)
	pos2 := Find(l1, 200)

	fmt.Println(pos1)
	fmt.Println(pos2)

	//Output:
}
