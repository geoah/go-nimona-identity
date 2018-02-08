package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/geoah/go-nimona-identity"
)

func main() {
	bl := &identity.Block{
		Event: identity.BlockEvent{
			Author: "geoah",
			Payload: identity.BlockEventPayload{
				Type:    "identity",
				Content: []byte("geoah"),
			},
			Type: string(identity.EventTypeGraphCreate),
		},
	}

	bs, err := bl.Marshal()
	if err != nil {
		log.Fatalf("Could not marshal block, %v", err)
	}

	ubl := &identity.Block{}
	if err := ubl.Unmarshal(bs); err != nil {
		log.Fatalf("Could not unmarshal bs, %v", err)
	}

	jbs, _ := json.MarshalIndent(ubl, "", "  ")
	fmt.Println(string(jbs))
}
