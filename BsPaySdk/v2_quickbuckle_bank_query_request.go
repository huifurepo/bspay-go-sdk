/**
 * 银行列表查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleBankQueryRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户汇付Id
    BizType string `json:"biz_type" structs:"biz_type"` // 业务类型
    DcType string `json:"dc_type" structs:"dc_type"` // 借贷类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleBankQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleBankQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleBankQueryRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleBankQueryRequest(reqParam V2QuickbuckleBankQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_BANK_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
