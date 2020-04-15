/*
Andrew Tong
Mr. Tomczak
4/20/20
AS2

This program gives the user a menu of options (1-5) which repeats until the user quits (option 5).
Option 1 will print an investment table compounded daily given an initial investment, number of years, and range of integer interest rates.
Option 2 will print a multiplication able given a range of starting and ending values.
Option 3 will print all prime numbers to an inputted integer.
Option 4 will convert an input seconds into days, hours, minutes, and seconds.
 */


package main

import (
	"fmt"
	"bufio" //this is the scanner
  "os" //to accept input from user
	)

func menu(inName string) {
	name:=inName
}

func main() {

	fmt.Println("Hi, this is a program that will give you a menu of options.")
  fmt.Println("First, it will print a table of investments compounded daily given an initial investment, number of years, and range of integer interest rates.")
  fmt.Println("Second, it will ask for a range of starting and ending values and print a multiplication table.")
  fmt.Println("Third, it will prompt for an integer and print out all prime numbers up to that integer.")
  fmt.Println("Fourth, it will convert inputted seconds into days, hours, minutes, and seconds.")
  fmt.Println()
	fmt.Println("Alright, let's begin! What is your name?")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
  name:=scanner.Text()

	//to address user
  fmt.Println(name+", you are now entering the menu. The menu will repeat until you input option 5 to quit.")

  menu(name)

}
