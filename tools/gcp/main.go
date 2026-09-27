package main

import (
	"log"
	"fmt"
	"os" 
	"gcp/vars/args" 
	"gcp/structs/network"
	"gcp/structs/eData"
	"gcp/vars/config" 
	"gcp/funcs/io" 
	"gcp/funcs/benchmark"
)


func main() {
	f, err := os.Create("gcp.log") 
	if err != nil {
		fmt.Println("Unable to create gcp.log, printing to stdout") 
		return 
	} else {
		
		defer close(f) 
		log.SetOutput(f) 
	}


	var prior *network.Network 
	var allData *eData.AllData
	var err error 

	prior, allData, err = io.ParseInput(
		args.NetworkPath,
		args.ExpressionPath, 
	) 

	prior, nulls, err = benchmark.CalcLogLs(prior, allData) 
	if err != nil {
		log.Println(err)
		os.Exit(1) 
	}
	
	fmt.Println(pvalue) 
}










































