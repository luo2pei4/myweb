package exratedao

import (
	"fmt"
	"myweb/dao"
	"myweb/db"
	exratedto "myweb/dto/exrate"
)

var conn *db.Connection

// BankInfoMap 银行基础信息Map
var bankInfoMap map[int]*exratedto.BankInfo

const key = "exrate"

func init() {

	dao.RegistFunc(key, func(c *db.Connection) {
		conn = c
		fmt.Println("Set connection in exrateDao.")
		bankInfoMap = initBankInfoMap()
		fmt.Println("initBankInfoMap successful.")
	})
}

// initBankInfoMap 初始化银行表基础信息
func initBankInfoMap() map[int]*exratedto.BankInfo {

	rows, err := conn.Select(`select id, bank_name, bank_name_en, time_zone, table_name from banks`)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	bankInfoMap := make(map[int]*exratedto.BankInfo)

	for rows.Next() {

		bankInfo := exratedto.BankInfo{}
		err = rows.Scan(&bankInfo.BankID, &bankInfo.BankName, &bankInfo.BankNameEN, &bankInfo.TimeZone, &bankInfo.TableName)

		if err != nil {
			continue
		}

		bankInfoMap[bankInfo.BankID] = &bankInfo
	}

	return bankInfoMap
}

// GetRateInfoSlice 获取汇率信息切片
func GetRateInfoSlice(bankID, bankTableColumnID, currency int, startDate, endDate string) (riSlice []exratedto.RateInfo, err error) {

	bankName := bankInfoMap[bankID].TableName
	rateName := exratedto.BankTableColumnMap[bankTableColumnID]

	sql := `
		select 
			%v, date_format(date_add(release_time, INTERVAL 8 hour), '%%Y-%%m-%%d %%H:%%i') as releaseTime 
		from 
			%v
		where 
			date_add(release_time, INTERVAL 8 hour) between STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s') and STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s')
			and currencyid = %v
		order by 
			release_time`

	sql = fmt.Sprintf(sql, rateName, bankName, startDate, endDate, currency)
	rows, err := conn.Select(sql)

	if err != nil {
		return nil, err
	}

	riSlice = make([]exratedto.RateInfo, 0)

	for rows.Next() {

		rateInfo := exratedto.RateInfo{}

		if err = rows.Scan(&rateInfo.Rate, &rateInfo.ReleaseTime); err != nil {
			fmt.Println(err.Error())
		}

		riSlice = append(riSlice, rateInfo)
	}

	return
}

// GetMinRate 获取最小汇率值
func GetMinRate(bankID, bankTableColumnID, currency int, startDate, endDate string) (minrate float64, err error) {

	bankName := bankInfoMap[bankID].TableName
	rateName := exratedto.BankTableColumnMap[bankTableColumnID]

	sql := `
		select 
			min(%v) as minrate 
		from 
			%v 
		where 
			date_add(release_time, INTERVAL 8 hour) between STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s') and STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s')
			and currencyid = %v`

	sql = fmt.Sprintf(sql, rateName, bankName, startDate, endDate, currency)
	rows, err := conn.Select(sql)

	if err != nil {
		return 0, err
	}

	rows.Next()

	if err = rows.Scan(&minrate); err != nil {
		return 0, err
	}

	return
}

// GetMaxRate 获取最大汇率值
func GetMaxRate(bankID, bankTableColumnID, currency int, startDate, endDate string) (maxrate float64, err error) {

	bankName := bankInfoMap[bankID].TableName
	rateName := exratedto.BankTableColumnMap[bankTableColumnID]

	sql := `
		select 
			max(%v) as maxrate 
		from 
			%v 
		where 
			date_add(release_time, INTERVAL 8 hour) between STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s') and STR_TO_DATE('%v', '%%Y-%%m-%%d %%H:%%i:%%s')
			and currencyid = %v`

	sql = fmt.Sprintf(sql, rateName, bankName, startDate, endDate, currency)
	rows, err := conn.Select(sql)

	if err != nil {
		return 0, err
	}

	rows.Next()

	if err = rows.Scan(&maxrate); err != nil {
		return 0, err
	}

	return
}
