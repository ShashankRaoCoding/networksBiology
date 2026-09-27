package vertex

type Vertex struct {
	Name string 
	LogL float64
	Regulators map[*Vertex]float64
	Regulatees map[*Vertex]float64 
	Intercept float64 
}












































