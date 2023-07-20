package rest

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"testing"
)

var testLogger *zap.Logger

/*
GET请求
*/
func TestRESTAPIGet(t *testing.T) {
	testLogger, _ = zap.NewProduction()
	defer testLogger.Sync()
	rest := NewRESTAPI("https://www.okex.win", GET, "/api/v5/account/balance", nil)
	rest.SetSimulate(true).SetAPIKey("xxxx", "xxxx", "xxxx")
	rest.SetUserId("xxxxx")
	response, err := rest.Run(testLogger, context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response:")
	fmt.Println("\thttp code: ", response.Code)
	fmt.Println("\t总耗时: ", response.TotalUsedTime)
	fmt.Println("\t请求耗时: ", response.ReqUsedTime)
	fmt.Println("\t返回消息: ", response.Body)
	fmt.Println("\terrCode: ", response.V5Response.Code)
	fmt.Println("\terrMsg: ", response.V5Response.Msg)
	fmt.Println("\tdata: ", response.V5Response.Data)

	// 请求的另一种方式
	apikey := APIKeyInfo{
		ApiKey:     "xxxxx",
		SecKey:     "xxxxx",
		PassPhrase: "xxx",
	}

	cli := NewRESTClient("https://www.okex.win", &apikey, true)
	rsp, err := cli.Get(testLogger, context.Background(), "/api/v5/account/balance", nil)
	if err != nil {
		return
	}

	fmt.Println("Response:")
	fmt.Println("\thttp code: ", rsp.Code)
	fmt.Println("\t总耗时: ", rsp.TotalUsedTime)
	fmt.Println("\t请求耗时: ", rsp.ReqUsedTime)
	fmt.Println("\t返回消息: ", rsp.Body)
	fmt.Println("\terrCode: ", rsp.V5Response.Code)
	fmt.Println("\terrMsg: ", rsp.V5Response.Msg)
	fmt.Println("\tdata: ", rsp.V5Response.Data)
}

/*
POST请求
*/
func TestRESTAPIPost(t *testing.T) {
	testLogger, _ = zap.NewProduction()
	defer testLogger.Sync()

	param := make(map[string]interface{})
	param["greeksType"] = "PA"

	rest := NewRESTAPI("https://www.okex.win", POST, "/api/v5/account/set-greeks", &param)
	rest.SetSimulate(true).SetAPIKey("xxxx", "xxxx", "xxxx")
	response, err := rest.Run(testLogger, context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response:")
	fmt.Println("\thttp code: ", response.Code)
	fmt.Println("\t总耗时: ", response.TotalUsedTime)
	fmt.Println("\t请求耗时: ", response.ReqUsedTime)
	fmt.Println("\t返回消息: ", response.Body)
	fmt.Println("\terrCode: ", response.V5Response.Code)
	fmt.Println("\terrMsg: ", response.V5Response.Msg)
	fmt.Println("\tdata: ", response.V5Response.Data)

	// 请求的另一种方式
	apikey := APIKeyInfo{
		ApiKey:     "xxxx",
		SecKey:     "xxxxx",
		PassPhrase: "xxxx",
	}

	cli := NewRESTClient("https://www.okex.win", &apikey, true)
	rsp, err := cli.Post(testLogger, context.Background(), "/api/v5/account/set-greeks", &param)
	if err != nil {
		return
	}

	fmt.Println("Response:")
	fmt.Println("\thttp code: ", rsp.Code)
	fmt.Println("\t总耗时: ", rsp.TotalUsedTime)
	fmt.Println("\t请求耗时: ", rsp.ReqUsedTime)
	fmt.Println("\t返回消息: ", rsp.Body)
	fmt.Println("\terrCode: ", rsp.V5Response.Code)
	fmt.Println("\terrMsg: ", rsp.V5Response.Msg)
	fmt.Println("\tdata: ", rsp.V5Response.Data)
}
