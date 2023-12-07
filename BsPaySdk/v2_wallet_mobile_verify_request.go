/**
 * 钱包绑定手机号验证
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletMobileVerifyRequest struct {
    SysId string `json:"sys_id" structs:"sys_id"` // 系统号
    ProductId string `json:"product_id" structs:"product_id"` // 产品号
    Sign string `json:"sign" structs:"sign"` // 加签结果
    Data string `json:"data" structs:"data"` // 数据

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
