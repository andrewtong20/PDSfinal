/*
Andrew Tong
Mr. Tomczak
4/13/20
AS1

This program finds the location, area, circumference, volume and surface area of the circle/sphere given a radius and x/y coordinates.
It also asks the user for meters and converts it into the corresponding number of miles, feet, and inches. Then it prints the meters and the miles, feet, and inches to "meters.txt"
Finally, it reads a line of text from the user, outputs the line with the first word moved to the end of the line, and prints the length of the string
*/

package main

import (
  "fmt"
  "bufio" //this is the scanner
  "os" //to accept input from user
  "strconv" //string conversion for inputs
  "log"
  //"As1/circle"//import struct
  )

func main() {
  //Part 1: circle
  
  //Inputs
  scanner := bufio.NewScanner(os.Stdin)
  //User intro
  fmt.Println("Welcome to Andrew's Assignment 1.")
  fmt.Println("This program will: ")
  fmt.Println("1) take a radius  and x/y coordinates to find the location, area, circumference,\n volume and surface area of the circle/sphere")
  fmt.Println("2) takes a length in meters, converts it to miles, feet, and inches, \n and outputs it to \"meters.txt\"")
  fmt.Println("3) takes a line of text, outputs the line with the first word moved to the end of the line,\n and prints the length of the string. ")
  fmt.Println()
  fmt.Println("You are entering the first part of the program where we will construct a circle.")
  fmt.Println("The program will ask for x/y coordinates and a radius and output the location, area,\n circumference, volume, and surface area.")

  //Instructions ask to prompt user to input their name to address them later
  fmt.Print("What is your name?")
  scanner.Scan()
  user:= scanner.Text()

  fmt.Println(user+", what is the x coordinate of your circle?")
  scanner.Scan()
  xCoordStr:= scanner.Text()

  fmt.Println(user+", what is the y coordinate of your circle?")
  scanner.Scan()
  yCoordStr:= scanner.Text()

  fmt.Println(user+", what is the radius of your circle?")
  scanner.Scan()
  radiusStr:= scanner.Text()

  //convert strings to floats
  xCoord, err := strconv.ParseFloat(xCoordStr, 64)
  if err != nil {
    log.Fatal(err)
  }
  yCoord, err := strconv.ParseFloat(yCoordStr, 64)
  if err != nil {
    log.Fatal(err)
  }
  radius, err := strconv.ParseFloat(radiusStr, 64)
  if err != nil {
    log.Fatal(err)
  }
  //go uses structs instead of objects
  c:=Circle {
    x: xCoord,
    y: yCoord,
    radius: radius,
  }

  fmt.Println()

  //Outputs

  fmt.Println(user+", your circle has the following values.")
  //instructions do not require rounding on location; best to be exact
  fmt.Printf("X Coordinate: %f \nY Coordinate: %f \n", xCoord, yCoord)

  //Calculated values: these are rounded to 2 digits for readability
  fmt.Printf("Area (rounded to 2 decimal places): %.2f \n", c.getArea())
  fmt.Printf("Circumference (rounded to 2 decimal places): %.2f \n", c.getCircum())
  fmt.Println()

  fmt.Println(user+", if your circle was a sphere with the same radius, it would have these following values.")
  fmt.Printf("Volume (rounded to 2 decimal places): %.2f \n", c.getVolume())
  fmt.Printf("Surface area (rounded to 2 decimal places): %.2f \n", c.getSurfaceA())

}
