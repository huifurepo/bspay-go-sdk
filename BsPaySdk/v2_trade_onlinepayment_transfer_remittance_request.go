/**
 * 汇付入账确认
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferRemittanceRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    OrgRemittanceOrderId string `json:"org_remittance_order_id" structs:"org_remittance_order_id"` // 原汇付通道流水号
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferRemittanceRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferRemittanceRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferRemittanceRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferRemittanceRequest(reqParam V2TradeOnlinepaymentTransferRemittanceRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_REMITTANCE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
