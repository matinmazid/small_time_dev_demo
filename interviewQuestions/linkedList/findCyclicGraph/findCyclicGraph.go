package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	NextNode *node
	Value    string
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
		fmt.Println(headPtr.Value + " ->")
		headPtr = headPtr.NextNode
	}
}
func generateGraph(fileData string) *node {

	listOfNames := strings.Split(fileData, ",")
	listOfNameSize := len(listOfNames)

	headPtr := &node{}
	tail := headPtr
	for nameIndex := 0; nameIndex < listOfNameSize; {
		tail.Value = listOfNames[nameIndex]
		if nameIndex+1 < listOfNameSize {
			tail.NextNode = &node{}
			tail = tail.NextNode
		}

		nameIndex++
	}
	return headPtr
}

func openFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	return string(data)
}
