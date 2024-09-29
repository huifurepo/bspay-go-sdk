/**
 * 快捷绑卡确认接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3QuickbuckleConfirmRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户Id
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原申请流水号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原申请日期
    VerifyCode string `json:"verify_code" structs:"verify_code"` // 验证码

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3QuickbuckleConfirmRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3QuickbuckleConfirmRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3QuickbuckleConfirmRequest(reqParam)
}

func (bp *BsPay) V3QuickbuckleConfirmRequest(reqParam V3QuickbuckleConfirmRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_QUICKBUCKLE_CONFIRM
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
