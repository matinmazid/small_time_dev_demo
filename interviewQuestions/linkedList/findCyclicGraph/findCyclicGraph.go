package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	NextNode *node
	Name     string
}

func main() {

	fmt.Printf("starting")
	fileData := openFile("/home/matin/src/go/src/github.com/matinmazid/small_time_dev_demo/interviewQuestions/linkedList/findCyclicGraph/data/nonCyclic.dat")

	fmt.Printf("generating graph")
	headPtr := generateGraph(fileData)

	printLinkedList(headPtr)
	// mySolution(headPtr, 7)
	// expectedSolution(headPtr, 7)
}

func printLinkedList(headPtr *node) {

	for headPtr != nil {
		fmt.Println(headPtr.Name + " ->")
		headPtr = headPtr.NextNode
	}
}
func generateGraph(fileData string) *node {

	listOfNames := strings.Split(fileData, ",")

	var nodeDescr string
	var headNode node
	// pop off array
	doLoop := true

	for doLoop {
		nodeDescr, listOfNames = listOfNames[0], listOfNames[1:]

		if len(listOfNames) == 0 {
			doLoop = false
		}

		nodeParts := strings.Split(nodeDescr, ":")
		newNode := node{nil, nodeParts[0]}
		addToGraph(&headNode, newNode)
		fmt.Println("kdkd ", headNode)

	}

	return nil
}

func addToGraph(headNode *node, nodeToAdd node) {
	if headNode.Name == "" {
		headNode = &nodeToAdd
	}
}
func openFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	return string(data)
}
