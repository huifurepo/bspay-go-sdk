/**
 * 银行大额支付差错申请
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradeOnlinepaymentTransferBankmistakeApplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    TransAmt string `json:"trans_amt" structs:"trans_amt"` // 交易金额
    OrderType string `json:"order_type" structs:"order_type"` // 订单类型
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原请求流水号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原请求日期
    NotifyUrl string `json:"notify_url" structs:"notify_url"` // 异步通知地址
    GoodsDesc string `json:"goods_desc" structs:"goods_desc"` // 商品描述
    CertificateFileId string `json:"certificate_file_id" structs:"certificate_file_id"` // 汇款凭证文件id

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradeOnlinepaymentTransferBankmistakeApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradeOnlinepaymentTransferBankmistakeApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradeOnlinepaymentTransferBankmistakeApplyRequest(reqParam)
}

func (bp *BsPay) V2TradeOnlinepaymentTransferBankmistakeApplyRequest(reqParam V2TradeOnlinepaymentTransferBankmistakeApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_ONLINEPAYMENT_TRANSFER_BANKMISTAKE_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
