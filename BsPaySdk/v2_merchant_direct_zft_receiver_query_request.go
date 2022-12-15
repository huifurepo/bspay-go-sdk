/**
 * 直付通分账关系查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectZftReceiverQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    PageSize string `json:"page_size" structs:"page_size"` // 每页数目
    PageNum string `json:"page_num" structs:"page_num"` // 页数

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectZftReceiverQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectZftReceiverQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectZftReceiverQueryRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectZftReceiverQueryRequest(reqParam V2MerchantDirectZftReceiverQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_ZFT_RECEIVER_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
