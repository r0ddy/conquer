package hexagon

type HexagonGridBuilder interface {
	AddHexagon(q int, r int, value ...interface{})
	Build() HexagonGrid
}

type rawHexagonGridBuilder struct {
	Grid *rawHexagonGrid
}

func (builder *rawHexagonGridBuilder) AddHexagon(q int, r int, value ...interface{}) {
	if _, qExists := builder.Grid.RawHexagons[q]; !qExists {
		builder.Grid.RawHexagons[q] = make(map[int]*rawHexagon)
	}
	wv := wrappedValue{}
	if len(value) == 1 {
		wv.HasValue = true
		wv.RawValue = value[0]
	}
	builder.Grid.RawHexagons[q][r] = &rawHexagon{Q: q, R: r, GridRef: builder.Grid, Value: wv}
}

func (builder *rawHexagonGridBuilder) Build() HexagonGrid {
	return builder.Grid
}

func NewHexagonGridBuilder() HexagonGridBuilder {
	return &rawHexagonGridBuilder{
		Grid: &rawHexagonGrid{
			RawHexagons: make(map[int]map[int]*rawHexagon),
		},
	}
}
