package exratedto

import "time"

// RateInfo 汇率信息
type RateInfo struct {
	BankID      int
	CurrencyID  int
	Rate        float64
	ReleaseTime string
}

// RateInfoList 汇率信息，含最大最小值
type RateInfoList struct {
	Min           float64
	Max           float64
	RateInfoSlice []RateInfo
}

// BankInfo 银行基础表信息
type BankInfo struct {
	BankID     int
	BankName   string
	BankNameEN string
	TimeZone   string
	TableName  string
}

// BankTableColumnMap 汇率表字段字典
var BankTableColumnMap = map[int]string{
	1: "buying_rate",
	2: "cash_buying_rate",
	3: "selling_rate",
	4: "cash_selling_rate",
}

// UserInfo 用户表信息
type UserInfo struct {
	ID         int
	UserID     string
	UserPwd    string
	UserName   string
	Email      string
	CreateTime time.Time
}
