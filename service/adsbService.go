package service

import (
	adsbdao "myweb/dao/adsb"
	adsbdto "myweb/dto/adsb"
)

// ReadCoordInfo 读取坐标信息
func ReadCoordInfo(startDate, endDate, arrDep string) (coordList []adsbdto.Coordinate, err error) {
	return adsbdao.GetCoordInfo(startDate, endDate, arrDep)
}
