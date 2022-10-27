package main

import (
    "github.com/mixedmachine/go-graphs/graphs"
)

func main() {
	graphStructure := graphs.NewGraph()
	busNetwork := graphStructure.NewGraphWeightedAndDirected()
    cliftonStation := busNetwork.AddVertex("Clifton")
    capeMayStation := busNetwork.AddVertex("Cape May")
	georgetownStation := busNetwork.AddVertex("Georgetown")

	busNetwork.AddEdge(cliftonStation, capeMayStation, 1000)
	busNetwork.AddEdge(cliftonStation, georgetownStation, 1500)
	busNetwork.AddEdge(georgetownStation, capeMayStation, 500)
	busNetwork.Print()
}
