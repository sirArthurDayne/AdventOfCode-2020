package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    /*AOC day 11. solution by: sirArthurDayne*/
    inputData, err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(inputData)
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
