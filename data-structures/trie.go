package datastructures

import "fmt"

const ALPHABET_SIZE = 26

type Trie struct {
	Root *Node
}

type Node struct {
	isEnd bool
	children [ALPHABET_SIZE]*Node
}

func InitTrie() *Trie {
	trie := &Trie{Root: &Node{}}
	return trie
}

// Insert word into trie
func (trie *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := trie.Root
	for i := 0; i < wordLength; i++ {
		// 'a' is a rune equal to 97
		// we want 'a' to represent 0, 'b' to represent 1 etc (hence offset below) => code - 94 = range{0, 25}
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			// create child node off current node with char
			currentNode.children[charIndex] = &Node{}
		}

		// Update current node being inspected in preparation for next loop
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search for word in trie (return true or false depending on presence)
func (trie *Trie) Search(word string) bool {
 	wordLength := len(word)
 	currentNode := trie.Root

 	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
 	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func TestTrie() {
	trie := InitTrie()

	toAdd := []string{
		"orc",
		"orange",
	}

	for _, v := range toAdd {
		trie.Insert(v)
	}

	fmt.Printf("or: %v \n", trie.Search("or")) // false
	fmt.Printf("orc: %v \n", trie.Search("orc")) // true
	fmt.Printf("ork: %v \n", trie.Search("ork")) // false
	fmt.Printf("oran: %v \n", trie.Search("oran")) // false
	fmt.Printf("orange: %v \n", trie.Search("orange")) // true


	x := trie.Root.children
	fmt.Println(x)
	y := (x[14]).children
	fmt.Println(y)
}