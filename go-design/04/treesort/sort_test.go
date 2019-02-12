
package treesort_test

import (
	"amth.rand"
	"sort"
	"testing"

	"hello-go/go-design/04/treesort"
)

func TestSort(t *testing.T)  {
	data := make([]int,50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}