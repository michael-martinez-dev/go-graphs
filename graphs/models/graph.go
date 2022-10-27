package models

type Graph struct {
	vertices *[]Vertex
	isWeighted bool
	isDirected bool
}

func NewGraph(weighted, directed bool ) *Graph {
	return &Graph{
		vertices: new([]Vertex),
		isWeighted: weighted,
		isDirected: directed,
	}
}

func (g *Graph) IsWeighted() bool {
	return g.isWeighted
}

func (g *Graph) IsDirected() bool {
	return g.isDirected
}

func (g *Graph) GetVertices() *[]Vertex {
	return g.vertices
}

func (g *Graph) GetVertexByValue(data any) *Vertex {
	for i, v := range *g.vertices {
		if v.data == data {
			return &(*g.vertices)[i]
		}
	}
	return nil
}

func (g *Graph) AddVertex(data string) *Vertex {
	newVertex := NewVertex(data)
	newVertextList := append(*g.GetVertices(), *newVertex)
	g.vertices = &newVertextList
	return &(*g.vertices)[len(*g.vertices) - 1]
}

func (g *Graph) AddEdge(vertex1, vertex2 *Vertex, weight int) {
	pWeight := &weight
	if !g.isWeighted {
		pWeight = nil
	}
	g.GetVertexByValue(vertex1.data).AddEdge(*vertex2, pWeight)
	if !g.isDirected {
		vertex2.AddEdge(*vertex1, pWeight)
	}
}

func (g *Graph) RemoveEdge(vertex1, vertex2 *Vertex) {
	vertex1.RemoveEdge(*vertex2)
	if !g.isDirected {
		vertex2.RemoveEdge(*vertex1)
	}
}

func (g *Graph) RemoveVertex(vertex Vertex) {
	for i, v := range *g.vertices {
		if v == vertex {
			newVerticesList := append((*g.vertices)[:i], (*g.vertices)[i+1:]...)
			g.vertices = &newVerticesList
		}
	}
}

func (g *Graph) Print() {
	for _, v := range *g.vertices {
		v.Print(g.isWeighted)
	}
}