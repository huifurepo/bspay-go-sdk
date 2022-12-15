/**
 * 新增快捷/代扣解绑接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2QuickbuckleUnbindRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付商户Id
    OutCustId string `json:"out_cust_id" structs:"out_cust_id"` // 用户ID
    TokenNo string `json:"token_no" structs:"token_no"` // 卡令牌

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2QuickbuckleUnbindRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2QuickbuckleUnbindRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2QuickbuckleUnbindRequest(reqParam)
}

func (bp *BsPay) V2QuickbuckleUnbindRequest(reqParam V2QuickbuckleUnbindRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_QUICKBUCKLE_UNBIND
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
