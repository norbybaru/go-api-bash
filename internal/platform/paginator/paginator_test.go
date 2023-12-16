package paginator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPaginator(t *testing.T) {
	tests := []struct {
		tag                                                                    string
		page, perPage, total                                                   int
		expectedPage, expectedPerPage, expectedTotal, pageCount, offset, limit int
	}{
		// varying page
		{"t1", 1, 20, 50, 1, 20, 50, 3, 0, 20},
		{"t2", 2, 20, 50, 2, 20, 50, 3, 20, 20},
		{"t3", 3, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t4", 4, 20, 50, 4, 20, 50, 3, 60, 20},
		{"t5", 0, 20, 50, 1, 20, 50, 3, 0, 20},

		// varying perPage
		{"t6", 1, 0, 50, 1, 15, 50, 4, 0, 15},
		{"t7", 1, -1, 50, 1, 15, 50, 4, 0, 15},
		{"t8", 1, 100, 50, 1, 100, 50, 1, 0, 100},
		{"t9", 1, 1001, 50, 1, 100, 50, 1, 0, 100},

		// varying total
		{"t10", 1, 20, 0, 1, 20, 0, 0, 0, 20},
		{"t11", 1, 20, -1, 1, 20, -1, 1, 0, 20},
	}

	for _, test := range tests {
		p := NewPaginator(test.page, test.perPage, test.total)
		assert.Equal(t, test.expectedPage, p.CurrentPage, test.tag)
		assert.Equal(t, test.expectedPerPage, p.PerPage, test.tag)
		assert.Equal(t, test.expectedTotal, p.TotalRecords, test.tag)
		assert.Equal(t, test.pageCount, p.TotalPages, test.tag)
		assert.Equal(t, test.offset, p.Offset(), test.tag)
		assert.Equal(t, test.limit, p.Limit(), test.tag)
	}
}
