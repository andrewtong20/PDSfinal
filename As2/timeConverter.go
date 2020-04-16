//This method takes user's inputted time in seconds and converts it to days, hours, minutes, and the remaining seconds.

import sys
def time(inSeconds):
    //global variables, initialized
    seconds=inSeconds

    //ints for rounding final output, floats for accuracy in calculations
    days=0.0
    hours=0.0
    minutes=0.0
    secondsRemain=0.0
    secondsRemain2=0.0
    secondsRemain3=0.0
    daysFinal=0
    hoursFinal=0
    minutesFinal=0
    secondsFinal=0.0

    //constants, conversion factors from Google
    seconds2days=86400
    seconds2hours=3600;
    seconds2minutes=60;

    //finding days
    def getDays():
        secondsRemain = seconds%seconds2days//remainder division to get seconds left over that could not be converted to whole days
        days = (seconds-secondsRemain)/seconds2days //integer division to get the days in integer form from seconds
        daysFinal= int(days) //convert to integer so that the miles won't have the unnecessary decimal of .0 when displaying it
        print(str(daysFinal)+" days")
        return secondsRemain //returns the secondsRemain to pass into next method (getHours)

    //finding hours
    def getHours(secondsRemain):
        secondsRemain2=secondsRemain%seconds2hours//remainder division to get seconds left over not divisible into hours
        hours= (secondsRemain-secondsRemain2)/seconds2hours//integer division to get hours in integer form from seconds
        hoursFinal= int(hours)//removes unnecessary .0 for visual pleasure
        print(str(hoursFinal)+" hours")
        return secondsRemain2 //returns the secondsRemain2 to pass into next method
    //finding minutes
    def getMinutes(secondsRemain2):
        secondsRemain3=secondsRemain2%seconds2minutes//remainder divisiont to get seconds left over not divisible to minutes
        minutes=(secondsRemain2-secondsRemain3)/seconds2minutes//integer division to get minutes in integer form from seconds
        minutesFinal=int(minutes)//removes unnecessary .0
        print(str(minutesFinal)+" minutes")
        return secondsRemain3 //returns to pass into next method
    //finding seconds
    def getSeconds(secondsRemain3):
        secondsFinal= secondsRemain3
        //round to 2 decimal places for balance between both accuracy and clutter (in case user enters in a float number of seconds)
        sys.stdout.write("%.2f seconds\n"%(secondsFinal))


    print("Your time of "+str(seconds) +" seconds is equivalent to:")
    //prints all values of time conversion
    getSeconds(getMinutes(getHours(getDays()))) //these methods are nested due to their return values
