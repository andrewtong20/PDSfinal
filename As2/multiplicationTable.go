//This method takes in a starting and ending factor and creates the corresponding multiplication table.
//Formatting of table will be off for large factors
import sys

def multiplicationTable(start,end):
    //so that it does not matter that start has to be less than end
    if (start>end):
        temp=start
        start=end
        end=temp

    result=0

    //print table column heading
    for column in range(start,end+1):
        sys.stdout.write("%10d"% column)
    print()

    //print table columns

    for row in range(start, end+1):
        //print table row heading
        sys.stdout.write(str(row)) #takes account for top left corner space

        for column in range(start, end+1):
            result=row*column
            sys.stdout.write("%10d"% result) #formatting may be off for large factors
        print()
