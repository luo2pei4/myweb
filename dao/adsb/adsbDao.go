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

// GetCoordInfo 获取坐标信息
func GetCoordInfo(startDate, endDate, arrDep string) (coordSlice []adsbdto.Coordinate, err error) {

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

	coordSlice = make([]adsbdto.Coordinate, 0)

	for rows.Next() {

		coordInfo := adsbdto.Coordinate{}

		if err = rows.Scan(&coordInfo.Lng, &coordInfo.Lat, &coordInfo.Count); err != nil {
			return nil, err
		}

		coordSlice = append(coordSlice, coordInfo)
	}

	return
}
