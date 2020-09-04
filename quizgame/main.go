package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

//➜  quizgame go build . && ./quizgame -csv=problems.csv
func main() {
	csvFilename := flag.String("csv", "smallproblems.csv", "A csv file in the format of 'question,answer' ")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		osExit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
	}

	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()
	if err != nil {
		osExit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	//fmt.Println(problems)
	correctCount := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correctCount++
			//fmt.Println("Correct")
		}
	}
	fmt.Printf("You scored %d of %d questions. \n", correctCount, len(problems))

}

func parseLines(lines [][]string) []problem {

	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return result
}

func osExit(message string) {

	fmt.Println(message)
	os.Exit(1)

}

/* The objective of this first iteration is to
	- presents the question
	- accepts  the user input
	- confirms the answer

1.	Get the name of the file that would be storing the questions and solutions.
	Here we use the flag.String() library since we know it will
	- be accepting a string.
	- takes the name (csv)
	- default value (problem.csv)
	-  and helptext sayinh "  . . ."
2. Parse the the file

3.	Read the file. new object
	NOTE that in reading the is file we are using an io.reader() in os.Open() which will be handle later
	handle error( os.Exit(1)). In handling the error let the user know the actual file u tried to open.
	Note to use pointer with the filename to get that actual file and not a copy.
	Altaernativly use a seperate function in the error handling  that will take a message ,print it out and the OS exit
	NOTE: the use of fmt.Sprintf to handle the string literals i the message passed to the osExit function. Very Cool

4.	Create a csv reader object with the csv.NewReader(file). The New Reader file takes in an io.Reader which happens to be
	 what we used earlier in the reading of the file.
	 Note that io.Reader is very often used in Go programming.

5.	Read/Parse the lines of the csv in entirety . This is possible since we know the file is not going to the that large.
	Handle error. exit("Message")

6.	Print the lines . This should give us a 2 dimensional array of bytes[]
	NOTE. Alaternatively a  SMART THING TO DO IS TO CREATE A STRUCT TYPE TO MAP TO OUR 2D SLICE FROM THE LINES OF THE CSV.

7. Create the struct (problem) to represent the question/answer nested array pairs or lines of the CSV

8. Create the function to parse the lines of the csv
	now were have a slice of problem strucst
	Note : An edge case would be such that the questions or answeers provided may include invalid spaces
		it would be mindfull to use the strainga.TrimSpace()

9.	Declare the problesms variable to hold the problem structs data
10.	create for range loop to go through the problems and print out in formated manner
	Also create a counter to store correct answers]
	While in for loops :
	- Create local variable to store the answer
	- Compare user input to answer variable within the loop
	- if correct increase correct answer counter value
11.

*/
