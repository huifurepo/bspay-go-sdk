/**
 * 银行大额转账
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferAccountRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 收款方商户号
    CertificateName string `json:"certificate_name" structs:"certificate_name"` // 付款方名称
    BankCardNo string `json:"bank_card_no" structs:"bank_card_no"` // 付款方银行卡号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferAccountRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferAccountRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferAccountRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferAccountRequest(reqParam V2TradeOnlinepaymentTransferAccountRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_ACCOUNT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
