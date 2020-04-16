//This method constructs an interest table given a range of interest rates, number of years, and an initial investment
//formatting will be off for large investments or large amounts of years

package main

import (
  "fmt"
  "math"
  "strconv"
)


func interest(inInitial float64, inLow int, inHigh int, inMAXyear int) {
    initial:=inInitial
    low:=inLow
    high:=inHigh
    MAXyear:=inMAXyear
    daysInYear:=365.0
    percentToDecimal:=100.0

    //Column heading
    //Spaces to center
    fmt.Println("                                   Years")


    for year:=1; year<=MAXyear; year++ {
      fmt.Printf("%16d",year)
    }

    fmt.Println()

    fmt.Println("Interest Rate: ")

    for rate:=low; rate<=high; rate++ {
         //row heading
         fmt.Print(strconv.Itoa(rate)+"%     ")

         for year:=1; year<=MAXyear; year++ {
             balance:=initial*math.Pow((1+float64(rate)/percentToDecimal/daysInYear), daysInYear*float64(year))
             //2 decimal places because money only goes to 2 decimal places
             fmt.Printf("$%10.2f     ",balance)

           }
         fmt.Println()

      }


}
