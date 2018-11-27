package namespace

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var cost Namespace

var (
	NamespaceCost = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mk_namespace_cost",
		Help: "ManagedKube - Cost of the namespace.",
	},
		[]string{"namespace_name", "duration"},
	)
)

// Registers the Prometheus metrics
func Register() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(NamespaceCost)
}

func Add(namespace string, newCost float64) {
	for _, n := range cost.cost {
		if namespace == n.Name {
			n.Cost += newCost
		}
	}
}

func Subtract(namespace string, newCost float64) {
	for _, n := range cost.cost {
		if namespace == n.Name {
			n.Cost -= newCost
		}
	}
}

func Export() {
	for _, n := range cost.cost {
		fmt.Println(n.Name)
		NamespaceCost.With(prometheus.Labels{"namespace_name": n.Name, "duration": "minute"}).Add(n.Cost)
	}
}
