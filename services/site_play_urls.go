package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"telebotV2/global"
	"telebotV2/model"

	"go.uber.org/zap"
)

type Response struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
	Msg  string          `json:"msg"`
}

type IntermediateResponse struct {
	SiteConfig model.SiteVideoUrls `json:"siteConfig"`
}

func GetSiteVideoUrls(sideId int) (sitePlayUrls model.SiteVideoUrls, err error) {
	reqBody := map[string]interface{}{
		"siteID":     sideId,
		"parentName": "桃红",
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		global.LOG.Error("marshal request body error: ", zap.Error(err))
		return
	}

	req, err := http.NewRequest("POST", "https://3yzt.com/siteConfig/getByID", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		global.LOG.Error("new request error: ", zap.Error(err))
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		global.LOG.Error("do request error: ", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected HTTP status: %v", resp.StatusCode)
		global.LOG.Error("response status code error: ", zap.Int("code", resp.StatusCode))
		return
	}

	var body Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		global.LOG.Error("decode response body error: ", zap.Error(err))
		return
	}

	if body.Code != 0 {
		err = fmt.Errorf("unexpected response code: %v", body.Code)
		global.LOG.Error("response code error: ", zap.Int("code", body.Code))
		return
	}

	var interResp IntermediateResponse
	err = json.Unmarshal(body.Data, &interResp)
	if err != nil {
		global.LOG.Error("unmarshal response data error: ", zap.Error(err))
		return
	}

	sitePlayUrls = interResp.SiteConfig
	return
}
