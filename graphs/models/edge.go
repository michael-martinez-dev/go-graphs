package models

type Edge struct {
	weight *int
	start Vertex
	end Vertex

}

func NewEdge(startV Vertex, endV Vertex, weight *int) *Edge {
	return &Edge{
		weight: weight,
		start: startV,
		end: endV,
		}
}

func (e *Edge) GetWeight() int {
	return *e.weight
}

func (e *Edge) GetStartingVertex() Vertex {
	return e.start
}

func (e *Edge) GetEndingVertex() Vertex {
	return e.end
}
