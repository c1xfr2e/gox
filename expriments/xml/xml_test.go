package main

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Detail struct {
	XMLName xml.Name `xml:"detail"`
	Data    string   `xml:",cdata"`
}

type Foo struct {
	Hello string `xml:"hello"`
	World int    `xml:"world,omitempty"`
}

type Bar struct {
	*Foo
	XMLName xml.Name `xml:"bar"`
	App     string   `xml:"app"`
	Value   int      `xml:"value"`
	Detail  *Detail  `xml:"detail"`
}

func doMarshal(data interface{}) error {
	b, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func Test_Marshal(t *testing.T) {
	bar := &Bar{
		Foo:    &Foo{Hello: "hello", World: 123456},
		App:    "MyApp",
		Value:  1000,
		Detail: &Detail{Data: `<a href="http://example.org">My Example Website</a>`},
	}
	err := doMarshal(bar)
	if err != nil {
		t.Error(err)
	}
}

// responseStatus represents result status from all API functions.
type responseStatus struct {
	// 返回结果 (SUCCESS/FAIL), 此字段是通信标识， 非交易标识
	Result string `xml:"return_code"`

	// 返回信息， 当 return_code 为 FAIL 时返回信息为错误原因例如签名失败
	Message string `xml:"return_msg"`

	// 业务结果 （SUCCESS/FAIL）
	APIResult string `xml:"result_code"`

	// 业务错误码
	APIErrorCode string `xml:"err_code"`

	// 业务错误码描述
	APIErrorMessage string `xml:"err_code_des"`
}

// createPaymentOrderResponse 用于表示微信支付 `统一下单` 接口的返回
type createPaymentOrderResponse struct {
	responseStatus

	// 预支付交易会话标识
	PrepayID string `xml:"prepay_id"`

	// 支付二维码链接
	QRCodeURL string `xml:"code_url"`
}

func doUnmarshal() error {
	str := `
	<xml>
		<return_code><![CDATA[SUCCESS]]></return_code>
		<return_msg><![CDATA[OK]]></return_msg>
		<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
		<mch_id><![CDATA[10000100]]></mch_id>
		<nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
		<openid><![CDATA[oUpF8uMuAJO_M2pxb1Q9zNjWeS6o]]></openid>
		<sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
		<result_code><![CDATA[SUCCESS]]></result_code>
		<prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
		<trade_type><![CDATA[JSAPI]]></trade_type>
	</xml>`
	var resp createPaymentOrderResponse
	if err := xml.Unmarshal([]byte(str), &resp); err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}

func Test_Unmarshal(t *testing.T) {
	if err := doUnmarshal(); err != nil {
		t.Error(err)
	}
}
