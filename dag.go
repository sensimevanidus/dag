package dag

import "fmt"

type DAG struct {
	Vertices map[string]*vertex
}

func (dag *DAG) AddVertex(id string, value interface {}) error {
	if _, ok := dag.Vertices[id]; ok {
		return fmt.Errorf("dag already contains a vertex with the id %v, so it won't be added", id)
	}

	dag.Vertices[id] = newVertex(id, value)
	return nil
}

func (dag *DAG) AddEdge(fromVertexId, toVertexId string) error {
	var from, to *vertex
	var ok bool

	if from, ok = dag.Vertices[fromVertexId]; !ok {
		return fmt.Errorf("vertex with the id %v not found, so it can't be added", fromVertexId)
	}

	if to, ok = dag.Vertices[toVertexId]; !ok {
		return fmt.Errorf("vertex with the id %v not found, so it can't be added", toVertexId)
	}

	for _, childVertex := range from.Children {
		if childVertex == to {
			return fmt.Errorf("edge (%v,%v) already exists, so it can't be added", fromVertexId, toVertexId)
		}
	}

	from.Children = append(from.Children, to)
	return nil
}

func (dag *DAG) EdgeExists(fromVertexId, toVertexId string) (bool, error) {
	var from, to *vertex
	var ok bool

	if from, ok = dag.Vertices[fromVertexId]; !ok {
		return false, fmt.Errorf("vertex with the id %v not found, so it can't be added", fromVertexId)
	}

	if to, ok = dag.Vertices[toVertexId]; !ok {
		return false, fmt.Errorf("vertex with the id %v not found, so it can't be added", toVertexId)
	}

	for _, childVertex := range from.Children {
		if childVertex == to {
			return true, nil
		}
	}

	return false, nil
}

func (dag *DAG) GetVertex(id string) (*vertex) {
	if v, ok := dag.Vertices[id]; ok {
		return v
	}

	return nil
}

func NewDAG() *DAG {
	return &DAG{
		Vertices: make(map[string]*vertex, 0),
	}
}