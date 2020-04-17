//This method takes the user's inputted integer, creates a list of numbers up to that integer, and tests to see if these numbers are prime.
package main

import (
  "fmt"
  "strconv"
)

func isPrime(inRunningInteger int) bool {
    runningInteger:=inRunningInteger
    for divisor:=2; divisor<runningInteger; divisor++ {
    //should not divide by itself so does not include runningInteger
        if (runningInteger%divisor==0) {
            return false
          }
    }
    return true //when no other factors from 2 to one minus the entered integer are not divisors of the integer
}
func primeList(inPrimeTarget int) {
    primeTarget:=inPrimeTarget
    for runningInteger:=2; runningInteger<=primeTarget; runningInteger++ {
        //runningInteger is potential prime number ot be printed
        //start at 2 because 1 is not a prime number

        if (isPrime(runningInteger)) {
          fmt.Print(strconv.Itoa(runningInteger)+" ")
        }

    }
    fmt.Println()//adds next line for the table
}
