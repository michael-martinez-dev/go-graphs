package models

import "fmt"

type Vertex struct {
	data any
	edges *[]Edge
}

func NewVertex(data any) *Vertex {
	return &Vertex{
		data: data,
		edges: new([]Edge),
	}
}

func (v *Vertex) GetData() any {
	return v.data
}


func (v *Vertex) GetEdges() *[]Edge {
	return v.edges
}

func (v *Vertex) AddEdge(vertex Vertex, weight *int) {
	newEdge := NewEdge(*v, vertex, weight)
	newEdgeList := append(*v.edges, *newEdge)
	v.edges = &newEdgeList
}

func (v * Vertex) RemoveEdge(vertex Vertex) {
	for i, e := range *v.edges {
		if e.GetEndingVertex() == vertex {
			newEdgeList := append((*v.edges)[:i], (*v.edges)[i+1:]...)
			v.edges = &newEdgeList
		}
	}
}

func (v *Vertex) Print(showWeight bool) {
	message := ""

	if len(*v.edges) == 0 {
		fmt.Printf("%v -->\n", v.data)
		return
	}

	for i, e := range *v.edges {
		if i == 0 {
			message = fmt.Sprintf("%v --> ", e.GetStartingVertex().data)
		}

		message = fmt.Sprintf("%s %v", message, e.GetEndingVertex().data)

		if showWeight {
			message = fmt.Sprintf("%s (%d)", message, e.GetWeight())
		}

		if i != len(*v.edges) - 1 {
			message += ", "
		}
	}
	println(message)
}
