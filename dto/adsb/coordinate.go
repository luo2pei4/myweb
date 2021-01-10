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

// SimpleAdsbInfo 简单的ADSB信息
type SimpleAdsbInfo struct {
	Icao       string
	Alt        string
	Lat        string
	Lng        string
	Spd        string
	Createtime string
}
