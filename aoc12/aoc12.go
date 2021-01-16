package main

import (
	"bufio"
	"fmt"
	"math"
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
    //part1
    // fmt.Println("1.total distance:",calculateDistance(directionsMap))
    //part2
    fmt.Println("2.total distance:",calculateWaypointDistance(directionsMap, 10,1))
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
func calculateDistance(directions []map[string]int) int {
    wayX,wayY := 0, 0
    shipFace := []string{"E", "S", "W", "N"}
    facing := 0//east
    for _, dir := range directions {
        for action, value := range dir {
            currentFace := action
            switch action {
                case "R"://cw
                    facing += (value / 90)
                    facing %= 4//stay on range
                case "L"://ccw
                    facing -= (value / 90)
                    facing += 4//make positive
                    facing %= 4//stay on range of index
                case "F"://forward
                    currentFace = shipFace[facing]
            }
            switch currentFace {
                case "N": wayY+=value;break
                case "S": wayY-=value;break
                case "E": wayX+=value;break
                case "W": wayX-=value;break
            }
        }
    }
    fmt.Printf("x:%v,y:%v\n",wayX,wayY)
    wayX = int(math.Abs(float64(wayX)))
    wayY = int(math.Abs(float64(wayY)))
    return wayX+wayY
}
//PART2
func calculateWaypointDistance(directions []map[string]int, wayX, wayY int) int {
    posX, posY := 0, 0
    for _, dir := range directions {
        for action, value := range dir {
            switch action {
            case "N":
                wayY += value
                break
            case "S":
                wayY -= value
                break
            case "E":
                wayX += value
                break
            case "W":
                wayX -= value
                break
            case "L"://ccw
                for i := 0; i < (value/90); i++ {//cool trick for dealing with orotations
                //ccw means: swap Y with X and change X with inverse of Y
                    temp:= wayY
                    wayY = wayX
                    wayX = temp * -1
                }
                break
            case "R"://cw
                for i := 0; i < (value/90); i++ {
                //cw means: swap X with Y and change Y with inverse of X
                    temp := wayY
                    wayY = wayX * -1
                    wayX = temp
                }
                break
            case "F":
                posX += wayX * value
                posY += wayY * value
                break
            default: fmt.Println("unknow action ",action)
            }
        }
        fmt.Println("wx:", wayX, "wayY:",wayY, "posX:", posX, "posY:",posY)
    }
    if posX < 0 {
        posX *= -1
    }
    if posY < 0 {
        posY *= -1
    }
    return posX+posY
}
