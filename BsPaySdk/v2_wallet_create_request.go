/**
 * 钱包开户
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletCreateRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    Name string `json:"name" structs:"name"` // 个人姓名钱包账户开户人的本人真实姓名；wallet_type&#x3D;1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 钱包绑定手机号
    VerifyCode string `json:"verify_code" structs:"verify_code"` // 手机短信验证码
    VerifySeqId string `json:"verify_seq_id" structs:"verify_seq_id"` // 短信验证流水号
    FrontUrl string `json:"front_url" structs:"front_url"` // 跳转地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletCreateRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletCreateRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletCreateRequest(reqParam)
}

func (bp *BsPay) V2WalletCreateRequest(reqParam V2WalletCreateRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_CREATE
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
