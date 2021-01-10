package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    /*AOC day 12. solution by: sirArthurDayne */
    inputData,err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    directionsMap := parseDirections(inputData)
    for _,d := range directionsMap {
        fmt.Println(d)
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

func parseDirections(data []string)  []map[string]int {
    var output []map[string]int
    for index, dir := range data {
        directionSplit := strings.SplitAfter(dir, string(dir[0]))
        letter := directionSplit[0]
        val, castErr := strconv.Atoi(directionSplit[1])
        if castErr != nil {
            panic(castErr)
        }
        output = append(output, make(map[string]int))
        output[index][letter] = val
    }
    return output
}

//PART1
