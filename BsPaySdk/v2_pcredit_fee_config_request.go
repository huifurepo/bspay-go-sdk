/**
 * 商户分期配置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2PcreditFeeConfigRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2PcreditFeeConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2PcreditFeeConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2PcreditFeeConfigRequest(reqParam)
}

func (bp *BsPay) V2PcreditFeeConfigRequest(reqParam V2PcreditFeeConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_PCREDIT_FEE_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
