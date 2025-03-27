/**
 * 新增绑定卡
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletCardAddRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID
    FrontUrl string `json:"front_url" structs:"front_url"` // 跳转地址
    TrxDeviceInfo  string `json:"trx_device_info " structs:"trx_device_info "` // 设备信息域

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletCardAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletCardAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletCardAddRequest(reqParam)
}

func (bp *BsPay) V2WalletCardAddRequest(reqParam V2WalletCardAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_CARD_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
