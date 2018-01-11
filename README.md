# dag

DAG implementation in **Go**.

## Usage

Let's suppose you want to represent the following DAG (represented as a JSON dictionary):

```json
{
    "1": ["2"],
    "2": ["3", "4"],
    "4": ["3"]
}
```

You should do the following:

```go
// Initialize the graph
graph := dag.NewDAG()

// Add the vertices (Here, the first parameter is the id of the vertex,
// and the second one is its value. Right now no value is set for the
// sake of simplicity)
graph.AddVertex("1", nil)
graph.AddVertex("2", nil)
graph.AddVertex("3", nil)
graph.AddVertex("4", nil)

// Add the edges (Note that given vertices must exist before adding an
// edge between them)
graph.AddEdge("1", "2")
graph.AddEdge("2", "3")
graph.AddEdge("2", "4")
graph.AddEdge("4", "3")
```

Voila, that's all about it. Now you can start using the DAG as you wish.

```go
// Get a vertex
graph.Vertex("1") // returns nil if not found

// Get children of a vertex
graph.Vertex("1").Children

// Verify that an edge exists
graph.EdgeExists("1", "2") // returns true
```