package main

import (
	"fmt"
)

// User represents a user with an ID and a Name
type User struct {
	ID   int
	Name string
}

// Node represents a node in the binary search tree
type Node struct {
	User  User
	Left  *Node
	Right *Node
}

// Insert a new user into the tree
func (n *Node) Insert(user User) {
	if n == nil {
		return // Ideally, this should never happen
	}

	if user.ID < n.User.ID {
		if n.Left == nil {
			n.Left = &Node{User: user}
		} else {
			n.Left.Insert(user)
		}
	} else if user.ID > n.User.ID {
		if n.Right == nil {
			n.Right = &Node{User: user}
		} else {
			n.Right.Insert(user)
		}
	}
	// If user.ID == n.User.ID, we won't insert it to avoid duplicates
}

// Search for a user by ID
func (n *Node) Search(userID int) (*User, bool) {
	if n == nil {
		return nil, false
	}

	if userID < n.User.ID {
		return n.Left.Search(userID)
	} else if userID > n.User.ID {
		return n.Right.Search(userID)
	}

	return &n.User, true
}

func main() {
	// Creating a new BST with a root user
	root := &Node{User: User{ID: 100, Name: "John Doe"}}

	// Inserting new users
	root.Insert(User{ID: 50, Name: "Jane Smith"})
	root.Insert(User{ID: 150, Name: "Alice Johnson"})
	root.Insert(User{ID: 25, Name: "Bob Brown"})

	// Searching for users
	user, found := root.Search(50)
	if found {
		fmt.Printf("User found: %v\n", user)
	} else {
		fmt.Println("User not found")
	}

	user, found = root.Search(99)
	if found {
		fmt.Printf("User found: %v\n", user)
	} else {
		fmt.Println("User not found")
	}
}

