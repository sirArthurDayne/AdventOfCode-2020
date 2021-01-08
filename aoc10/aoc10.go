package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
    /*AOC day 10. Solution by: sirArthurDayne */
    inputData, err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    voltageList := parseStringsToInts(inputData)
    //sort the list
    sort.Ints(voltageList)
    //PART1.
    // diffVals := make(map[int]int)
    // diffVals[1] = 0
    // diffVals[3] = 0
    // fmt.Printf("%v",getDifferenceOfVoltages(voltageList, diffVals))
    //part2.
    fmt.Printf("%v\n", getPossibleCombinations(voltageList,0, make(map[int]int)))
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

func parseStringsToInts(inputData []string) (output []int){
    output = append(output, 0)
    for _, numString := range inputData {
        val,err := strconv.Atoi(numString)
        if err != nil {
            panic(err)
        }
        output = append(output, val)
    }
    return
}

//PART1
func getDifferenceOfVoltages(voltageList []int, diffVals map[int]int) (product int) {
    //add sum of the greater +3
    lastVal := voltageList[len(voltageList)-1] +3
    voltageList = append(voltageList, lastVal)
    //traverse the list and find the difference base on your list
    for i:=0; i < len(voltageList)-1; i++ {
            //calculate difference in voltage
            diff := voltageList[i+1] - voltageList[i]
            //traverse the list of valid diff Voltages
            for keyName := range diffVals {
                if diff == keyName {
                    diffVals[diff]++
                }
            }
    }
    fmt.Println(diffVals)

    //multiply the result of the differences and return
    product = diffVals[3] * diffVals[1]
    return
}

//PART2
func getPossibleCombinations(voltages []int, index int, resultSet map[int]int) int {
    //check for existing tree base on current index
    if value, keyExits := resultSet[index]; keyExits {
        return value
    }
    //found a valid route
    if index == len(voltages) -1 {
        return 1
    }
    //make calls of child nodes
    total := 0
    if index+1 <= len(voltages)-1 && voltages[index+1] - voltages[index] <= 3 {
        total += getPossibleCombinations(voltages,index+1, resultSet)
    }
    if index+2 <= len(voltages)-1 && voltages[index+2] - voltages[index]<= 3 {
        total += getPossibleCombinations(voltages,index+2, resultSet)
    }
    if index+3 <= len(voltages)-1 && voltages[index+3] - voltages[index] <= 3 {
        total += getPossibleCombinations(voltages,index+3, resultSet)
    }
    //save the value in map
    resultSet[index] = total
    return total
}
