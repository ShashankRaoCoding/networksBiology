package io

import (
	// "bufio" 
	// "os"
	"gcp/structs/network"
	"gcp/structs/eData"
	"fmt" 
	"log" 
)

func ParseInput(networkPath string, expressionPath string) (network.Network, eData.AllData, error) {
	var n network.Network
	var e eData.AllData 
	var err error 

	log.Println("Parsing inputs") 
	
	n, err = network.NetworkFromCsv(networkPath, ",") 
	if err != nil {
		return n, e, err
	}

	log.Println("Successfully parsed network date") 
	
	e, err = eData.AllDataFromCsv(expressionPath, ",") 
	if err != nil {
		return n, e, err 
	}

	log.Println("Successfully parsed expression data") 
	
	return n, e, nil 
}










































