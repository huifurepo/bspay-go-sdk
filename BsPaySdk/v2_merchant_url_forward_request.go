/**
 * 商户统一进件（页面版）
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantUrlForwardRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商号汇付分配的渠道商号，该商户进件完成后将隶属该渠道商/代理商&lt;br/&gt;sys_id主体为渠道商时填写； &lt;br/&gt;sys_id主体为总部商户时选填；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    StoreId string `json:"store_id" structs:"store_id"` // 门店号
    HeadHuifuId string `json:"head_huifu_id" structs:"head_huifu_id"` // 上级汇付Id如果head_office_flag&#x3D;0，则字段必填，如果head_office_flag&#x3D;1，上级汇付Id不可传

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantUrlForwardRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantUrlForwardRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantUrlForwardRequest(reqParam)
}

func (bp *BsPay) V2MerchantUrlForwardRequest(reqParam V2MerchantUrlForwardRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_URL_FORWARD
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
