/**
 * 钱包密码重置
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletPasswordResetRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID
    CustMobile string `json:"cust_mobile" structs:"cust_mobile"` // 钱包绑定手机号
    VerifyNo string `json:"verify_no" structs:"verify_no"` // 手机短信验证码
    VerifySeqId string `json:"verify_seq_id" structs:"verify_seq_id"` // 短信验证流水号
    FrontUrl string `json:"front_url" structs:"front_url"` // 跳转地址

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletPasswordResetRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletPasswordResetRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletPasswordResetRequest(reqParam)
}

func (bp *BsPay) V2WalletPasswordResetRequest(reqParam V2WalletPasswordResetRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_PASSWORD_RESET
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
