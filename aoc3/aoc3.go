package main

import (
	"bufio"
	"fmt"
	"os"
)

type Movement struct {
    rows,cols int
}

func main() {
    /*AOC day 3. Solution by: sirArthurDayne*/
    inputData, err:= scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    //1. FIRST PART
    // fmt.Println("total trees: ", getTotalTrees(inputData, Movement{1,3}))
    //2. SECOND PART

    //2.1 Setup movements
    var movementList = []Movement{
        {1,1},
        {1,3},
        {1,5},
        {1,7},
        {2,1},
    }
    //2.2 getResult
    treesProduct := 1
    for i:=0; i<len(movementList); i++ {
        treesProduct *= getTotalTrees(inputData, movementList[i])
    }
    fmt.Println("total product of trees: ", treesProduct)
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

func getTotalTrees(data []string, movement Movement) int {
    maxColum := len(data[0])
    currentCol := 0
    treesCount := 0
    for row:=0; row < len(data); row+=movement.rows {
        //calculate current pos + get char
        if data[row][currentCol % maxColum] == '#' {
            treesCount++
        }
        currentCol += movement.cols
    }
    return treesCount
}
