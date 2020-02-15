package main

type vertex struct {
	value      string
	visited    bool
	neighbours []*vertex
}

func newVertex(value string) *vertex {
	return &vertex{
		value:      value,
		visited:    false,
		neighbours: nil,
	}
}

func (v *vertex) connect(vertex ...*vertex) {
	v.neighbours = append(v.neighbours, vertex...)
}

func main() {
	v1 := newVertex("A")
	v2 := newVertex("B")
	v3 := newVertex("C")
	v4 := newVertex("D")
	v5 := newVertex("E")
	v3.connect(v4, v5)
	v1.connect(v2, v3)

}
