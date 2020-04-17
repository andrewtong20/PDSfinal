//This method takes user's inputted time in seconds and converts it to days, hours, minutes, and the remaining seconds.
package main

import (
  "fmt"
  "strconv"
  "math"
)

const (
  //constants, conversion factors from Google
  seconds2days=86400
  seconds2hours=3600
  seconds2minutes=60
)

func timeConvert(inSeconds float64) {
    //global variables, initialized
    seconds:=inSeconds

    //ints for rounding final output, floats for accuracy in calculations
    //days:=0.0
    //secondsRemain:=0.0
    //daysFinal:=0

    // hours:=0.0
    // secondsRemain2:=0.0
    // hoursFinal:=0
    //
    // minutes:=0.0
    // secondsRemain3:=0.0
    // minutesFinal:=0
    //
    // secondsFinal:=0.0



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
