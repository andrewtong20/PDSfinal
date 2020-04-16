//This method takes the user's inputted integer, creates a list of numbers up to that integer, and tests to see if these numbers are prime.
import sys

def isPrime(runningInteger):
    for divisor in range(2, runningInteger): #should not divide by itself so does not include runningInteger
        if (runningInteger%divisor==0):
            return False

    return True //when no other factors from 2 to one minus the entered integer are not divisors of the integer

def primeList(primeTarget):

    for runningInteger in range(2, primeTarget+1):
        //runningInteger is potential prime number ot be printed
        //start at 2 because 1 is not a prime number

        if (isPrime(runningInteger)):
            print(str(runningInteger), end=" ")
    print()//adds next line for the table
