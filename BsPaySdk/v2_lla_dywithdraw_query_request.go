/**
 * 提现记录查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2LlaDywithdrawQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    AgencyHuifuId string `json:"agency_huifu_id" structs:"agency_huifu_id"` // 代运营汇付id
    MerchantHuifuId string `json:"merchant_huifu_id" structs:"merchant_huifu_id"` // 商家汇付id
    PlatformType string `json:"platform_type" structs:"platform_type"` // 平台
    StartDate string `json:"start_date" structs:"start_date"` // 提现发起开始日期
    Cursor string `json:"cursor" structs:"cursor"` // 查询游标
    Size string `json:"size" structs:"size"` // 页大小

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2LlaDywithdrawQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2LlaDywithdrawQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2LlaDywithdrawQueryRequest(reqParam)
}

func (bp *BsPay) V2LlaDywithdrawQueryRequest(reqParam V2LlaDywithdrawQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_LLA_DYWITHDRAW_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
