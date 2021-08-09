package source

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hb0730/go-request"
	"strconv"
	"strings"
	"time"
)

type GovCN struct {
	Appid, Token, Nonce, Passid, Key string
	WifNonce, WifPaasid, WifToken    string
	startTime                        time.Time
	result                           Result
}

func NewGovCN() *GovCN {
	cn := new(GovCN)
	cn.Token = "23y0ufFl5YxIyGrI8hWRUZmKkvtSjLQA"
	cn.Nonce = "123456789abcdefg"
	cn.Passid = "zdww"
	cn.Appid = "NcApplication"
	cn.Key = "3C502C97ABDA40D0A60FBEE50FAAD1DA"
	cn.WifNonce = "QkjjtiLM2dCratiA"
	cn.WifPaasid = "smt-application"
	cn.WifToken = "fTN2pfuisxTavbTuYVSsNJHetwq5bJvC"
	return cn
}

func (c *GovCN) Time() string {
	r := c.getResult()
	if r.Code != 0 {
		return ""
	}
	return r.Data.EndUpdateTime
}

func (c *GovCN) HighRisk() []Risk {
	result := c.getResult()
	if result.Code != 0 {
		return nil
	}
	high := result.Data.Highlist
	risks := []Risk{}
	for _, v := range high {
		r := Risk{
			Type:       v.Type,
			Province:   v.Province,
			City:       v.City,
			County:     v.County,
			AreaName:   v.AreaName,
			Communitys: v.Communitys,
		}
		risks = append(risks, r)
	}
	return risks
}

func (c *GovCN) MiddleRisk() []Risk {
	result := c.getResult()
	if result.Code != 0 {
		return nil
	}

	high := result.Data.Middlelist
	risks := []Risk{}
	for _, v := range high {
		r := Risk{
			Type:       v.Type,
			Province:   v.Province,
			City:       v.City,
			County:     v.County,
			AreaName:   v.AreaName,
			Communitys: v.Communitys,
		}
		risks = append(risks, r)
	}
	return risks
}

func (r *GovCN) Close() error {
	return nil
}

func (c *GovCN) request() (Result, error) {
	var r Result
	data := c.generateAjaxParams()
	headers := c.headers(data.TimestampHeader, c.WifNonce, c.WifPaasid, c.WifToken)
	paramsByte, err := json.Marshal(data)
	if err != nil {
		return r, err
	}
	req, err := request.CreateRequest("POST", "http://103.66.32.242:8005/zwfwMovePortal/interface/interfaceJson", string(paramsByte))
	if err != nil {
		return r, err
	}
	req.SetHeaders(headers)
	err = req.Do()
	if err != nil {
		return r, err
	}
	body, err := req.GetBody()
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}
	return r, err
}

//src/common/ajax.js
func (c *GovCN) generateAjaxParams() *AjaxParams {
	timestamp := time.Now().Local().Unix()
	sha := crypto.SHA256.New()
	sha.Write([]byte(fmt.Sprintf("%d%s%s%d", timestamp, c.Token, c.Nonce, timestamp)))
	signatureHeader := strings.ToUpper(hex.EncodeToString(sha.Sum(nil)))
	p := new(AjaxParams)
	p.AppId = c.Appid
	p.PassHeader = c.Passid
	p.TimestampHeader = strconv.FormatInt(timestamp, 10)
	p.NonceHeader = c.Nonce
	p.SignatureHeader = signatureHeader
	p.Key = c.Key
	return p
}

func (c *GovCN) headers(timestamp, wifNonce, wifPaasid, wifToken string) map[string]string {
	sha := crypto.SHA256.New()
	sha.Write([]byte(fmt.Sprintf("%s%s%s%s", timestamp, wifToken, wifNonce, timestamp)))
	signature := hex.EncodeToString(sha.Sum(nil))
	headers := map[string]string{
		"x-wif-nonce":     wifNonce,
		"x-wif-paasid":    wifPaasid,
		"x-wif-signature": strings.ToUpper(signature),
		"x-wif-timestamp": timestamp,
		"Content-Type":    "application/json; charset=utf-8",
	}
	return headers
}

func (c *GovCN) getResult() (r Result) {
	if c.startTime.IsZero() {
		r, _ = c.try(4)
		return
	}
	if time.Now().Sub(c.startTime).Minutes() > 30 {
		r, _ = c.try(4)
		return
	}
	return c.result
}
func (c *GovCN) try(num int) (Result, error) {
	return func(num int) (Result, error) {
		result, err := c.request()
		if err != nil || (num > 0 && result.Code != 0) {
			num--
			time.Sleep(2 * time.Second)
			result, err = c.try(num)
		}
		c.result = result
		c.startTime = time.Now()
		return result, err
	}(num)
}

func init() {
	Instance.Put("gov", NewGovCN())
}

type AjaxParams struct {
	AppId           string `json:"appId"`
	Key             string `json:"key"`
	NonceHeader     string `json:"nonceHeader"`
	PassHeader      string `json:"paasHeader"`
	SignatureHeader string `json:"signatureHeader"`
	TimestampHeader string `json:"timestampHeader"`
}

type Result struct {
	Data struct {
		EndUpdateTime string `json:"end_update_time"`
		Hcount        int    `json:"hcount"`
		Mcount        int    `json:"mcount"`
		Highlist      []struct {
			Type       string   `json:"type"`
			Province   string   `json:"province"`
			City       string   `json:"city"`
			County     string   `json:"county"`
			AreaName   string   `json:"area_name"`
			Communitys []string `json:"communitys"`
		} `json:"highlist"`
		Middlelist []struct {
			Type       string   `json:"type"`
			Province   string   `json:"province"`
			City       string   `json:"city"`
			County     string   `json:"county"`
			AreaName   string   `json:"area_name"`
			Communitys []string `json:"communitys"`
		} `json:"middlelist"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
