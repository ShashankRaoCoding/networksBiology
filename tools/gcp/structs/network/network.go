package network

import (
	"gcp/structs/eData" 
	"gcp/structs/vertex" 
	"sync" 
)

type Network struct {
	LogL float64
	Vertices map[string]*vertex.Vertex
	Prior bool
}

func (n *Network) Fit(a eData.AllData) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(n.Vertices)) 
	for name, vertex := range n.Vertices {
		wg.Add(1) 
		go func(v *vertex.Vertex) {
			defer wg.Done() 
			errCh <- v.Fit(a[v.Name])
		} (vertex) 
	}

	wg.Wait()
	for err := range errs {
		errs = append(errs, err) 
	}

	return fmt.Errorf("%v\n", errs) 
	// all data is of form map[geneName]geneData 
	// gene data is of form map[replicateName]expressionValue
}

// type Vertex struct {
	// Name string
	// Regulators []*Vertex
	// Regulatees []*Vertex
	// Coefficients map[*Vertex]float64 
	// Intercept float64 
	// LogL float64 
// }

// func (v *Vertex) Fit(g eData.GeneData) error {
	
// }






































