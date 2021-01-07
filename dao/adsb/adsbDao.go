package adsbdao

import (
	"fmt"
	"myweb/db"
	adsbdto "myweb/dto/adsb"
)

var conn *db.Connection

func init() {

	c, err := db.NewConnection("adsb", "mysql", "dbo:caecaodb@tcp(192.168.3.169:3306)/adsb?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	conn = c
}

// GetCoordInfo 获取坐标信息
func GetCoordInfo(startDate, endDate, arrDep string) (coordList []adsbdto.Coordinate, err error) {

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

	var coordSlice = make([]adsbdto.Coordinate, 0)

	for rows.Next() {

		coordInfo := adsbdto.Coordinate{}

		if err = rows.Scan(&coordInfo.Lng, &coordInfo.Lat, &coordInfo.Count); err != nil {
			return nil, err
		}

		coordSlice = append(coordSlice, coordInfo)
	}

	return coordSlice, nil
}
