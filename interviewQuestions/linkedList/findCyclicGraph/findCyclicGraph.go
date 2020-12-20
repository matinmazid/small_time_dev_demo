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

	var nameOfDestNode string
	var headNode node

	headNode, nameOfDestNode, listOfNames = makeNode(listOfNames)
	for true {

		if len(listOfNames) == 0 {
			break
		}

		var newNode node
		newNode, nameOfDestNode, listOfNames = makeNode(listOfNames)
		addToGraph(&headNode, nameOfDestNode, newNode)

	}

	return nil
}

func makeNode(listOfNames []string) (node, string, []string) {

	nodeDescr, listOfNames := listOfNames[0], listOfNames[1:]
	nodeParts := strings.Split(nodeDescr, ":")
	createdNode := node{nil, nodeParts[0]}
	// xx := make([]string, len(listOfNames))
	// copy(xx, listOfNames)
	return createdNode, nodeParts[1], listOfNames
}

func addToGraph(headNode *node, nameOfDestNode string, nodeToAdd node) {

	for headNode.NextNode != nil {
		if headNode.Name == nameOfDestNode {

			nodeToAdd.NextNode = headNode
			break
		} else {

			headNode = headNode.NextNode
		}
	}
	headNode.NextNode = &nodeToAdd
}

func openFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	return string(data)
}
