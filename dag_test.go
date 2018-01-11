package dag_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sensimevanidus/dag"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail) // Sole connection point between Ginkgo and Gomega.
	RunSpecs(t, "DAG Testing Suite")
}

var _ = Describe("DAG", func() {
	Context("initially", func() {
		var graph *dag.DAG

		Specify("DAG must exist", func() {
			graph = dag.NewDAG()

			Expect(graph).ShouldNot(BeNil())
		})

		Specify("Edge should not exist", func() {
			edgeExists, err := graph.EdgeExists("1", "2")

			Expect(err).ShouldNot(BeNil())
			Expect(edgeExists).Should(BeFalse())
		})
	})

	/*
	representationOfDAG := map[string][]string {
		"1": {"2"},
		"2": {"3", "4"},
		"4": {"3"},
	}
	*/
	Context("A newly created DAG", func() {
		graph := dag.NewDAG()

		It("should not contain any vertices", func() {
			Expect(graph.Vertices).ShouldNot(BeNil())
			Expect(len(graph.Vertices)).Should(Equal(0))
		})

		It("should contain 1 vertex when a new vertex is added", func() {
			err := graph.AddVertex("1", nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(graph.Vertices)).Should(Equal(1))
		})

		It("should represent a DAG structure as expected", func() {
			err := graph.AddVertex("2", nil)
			Expect(err).ShouldNot(HaveOccurred())

			err = graph.AddVertex("3", nil)
			Expect(err).ShouldNot(HaveOccurred())

			err = graph.AddVertex("4", nil)
			Expect(err).ShouldNot(HaveOccurred())

			err = graph.AddEdge("1", "2")
			Expect(err).ShouldNot(HaveOccurred())


			err = graph.AddEdge("2", "3")
			Expect(err).ShouldNot(HaveOccurred())

			err = graph.AddEdge("2", "4")
			Expect(err).ShouldNot(HaveOccurred())

			err = graph.AddEdge("4", "3")
			Expect(err).ShouldNot(HaveOccurred())

			var ok bool
			ok, err = graph.EdgeExists("1", "2")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())

			ok, err = graph.EdgeExists("2", "3")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())

			ok, err = graph.EdgeExists("2", "4")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())

			ok, err = graph.EdgeExists("4", "3")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())

			Expect(graph.GetVertex("1")).ShouldNot(BeNil())
			Expect(graph.GetVertex("100")).Should(BeNil())
			Expect(len(graph.GetVertex("1").Children)).Should(Equal(1))
			Expect(len(graph.GetVertex("2").Children)).Should(Equal(2))
			Expect(len(graph.GetVertex("3").Children)).Should(Equal(0))
			Expect(len(graph.GetVertex("4").Children)).Should(Equal(1))
		})
	})
})
