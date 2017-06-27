package UcGoSdk

import (
	"bytes"
	"errors"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	softVersion   = "2014-06-30"
	baseUrl       = "https://api.ucpaas.com"
	appId         = "appId"
	accountSId    = "accountSId"
	token         = "token"
)

var timestamp string

type UcPaaS struct {
	TmplSMS templateSMS `json:"templateSMS"`
}

type templateSMS struct {
	AppId      string `json:"appId"`
	TemplateId string `json:"templateId"`
	PhoneNum   string `json:"to"`
	Param      string `json:"param"`
}

type result struct {
	Resp response `json:"resp"`
}

type response struct {
	RespCode      string        `json:"respCode"`
	Failure       int           `josn:"failure,omitempty"`
	ReturnTmplSMS returnTmplSMS `json:"templateSMS"`
}

type returnTmplSMS struct {
	CreateDate string `json:"createDate"`
	SMSId      string `json:"smsId"`
}

/**
 * Name: initUc
 * Description: Create the instance of UcPaaS
 */
func initUc(phoneNum, tempId string, params ...string) *UcPaaS {
	curTime := time.Now().Format("20060102150405000")
	intCurTime, _ := strconv.Atoi(curTime)
	timestamp = strconv.Itoa(intCurTime)
	catParam := ""
	for index, param := range params {
		if index == 0 {
			catParam = param
		} else {
			catParam = catParam + "," + param
		}
	}
	//param := verifyCode + "," + expireTime

	return &UcPaaS{
		TmplSMS: templateSMS{
			AppId:      appId,
			TemplateId: tempId,
			PhoneNum:   phoneNum,
			Param:      catParam,
		},
	}
}

/**
 * Name: getAuthorization
 * Descriotion: 包头验证信息,使用Base64编码（账户Id:时间戳）
 */
func (res *result) getAuthorization(accountSId string) string {
	src := accountSId + ":" + timestamp
	return base64.StdEncoding.EncodeToString([]byte(src))
}

/**
 * Name: getSigParameter
 * Description: 验证参数,URL后必须带有sig参数，sig= MD5（账户Id + 账户授权令牌 + 时间戳，共32位）(注:转成大写)
 */
func (res *result) getSigParameter(accountSId, token string) string {
	sig := accountSId + token + timestamp
	//获取MD5
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(sig))))
}

/**
 * Name: getResult
 * @return respcode 000000表示发送成功
 *
 */
func (res *result) getResult(url, dataType, method string, data []byte) {
	resp := res.connection(url, dataType, method, data)
	_ = json.Unmarshal([]byte(resp), res)
}

/**
 * Name: connection
 * @param url
 * @param dataType
 * @param method post或get
 * @param body POST数据
 * @return string
 */
func (res *result) connection(url, dataType, method string, data []byte) string {
	mine := ""
	if dataType == "json" {
		mine = "application/json"
	} else {
		mine = "application/xml"
	}
	client := &http.Client{}
	//Get data
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return err.Error()
	}
	req.Header.Add("Accept", mine)
	req.Header.Add("Content-Type", (mine + ";charset=utf-8"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", res.getAuthorization(accountSId))
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return ""
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respBody)
}

/**
 * Name:SendSMS
 * Description: Send the text message.
 * @param phoneNum: The phone number
 * @param tempId: The message template id
 * @param params... : The parameter list.
 */
func SendSMS(phoneNum, tempId string, params ...string) (respCode string, err error) {
	var rs result
	//初始化
	uc := initUc(phoneNum, tempId, params...)
	//获取Url
	resource := "/" + softVersion + "/Accounts/" + accountSId + "/Messages/templateSMS"
	parseUrl, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return "", errors.New("请求地址出错")
	}
	parseUrl.Path = resource
	urlStr := fmt.Sprintf("%v?sig=%s", parseUrl, rs.getSigParameter(accountSId, token))
	//获取post数据
	data, err := json.MarshalIndent(uc, "", "  ")
	if err != nil {
		return "", err
	}
	//POST必须大写
	rs.getResult(urlStr, "json", "POST", data)
	respCode = rs.Resp.RespCode
	if respCode != "000000" {
		return respCode, errors.New("发送验证码失败。")
	}
	return "", nil
}

