package hexagon

func GetRadialHexGrid(radius int) HexagonGrid {
	builder := NewHexagonGridBuilder()
	diameter := 1 + 2*radius
	for q := 0; q < diameter; q += 1 {
		for r := 0; r < diameter; r += 1 {
			if radius <= q+r && q+r <= 3*radius {
				builder.AddHexagon(q, r)
			}
		}
	}
	return builder.Build()
}
