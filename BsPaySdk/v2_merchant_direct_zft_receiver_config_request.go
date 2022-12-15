/**
 * 直付通分账关系绑定解绑
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantDirectZftReceiverConfigRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付ID
    AppId string `json:"app_id" structs:"app_id"` // 开发者的应用ID
    SplitFlag string `json:"split_flag" structs:"split_flag"` // 分账开关
    ZftSplitReceiverList string `json:"zft_split_receiver_list" structs:"zft_split_receiver_list"` // 分账接收方列表
    Status string `json:"status" structs:"status"` // 状态

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantDirectZftReceiverConfigRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantDirectZftReceiverConfigRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantDirectZftReceiverConfigRequest(reqParam)
}

func (bp *BsPay) V2MerchantDirectZftReceiverConfigRequest(reqParam V2MerchantDirectZftReceiverConfigRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_DIRECT_ZFT_RECEIVER_CONFIG
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
