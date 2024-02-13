README.md

# LRU Cache Implementation in Go

This Go program implements a simple LRU (Least Recently Used) cache using a doubly linked list and a hash map. The cache has a fixed size and evicts the least recently used item when it reaches its maximum capacity.

## Files

- `main.go`: Contains the main program implementing the LRU cache logic.
- `README.md`: This file, providing information about the program.

## Components

### Node
- Represents a node in the doubly linked list.
- Each node contains a string value and pointers to its left and right neighbors.

### Queue
- Represents the doubly linked list.
- Contains a head and tail node to maintain the structure.
- Provides methods to display the elements in the queue.

### Hash
- Represents a hash map mapping string keys to Node pointers.
- Used for quick access to nodes based on their values.

### Cash (LRU Cache)
- Combines the Queue and Hash components.
- Provides methods to add, remove, and check nodes.
- Maintains the LRU cache policy by evicting the least recently used item when the cache size exceeds the specified limit.

## Usage

1. Ensure you have Go installed on your system.
2. Clone the repository or copy the `main.go` file to your local machine.
3. Run the program using the command `go run main.go`.

## Example

The `main` function in `main.go` demonstrates the usage of the LRU cache. It adds and checks several words, displaying the cache after each operation.

```go
func main() {
	cash := NewCash()
	for _, word := range []string{"avocado", "bob", "sun", "avocado", "dog", "sun", "person", "cat", "dog", "grass"} {
		cash.Check(word)
		cash.Display()
	}
}
```

