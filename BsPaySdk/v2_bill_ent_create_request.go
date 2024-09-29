/**
 * 创建企业账单
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2BillEntCreateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求时间
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    PayerId string `json:"payer_id" structs:"payer_id"` // 付款人
    BillName string `json:"bill_name" structs:"bill_name"` // 账单名称
    BillAmt string `json:"bill_amt" structs:"bill_amt"` // 账单金额
    SupportPayType string `json:"support_pay_type" structs:"support_pay_type"` // 可支持的付款方式
    BillEndDate string `json:"bill_end_date" structs:"bill_end_date"` // 账单截止日期
    PayeeInfo string `json:"payee_info" structs:"payee_info"` // 收款人信息

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2BillEntCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2BillEntCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2BillEntCreateRequest(reqParam)
}

func (bp *BsPay) V2BillEntCreateRequest(reqParam V2BillEntCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_BILL_ENT_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
