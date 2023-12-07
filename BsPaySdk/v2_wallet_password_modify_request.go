/**
 * 钱包密码修改
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2WalletPasswordModifyRequest struct {
    SysId string `json:"sys_id" structs:"sys_id"` // 系统号
    ProductId string `json:"product_id" structs:"product_id"` // 产品号
    Sign string `json:"sign" structs:"sign"` // 加签结果
    Data string `json:"data" structs:"data"` // 数据

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2WalletPasswordModifyRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2WalletPasswordModifyRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2WalletPasswordModifyRequest(reqParam)
}

func (bp *BsPay) V2WalletPasswordModifyRequest(reqParam V2WalletPasswordModifyRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_WALLET_PASSWORD_MODIFY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
