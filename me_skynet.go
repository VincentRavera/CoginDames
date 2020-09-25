// https://www.codingame.com/training/medium/skynet-revolution-episode-1
package main

import (
	"errors"
	"fmt"
	"os"
)

func Find(slice []int, val int) (int, bool){
    for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func printArray(target []int) {
	for _, value := range target {
		fmt.Fprintf(os.Stderr, "out>%d\n", value)
	}
}

func printLinks(links map[int][]int) {
	for key, val := range links {
		msg := fmt.Sprintf("%d:", key)
		fmt.Fprintln(os.Stderr, msg)
		for i := 0; i < len(val); i++ {
			msg2 := fmt.Sprintf("->%d", val[i])
			fmt.Fprintln(os.Stderr, msg2)
		}
	}
}

func cpmap(links map[int][]int) map[int][]int{
    output := make(map[int][]int)
	for key, value := range links {
		output[key] = value
	}
	return output
}

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}

func removeALink(links map[int][]int, n1 int, n2 int) map[int][]int {
	output := cpmap(links)
	for key, value := range links {
		newVal := value
		if key == n1 || key == n2 {
			for i := 0; i < len(value); i++ {
				if value[i] == n1 || value[i] == n2 {
					newVal = remove(newVal, i)
				}
			}
			output[key] = newVal
		}
	}
	return output

}

func BFS(links map[int][]int, src int, exits []int) (int, int, error) {
	// 1
	var file []int
	var ferme []int
	previousNode := src
	file = append(file, src)

	// 2
	for len(file) >= 0 {
		curent_node := file[0]
		file = file[1:]
		ferme = append(ferme, curent_node)
		// 3
		for i := 0; i < len(links[curent_node]); i++ {
			_, found := Find(ferme, links[curent_node][i])
			if !found {
				file = append(file, links[curent_node][i])
			}
		}
		_, isExit := Find(exits, curent_node)
		if isExit {
			return previousNode, curent_node, nil
		}
		previousNode = curent_node
	}
	return -1, -1, errors.New("Not solvable")

}

type Node struct {
	id int
	distance int
	visited bool
	previousNodeId int
}

func FindNode(list []Node, val int) (int, bool) {
    for i, item := range list {
		if item.id == val {
			return i, true
		}
	}
	return -1, false
}

func removeNode(s []Node, i int) []Node {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}


func findDistance(a int, b int) int {
    return 1
}

func FindClosesteNode(list []Node) (int) {
	nodeId := -1
	minDist := 99999
    for _, item := range list {
		if item.distance < minDist {
			nodeId = item.id
			minDist = item.distance
		}
	}
	return nodeId
}

func Djkstra(links map[int][]int, src int, exits []int, NumberOfNode int) (int, int) {
	// 1
	unvisited := make([]Node, NumberOfNode)
	for i := 0; i < NumberOfNode; i++ {
		unvisited[i] = Node{id: i, distance: 9999, visited: false, previousNodeId: -2}
	}
	// all_nodes := unvisited

	// 2
	start_id, _ := FindNode(unvisited, src)
	currentNode := unvisited[start_id]
	currentNode.distance=0

	for len(unvisited) >= 0 {
		// 3
		msg := fmt.Sprintf("Visiting: %d", currentNode.id)
		fmt.Fprintln(os.Stderr, msg)
		for i := 0; i < len(links[currentNode.id]); i++ {
			nextNodeId := links[currentNode.id][i]
			nextNodeIndex, isUnvis := FindNode(unvisited, nextNodeId)
			if isUnvis {
				nextNode := unvisited[nextNodeIndex]
				newDistance := currentNode.distance + findDistance(currentNode.id, nextNode.id)
				fmt.Fprintf(os.Stderr, "-> node:%d;distance:%d;\n", nextNode.id, newDistance)
				if newDistance < nextNode.distance {
					nextNode.distance = newDistance
					nextNode.previousNodeId = currentNode.id
					unvisited[nextNodeIndex] = nextNode
				}
			}
		}
		// 4
		currentNode.visited = true
		cnodeindex, _ := FindNode(unvisited, currentNode.id)
		unvisited = removeNode(unvisited, cnodeindex)

		// 5
		_, isExit := Find(exits, currentNode.id)
		if isExit {
			fmt.Fprintf(os.Stderr, "Found exit node: %d\n", currentNode.id)
			return currentNode.id, currentNode.previousNodeId
		}
		// 6
		currentNodeid := FindClosesteNode(unvisited)
		currentNodeindex, _ := FindNode(unvisited, currentNodeid)
		currentNode = unvisited[currentNodeindex]
	}

	return -1, -1

}

func process(links map[int][]int, gates []int, si int, numberOfNode int) (map[int][]int, int, int) {
	fmt.Fprintln(os.Stderr, "------------------------------")
	// n1, n2, _ := BFS(links, si, gates)
	n1, n2 := Djkstra(links, si, gates, numberOfNode)
	printLinks(links)
	fmt.Fprintf(os.Stderr, "Removing: %d-%d \n", n1, n2)
	links = removeALink(links, n1, n2)
	printLinks(links)
	fmt.Fprintln(os.Stderr, "------------------------------")
	return links, n1, n2
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	fmt.Fprintln(os.Stderr, "-----------INIT---------------")
	var N, L, E int
	fmt.Scan(&N, &L, &E)
	links := make(map[int][]int)
	exits := make([]int,0)

	var msg string
	msg = fmt.Sprintf("There is %d nodes", N)
	fmt.Fprintln(os.Stderr, msg)


	fmt.Fprintln(os.Stderr, "LINKS:")
	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		msg = fmt.Sprintf("%d:%d", N1, N2)
		fmt.Fprintln(os.Stderr, msg)
		links[N1] = append(links[N1], N2)
		links[N2] = append(links[N2], N1)

	}
	fmt.Fprintln(os.Stderr, "EXITS:")
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		exits = append(exits, EI)
		msg = fmt.Sprintf("%d:", EI)
		fmt.Fprintln(os.Stderr, msg)
	}
	printArray(exits)
	fmt.Fprintln(os.Stderr, "-----------MAIN---------------")
	fmt.Fprintln(os.Stderr, "------------------------------")
	printLinks(links)
	fmt.Fprintln(os.Stderr, "------------------------------")
	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)
		var out string
		var N1 int
		var N2 int

		msg := fmt.Sprintf(">>>>SI: %d|", SI)
		fmt.Fprintln(os.Stderr, msg)

		links, N1, N2 = process(links, exits, SI, N)

		out = fmt.Sprintf("%d %d", N1, N2)
		fmt.Println(out)
		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// Example: 0 1 are the indices of the nodes you wish to sever the link between
		// fmt.Println("0 1")
	}
}
