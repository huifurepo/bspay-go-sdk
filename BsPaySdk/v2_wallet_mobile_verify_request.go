/**
 * 钱包绑定手机号验证
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletMobileVerifyRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    UserHuifuId string `json:"user_huifu_id" structs:"user_huifu_id"` // 钱包用户ID斗拱系统生成的钱包用户ID。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123122343&lt;/font&gt;&lt;br/&gt;验证类型为2-密码修改和3-密码重置时，必须提供钱包用户的汇付ID。
    MobileNo string `json:"mobile_no" structs:"mobile_no"` // 用户手机号
    Type string `json:"type" structs:"type"` // 验证类型

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletMobileVerifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletMobileVerifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletMobileVerifyRequest(reqParam)
}

func (bp *BsPay) V2WalletMobileVerifyRequest(reqParam V2WalletMobileVerifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_MOBILE_VERIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
