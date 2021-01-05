package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
    destination string;
    sources []string;
}

func main() {
    /*AOC day7. Solution by: sirArthurDayne*/
    inputData,err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    //PART1.
    nodeList := parseData(inputData)
    fmt.Printf("total shiny bags: %v", totalShinyBags(nodeList))
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

//parseLines
func parseData(inputData []string) (output []Node) {
    for _, line := range inputData {
        splitLine := strings.Split(line, "s contain ")
        destination := strings.TrimSpace(splitLine[0])
        sources := strings.Split(splitLine[1], ", ")
        for i:=0; i < len(sources); i++ {
            sources[i] = sources[i][2:len(sources[i])]
            sources[i] = strings.ReplaceAll(sources[i],"bags","bag")
            if i == len(sources)-1 {
                sources[i] = strings.TrimRight(sources[i], ".")
            }
        }
        output = append(output, Node{destination, sources})
    }
    return
}

func totalShinyBags(nodeList []Node) int {
    graph := make(map[string][]string)
    for _, node := range nodeList {
        destination, sources := node.destination,node.sources
        if _,keyExits := graph[destination]; !keyExits {
            //add node to addyacency list
            graph[destination] = []string{}
        }
        for _, sou := range sources {
            //make sure all nodes are present in the adyacency list
            if _,keyExits := graph[sou]; !keyExits {
                graph[sou] = []string{}
            }
            //add destination to the souce-element list
            graph[sou] = append(graph[sou], destination)
        }
    }
    //traverse graph to get total(from bottom to Up)
    // fmt.Print(fmt.Sprintln(graph))
    //find answer and reduce by 1
    return TraverseGraph(graph, "shiny gold bag", make(map[string]bool)) -1
}

func TraverseGraph(graph map[string][]string, root string, visited map[string]bool) int {
    //check if the node was visited or count
    if _, nodeWasVisited := visited[root]; nodeWasVisited {
        return 0
    }
    //add node to visited list
    visited[root] = true
    count := 1
    for _, neighbor := range graph[root] {
        count += TraverseGraph(graph,neighbor, visited)
    }
    return count
}
