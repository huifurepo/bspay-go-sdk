/**
 * 银行大额支付固定转账标识申请接口
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferFixedflagApplyRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    UniqueNo string `json:"unique_no" structs:"unique_no"` // 唯一标识号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferFixedflagApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferFixedflagApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferFixedflagApplyRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferFixedflagApplyRequest(reqParam V2TradeOnlinepaymentTransferFixedflagApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_FIXEDFLAG_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
