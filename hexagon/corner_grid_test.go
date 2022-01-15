package hexagon

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func AssertCornerGridEquals(t *testing.T, expected, actual CornerGrid) bool {
	reflessExpected := expected.removeRef()
	reflessActual := actual.removeRef()
	return assert.True(t,
		cmp.Equal(reflessExpected, reflessActual),
		cmp.Diff(reflessExpected, reflessActual),
	)
}

func Test_CornerGrid_GetCornerGrid(t *testing.T) {
	builder := NewHexagonGridBuilder()
	builder.AddHexagon(2, 3)
	builder.AddHexagon(3, 3)
	hexGrid := builder.Build()
	actual_corner_grid := GetCornerGrid(hexGrid)
	firstHex, _ := hexGrid.GetHexagon(2, 3)
	secondHex, _ := hexGrid.GetHexagon(3, 3)
	expected_corners := []Corner{
		// the four corners on the first hex that
		// are not on the edge shared with the second hex
		{HexCorners: []HexCorner{{Hex: firstHex, CornerDirection: N}}},
		{HexCorners: []HexCorner{{Hex: firstHex, CornerDirection: NW}}},
		{HexCorners: []HexCorner{{Hex: firstHex, CornerDirection: S}}},
		{HexCorners: []HexCorner{{Hex: firstHex, CornerDirection: SW}}},
		// the four corners on the second hex that
		// are not on the edge shared with the first hex
		{HexCorners: []HexCorner{{Hex: secondHex, CornerDirection: N}}},
		{HexCorners: []HexCorner{{Hex: secondHex, CornerDirection: NE}}},
		{HexCorners: []HexCorner{{Hex: secondHex, CornerDirection: S}}},
		{HexCorners: []HexCorner{{Hex: secondHex, CornerDirection: SE}}},
		// the two corners on the shared edge
		{HexCorners: []HexCorner{{Hex: firstHex, CornerDirection: NE}, {Hex: secondHex, CornerDirection: NW}}},
		{HexCorners: []HexCorner{{Hex: firstHex, CornerDirection: SE}, {Hex: secondHex, CornerDirection: SW}}},
	}
	expected_corner_grid := CornerGrid{
		Corners: expected_corners,
		HexCornerToCornerIndex: map[int]map[int]map[Direction]int{
			2: {
				3: {
					"N":  0,
					"NE": 8,
					"NW": 1,
					"S":  2,
					"SE": 9,
					"SW": 3,
				},
			},
			3: {
				3: {
					"N":  4,
					"NE": 5,
					"NW": 8,
					"S":  6,
					"SE": 7,
					"SW": 9,
				},
			},
		},
		Sides: []Side{
			{CornerIndices: []int{0, 1}, Angle: Up},
			{CornerIndices: []int{0, 8}, Angle: Down},
			{CornerIndices: []int{1, 3}, Angle: Vertical},
			{CornerIndices: []int{2, 3}, Angle: Down},
			{CornerIndices: []int{2, 9}, Angle: Up},
			{CornerIndices: []int{4, 5}, Angle: Down},
			{CornerIndices: []int{4, 8}, Angle: Up},
			{CornerIndices: []int{5, 7}, Angle: Vertical},
			{CornerIndices: []int{6, 7}, Angle: Up},
			{CornerIndices: []int{6, 9}, Angle: Down},
			{CornerIndices: []int{8, 9}, Angle: Vertical},
		},
	}
	AssertCornerGridEquals(t, expected_corner_grid, actual_corner_grid)
}
