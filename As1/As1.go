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
  "log"//for scanner errors
  "math"//for math
  "io"//for file writer
  "strings"//for substrings and manipulation
  )

//Part 1: circle
func part1() string{
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
  User:= scanner.Text()

  fmt.Println(User+", what is the x coordinate of your circle?")
  scanner.Scan()
  xCoordStr:= scanner.Text()

  fmt.Println(User+", what is the y coordinate of your circle?")
  scanner.Scan()
  yCoordStr:= scanner.Text()

  fmt.Println(User+", what is the radius of your circle?")
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

  fmt.Println(User+", your circle has the following values.")
  //instructions do not require rounding on location; best to be exact
  fmt.Printf("X Coordinate: %f \nY Coordinate: %f \n", xCoord, yCoord)

  //Calculated values: these are rounded to 2 digits for readability
  fmt.Printf("Area (rounded to 2 decimal places): %.2f \n", c.getArea())
  fmt.Printf("Circumference (rounded to 2 decimal places): %.2f \n", c.getCircum())
  fmt.Println()

  fmt.Println(User+", if your circle was a sphere with the same radius, it would have these following values.")
  fmt.Printf("Volume (rounded to 2 decimal places): %.2f \n", c.getVolume())
  fmt.Printf("Surface area (rounded to 2 decimal places): %.2f \n", c.getSurfaceA())

  return User
}

//Part 2: meter conversion
//conversion factors from google
const (
  meters2miles=1609.34 //conversion factors from Google
  meters2feet=.3048
  meters2inches=.0254
)

func part2(inUser string) string{
  User:=inUser
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println()
  fmt.Println()
  fmt.Println(User+", you are now entering the second part of the program.")
  fmt.Println("The program will ask for a length in meters and convert it to its equivalent miles, \n feet, and inches and output these values to \"meters.txt\"")

  fmt.Println("What is your length in meters (do not input units)?")
  scanner.Scan()
  metersStr:= scanner.Text()

  //str to float conversion
  meters, err := strconv.ParseFloat(metersStr, 64)
  if err != nil {
    log.Fatal(err)
  }

  //dividing by miles
  metersRemain := math.Mod(meters,meters2miles);//remainder division to get meters left over that could not be converted to whole miles
  miles := (meters-metersRemain)/meters2miles; //integer division to get the miles in integer form from meters

  //convert to integer so that the miles won't have the unnecessary decimal of .0 when displaying it
  milesFinal:= int(miles);

  //dividing by feet
  metersRemain2:=math.Mod(metersRemain,meters2feet);//remainder division to get meters left over not divisible into feet
  feet:=(metersRemain-metersRemain2)/meters2feet; //integer division to get feet in integer form from meters
  feetFinal:= int(feet); //removes unnecessary .0 for visual pleasure

  //dividing by inches
  inches:=metersRemain2/meters2inches;//double division because the inches division will not be an exact integer.

  //conversion

  fmt.Printf(User+", your length of %f meters is equivalent to:\n", meters);
  fmt.Printf("%d miles\n", milesFinal);//units required for user experience
  fmt.Printf("%d feet\n", feetFinal);
  fmt.Printf("%.2f inches \n", inches);//only inches should be a double and for ease of use, rounded to 2 decimal places.


  //format data for output

  metersOut:=fmt.Sprintf("%f",meters)
  milesOut:=strconv.Itoa(milesFinal)
  feetOut:=strconv.Itoa(feetFinal)
  inchesOut:=fmt.Sprintf("%.2f",inches)
  finalOutput:=metersOut+"\n"+milesOut+"\n"+feetOut+"\n"+inchesOut+"\n"

  //output to file
  error := WriteToFile("meters.txt", finalOutput)
  if error != nil {
      log.Fatal(error)
  }
  return User
}

func WriteToFile(fileName string, fileText string) error {
  file, err := os.Create(fileName)
  if err != nil {
      return err
  }
  defer file.Close()

  //uses _ for null variables because golang does not allow unused variables
  _, err = io.WriteString(file, fileText)
  if err != nil {
      return err
  }
  return file.Sync()
}

//Part 3: String Manipulation
func part3(inUser string) {
   User:=inUser
   scanner := bufio.NewScanner(os.Stdin)
   fmt.Println()
   fmt.Println()
   fmt.Println(User+", you are now entering the third part of the program.")
   fmt.Println("This part will take a line of text and move the first word to the end,\n changing the capitalizations of the first two words to \n match the creation of a new sentence, and finds the length of the string")
   fmt.Println("Enter a line of text.  No punctuation please.")
   scanner.Scan()
   userInput:= scanner.Text()

   fmt.Println("I have rephrased the line to read: ")

   //identifies the position of the first character in first word/substring
   endPosition:=strings.Index(userInput," ")
   firstWord:=userInput[0:endPosition]
   startPosition:=endPosition+1
   restOfSentence:=userInput[startPosition:len(userInput)]

   //capitalizations
   char1:=string(firstWord[0])
   //instructions require the first word's first character to be lower case because it will be at the end of a sentence.
   char1Down:=strings.ToLower(char1)
   firstWordTrunc:=firstWord[1:endPosition]
   firstWordOutput:=char1Down+firstWordTrunc

   char2:=string(restOfSentence[0])
   //instructions require the second word's first character to be upper case because it begins a sentence.
   char2Up:=strings.ToUpper(char2)
   restOfSentenceTrunc:=restOfSentence[1:len(restOfSentence)]
   restOfSentenceOutput:=char2Up+restOfSentenceTrunc

   fmt.Println(restOfSentenceOutput+" "+firstWordOutput)

   //finding length
   fmt.Printf("The line of text has %d characters in it.",len(userInput))

   fmt.Println()
   fmt.Println(User+", thank you for using Andrew's As1!")

}

func main() {
   part3(part2(part1()))
}
