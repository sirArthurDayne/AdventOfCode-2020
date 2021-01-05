package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
    /*AOC day 5. solution by: sirArthurDayne*/
    inputData, err:= scanLines("data.txt")
    if err != nil {
        panic(err)
    }

    //PART1
    //1.1 convert to binary string array
    var binaryData = []int{}
    for _, passport := range inputData {
        currentString := ""
        number := getTicketID(passport, 0, &currentString)
        binaryData = append(binaryData, number)
    }

    //1.2 convert from binary to int and sort array to get greater id
    sort.Ints(binaryData)
    // fmt.Println(binaryData)
    //PART 2. find the id of your seat
    counter := 40 + 1
    for i:=0; i<len(binaryData); i++ {
        if i == 0 {continue}
        if binaryData[i] - binaryData[i-1] != 1 {
            fmt.Println("Your seat is:", counter)
            break
        }
        counter++
    }
}
func scanLines(path string) ([]string, error) {

  file, err := os.Open(path)
  if err != nil {
     return nil, err
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanLines)

  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, nil
}

//PART1
func getTicketID(ticket string, index int, currentString *string) int {
    //base case
    if index >len(ticket)-1 {
        return binaryToInt(*currentString)
    }

    letter := ticket[index]
    if letter == 'F' || letter == 'L' {
        *currentString += "0"
        getTicketID(ticket, index+1, currentString)
    } else if letter == 'B' || letter == 'R' {
        *currentString += "1"
        getTicketID(ticket, index+1, currentString)
    }

    return binaryToInt(*currentString)
}

func binaryToInt(id string) int{
    i, err := strconv.ParseInt(id, 2, 0)
    if err != nil {
        panic(err)
    }
    return int(i)
}
