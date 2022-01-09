package hexagon

import "sort"

type HexagonGrid interface {
	GetHexagon(q, r int) (Hexagon, error)
	GetHexagons() []Hexagon
	removeRefs() HexagonGrid
}

type rawHexagonGrid struct {
	RawHexagons map[int]map[int]*rawHexagon
}

func (grid rawHexagonGrid) GetHexagon(q, r int) (Hexagon, error) {
	hexExists := false
	if _, qExists := grid.RawHexagons[q]; qExists {
		if _, rExists := grid.RawHexagons[q][r]; rExists {
			hexExists = true
		}
	}
	if !hexExists {
		return nil, HexagonNotFoundError
	}
	return grid.RawHexagons[q][r], nil
}

func sortHexagons(hexes []Hexagon) {
	sort.Slice(hexes, func(i, j int) bool {
		i_q, i_r := hexes[i].GetCoordinates()
		j_q, j_r := hexes[j].GetCoordinates()
		if i_q != i_r {
			return i_q < j_q
		}
		return i_r < j_r
	})
}

func (grid rawHexagonGrid) GetHexagons() []Hexagon {
	hexes := make([]Hexagon, 0)
	for _, rHexes := range grid.RawHexagons {
		for _, hex := range rHexes {
			hexes = append(hexes, hex)
		}
	}
	sortHexagons(hexes)
	return hexes
}

func (grid rawHexagonGrid) removeRefs() HexagonGrid {
	for q, hexes := range grid.RawHexagons {
		for r, hex := range hexes {
			refless_hex := hex.removeRef().(rawHexagon)
			grid.RawHexagons[q][r] = &refless_hex
		}
	}
	return grid
}
