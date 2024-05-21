package model

type SiteInfo struct {
	Result int8     `json:"result"`
	Data   SiteData `json:"data"`
}

type SiteData struct {
	List []List `json:"list"`
}

type List struct {
	Id string `json:"id"`
}
