package part1

import (
	"os"
	"strconv"
	"strings"
)

type Node struct {
	next  *Node
	value int
}

func (n *Node) insert(value int, ll *LinkedList) {
	store := n.next
	n.next = &Node{value: value, next: store}
	ll.length++
}

type LinkedList struct {
	head   *Node
	length int
	tail   *Node
}

func (ll *LinkedList) append(value int) {
	newNode := Node{value: value}
	if ll.length == 0 {
		ll.head = &newNode
	} else {
		ll.tail.next = &newNode
	}
	ll.tail = &newNode
	ll.length++
}

func parseText(text string) LinkedList {
	list := strings.Split(text, " ")

	// create linked list
	var linkedList LinkedList = LinkedList{head: nil, length: 0, tail: nil}

	// create each successive element
	for i := 0; i < len(list); i++ {
		integer, _ := strconv.Atoi(list[i])
		linkedList.append(integer)
	}

	return linkedList
}

func (ll *LinkedList) applyRules(n int) {
	for i := 0; i < n; i++ {
		currentNode := ll.head
		for j := 0; j < ll.length; j++ {
			if currentNode.value == 0 {
				currentNode.value = 1
				if currentNode.next != nil {
				}
				currentNode = currentNode.next
				continue
			}
			if currentNode.value == 1 {
				currentNode.value = 2024
				currentNode = currentNode.next
				continue
			}
			str := strconv.Itoa(currentNode.value)
			if len(str)%2 == 0 {
				p1, _ := strconv.Atoi(str[:len(str)/2])
				p2, _ := strconv.Atoi(str[len(str)/2:])

				currentNode.value = p1
				currentNode.insert(p2, ll)
				currentNode = currentNode.next
				currentNode = currentNode.next
				j++
				continue
			}
			currentNode.value = currentNode.value * 2024
			currentNode = currentNode.next
		}
	}
}

func DoPuzzle(file string, iterations int) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	linkedList := parseText(string(text))

	linkedList.applyRules(iterations)

	return linkedList.length
}
