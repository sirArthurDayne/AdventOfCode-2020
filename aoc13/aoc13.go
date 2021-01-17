package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    /*AOC day 13. Solution by: sirArthurDayne */
    inputData, err := scanLines("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    //part1
    // timestamp, busList := parseData(inputData)
    // fmt.Printf("val:%v\n",getBusIdTimeValue(timestamp, busList))
    //part2
    busList := parseData2(inputData)
    fmt.Printf("val2:%v\n", getTimeStamp(busList))
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

func parseData(inputData []string) (timestamp int, busList []int) {
    timestamp, _ = strconv.Atoi(inputData[0])
    rawBusList := strings.Split(inputData[1], ",")

    for _, testBus := range rawBusList {
        if testBus == "x" {
            continue
        }
        strToInt, err := strconv.Atoi(testBus)
        if err != nil {
            log.Fatal(err)
        }
        busList = append(busList, strToInt)
    }
    return
}

//PART 1
func getBusIdTimeValue(timestamp int, busList []int) int {

    bustTimeMap := make(map[int]int)
    for _, busID := range busList {
        timeCounter := 0
        for timeCounter <= timestamp {
            timeCounter += busID
        }
        bustTimeMap[busID] = timeCounter - timestamp
    }
    //get smaller bustime
    min := 1000
    minBusID := -1
    for busID, timeVal := range bustTimeMap {
        if  timeVal < min {
            min = timeVal
            minBusID = busID
        }
    }
    return minBusID * min
}

//PART2

func parseData2(inputData []string) (busList []int) {
    rawBusList := strings.Split(inputData[1], ",")
    for _, testBus := range rawBusList {
        if testBus == "x" {
            busList = append(busList, 1)
            continue
        }
        strToInt, err := strconv.Atoi(testBus)
        if err != nil {
            log.Fatal(err)
        }
        busList = append(busList, strToInt)
    }
    return
}

func getTimeStamp(buslist []int) int {
    tcounter := 0
    stepsize := 0
    for iter, busID := range buslist {
        if iter == 0 {
            stepsize = busID
            continue
        }
        for (tcounter + iter) % busID != 0 {
            tcounter += stepsize
        }
        stepsize *= busID
    }
    return tcounter
}
