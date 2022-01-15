package hexagon

import "errors"

var HexagonNotFoundError = errors.New("hexagon not found")

var DirectionNotFoundError = errors.New("direction not found")

var HexagonHasNoValueError = errors.New("hexagon has no value")

var NotAValidSideDirection = errors.New("not a valid side direction")
