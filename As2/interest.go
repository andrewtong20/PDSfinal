//This method constructs an interest table given a range of interest rates, number of years, and an initial investment
//formatting will be off for large investments or large amounts of years

import (
  "fmt"
)


func interest(initial, low, high, MAXyear):
    balance=0
    daysInYear=365.0
    percentToDecimal=100.0

    //Column heading
    //Spaces to center
    print("                                   Years")


    for year in range(1,MAXyear+1):
        sys.stdout.write("%16d"% (year))
    print()

    print("Interest Rate: ")

    for rate in range(low, high+1):
         //row heading
        sys.stdout.write(str(rate)+"%     ")

        for year in range(1, MAXyear+1):
            balance=initial*pow((1+rate/percentToDecimal/daysInYear), daysInYear*year)
            //2 decimal places because money only goes to 2 decimal places
            sys.stdout.write("$%10.2f     "%(balance))
        print()
