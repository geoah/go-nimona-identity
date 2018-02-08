package identity

// EventType for the BlockEvents
//proteus:generate
type EventType string

const (
	// EventTypeGraphCreate is the first event for all graphs
	EventTypeGraphCreate EventType = "graph.create"
)

//go:generate proteus -p github.com/geoah/go-nimona-identity -f $GOPATH/src

// Block is our BlockEvent wrapper
//proteus:generate
type Block struct {
	Event     BlockEvent `json:"event"`
	Signature string     `json:"signature"`
	Hash      string     `json:"hash,omitempty"`
}

// BlockEvent holds all the information for each graph node
//proteus:generate
type BlockEvent struct {
	ACL struct {
		Read  []string `json:"read,omitempty"`
		Write []string `json:"write,omitempty"`
	} `json:"acl,omitempty"`
	Author  string            `json:"author,omitempty"`
	Payload BlockEventPayload `json:"payload,omitempty"`
	Nonce   string            `json:"nonce,omitempty"`
	Parents []string          `json:"parents,omitempty"`
	Type    string            `json:"type,omitempty"`
}

// BlockEventPayload is the arbitrary payload for each event
//proteus:generate
type BlockEventPayload struct {
	Type    string `json:"type,omitempty"`
	Content []byte `json:"content,omitempty"`
}
