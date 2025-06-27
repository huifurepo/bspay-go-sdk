/**
 * 商户业务开通修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBusiModifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 汇付客户Id
    OnlineBusiType string `json:"online_busi_type" structs:"online_busi_type"` // *线上业务类型编码*开通快捷、网银、余额支付、分账必填；参见[线上业务类型编码及补充材料说明](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/download_file/kyc/KYC-%E7%BA%BF%E4%B8%8A%E4%B8%9A%E5%8A%A1%E7%B1%BB%E5%9E%8B%E7%BC%96%E7%A0%81%E5%8F%8A%E8%A1%A5%E5%85%85%E6%9D%90%E6%96%99%E8%AF%B4%E6%98%8E.xlsx)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：H7999AL&lt;/font&gt;
    SignUserInfo string `json:"sign_user_info" structs:"sign_user_info"` // 签约人jsonObject格式；agreement_info中选择电子签约时必填；个人商户填本人信息。

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBusiModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBusiModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBusiModifyRequest(reqParam)
}

func (bp *BsPay) V2MerchantBusiModifyRequest(reqParam V2MerchantBusiModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BUSI_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
