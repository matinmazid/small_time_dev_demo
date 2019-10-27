package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	Ptr   *node
	Value string
}

func main() {

	fileData := openFile("./linkedList/findNthFromTheEnd/data/findNthFromTheEnd.dat")

	headPtr := generateLinkedList(fileData)
	printLinkedList(headPtr)
	mySolution(headPtr, 7)
	expectedSolution(headPtr, 7)
}

func expectedSolution(headPtr *node, nth int) *node {

	ptr := headPtr
	ptrLookahead := headPtr

	//first move the look ahead forward n paces
	i := 0
	for ; i < nth && ptrLookahead.Ptr != nil; i++ {
		ptrLookahead = ptrLookahead.Ptr
	}

	// now move both the look ahead and the main pointer
	// forward on node at a time until the look head
	// reaches the end
	for ptrLookahead.Ptr != nil {
		ptrLookahead = ptrLookahead.Ptr
		ptr = ptr.Ptr
	}

	// time of this is 2(N)
	fmt.Printf("\n expected %+v", *ptr)
	return ptr
}

func mySolution(headPtr *node, nth int) {

	arrayPtr := make([]node, 5)

	// copy the linked list to a
	// growable array
	for headPtr.Ptr != nil {
		arrayPtr = append(arrayPtr, *headPtr)
		headPtr = headPtr.Ptr
	}

	//get the length of the array and subract
	//to get the index of the desired node
	element := len(arrayPtr) - nth
	someNode := node{}
	ptr := &someNode
	ptr.Value = "kdkdk"
	// time of this is cost of copying an array
	// and growing it.
	fmt.Printf("\n mySolution %+v", arrayPtr[element])
}

func printLinkedList(head *node) {

	for head.Ptr != nil {
		fmt.Printf("%s ", head.Value)
		head = head.Ptr
	}
}

func generateLinkedList(fileData string) *node {
	listOfNames := strings.Split(fileData, ",")

	newNOde := &node{}
	headPtr := newNOde

	arrayLength := len(listOfNames)
	for i := 0; i < arrayLength; i++ {
		newNOde.Value = listOfNames[i]
		if i+1 < arrayLength {
			nextNode := &node{}
			newNOde.Ptr = nextNode
			newNOde = nextNode
		}

	}
	return headPtr
}

func openFile(filePath string) string {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
