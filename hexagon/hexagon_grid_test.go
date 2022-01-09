package hexagon

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func AssertHexagonEquals(t *testing.T, expected, actual Hexagon) bool {
	reflessExpected := expected.removeRef()
	reflessActual := actual.removeRef()
	return assert.True(t,
		cmp.Equal(reflessExpected, reflessActual),
		cmp.Diff(reflessExpected, reflessActual),
	)
}

func removeRefs(hexes []Hexagon) []Hexagon {
	reflessHexes := make([]Hexagon, 0)
	for _, hex := range hexes {
		reflessHexes = append(reflessHexes, hex.removeRef())
	}
	return reflessHexes
}

func AssertHexagonsEquals(t *testing.T, expected, actual []Hexagon) bool {
	reflessExpected := removeRefs(expected)
	reflessActual := removeRefs(actual)
	return assert.True(t,
		cmp.Equal(reflessExpected, reflessActual),
		cmp.Diff(reflessExpected, reflessActual),
	)
}

func Test_HexagonGrid_GetHexagon(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(2, 3, "val")
	grid := builder.Build()
	actual_hex, err := grid.GetHexagon(2, 3)
	assert.NoError(t, err)
	expected_hex := rawHexagon{
		Q: 2,
		R: 3,
		Value: wrappedValue{
			HasValue: true,
			RawValue: "val",
		},
	}
	AssertHexagonEquals(t, expected_hex, actual_hex)

	_, err = grid.GetHexagon(100, 100)
	assert.ErrorIs(t, err, HexagonNotFoundError)
}

func Test_HexagonGrid_GetHexagons(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(2, 3, "val1")
	builder.AddHexagon(100, -1)
	builder.AddHexagon(3, 4, "val2")
	grid := builder.Build()
	actual_hexes := grid.GetHexagons()
	expected_hexes := []Hexagon{
		rawHexagon{Q: 2, R: 3, Value: wrappedValue{HasValue: true, RawValue: "val1"}},
		rawHexagon{Q: 3, R: 4, Value: wrappedValue{HasValue: true, RawValue: "val2"}},
		rawHexagon{Q: 100, R: -1},
	}
	AssertHexagonsEquals(t, expected_hexes, actual_hexes)
}
