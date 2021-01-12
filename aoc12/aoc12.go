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
    fmt.Println("2.total distance:",calculateWaypointDistance(directionsMap, 10.0,1.0))
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
func calculateWaypointDistance(directions []map[string]int, wayX, wayY float64) int {
    posX,posY := 0.0, 0.0//wayPointLocation
    for _, dir := range directions {
        for action, value := range dir {
            fmt.Printf("action:%v,val:%v |",action,value)
            switch action {
            case "R"://cw
                radians := float64(value) * math.Pi / 180.0
                newWayX := math.Floor(wayX * math.Cos(radians) - wayY * math.Sin(radians))
                newWayY := math.Floor(wayX * math.Sin(radians) + wayY * math.Cos(radians))
                wayX = newWayX
                wayY = newWayY
                break
            case "L"://ccw
                degrees := (360 + value) % 360
                switch degrees {
                case 90:
                    wayX, wayY = -wayY, wayX;break
                case 180:
                    wayX, wayY = -wayX, -wayY;break
                case 270:
                    wayX, wayY = wayY, -wayX;break
                default: fmt.Println("Unknow degrees", degrees)
                }
                break
            case "F"://forward
                posX += wayX * float64(value)
                posY += wayY * float64(value)
                break
                case "N": wayY += float64(value);break
                case "S": wayY -= float64(value);break
                case "E": wayX += float64(value);break
                case "W": wayX -= float64(value);break
                default: fmt.Println("ERROR!unknow action:",action)
            }
        }
        fmt.Printf("w_x:%v,w_y:%v | x:%v,y:%v\n",wayX,wayY,posX,posY)
    }
    return int(math.Abs(posX) + math.Abs(posY))
}
