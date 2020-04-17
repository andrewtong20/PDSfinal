//This method takes in a starting and ending factor and creates the corresponding multiplication table.
//Formatting of table will be off for large factors
package main

import (
  "fmt"
  "strconv"
)

func multiplicationTable(inStart int,inEnd int) {
    start:=inStart
    end:=inEnd

    //so that it does not matter that start has to be less than end
    if start>end {
      temp:=start
      start=end
      end=temp
    }

    result:=0

    //print table column heading
    for column:=start; column<=end; column++ {
        fmt.Printf("%10d", column)
    }

    fmt.Println()

    //print table columns

    for row:=start;row<=end; row++ {
        //print table row heading
        fmt.Print(strconv.Itoa(row)) //takes account for top left corner space

        for column:=start; column<=end; column++ {
            result=row*column
            fmt.Printf("%10d", result) //formatting may be off for large factors
        }
        fmt.Println()
    }
}
