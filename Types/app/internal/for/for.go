package internal

import "fmt"

func StandarFor() {
	sum := 0
	//declaration; condition; post-declaration
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

func ForLikeWhile() {
	sum := 1
	for sum < 10 {
		sum += sum
		println(sum)
	}

}

func ForWithRange() {
	//slice
	fruits := []string{"apple", "banana", "pear"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
	keysum := len(fruits)
	fmt.Println("There are", keysum, "elements in the slice")
}

//map[KeyType]ValueType

func ForWithContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}
}

func ForWithBreak() {
	sum := 0
	for {
		sum++
		if sum >= 100 {
			break
		}
	}
	fmt.Println(sum)
}
