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
	"strconv" //convert strings to ints/floats
	"math"//calculations
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

//Interest table method
//This method constructs an interest table given a range of interest rates, number of years, and an initial investment
//formatting will be off for large investments or large amounts of years
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

//multiplication table method
//This method takes in a starting and ending factor and creates the corresponding multiplication table.
//Formatting of table will be off for large factors
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

//prime generator methods
//This method takes the user's inputted integer, creates a list of numbers up to that integer, and tests to see if these numbers are prime.
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

//time converter methods
//This method takes user's inputted time in seconds and converts it to days, hours, minutes, and the remaining seconds.
const (
  //constants, conversion factors from Google
  seconds2days=86400
  seconds2hours=3600
  seconds2minutes=60
)

func timeConvert(inSeconds float64) {
    //global variables, initialized
    seconds:=inSeconds

    //Output
    fmt.Println("Your time of "+strconv.FormatFloat(seconds,'f', -1, 64) +" seconds is equivalent to:")
    //prints all values of time conversion
    getSeconds(getMinutes(getHours(getDays(seconds)))) //these methods are nested due to their return values

}
//finding days
func getDays(inSeconds float64) float64 {
    seconds:=inSeconds
    secondsRemain := math.Mod(seconds,seconds2days)//remainder division to get seconds left over that could not be converted to whole days
    days := (seconds-secondsRemain)/seconds2days //integer division to get the days in integer form from seconds
    daysFinal:= int(days) //convert to integer so that the miles won't have the unnecessary decimal of .0 when displaying it
    fmt.Println(strconv.Itoa(daysFinal)+" days")
    return secondsRemain //returns the secondsRemain to pass into next method (getHours)
  }
//finding hours
func getHours(inSecondsRemain float64) float64{
    secondsRemain:=inSecondsRemain
    secondsRemain2:=math.Mod(secondsRemain,seconds2hours)//remainder division to get seconds left over not divisible into hours
    hours:= (secondsRemain-secondsRemain2)/seconds2hours//integer division to get hours in integer form from seconds
    hoursFinal:= int(hours)//removes unnecessary .0 for visual pleasure
    fmt.Println(strconv.Itoa(hoursFinal)+" hours")
    return secondsRemain2 //returns the secondsRemain2 to pass into next method
  }
//finding minutes
func getMinutes(inSecondsRemain2 float64) float64{
    secondsRemain2:=inSecondsRemain2
    secondsRemain3:=math.Mod(secondsRemain2,seconds2minutes)//remainder divisiont to get seconds left over not divisible to minutes
    minutes:=(secondsRemain2-secondsRemain3)/seconds2minutes//integer division to get minutes in integer form from seconds
    minutesFinal:=int(minutes)//removes unnecessary .0
    fmt.Println(strconv.Itoa(minutesFinal)+" minutes")
    return secondsRemain3 //returns to pass into next method
  }
//finding seconds
func getSeconds(inSecondsRemain3 float64) {

    secondsFinal:= inSecondsRemain3
    //round to 2 decimal places for balance between both accuracy and clutter (in case user enters in a float number of seconds)
    fmt.Printf("%.2f seconds\n",(secondsFinal))
  }

//Menu method
func menu(inName string) {
	name:=inName

	var input int
	for ok := true; ok; ok = (input != 5) {
		fmt.Println("Option 1: Investment table.")
		fmt.Println("Option 2: Multiplication table.")
		fmt.Println("Option 3: Prime number generator.")
		fmt.Println("Option 4: Time conversion.")
		fmt.Println("Option 5: Quit.")
		fmt.Println("The table will repeat itself until you quit.")

    n, err := fmt.Scanln(&input)

    for n < 1 || err != nil {

      fmt.Println("Input has to be an integer within range.")
      break
    }

    switch input {
    case 1:
			fmt.Println(name+ ", you chose option 1, where you input an initial investment, lowest and highest \n integer interest rates, and the number of years invested to make a daily compound interest table.\nFormatting of the table will be off for large investments.")
      fmt.Println("Please enter your initial investment.")
      initial:=isFloat()

      //indefinite bounds check
      for initial<=0 {
          fmt.Println("Your investment can only be a positive amount! Please enter again.")
          initial=isFloat()
				}

      //no range bounds check for interest rates because user can enter any (even negative)
      fmt.Println("Enter lowest integer interest rate (ex: 3% would be 3 not .03)")
      low:=isInteger()
      fmt.Println("Enter highest integer interest rate (ex: 9% would be 9 not .09)")
      high:=isInteger()

      //upper bound cannot be lower than lower bound
      for high<=low {
          fmt.Println("Your input was the same or lower than the lowest integer interest rate. \nPlease enter a higher integer interest rate.")
          high=isInteger()
				}

      fmt.Println("How many whole years will you invest?")
      MAXyear:=isInteger()

      //realistic bounds restriction for time
      for MAXyear<0 {
          fmt.Println("Time cannot be negative. Please enter the whole amount of years you will invest.")
          MAXyear=isInteger()
				}

      //formatting of table will be off for large initial investments or large amount of years
      interest(initial, low, high, MAXyear);
		case 2:
			fmt.Println(name+", you chose Option 2, where you input a starting \n and ending integer values to construct a multiplication table.\n Formatting of the table will be off for large factors.")
      //instructions say integers only
      fmt.Println("If you enter a higher starting value than ending value, they will be flipped automatically.")
      //because of this feature, starting does not have to be less than end and so I don't have that bounds check
      fmt.Println("Enter your integer starting value.")

      start:=isInteger()
			//factors can't be negative bounds check
			for start<0 {
				fmt.Println("Factor cannot be negative. Please enter again.")
				start=isInteger()
			}
      fmt.Println("Enter your integer ending value.")
      end:=isInteger()

			//factors can't be negative bounds check
			for end<0 {
				fmt.Println("Factor cannot be negative. Please enter again.")
				end=isInteger()
			}

      //formatting of table will be off for large factors
      multiplicationTable(start,end);
		case 3:
			fmt.Println(name+", you chose Option 3, where you enter an integer to create a list of all prime numbers up to that integer.")

      fmt.Println("Please enter your integer.")
      primeTarget:=isInteger()

      //prime bounds check
      for primeTarget<2 {
          print("2 is the first prime number, so you have to enter an integer greater than 2. Please enter again")
          primeTarget=isInteger()
			}
      //if the primeTarget is a prime number, will include that value
      fmt.Print("The prime numbers up to "+strconv.Itoa(primeTarget)+" are ")

      primeList(primeTarget)

		case 4:
			fmt.Println(name+", you chose Option 4, where a time in seconds to convert to days, hours, minutes, and seconds.")
      fmt.Println("Please enter a time in seconds (decimals allowed)")
      seconds:=isFloat()

      //time can't be negative bounds check
      for (seconds<0) {
          fmt.Println("Time cannot be negative. Please enter your time as a positive value in seconds.")
          seconds=isFloat()
				}

      timeConvert(seconds)
    case 5:
			fmt.Println(name+", thank you for using this program!")
			break
    default:
      fmt.Println("Wrong input. Try again.")//This only works if you enter integers.

    }
	}

}

//main method (runs menu)
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
