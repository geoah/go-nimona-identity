package identity

// EventType for the BlockEvents
type EventType string

const (
	// EventTypeGraphCreate is the first event for all graphs
	EventTypeGraphCreate EventType = "graph.create"
)

// Block is our BlockEvent wrapper
// proteus:generate
type Block struct {
	Event     BlockEvent `json:"event"`
	Signature string     `json:"signature"`
	Hash      string     `json:"hash,omitempty"`
}

// BlockEvent holds all the information for each graph node
type BlockEvent struct {
	ACL struct {
		Read  []string `json:"read,omitempty"`
		Write []string `json:"write,omitempty"`
	} `json:"acl,omitempty"`
	Author  string `json:"author"`
	Payload struct {
		Type    string `json:"type,omitempty"`
		Content []byte `json:"content,omitempty"`
	} `json:"payload"`
	Nonce   string   `json:"nonce"`
	Parents []string `json:"parents,omitempty"`
	Type    string   `json:"type"`
}
