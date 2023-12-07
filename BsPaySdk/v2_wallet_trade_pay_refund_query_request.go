/**
 * 钱包支付退款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletTradePayRefundQueryRequest struct {
    SysId string `json:"sys_id" structs:"sys_id"` // 系统号
    ProductId string `json:"product_id" structs:"product_id"` // 产品号
    Sign string `json:"sign" structs:"sign"` // 加签结果
    Data string `json:"data" structs:"data"` // 数据

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletTradePayRefundQueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletTradePayRefundQueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletTradePayRefundQueryRequest(reqParam)
}

func (bp *BsPay) V2WalletTradePayRefundQueryRequest(reqParam V2WalletTradePayRefundQueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_TRADE_PAY_REFUND_QUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
