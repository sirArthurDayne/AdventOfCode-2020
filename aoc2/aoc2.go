package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
    /*ADVENT OF CODE DAY2(part1 & 2). Solution by: sirArthurDayne*/
    cleanData,err := scanLines("data.txt")
    if (err != nil) {
        panic(err)
    } else {
    //fmt.Printf("%v", totalOfValidPasswords(cleanData))
    fmt.Printf("%v", totalOfValidPasswordsByPos(cleanData))
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


func getData(rawLineData string) (int,int,string,string) {
    sliceByColon := strings.Split(rawLineData, ":")
    getPassword :=  strings.Trim(sliceByColon[1]," ")

    sliceBySpace := strings.Split(sliceByColon[0], " ")
    getChar := sliceBySpace[1]

    sliceByRule := strings.Split(sliceBySpace[0], "-")
    getLowerBound,getHighBound := getBoundaries(sliceByRule)

    return getLowerBound,getHighBound,getChar,getPassword
}

func getBoundaries(data []string) (int,int) {
    lower,err := strconv.Atoi(data[0])
    if err != nil {
        lower = -1
    }
    high,err2 := strconv.Atoi(data[1])
    if err2 != nil {
        high = -1
    }

    return lower,high
}

//FIRST PART OF CHALLENGE
func totalOfValidPasswords(inputData []string) int {
    validPass := 0
    for _,line := range inputData {
        lb,hb,char,pass := getData(line)
        // fmt.Printf("%v|%v|%s|%s\n",lb,hb,char,pass)
        count:=0
        for i:=0; i<len(pass); i++ {
            if (string(pass[i]) == char) {
                count++
            }
        }
        //check for boundaries
        if count >= lb && count <= hb {
            validPass++
        }
    }
    return validPass
}

//SECOND PART OF CHALLENGE
func totalOfValidPasswordsByPos(inputData []string) int {
    validPass := 0
    for _, line := range inputData {
        count := 0
        lb,hb,char,pass := getData(line)
        if string(pass[lb-1]) == char {
            count++
        }
        if string(pass[hb-1]) == char {
            count++
        }
        if count == 1 {
            validPass++
        }
    }
    return validPass
}
