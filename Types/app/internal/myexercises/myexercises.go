package internal

import "fmt"

func ExerciseOne() {
	va := "Hola"
	fmt.Println("The length of", va, "is", len(va))

	for _, i := range va {
		fmt.Printf("%c\n", i)
	}
}

func ExerciseTwo() {
	age := 24
	employee := true
	years := 2
	salary := 200000

	if age > 22 && employee && years >= 2 {
		if salary > 100000 {
			fmt.Println("Interest-free loan")
		} else {
			fmt.Println("Interest-bearing loan")
		}
	} else {
		print("Does not meet the requirements")
	}
}

func ExerciseThree() {
	months := map[int]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May",
		6: "June", 7: "July", 8: "August", 9: "September", 10: "Octuber", 11: "November", 12: "December"}
	number := 12
	fmt.Println(months[number])
}

func ExerciseFour() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 38}

	for k, v := range employees {
		if v > 20 {
			fmt.Println(k)
		}

	}
	employees["Federico"] = 60
	delete(employees, "Pedro")
	fmt.Println(employees)

}
