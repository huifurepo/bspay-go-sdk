/**
 * 商户活动报名
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantActivityAddRequest struct {
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    BlPhoto string `json:"bl_photo" structs:"bl_photo"` // 营业执照图片
    DhPhoto string `json:"dh_photo" structs:"dh_photo"` // 店内环境图片
    FeeType string `json:"fee_type" structs:"fee_type"` // 手续费类型
    MmPhoto string `json:"mm_photo" structs:"mm_photo"` // 整体门面图片（门头照）
    SytPhoto string `json:"syt_photo" structs:"syt_photo"` // 收银台照片
    PayWay string `json:"pay_way" structs:"pay_way"` // 支付通道

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantActivityAddRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantActivityAddRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantActivityAddRequest(reqParam)
}

func (bp *BsPay) V2MerchantActivityAddRequest(reqParam V2MerchantActivityAddRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_ACTIVITY_ADD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
