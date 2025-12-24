/**
 * 账单数据状态变更
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V3BillpayOrderChangestatRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    BillNo string `json:"bill_no" structs:"bill_no"` // 账单编号
    BillStat string `json:"bill_stat" structs:"bill_stat"` // 变更状态

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV3BillpayOrderChangestatRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V3BillpayOrderChangestatRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V3BillpayOrderChangestatRequest(reqParam)
}

func (bp *BsPay) V3BillpayOrderChangestatRequest(reqParam V3BillpayOrderChangestatRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V3_BILLPAY_ORDER_CHANGESTAT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
