package adsbdto

// Coordinate 坐标数据
type Coordinate struct {
	Lng string
	Lat string
}

// CoordWithCount 带计数的坐标数据
type CoordWithCount struct {
	Coord Coordinate
	Count int
}
