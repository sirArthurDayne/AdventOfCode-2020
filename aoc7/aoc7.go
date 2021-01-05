package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
    destination string;
    sources map[string]int;
}

func main() {
    /*AOC day7. Solution by: sirArthurDayne*/
    inputData,err := scanLines("data.txt")
    if err != nil {
        panic(err)
    }
    //parse the data to a nodelist
    nodeList := parseData(inputData)
    //PART1.
    // fmt.Printf("total shiny bags: %v", totalShinyBags(nodeList))
    //PART2
    fmt.Printf("total of bags inside shiny ones: %v", totalBagsInsideShinyOnes(nodeList))

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
        var sourceMap = make(map[string]int)
        for i:=0; i < len(sources); i++ {
            amount := 0
            //check for 'no other bags'
            if !strings.Contains(sources[i], "no other") {
                amount,_ = strconv.Atoi(string(sources[i][0]))
            }
            sources[i] = sources[i][2:len(sources[i])]
            sources[i] = strings.ReplaceAll(sources[i],"bags","bag")
            if i == len(sources)-1 {
                sources[i] = strings.TrimRight(sources[i], ".")
            }
            //load into map
            sourceMap[sources[i]] = amount
        }
        output = append(output, Node{destination, sourceMap})
    }
    return
}

/**PART1**/
func totalShinyBags(nodeList []Node) int {
    graph := make(map[string][]string)
    for _, node := range nodeList {
        destination, sources := node.destination,node.sources
        if _,keyExits := graph[destination]; !keyExits {
            //add node to addyacency list
            graph[destination] = []string{}
        }
        for sourceName,_ := range sources {
            //make sure all nodes are present in the adyacency list
            if _,keyExits := graph[sourceName]; !keyExits {
                graph[sourceName] = []string{}
            }
            //add destination to the souce-element list
            graph[sourceName] = append(graph[sourceName], destination)
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

/**PART2**/
func totalBagsInsideShinyOnes(nodeList []Node) int {
    graph := make(map[string] map[string]int)
    //recover each node and add it to the adyacency list
    for _, node := range nodeList {
        destination, sources := node.destination, node.sources
        graph[destination] = sources
    }
    return bfsTraversal(graph, "shiny gold bag") -1
}
func bfsTraversal(graph map[string] map[string]int, root string) int {
    //traverse the adyacency list and calculate weight
    count := 1
    for neighborName,neighborValue := range graph[root]{
        // fmt.Printf("%s : %v\n",neighborName, neighborValue)
        count += neighborValue * bfsTraversal(graph, neighborName)
    }
    return count
}
