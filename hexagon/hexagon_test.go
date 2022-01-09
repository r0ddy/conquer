package hexagon

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func removeRefsFromNeighbors(neighbors Neighbors) Neighbors {
	reflessNeighbors := make(Neighbors)
	for dir, hex := range neighbors {
		reflessNeighbors[dir] = hex.removeRef()
	}
	return reflessNeighbors
}

func AssertNeighborsEquals(t *testing.T, expected, actual Neighbors) bool {
	reflessExpected := removeRefsFromNeighbors(expected)
	reflessActual := removeRefsFromNeighbors(actual)
	return assert.True(t,
		cmp.Equal(reflessExpected, reflessActual),
		cmp.Diff(reflessExpected, reflessActual),
	)
}

func Test_Hexagon_GetCoordinates(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(1, 5)
	grid := builder.Build()
	hex, err := grid.GetHexagon(1, 5)
	assert.NoError(t, err)
	actual_q, actual_r := hex.GetCoordinates()
	assert.Equal(t, 1, actual_q)
	assert.Equal(t, 5, actual_r)
}

func Test_Hexagon_GetValue(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(1, 5, "value")
	builder.AddHexagon(1, 2)
	grid := builder.Build()

	hex, err := grid.GetHexagon(1, 5)
	assert.NoError(t, err)
	val, err := hex.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "value", val)

	hex, err = grid.GetHexagon(1, 2)
	assert.NoError(t, err)
	_, err = hex.GetValue()
	assert.ErrorIs(t, err, HexagonHasNoValueError)
}

func Test_Hexagon_GetNeighbors(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(3, 3, "center")
	builder.AddHexagon(4, 3, "east")
	builder.AddHexagon(3, 4, "southeast")
	builder.AddHexagon(2, 4, "southwest")
	builder.AddHexagon(2, 3, "west")
	builder.AddHexagon(3, 2, "northwest")
	grid := builder.Build()
	hex, err := grid.GetHexagon(3, 3)
	assert.NoError(t, err)
	actual_neighbors := hex.GetNeighbors()
	expected_neighbors := Neighbors{
		E:  rawHexagon{Q: 4, R: 3, Value: wrappedValue{HasValue: true, RawValue: "east"}},
		SE: rawHexagon{Q: 3, R: 4, Value: wrappedValue{HasValue: true, RawValue: "southeast"}},
		SW: rawHexagon{Q: 2, R: 4, Value: wrappedValue{HasValue: true, RawValue: "southwest"}},
		W:  rawHexagon{Q: 2, R: 3, Value: wrappedValue{HasValue: true, RawValue: "west"}},
		NW: rawHexagon{Q: 3, R: 2, Value: wrappedValue{HasValue: true, RawValue: "northwest"}},
	}
	AssertNeighborsEquals(t, expected_neighbors, actual_neighbors)
}
