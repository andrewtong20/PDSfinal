/*
Andrew Tong
Mr. Tomczak
4/28/20
Coding Bats

This file runs all of the coding bats given user inputs.
*/
package main

import (
  "fmt"
  "bufio" //this is the scanner
  "os" //to accept input from user
  "strconv"
)

//Float bounds check
func isFloat() float64 {
  scanner := bufio.NewScanner(os.Stdin)

  for true {

    fmt.Print("Input: ")
    scanner.Scan()

    input := scanner.Text()
    if inputConv, err := strconv.ParseFloat(input, 64); err == nil {
      return inputConv
    } else {
      fmt.Println("This is not a number. Please enter a valid number.")
    }


  }
  //function has to return something; 0 has no meaning
  return 0
}

//Integer bounds check
func isInteger() int {
  scanner := bufio.NewScanner(os.Stdin)

  for true {

    fmt.Print("Input: ")
    scanner.Scan()

    input := scanner.Text()
    if inputConv, err := strconv.Atoi(input); err == nil {
      return inputConv
    } else {
      fmt.Println("This is not an integer. Please enter a valid number.")
    }


  }
  //function has to return something; 0 has no meaning
  return 0
}

func menu(inName string) {
	name:=inName

	var input int
	for ok := true; ok; ok = (input != 4) {
		fmt.Println("Option 1: Array1 - front11")
		fmt.Println("Option 2: Logic2 - luckySum")
		fmt.Println("Option 3: Recursion1 - bunnyEars2")
		fmt.Println("Option 4: Quit")
		fmt.Println("The table will repeat itself until you quit.")

    n, err := fmt.Scanln(&input)

    for n < 1 || err != nil {

      fmt.Println("Input has to be an integer within range.")
      break
    }

    switch input {
    case 1:
      fmt.Println("You have decided to test front11.")
      fmt.Println("Instructions are: Given 2 int arrays, a and b, of any length, return a new array with \nthe first element of each array. If either array is length 0, ignore that array.")

		case 2:
      fmt.Println("You have decided to test luckySum")
      fmt.Println("Instructions are: Given 3 int values, a b c, return their sum. However, if one of the \nvalues is 13 then it does not count towards the sum and values to its right do not count. \nSo for example, if b is 13, then both b and c do not count.")

		case 3:
      fmt.Println("You have decided to test bunnyEars2")
      fmt.Println("We have bunnies standing in a line, numbered 1, 2, ... The odd bunnies (1, 3, ..) have \nthe normal 2 ears. The even bunnies (2, 4, ..) we'll say have 3 ears, because they each have \na raised foot. Recursively return the number of ears in the bunny line 1, 2, ... n \n(without loops or multiplication)."



    case 4:
			fmt.Println(name+", thank you for using this program!")
			break
    default:
      fmt.Println("Wrong input. Try again.")//This only works if you enter integers.

    }
	}

}

func main() {
  fmt.Println("Hi, this is a program that will give you a menu of options.")
  fmt.Println("First, it will allow you to test an Array1 codingbat.")
  fmt.Println("Second, it will allow you to test a Logic1 codingbat.")
  fmt.Println("Third, it will allow you to test a Recursion1 codingbat.")

  fmt.Println()
  fmt.Println("Alright, let's begin! What is your name?")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  name:=scanner.Text()

  //to address user
  fmt.Println(name+", you are now entering the menu. The menu will repeat until you input option 4 to quit.")

  menu(name)

}
