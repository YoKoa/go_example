package common_test

import (
	"fmt"
	common "go_example"
	cdshttp "go_example/http"
	"go_example/profile"
	"go_example/regions"
	"testing"
)

const (
	Version = "2019-08-08"
)

type (
	SearchHAProxyRequest struct {
		*cdshttp.BaseRequest
		StartTime *string `json:"StartTime" name:"StartTime"`
		EndTime   *string `json:"EndTime" name:"EndTime"`
	}

	SearchHAProxyResponse struct {
		*cdshttp.BaseResponse
		Message string `json:"Message,omitempty" name:"Message"`
		Code    string `json:"Code,omitempty" name:"Code"`
		Data    []struct {
			CloneServices   interface{} `json:"CloneServices"`
			CPU             int         `json:"Cpu"`
			CreatedTime     string      `json:"CreatedTime"`
			DisplayName     string      `json:"DisplayName"`
			IP              string      `json:"IP"`
			InstanceName    string      `json:"InstanceName"`
			InstanceUUID    string      `json:"InstanceUuid"`
			LinkType        string      `json:"LinkType"`
			LinkTypeStr     string      `json:"LinkTypeStr"`
			MasterInfo      string      `json:"MasterInfo"`
			Port            int         `json:"Port"`
			RAM             int         `json:"Ram"`
			RegionID        string      `json:"RegionId"`
			RelationService interface{} `json:"RelationService"`
			ResourceID      string      `json:"ResourceId"`
			RoGroups        interface{} `json:"RoGroups"`
			Status          string      `json:"Status"`
			StatusStr       string      `json:"StatusStr"`
			SubProductName  string      `json:"SubProductName"`
			VdcID           string      `json:"VdcId"`
			VdcName         string      `json:"VdcName"`
			Version         string      `json:"Version"`
			Vips            []struct {
				IP   string `json:"IP"`
				Type string `json:"Type"`
			} `json:"Vips"`
		} `json:"Data,omitempty" name:"Code"`
	}
)

func TestSearchHAProxyByTimeRange(t *testing.T) {
	tests := []struct {
		AK        string
		SK        string
		Method    string
		Service   string
		Action    string
		StartTime string
		EndTime   string
	}{
		{
			AK:        "",
			SK:        "",
			Method:    "GET",
			Service:   "lb",
			Action:    "DescribeLoadBalancers",
			StartTime: "2020-02-21 17:30:33",
			EndTime:   "2021-03-01 19:00:44",
		},
	}

	for _, testCase := range tests {
		credential := common.NewCredential(testCase.AK, testCase.SK)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.ReqMethod = testCase.Method
		cpf.SignMethod = "HMAC-SHA1"

		client := &common.Client{}
		client.Init(regions.Beijing).WithCredential(credential).WithProfile(cpf)

		request := &SearchHAProxyRequest{
			BaseRequest: &cdshttp.BaseRequest{},
		}
		request.Init().WithApiInfo(testCase.Service, Version, testCase.Action)

		request.StartTime = common.StringPtr(testCase.StartTime)
		request.EndTime = common.StringPtr(testCase.EndTime)

		response := &SearchHAProxyResponse{BaseResponse: &cdshttp.BaseResponse{}}
		err := client.Send(request, response)
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Log(fmt.Sprintf("result: %+v", response))
		}
	}
}
