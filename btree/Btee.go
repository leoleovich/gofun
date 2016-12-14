package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	left        *node
	right       *node
	val         int
	initialized bool
}

func insertBtree(head *node, value int) {
	if !head.initialized {
		head.val = value
		head.initialized = true
		return
	}
	tmpNode := node{}
	if value < head.val && head.left == nil {
		tmpNode.val = value
		tmpNode.initialized = true
		head.left = &tmpNode
	} else if value < head.val {
		insertBtree(head.left, value)
	} else if value >= head.val && head.right == nil {
		tmpNode.val = value
		tmpNode.initialized = true
		head.right = &tmpNode
	} else {
		insertBtree(head.right, value)
	}
}

func findPath(curHead *node, value int, path *[]int) bool {
	*path = append(*path, curHead.val)
	if curHead.val == value {
		return true
	} else if value < curHead.val && curHead.left != nil {
		if findPath(curHead.left, value, path) {
			return true
		}
	} else if value < curHead.val {
		return false
	} else if value >= curHead.val && curHead.right != nil {
		if findPath(curHead.right, value, path) {
			return true
		}
	} else {
		return false
	}
	return false
}

func checkLeft(n *node, res []string, directions []string, howDeep, position int) {
	if howDeep == 0 {
		res[position]=fmt.Sprintf("%v", n.val)
		if n.right != nil {
			directions[position+1]=fmt.Sprintf("\\")
		}
		return
	} else {
		if n.left != nil {
			checkLeft(n.left, res, directions, howDeep-1, position-2)
			checkRight(n.left, res, directions, howDeep-1, position-2)
		}
	}
}

func checkRight(n *node, res []string, directions []string, howDeep, position int) {
	if howDeep == 0 {
		res[position]=fmt.Sprintf("%v", n.val)
		if n.left != nil {
			directions[position-1]=fmt.Sprintf("/")
		}
		return
	} else {
		if n.right != nil {
			checkLeft(n.right, res, directions, howDeep-1, position+2)
			checkRight(n.right, res, directions, howDeep-1, position+2)
		}
	}
}

func generateArray(mapSize int, random bool) []int {
	res := make([]int, mapSize)
	if random {
		for i:=0; i<mapSize; i++ {
			res[i] = int(rand.Int63n(int64(mapSize*mapSize)))
		}
	} else {
		for i, val := range([]int{1,2,3,4,5}) {
			res[i] = val
		}
	}
	return res
}

func main() {
	mapSize := 10
	head := &node{}
	tmpArray := generateArray(mapSize, true)

	for _,val := range(tmpArray) {
		insertBtree(head, val)
	}

	pathVar := 4
	fmt.Println("Search for a path to", pathVar)
	var path []int
	if findPath(head, pathVar, &path) {
		fmt.Println("Found. The path is")
		fmt.Println(path)
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Whole map")
	maxCells := int(mapSize*4)
	startPosition := maxCells/2
	for i:=0; i<mapSize; i++ {
		//arr := generateLine(head, i, maxDots)
		line := make([]string, maxCells)
		directions := make([]string, maxCells)
		checkLeft(head, line, directions, i, startPosition)
		checkRight(head, line, directions, i, startPosition)

		for _, val := range(line){
			if val == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Print("\n")

		for _, val := range(directions){
			if val == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(val, " ")
			}
		}
		fmt.Print("\n")
	}
}
