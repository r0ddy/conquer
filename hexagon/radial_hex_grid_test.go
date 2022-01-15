package hexagon

import "testing"

func Test_RadialHexGrid_GetRadialHexGrid(t *testing.T) {
	zeroRadiusGrid := GetRadialHexGrid(0)
	expected := rawHexagonGrid{
		RawHexagons: map[int]map[int]*rawHexagon{
			0: {
				0: &rawHexagon{Q: 0, R: 0},
			},
		},
	}
	AssertHexagonGridEquals(t, zeroRadiusGrid, expected)

	oneRadiusGrid := GetRadialHexGrid(1)
	expected = rawHexagonGrid{
		RawHexagons: map[int]map[int]*rawHexagon{
			0: {
				1: &rawHexagon{Q: 0, R: 1},
				2: &rawHexagon{Q: 0, R: 2},
			},
			1: {
				0: &rawHexagon{Q: 1, R: 0},
				1: &rawHexagon{Q: 1, R: 1},
				2: &rawHexagon{Q: 1, R: 2},
			},
			2: {
				0: &rawHexagon{Q: 2, R: 0},
				1: &rawHexagon{Q: 2, R: 1},
			},
		},
	}
	AssertHexagonGridEquals(t, oneRadiusGrid, expected)

	threeRadiusGrid := GetRadialHexGrid(3)
	expected = rawHexagonGrid{
		RawHexagons: map[int]map[int]*rawHexagon{
			0: {
				3: &rawHexagon{Q: 0, R: 3},
				4: &rawHexagon{Q: 0, R: 4},
				5: &rawHexagon{Q: 0, R: 5},
				6: &rawHexagon{Q: 0, R: 6},
			},
			1: {
				2: &rawHexagon{Q: 1, R: 2},
				3: &rawHexagon{Q: 1, R: 3},
				4: &rawHexagon{Q: 1, R: 4},
				5: &rawHexagon{Q: 1, R: 5},
				6: &rawHexagon{Q: 1, R: 6},
			},
			2: {
				1: &rawHexagon{Q: 2, R: 1},
				2: &rawHexagon{Q: 2, R: 2},
				3: &rawHexagon{Q: 2, R: 3},
				4: &rawHexagon{Q: 2, R: 4},
				5: &rawHexagon{Q: 2, R: 5},
				6: &rawHexagon{Q: 2, R: 6},
			},
			3: {
				0: &rawHexagon{Q: 3, R: 0},
				1: &rawHexagon{Q: 3, R: 1},
				2: &rawHexagon{Q: 3, R: 2},
				3: &rawHexagon{Q: 3, R: 3},
				4: &rawHexagon{Q: 3, R: 4},
				5: &rawHexagon{Q: 3, R: 5},
				6: &rawHexagon{Q: 3, R: 6},
			},
			4: {
				0: &rawHexagon{Q: 4, R: 0},
				1: &rawHexagon{Q: 4, R: 1},
				2: &rawHexagon{Q: 4, R: 2},
				3: &rawHexagon{Q: 4, R: 3},
				4: &rawHexagon{Q: 4, R: 4},
				5: &rawHexagon{Q: 4, R: 5},
			},

			5: {
				0: &rawHexagon{Q: 5, R: 0},
				1: &rawHexagon{Q: 5, R: 1},
				2: &rawHexagon{Q: 5, R: 2},
				3: &rawHexagon{Q: 5, R: 3},
				4: &rawHexagon{Q: 5, R: 4},
			},
			6: {
				0: &rawHexagon{Q: 6, R: 0},
				1: &rawHexagon{Q: 6, R: 1},
				2: &rawHexagon{Q: 6, R: 2},
				3: &rawHexagon{Q: 6, R: 3},
			},
		},
	}
	AssertHexagonGridEquals(t, threeRadiusGrid, expected)
}
