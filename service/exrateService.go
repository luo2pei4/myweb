package service

import (
	exratedao "myweb/dao/exrate"
	exratedto "myweb/dto/exrate"
)

// GetRateInfoList 获取汇率信息列表
func GetRateInfoList(bankID, bankTableColumnID, currency int, startDate, endDate string) (rateInfoList *exratedto.RateInfoList, err error) {

	rateInfoSlice, err := exratedao.GetRateInfoSlice(bankID, bankTableColumnID, currency, startDate, endDate)

	if err != nil {
		return nil, err
	}

	if len(rateInfoSlice) != 0 {

		minrate, err := exratedao.GetMinRate(bankID, bankTableColumnID, currency, startDate, endDate)

		if err != nil {
			return nil, err
		}

		maxrate, err := exratedao.GetMaxRate(bankID, bankTableColumnID, currency, startDate, endDate)

		if err != nil {
			return nil, err
		}

		rateInfoList = &exratedto.RateInfoList{}
		rateInfoList.RateInfoSlice = rateInfoSlice
		rateInfoList.Min = minrate
		rateInfoList.Max = maxrate
	}

	return
}
