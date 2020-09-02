你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
