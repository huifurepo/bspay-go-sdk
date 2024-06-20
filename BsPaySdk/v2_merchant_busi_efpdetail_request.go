/**
 * 全渠道资金管理配置查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiEfpdetailRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付id
    OutFundsGateId string `json:"out_funds_gate_id" structs:"out_funds_gate_id"` // 银行类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiEfpdetailRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiEfpdetailRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiEfpdetailRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiEfpdetailRequest(reqParam V2MerchantBusiEfpdetailRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_EFPDETAIL
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
