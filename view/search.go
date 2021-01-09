package view

import (
	"encoding/json"
	"fmt"
	authenticate "myweb/auth"
	exratedto "myweb/dto/exrate"
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

	var rateInfoList *exratedto.RateInfoList

	request.ParseForm()
	bankID := utils.StringToInt(request.Form.Get("bank"))
	bankTableColumnID := utils.StringToInt(request.Form.Get("price"))
	startDate := request.Form.Get("startDate")
	endDate := request.Form.Get("endDate")
	currencyID := utils.StringToInt(request.Form.Get("currency"))

	if startDate == "" {

		rateInfoList = &exratedto.RateInfoList{
			ErrMsg: "请输入开始和结束日期.",
		}

	} else {

		// 如果没有输入结束日期, 默认以当天为结束日期
		if endDate == "" {
			now := time.Now().Format(time.RFC3339)
			endDate = now[:10]
		}

		rilist, err := service.GetRateInfoList(bankID, bankTableColumnID, currencyID, startDate+" 00:00:00", endDate+" 23:59:59")

		if err != nil {
			fmt.Println(err.Error())
			rilist.ErrMsg = "系统异常, 请稍后再试."
		}

		if len(rilist.RateInfoSlice) == 0 {
			rilist.ErrMsg = "没有数据, 请重新查询."
		}

		rateInfoList = rilist
	}

	obj, err := json.Marshal(rateInfoList)

	if err != nil {
		fmt.Println(err.Error())
		rateInfoList.ErrMsg = "系统异常, 请稍后再试."
	}

	authenticate.RedistributeToken(response, claims.UserID)
	response.Write(obj)
}

func adsbCoverRange(response http.ResponseWriter, request *http.Request) {

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

	coordList, err := service.ReadCoordInfo(startDate+" 00:00:00", endDate+" 23:59:59", arrDep)

	if err != nil {
		fmt.Println(err.Error())
	}

	if len(coordList) == 0 {
		fmt.Println("No data...")
	}

	obj, err := json.Marshal(coordList)

	if err != nil {
		fmt.Println(err.Error())
	}

	authenticate.RedistributeToken(response, claims.UserID)
	response.Write(obj)
}
