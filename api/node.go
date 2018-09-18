package api

// Node allows to communicate with an EOS node via its HTTP API
type Node struct {
	APIEndpoint string
}

// NewNode returns a Node
func NewNode(endpointURL string) *Node {
	return &Node{
		APIEndpoint: endpointURL,
	}
}
