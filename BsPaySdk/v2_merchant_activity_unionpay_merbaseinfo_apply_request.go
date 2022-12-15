/**
 * 银联活动商户入驻
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantActivityUnionpayMerbaseinfoApplyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    MerType string `json:"mer_type" structs:"mer_type"` // 商户类型
    DealType string `json:"deal_type" structs:"deal_type"` // 经营类型
    Mcc string `json:"mcc" structs:"mcc"` // 所属行业（MCC）MERCHANT_01-自然人 需要传入，参考[银联MCC编码](https://paas.huifu.com/partners/api/#/csfl/api_csfl_ylmccbm) ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：5311&lt;/font&gt;
    LegalMobile string `json:"legal_mobile" structs:"legal_mobile"` // 负责人手机号
    ContractIdNo string `json:"contract_id_no" structs:"contract_id_no"` // 联系人身份证号

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantActivityUnionpayMerbaseinfoApplyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantActivityUnionpayMerbaseinfoApplyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantActivityUnionpayMerbaseinfoApplyRequest(reqParam)
}

func (bp *BsPay) V2MerchantActivityUnionpayMerbaseinfoApplyRequest(reqParam V2MerchantActivityUnionpayMerbaseinfoApplyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_ACTIVITY_UNIONPAY_MERBASEINFO_APPLY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
