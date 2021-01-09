package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Movement struct {
    r,c int;
}

func main() {
    /*AOC day 11. solution by: sirArthurDayne*/
    inputData, err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    total := 0
    for stable:=false; !stable; {
        stable = generateMap(inputData)
        //count for total #
        printMap(inputData)
        total = getTotalSeats(inputData)
    }
    fmt.Println(total)
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

//PART1 x*w+y
func printMap(mapa []string) {
    for _,seats := range mapa {
        fmt.Println(seats)
    }
    fmt.Printf("\n")
}
func generateMap(data []string) (bool) {
    /*FIXME: not generating the correct map {37 is the right value}*/
    isStable := true
    originalMap := make([]string, len(data))
    copy(originalMap, data)
    for row,seats := range originalMap {
        for col, s := range seats {
            //count neighbors
            location := Movement{r:row,c:col}
            maxLocation := Movement{r:len(originalMap)-1, c:len(seats)-1}
            neighbors := getNeighborsOnMap(location, maxLocation, originalMap)
            //change state base on rules and verify stability
            if s == 'L' && neighbors == 0 {
                smod := []rune(seats)
                smod[col] = '#'
                data[row] = string(smod)
                isStable = false
            } else if s == '#' && neighbors >= 4 {
                smod := []rune(seats)
                smod[col] = 'L'
                data[row] = string(smod)
                isStable = false
            }
        }
    }
    return isStable
}

func getNeighborsOnMap(location, maxLocation Movement, currentMap []string) (amount int) {
    for i:=-1; i<=1; i++{
        for j:=-1; j<=1; j++ {
            if (i==0 && j==0) || !isValid(Movement{location.r+i,location.c+j}, maxLocation) {
                continue
            }
            if currentMap[location.r+i][location.c+j] == '#' {
                amount++
            }
        }
    }
    return
}

func isValid(location, maxLocation Movement) bool {
    if location.r < 0 || location.c < 0 || location.r > maxLocation.r || location.c > maxLocation.c {
        return false
    }
    return true
}


func getTotalSeats(dataMap []string) int {
    total := 0
    for _,seats := range dataMap {
        total += strings.Count(seats, "#")
    }
    return total
}
