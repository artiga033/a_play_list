package alist

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ListReq struct {
	PageReq
	Path     string `json:"path" form:"path"`
	Password string `json:"password" form:"password"`
	Refresh  bool   `json:"refresh"`
}

type DirReq struct {
	Path      string `json:"path" form:"path"`
	Password  string `json:"password" form:"password"`
	ForceRoot bool   `json:"force_root" form:"force_root"`
}

type ObjResp struct {
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	IsDir    bool      `json:"is_dir"`
	Modified time.Time `json:"modified"`
	Sign     string    `json:"sign"`
	Thumb    string    `json:"thumb"`
	Type     int       `json:"type"`
}

type FsListResp struct {
	Content  []ObjResp `json:"content"`
	Total    int64     `json:"total"`
	Readme   string    `json:"readme"`
	Write    bool      `json:"write"`
	Provider string    `json:"provider"`
}

func (c *Client) PostJson(api string, data interface{}) (*http.Response, error) {
	jsonstr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return http.Post(c.ApiUrl(api), "application/json;charset=UTF-8", bytes.NewReader(jsonstr))
}
func (c *Client) FsList(req ListReq) (*Resp[FsListResp], error) {
	httpresp, err := c.PostJson("fs/list", req)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()
	bodyText, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}
	var resp Resp[FsListResp]
	err = json.Unmarshal(bodyText, &resp)
	return &resp, err
}
