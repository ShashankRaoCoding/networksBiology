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

	line, err := reader.ReadString('\n')
	if err == io.EOF {
		log.Printf("Network data is empty\n") 
		return n, fmt.Errorf("Network data is empty")
	} else if err != nil {
		log.Printf("Unexpected Error: Could not read network data: %s\n", err) 
		return n, err 
	}
	
	columns := map[string]int{
		"from": -1,
		"to": -1,
	} 
	
	values := strings.Split(line, ",") 
	if len(values) < 2 {
		log.Printf("Fewer than 2 columns\n") 
		return n, fmt.Errorf("Fewer than 2 coloumns") 
	}
	
	for i, v := range values {
		_, o := columns[v]
		if o {
			columns[v] = i 
		}
	}

	for c, i := range columns {
		if i == -1 {
			log.Printf("%s with index of -1", c) 
			return n, fmt.Errorf("Data must have column %s", c) 
		}
	}

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










































