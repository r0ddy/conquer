package hexagon

type Direction string

const (
	N  Direction = "N"
	NE           = "NE"
	E            = "E"
	SE           = "SE"
	S            = "S"
	SW           = "SW"
	W            = "W"
	NW           = "NW"
)

var SideDirections = []Direction{NE, E, SE, SW, W, NW}

var CornerDirections = []Direction{N, NE, SE, S, SW, NW}

type Angle string

const (
	Down     Angle = "DOWN"
	Vertical       = "VERTICAL"
	Up             = "UP"
)

var Angles = []Angle{Down, Vertical, Up}

func sideToCornersAndAngle(dir Direction) (Direction, Direction, Angle, error) {
	switch dir {
	case NE:
		return N, NE, Down, nil
	case E:
		return NE, SE, Vertical, nil
	case SE:
		return SE, S, Up, nil
	case SW:
		return S, SW, Down, nil
	case W:
		return SW, NW, Vertical, nil
	case NW:
		return NW, N, Up, nil
	default:
		return N, N, Down, NotAValidSideDirection
	}
}

func directionToQR(dir Direction) (q, r int, err error) {
	switch dir {
	case NE:
		return 1, -1, nil
	case E:
		return 1, 0, nil
	case SE:
		return 0, 1, nil
	case SW:
		return -1, 1, nil
	case W:
		return -1, 0, nil
	case NW:
		return 0, -1, nil
	default:
		return 0, 0, DirectionNotFoundError
	}
}

type Neighbors map[Direction]Hexagon

type Hexagon interface {
	GetCoordinates() (q, r int)
	GetNeighbors() Neighbors
	GetValue() (interface{}, error)
	removeRef() Hexagon
}

type wrappedValue struct {
	HasValue bool
	RawValue interface{}
}

type rawHexagon struct {
	Q       int
	R       int
	Value   wrappedValue
	GridRef *rawHexagonGrid
}

func (hex rawHexagon) GetCoordinates() (q, r int) {
	return hex.Q, hex.R
}

func (hex rawHexagon) GetNeighbors() Neighbors {
	nei := make(map[Direction]Hexagon)
	for _, dir := range SideDirections {
		q_diff, r_diff, err := directionToQR(dir)
		if err != nil {
			continue
		}
		nei_hex, err := hex.GridRef.GetHexagon(hex.Q+q_diff, hex.R+r_diff)
		if err != nil {
			continue
		}
		nei[dir] = nei_hex
	}
	return nei
}

func (hex rawHexagon) GetValue() (interface{}, error) {
	if !hex.Value.HasValue {
		return nil, HexagonHasNoValueError
	}
	return hex.Value.RawValue, nil
}

func (hex rawHexagon) removeRef() Hexagon {
	hex.GridRef = nil
	return hex
}
