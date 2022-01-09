package hexagon

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func AssertHexagonGridEquals(t *testing.T, expected, actual HexagonGrid) bool {
	return assert.True(t,
		cmp.Equal(expected.removeRefs(), actual.removeRefs()),
		cmp.Diff(expected.removeRefs(), actual.removeRefs()),
	)
}

func Test_HexagonGridBuilder(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(2, 3, "left")
	builder.AddHexagon(3, 3, "center")
	actual_grid := builder.Build()
	expected_grid := rawHexagonGrid{
		RawHexagons: map[int]map[int]*rawHexagon{
			2: {
				3: &rawHexagon{Q: 2, R: 3, Value: wrappedValue{HasValue: true, RawValue: "left"}},
			},
			3: {
				3: &rawHexagon{Q: 3, R: 3, Value: wrappedValue{HasValue: true, RawValue: "center"}},
			},
		},
	}
	AssertHexagonGridEquals(t, expected_grid, actual_grid)
}
