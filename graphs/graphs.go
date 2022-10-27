package graphs

import (
	"github.com/mixedmachine/go-graphs/graphs/models"
)

type GraphInterface interface {
	NewGraphWeightedAndDirected() *models.Graph
	NewGraphWeighted() *models.Graph
	NewGraphDirected() *models.Graph
	NewGraphSimple() *models.Graph
}

type graphBuilder struct {
}

func NewGraph() GraphInterface {
	return &graphBuilder{}
}

func (g *graphBuilder) NewGraphWeightedAndDirected() *models.Graph {
	return models.NewGraph(true, true)
}

func (g *graphBuilder) NewGraphWeighted() *models.Graph {
	return models.NewGraph(true, false)
}

func (g *graphBuilder) NewGraphDirected() *models.Graph {
	return models.NewGraph(false, true)
}

func (g *graphBuilder) NewGraphSimple() *models.Graph {
	return models.NewGraph(false, false)
}
