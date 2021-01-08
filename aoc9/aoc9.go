package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
    /*AOC day9. solution: sirArthurDayne*/
    inputData, err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    numberList := parseStringToNumbers(inputData)
    //part1
    //fmt.Printf("number: %#v", searchNumber(numberList,25))
    //part 2
    fmt.Printf("number: %#v",findSumOfSet(numberList,searchNumber(numberList,25)))
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

func parseStringToNumbers(inputData []string) []int {
    var convert []int
    for i:=0; i <len(inputData); i++ {
        number, err := strconv.Atoi(inputData[i])
        if err != nil {
            panic(err)
        }
        convert = append(convert, number)
    }
    return convert
}
//part1
func searchNumber(data []int, preamble int) int {
    addlist := []int{}
    for i:=0; i < len(data); i++ {
        if preamble+i == len(data) {
            return -1
        }
        left:= i
        right:= (preamble)+i
        list:= data[left:right]
        testNumber:= data[preamble+i]
        if !findSum(list, testNumber) {
            return testNumber
        }
        addlist = append(addlist, testNumber)
    }
    return -1
}

func findSum(list []int, testNumber int) bool {
    for i:=0; i< len(list); i++ {
        for j:=i+1; j <len(list); j++ {
            if list[i] + list[j] == testNumber {
                return true
            }
        }
    }
    return false
}

//part2
func findSumOfSet(data []int,target int) int {
    for i:=0; i<len(data); i++ {
    testSet := []int{data[i]}
        for j:=i+1; j<len(data); j++ {
            //add new element to list
            testSet = append(testSet,data[j])
            acc := 0
            for _, e := range testSet{
                acc += e
            }
            if acc == target {
                sort.Ints(testSet)
                return testSet[0] + testSet[len(testSet)-1]
            }
        }
    }
    return -1
}

