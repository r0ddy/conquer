package hexagon

import "sort"

var CornerDirections = []Direction{N, NE, SE, S, SW, NW}

type HexCorner struct {
	Hex             Hexagon
	CornerDirection Direction
}

type Corner struct {
	HexCorners []HexCorner
}

type CornerGrid struct {
	Corners []Corner
}

type VisitedMap map[int]map[int]map[Direction]bool

func hasVisited(hex Hexagon, dir Direction, visited VisitedMap) bool {
	q, r := hex.GetCoordinates()
	if _, qExists := visited[q]; qExists {
		if _, rExists := visited[q][r]; rExists {
			if _, dirExists := visited[q][r][dir]; dirExists {
				return true
			}
		}
	}
	return false
}

func visit(hex Hexagon, dir Direction, visited VisitedMap) {
	q, r := hex.GetCoordinates()
	if _, qExists := visited[q]; !qExists {
		visited[q] = make(map[int]map[Direction]bool)
	}
	if _, rExists := visited[q][r]; !rExists {
		visited[q][r] = make(map[Direction]bool)
	}
	visited[q][r][dir] = true
}

func getCorner(hex Hexagon, dir Direction, visited VisitedMap) Corner {
	corner := Corner{
		HexCorners: []HexCorner{
			{Hex: hex, CornerDirection: dir},
		},
	}
	neighbors := hex.GetNeighbors()
	for _, sideCorner := range getNeighboringSideAndCorners(dir) {
		if neighbor, neighborExists := neighbors[sideCorner.Side]; neighborExists {
			visit(neighbor, sideCorner.Corner, visited)
			corner.HexCorners = append(corner.HexCorners,
				HexCorner{Hex: neighbor, CornerDirection: sideCorner.Corner},
			)
		}
	}
	return corner
}

type SideAndCorner struct {
	Side   Direction
	Corner Direction
}

func getNeighboringSideAndCorners(dir Direction) []SideAndCorner {
	switch dir {
	case N:
		return []SideAndCorner{{Side: NW, Corner: SE}, {Side: NE, Corner: SW}}
	case NE:
		return []SideAndCorner{{Side: NE, Corner: S}, {Side: E, Corner: NW}}
	case SE:
		return []SideAndCorner{{Side: E, Corner: SW}, {Side: SE, Corner: N}}
	case S:
		return []SideAndCorner{{Side: SE, Corner: NW}, {Side: SW, Corner: NE}}
	case SW:
		return []SideAndCorner{{Side: SW, Corner: N}, {Side: W, Corner: SE}}
	case NW:
		return []SideAndCorner{{Side: W, Corner: NE}, {Side: NW, Corner: S}}
	default:
		return []SideAndCorner{}
	}
}

func GetCornerGrid(hexGrid HexagonGrid) CornerGrid {
	grid := CornerGrid{Corners: make([]Corner, 0)}
	visited := make(map[int]map[int]map[Direction]bool)
	for _, hex := range hexGrid.GetHexagons() {
		for _, dir := range CornerDirections {
			if !hasVisited(hex, dir, visited) {
				visit(hex, dir, visited)
				corner := getCorner(hex, dir, visited)
				sortHexCorners(corner.HexCorners)
				grid.Corners = append(grid.Corners, corner)
			}
		}
	}
	sortCorners(grid.Corners)
	return grid
}

func isNotEqual(i, j HexCorner) bool {
	i_q, i_r := i.Hex.GetCoordinates()
	j_q, j_r := j.Hex.GetCoordinates()
	i_dir := i.CornerDirection
	j_dir := j.CornerDirection
	if i_q != j_q {
		return true
	}
	if i_r != j_r {
		return true
	}
	if i_dir != j_dir {
		return true
	}
	return false
}

func isLessThan(i, j HexCorner) bool {
	i_q, i_r := i.Hex.GetCoordinates()
	j_q, j_r := j.Hex.GetCoordinates()
	i_dir := i.CornerDirection
	j_dir := j.CornerDirection
	if i_q != j_q {
		return i_q < j_q
	}
	if i_r != j_r {
		return i_r < j_r
	}
	return i_dir < j_dir
}

func sortHexCorners(hexCorners []HexCorner) {
	sort.Slice(hexCorners, func(i, j int) bool {
		return isLessThan(hexCorners[i], hexCorners[j])
	})
}

func sortCorners(corners []Corner) {
	sort.Slice(corners, func(i, j int) bool {
		i_hexCorners := corners[i].HexCorners
		j_hexCorners := corners[j].HexCorners
		if len(i_hexCorners) != len(j_hexCorners) {
			return len(i_hexCorners) < len(j_hexCorners)
		}
		for index, _ := range i_hexCorners {
			if isNotEqual(i_hexCorners[index], j_hexCorners[index]) {
				return isLessThan(i_hexCorners[index], j_hexCorners[index])
			}
		}
		return false
	})
}

func (grid CornerGrid) removeRef() CornerGrid {
	refLessCorners := CornerGrid{Corners: make([]Corner, 0)}
	for _, corner := range grid.Corners {
		refLessCorner := Corner{HexCorners: make([]HexCorner, 0)}
		for _, hexCorner := range corner.HexCorners {
			refLessCorner.HexCorners = append(
				refLessCorner.HexCorners,
				HexCorner{
					Hex:             hexCorner.Hex.removeRef(),
					CornerDirection: hexCorner.CornerDirection,
				},
			)
		}
		refLessCorners.Corners = append(refLessCorners.Corners, refLessCorner)
	}
	return refLessCorners
}
