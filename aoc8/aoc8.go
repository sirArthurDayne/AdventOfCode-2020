package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Instruction struct {
    action string;
    value int;
}

func main() {
    /**AOC day 8. Solution by: sirArthurDayne*/
    inputData, err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    instructionSet := parseInstruction(inputData)
    //PART 1
    val,_ := getAccumulator(instructionSet)
    fmt.Printf("total:%v\n",val)
    //PART 2
    fmt.Printf("total:%v\n", getAccumulatorByChange(instructionSet))
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

func parseInstruction(input []string) (output []Instruction) {
    for _, instruction := range input {
        //get instruction
        segments := strings.Split(instruction, " ")
        action := segments[0]
        val := -10000
        //get value
        if strings.Contains(segments[1],"-") {
            valString := strings.ReplaceAll(segments[1],"-","")
            conv, err := strconv.Atoi(valString)
            if err != nil {
                panic(err)
            }
            val = -conv
        } else {
            valString := strings.ReplaceAll(segments[1],"+","")
            conv, err := strconv.Atoi(valString)
            if err != nil {
                panic(err)
            }
            val = conv
        }
        output = append(output, Instruction{action, val})
    }
    return
}


func getAccumulator(instructionSet []Instruction) (int,bool) {
    visited := make(map[int]bool)
    accumulator := 0
    jump :=1
    for index:=0; index < len(instructionSet); index +=jump {
        //check for visited instruction
        if _, wasVisited := visited[index]; wasVisited {
            return accumulator,false
        }
        //add instruction to visited list
        visited[index] = true
        //get action and perform it
        currentInstruction := instructionSet[index]
        switch action := currentInstruction.action; action {
            case "nop": jump = 1; //fmt.Println(action,"executed");
            continue
        case "jmp":
            jump = currentInstruction.value; //fmt.Println(action,"executed")
            break
        case "acc":
            accumulator += currentInstruction.value; //fmt.Println(action,"executed ",accumulator)
            jump = 1
        break
        default: fmt.Println("ERROR: unknown instruction", action)
        }
    }
    return accumulator,true
}

//PART2
func getAccumulatorByChange(instructionSet []Instruction) (acc int){
    found := false
    modSet := instructionSet
    for index:=0; index < len(instructionSet) && !found; index++ {
        modSet[index]= flipInstruction(instructionSet[index])
        acc,found = getAccumulator(modSet)
        //reset before next iter
        modSet[index] = flipInstruction(modSet[index])
    }
    return
}

func flipInstruction(current Instruction) Instruction {
    if current.action == "jmp" {
        current.action = "nop"
    } else if current.action == "nop" {
        current.action = "jmp"
    }
    return current
}
