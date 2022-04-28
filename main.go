package main

import "fmt"

func main() {

	if 7%2 == 0 {

		fmt.Println("7 is even")

	} else {

		fmt.Println("7 is odd")

	}

	//You can have an 'if' statement without an else

	if 8%4 == 0 {

		fmt.Println("8 is divisible by 4")

	}

	//A statement can precede conditionals; any variables declared in this statement are available in all branches.

	if num := 9; num < 0 {

		fmt.Println(num, "is negative")

	} else if num < 10 {

		fmt.Println(num, "has 1 digit")

	} else {

		fmt.Println(num, "has multiple digits")

	}

}

// output :
// 7 is odd
// 8 is divisible by 4
// 9 has 1 digit

// more details :
//In the first example, we simply measured a condition and vice versa.
//In the second example we saw that a condition can be used without else.
//The third example shows how to use multiple conditions using if else,
// and the variable defined in the condition is available inside all conditional spaces.
//In Golang, you do not need to use parentheses around the bet, but the content inside it should be in a space {}
