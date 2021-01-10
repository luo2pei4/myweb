package adsbdao

import (
	"fmt"
	"myweb/dao"
	"myweb/db"
	adsbdto "myweb/dto/adsb"
)

var conn *db.Connection

const key = "adsb"

func init() {

	dao.RegistFunc(key, func(c *db.Connection) {
		conn = c
		fmt.Println("Set connection in adsbDao.")
	})
}

// GetCoordWithCount 获取坐标信息
func GetCoordWithCount(startDate, endDate, arrDep string) (coordSlice []adsbdto.CoordWithCount, err error) {

	sql := `
	select firstlng as lng, firstlat as lat , count(1) as count from adsb.flightinfo f 
	where 
	firstlat <> 0 
	and firstlng <> 0 
	%v 
	and firstsignaltime BETWEEN STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s') and STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s')
	group by firstlat , firstlng`

	condition := "and arrdep = '%v'"

	if arrDep == "" {
		sql = fmt.Sprintf(sql, "", startDate, endDate)
	} else {
		condition := fmt.Sprintf(condition, arrDep)
		sql = fmt.Sprintf(sql, condition, startDate, endDate)
	}
	rows, err := conn.Select(sql)

	if err != nil {
		return nil, err
	}

	coordSlice = make([]adsbdto.CoordWithCount, 0)

	for rows.Next() {

		coordInfo := adsbdto.CoordWithCount{}

		if err = rows.Scan(&coordInfo.Coord.Lng, &coordInfo.Coord.Lat, &coordInfo.Count); err != nil {
			return nil, err
		}

		coordSlice = append(coordSlice, coordInfo)
	}

	return
}

// GetCoordsWithIcao 获取坐标信息
func GetCoordsWithIcao(actualDate, callsign string) (icaoCoordMap map[string][]adsbdto.Coordinate, err error) {

	sql := `
	select 
		t1.icao, t1.Longitude, t1.Latitude 
	from 
		adsbinfohis t1, 
		(select icao, firstsignaltime, lastsignaltime from flightinfo where callsign = '%v' and actualdate = '%v') t2
	where 
		t1.icao = t2.icao
		and t1.createtime >= t2.firstsignaltime
		and t1.createtime <= t2.lastsignaltime
		and t1.Longitude <> 0
		and t1.Latitude <> 0
	order by 
		t1.createtime asc`

	sql = fmt.Sprintf(sql, callsign, actualDate)
	fmt.Println(sql)
	rows, err := conn.Select(sql)

	if err != nil {
		return nil, err
	}

	icaoCoordMap = make(map[string][]adsbdto.Coordinate)

	for rows.Next() {

		coord := adsbdto.Coordinate{}
		icao := ""

		if err = rows.Scan(&icao, &coord.Lng, &coord.Lat); err != nil {
			return nil, err
		}

		if icaoCoordMap[icao] == nil {
			coordSlice := make([]adsbdto.Coordinate, 0)
			coordSlice = append(coordSlice, coord)
			icaoCoordMap[icao] = coordSlice
		} else {
			coordSlice := icaoCoordMap[icao]
			coordSlice = append(coordSlice, coord)
			icaoCoordMap[icao] = coordSlice
		}
	}

	return
}
