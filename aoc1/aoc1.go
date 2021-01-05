package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
    /* AOC 1. Solution by sirArthurDayne */
    //1. clean input data
    inputData,err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    //2.cast to int and sort it
    numberData := []int{}
    for _,i := range inputData {
        num, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        numberData = append(numberData, num)
    }
    sort.Ints(numberData)
    //3. Calculate result (PART1)
    // x,y,sum,result := getResult(numberData)
    // fmt.Printf("%v * %v = %v || sum:%v\n",x,y,result,sum)
    //3. Calculate result (PART2)
    x,y,z,result := getResult2(numberData)
    fmt.Printf("%v * %v * %v = %v || sum:%v",x,y,z,result, x+y+z)

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

func getResult(data []int) (x,y,sum,result int) {

    found := false
    for i:=0; i< len(data); i++ {
        for j:=i; j <len(data); j++ {
            if data[i] + data[j] == 2020 {
                x = data[i]
                y = data[j]
                found= true
                break
            }
        }
        if found == true {
            break
        }
    }
    sum = x+y
    result = x * y
    return
}


func getResult2(data []int) (x,y,z,result int) {

    found := false
    for i:=0; i< len(data); i++ {
        if found == true {
            break
        }
        for j:=i; j <len(data); j++ {
            if found == true {
                break
            }
            for k:=j; k <len(data); k++ {
                if data[i] + data[j] + data[k] == 2020 {
                    x = data[i]
                    y = data[j]
                    z = data[k]
                    found= true
                    break
                }
            }
        }
    }
    result = x * y * z
    return
}
