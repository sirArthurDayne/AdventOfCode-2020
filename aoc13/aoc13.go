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
    inputData, err := scanLines("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    timestamp, busList := parseData(inputData)
    fmt.Printf("val:%v",getBusIdTimeValue(timestamp, busList))
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
        fmt.Printf("busID:%v time:%v\n",busID,timeCounter)
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
