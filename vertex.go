package dag

type vertex struct {
	Id       string
	Value    interface{}
	Children []*vertex
}

func newVertex(id string, value interface{}) *vertex {
	return &vertex{
		Id:       id,
		Value:    value,
		Children: make([]*vertex, 0),
	}
}
