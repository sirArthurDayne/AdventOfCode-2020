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
	/*AOC day 14. solution by sirArthurDayne */
	// inputData, err := scanLines("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
    //PART1
    //fmt.Println("total",getSumOfAdressesInMemory(inputData))
    //PART2.
    // fmt.Println("total",getSumOfAdressesInMemory2(inputData))
    list := getAllPosiblesCombinations("1XX0")
    fmt.Println(list)
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

//PART 1
func getSumOfAdressesInMemory(data []string) int {
    mem := make(map[int]int)
    currentMask := ""
    for _, line := range data {
        if strings.Contains(line, "mask") {
            currentMask = strings.Trim(line, "mask = ")
        } else if strings.Contains(line, "mem") {
            memSplit := strings.Split(line," = ")
            //get key and value int
            key := StringToInt(strings.Trim(memSplit[0], "mem[]"))
            valStringToInt := StringToInt(strings.TrimSpace(memSplit[1]))
            //cast value from string to Binary 36bits
            valIntToBinary := padStart(IntToBinary(int64(valStringToInt)),36,"0")
            //Apply the mask
            newBinaryValue := applyMask(valIntToBinary, currentMask)
            //cast back from Binary to Int
            value := BinaryToInt(newBinaryValue)
            //save in map
            mem[key] = int(value)
        }
    }

    //add values and return
    total := 0
    for _,val := range mem {
        total += val
    }
    return total
}

//PART 2
func getSumOfAdressesInMemory2(data []string) int {
    mem := make(map[int]int)
    currentMask := ""
    for _, line := range data {
        if strings.Contains(line, "mask") {
            currentMask = strings.Trim(line, "mask = ")
        } else if strings.Contains(line, "mem") {
            memSplit := strings.Split(line," = ")
            //get key and value int
            key := StringToInt(strings.Trim(memSplit[0], "mem[]"))
            valStringToInt := StringToInt(strings.TrimSpace(memSplit[1]))
            //cast value from string to Binary 36bits
            valIntToBinary := padStart(IntToBinary(int64(valStringToInt)),36,"0")

            //apply mask rules
            valMask := applyMask2(valIntToBinary, currentMask)
            //GetAllAddresses

            //add result to mem[key]
            mem[key] = getAllPosiblesCombinations(valMask)
        }
    }

    //sum all in memory
    total:=0
    return total
}

func StringToInt(n string) int {
    val, err := strconv.Atoi(n)
    if err != nil {
        log.Fatal(err)
    }
    return val
}

func IntToBinary(n int64) string {
    return strconv.FormatInt(n,2)
}

func BinaryToInt(n string) int64 {
    val,err :=  strconv.ParseInt(n, 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    return val
}

func padStart(n string, length int, char string) string {
    container :=  ""
    for i:=0; i < length; i++ {
        container+=char
    }
    outputVal := container + n
    return outputVal[len(outputVal)-length:]
}

func applyMask(val, currentMask string) string {
    output:= ""
    for i:=0; i < len(currentMask); i++ {
        if currentMask[i] == 'X' {
            output += string(val[i])
        } else if currentMask[i] == '1' {
            output += "1"
        } else if currentMask[i] == '0' {
            output += "0"
        }
    }
    return output
}

func applyMask2(val, currentMask string) string {
    output:= ""
    for i:=0; i < len(currentMask); i++ {
        if currentMask[i] == 'X' {
            output += "X"
        } else if currentMask[i] == '1' {
            output += "1"
        } else if currentMask[i] == '0' {
            output += string(val[i])
        }
    }
    return output
}

func getAllPosiblesCombinations(value string) int {
return -1
}

