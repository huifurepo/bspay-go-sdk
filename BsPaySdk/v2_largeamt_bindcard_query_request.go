/**
 * 银行大额支付绑卡查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LargeamtBindcardQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    CardNo string `json:"card_no" structs:"card_no"` // 银行卡号密文
    PageSize string `json:"page_size" structs:"page_size"` // 每页条数
    PageNum string `json:"page_num" structs:"page_num"` // 分页页码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LargeamtBindcardQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LargeamtBindcardQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LargeamtBindcardQueryRequest(reqParam)
}

func (bp *BsPay) V2LargeamtBindcardQueryRequest(reqParam V2LargeamtBindcardQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LARGEAMT_BINDCARD_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
