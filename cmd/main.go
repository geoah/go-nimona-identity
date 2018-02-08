package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/geoah/go-nimona-identity"
)

func main() {
	bls := identity.Blocks{
		Blocks: []identity.Block{
			identity.Block{
				Event: identity.BlockEvent{
					Author: "geoah",
					Payload: identity.BlockEventPayload{
						Type:    "identity",
						Content: []byte("geoah"),
					},
					Type: string(identity.EventTypeGraphCreate),
				},
			},
		},
	}

	bs, err := bls.Marshal()
	if err != nil {
		log.Fatalf("Could not marshal block, %v", err)
	}

	ubls := &identity.Blocks{}
	if err := ubls.Unmarshal(bs); err != nil {
		log.Fatalf("Could not unmarshal bs, %v", err)
	}

	jbs, _ := json.MarshalIndent(ubls, "", "  ")
	fmt.Println(string(jbs))
}
