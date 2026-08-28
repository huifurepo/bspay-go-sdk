/*
 * @Author: huaze.fan
 * @LastEditTime: 2022-11-10 20:59:50
 * @LastEditors: huaze.fan
 * @Description: Good Good Study！Day Day Up
 */
package BsPaySdk

type MerchSysConfig struct {
	ProductId          string `json:"product_id"`
	SysId              string `json:"sys_id"`
	RsaMerchPrivateKey string `json:"rsa_merch_private_key"`
	RsaHuifuPublicKey  string `json:"rsa_huifu_public_key"`
	Sm2MerchPrivateKey string `json:"sm2_merch_private_key"`
	Sm2HuifuPublicKey  string `json:"sm2_huifu_public_key"`
	SignType           string `json:"sign_type"`
}

func (msc *MerchSysConfig) IsEmpty() bool {
	return msc.ProductId == "" || msc.SysId == "" || msc.RsaMerchPrivateKey == "" || msc.RsaHuifuPublicKey == ""
}

func (msc *MerchSysConfig) IsSm2Empty() bool {
	return msc.ProductId == "" || msc.SysId == "" || msc.SignType == "" || msc.Sm2MerchPrivateKey == "" || msc.Sm2HuifuPublicKey == ""
}