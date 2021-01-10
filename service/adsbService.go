package service

import (
	adsbdao "myweb/dao/adsb"
	adsbdto "myweb/dto/adsb"
)

// GetCoordWithCount 读取坐标信息
func GetCoordWithCount(startDate, endDate, arrDep string) (coordList []adsbdto.CoordWithCount, err error) {
	return adsbdao.GetCoordWithCount(startDate, endDate, arrDep)
}

// GetAircraftCoordMap 获取飞机的坐标Map
func GetAircraftCoordMap(actualDate, callsign string) (simpleAdsbMap map[string][]adsbdto.SimpleAdsbInfo, err error) {

	simpleAdsbMap, err = adsbdao.GetCoordsWithIcao(actualDate, callsign)

	if err != nil {
		return nil, err
	}

	return
}
