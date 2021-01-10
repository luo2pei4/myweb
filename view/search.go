package view

import (
	"encoding/json"
	"fmt"
	authenticate "myweb/auth"
	"myweb/dto"
	"myweb/service"
	"myweb/utils"
	"net/http"
	"time"
)

// 查询汇率
func searchExRate(response http.ResponseWriter, request *http.Request) {

	method := request.Method

	// 如果是直接从网页输入地址, 直接跳转到登录界面.
	if method == "GET" {
		renderTempalate(response, "static/html/index.html", LoginPage{})
		return
	}

	// 验证Token并返回用户信息
	claims, returnCode := authenticate.AuthToken(request)

	if returnCode != http.StatusOK {
		response.WriteHeader(returnCode)
		return
	}

	request.ParseForm()
	bankID := utils.StringToInt(request.Form.Get("bank"))
	bankTableColumnID := utils.StringToInt(request.Form.Get("price"))
	startDate := request.Form.Get("startDate")
	endDate := request.Form.Get("endDate")
	currencyID := utils.StringToInt(request.Form.Get("currency"))

	pageInfo := dto.PageInfo{}

	if startDate == "" {

		pageInfo.ErrMsg = "请输入开始和结束日期."

	} else {

		// 如果没有输入结束日期, 默认以当天为结束日期
		if endDate == "" {
			now := time.Now().Format(time.RFC3339)
			endDate = now[:10]
		}

		rilist, err := service.GetRateInfoList(bankID, bankTableColumnID, currencyID, startDate+" 00:00:00", endDate+" 23:59:59")

		if err != nil {
			fmt.Println(err.Error())
			pageInfo.ErrMsg = "系统异常, 请稍后再试."
		}

		if len(rilist.RateInfoSlice) == 0 {
			pageInfo.ErrMsg = "没有数据, 请重新指定查询条件."
		}

		pageInfo.Data = rilist
	}

	obj, err := json.Marshal(pageInfo)

	if err != nil {
		fmt.Println(err.Error())
		pageInfo.ErrMsg = "系统异常, 请稍后再试."
	}

	authenticate.RedistributeToken(response, claims.UserID)
	response.Write(obj)
}

func adsbCoverage(response http.ResponseWriter, request *http.Request) {

	// 如果是直接从网页输入地址, 直接跳转到登录界面.
	if request.Method == "GET" {
		renderTempalate(response, "static/html/index.html", LoginPage{})
		return
	}

	// 验证Token并返回用户信息
	claims, returnCode := authenticate.AuthToken(request)

	if returnCode != http.StatusOK {
		response.WriteHeader(returnCode)
		return
	}

	request.ParseForm()
	startDate := request.Form.Get("startDate")
	endDate := request.Form.Get("endDate")
	arrDep := request.Form.Get("arrDep")
	pageInfo := dto.PageInfo{}

	if startDate == "" {

		pageInfo.ErrMsg = "请输入开始和结束日期."

	} else {

		// 如果没有输入结束日期, 默认以当天为结束日期
		if endDate == "" {
			now := time.Now().Format(time.RFC3339)
			endDate = now[:10]
		}

		coordList, err := service.GetCoordWithCount(startDate+" 00:00:00", endDate+" 23:59:59", arrDep)

		if err != nil {
			fmt.Println(err.Error())
		}

		if len(coordList) == 0 {
			pageInfo.ErrMsg = "没有数据, 请重新指定查询条件."
		} else {
			pageInfo.Data = coordList
		}
	}

	obj, err := json.Marshal(pageInfo)

	if err != nil {
		fmt.Println(err.Error())
		pageInfo.ErrMsg = "系统异常, 请稍后再试."
	}

	authenticate.RedistributeToken(response, claims.UserID)
	response.Write(obj)
}

func adsbTrail(response http.ResponseWriter, request *http.Request) {

	// 如果是直接从网页输入地址, 直接跳转到登录界面.
	if request.Method == "GET" {
		renderTempalate(response, "static/html/index.html", LoginPage{})
		return
	}

	// 验证Token并返回用户信息
	claims, returnCode := authenticate.AuthToken(request)

	if returnCode != http.StatusOK {
		response.WriteHeader(returnCode)
		return
	}

	request.ParseForm()
	actualDate := request.Form.Get("actualDate")
	callsign := request.Form.Get("callsign")
	pageInfo := dto.PageInfo{}

	if actualDate == "" {

		pageInfo.ErrMsg = "请输入日期."

	} else {

		acMap, err := service.GetAircraftCoordMap(actualDate, callsign)

		if err != nil {
			pageInfo.ErrMsg = "系统异常, 请稍后再试."
		}

		if len(acMap) == 0 {
			pageInfo.ErrMsg = "没有数据, 请重新指定查询条件."
		} else {
			pageInfo.Data = acMap
		}
	}

	obj, err := json.Marshal(pageInfo)

	if err != nil {
		fmt.Println(err.Error())
		pageInfo.ErrMsg = "系统异常, 请稍后再试."
	}

	authenticate.RedistributeToken(response, claims.UserID)
	response.Write(obj)
}
