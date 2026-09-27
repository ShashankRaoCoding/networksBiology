package network

import (
	"os"
	"strings" 
	"bufio"
	"log"
	"fmt" 
)

// from, to, weight 

func NetworkFromCsv(networkPath string, sep string) (*Network, error) {
	n = New() 
	var err error 

	log.Println("Starting to read network data") 

	f, err := os.Open(networkPath) 
	if err != nil {
		log.Printf("There was an error reading the network data at %s: %s", networkPath, err) 
		return n, err 
	}
	defer close(f) 

	reader := bufio.NewReader(f) 
	for {
		line , err := reader.ReadString('\n') 
		line = strings.TrimSuffix(line, "\n")
		values := strings.Split(line, ",") 
		if err != nil {
			break 
		}
	}

	if err == io.EOF {
		
	} else {
		log.Printf("There was an error reading the network data: %s", err) 
		return n, err 
	}
}










































