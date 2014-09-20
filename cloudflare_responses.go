package main

type CfLoadAllResponse struct {
	Request  map[string]string `json:"request"`
	Response CfRecordsResponse `json:"response"`
}

type CfRecordsResponse struct {
	Recs   CfRecordsData `json:"recs"`
	Result string        `json:"result"`
	Msg    string        `json:"msg"`
}

type CfRecordsData struct {
	Has_more bool                     `json:"has_more"`
	Count    int                      `json:"count"`
	Objs     []map[string]interface{} `json:"objs"`
}
