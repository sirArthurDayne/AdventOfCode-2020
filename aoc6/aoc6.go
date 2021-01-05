package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    /*AOC day6. solution by: sirArthurDayne */

    inputData, err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    //PART 1 && PART 2
    totalAnswers,totalSameAnswers := getTotalEqualsAnswers(inputData)
    fmt.Printf("total: %v | sameAnwers: %v", totalAnswers, totalSameAnswers)
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

//RESOLUTION
func getTotalEqualsAnswers(data []string) (int,int) {
    groupAmount := 0
    groupSize := []int{0}
    groupList := []map[rune]int{make(map[rune]int),}
    for _, person := range data {
        //group end
        if person == "" {
            groupList = append(groupList, make(map[rune]int))
            groupSize = append(groupSize, 0)
            groupAmount++
            continue
        }
        //traverse person answers
        for _, answer := range person {
            //add to that groupAmount their answer amount
            groupList[groupAmount][answer]++
        }
        //increase userCount before traverse next person
        groupSize[groupAmount]++
    }
    sum,total := 0,0
    //traverse each group
    for index, gr := range groupList {
        //add all of their answer
        sum += len(gr)
        for _, count := range gr{
            //found a answer that everyone on the group has
            if count == groupSize[index] {
                total++
            }
        }
    }

    return sum,total
}
