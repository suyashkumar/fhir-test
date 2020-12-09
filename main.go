package main

import (
	"fmt"
	"log"

	"github.com/google/fhir/go/jsonformat"
)

func main() {
	json := `{"resourceType":"Patient", "id": "exampleID1"}`
	um, err := jsonformat.NewUnmarshaller("America/Los_Angeles", jsonformat.R4)
	if err != nil {
		log.Fatalf("Failed to create unmarshaller %v", err)
	}
	unmarshalled, err := um.UnmarshalR4([]byte(json))
	fmt.Println(unmarshalled.GetPatient().Id.GetValue())
}
