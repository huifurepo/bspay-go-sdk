/**
 * 企业商户进件
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataEntRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 渠道商号
    RegName string `json:"reg_name" structs:"reg_name"` // 商户名称
    ShortName string `json:"short_name" structs:"short_name"` // 商户简称
    ReceiptName string `json:"receipt_name" structs:"receipt_name"` // 小票名称
    EntType string `json:"ent_type" structs:"ent_type"` // 公司类型
    Mcc string `json:"mcc" structs:"mcc"` // 所属行业参考[汇付MCC编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_hfmccbm) ；当use_head_info_flag&#x3D;Y时不填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：5311&lt;/font&gt;
    BusiType string `json:"busi_type" structs:"busi_type"` // 经营类型1：实体，2：虚拟 ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1&lt;/font&gt; ；当use_head_info_flag&#x3D;Y时不填
    SceneType string `json:"scene_type" structs:"scene_type"` // 场景类型
    LicensePic string `json:"license_pic" structs:"license_pic"` // 证照图片
    LicenseCode string `json:"license_code" structs:"license_code"` // 证照编号
    LicenseValidityType string `json:"license_validity_type" structs:"license_validity_type"` // 证照有效期类型
    LicenseBeginDate string `json:"license_begin_date" structs:"license_begin_date"` // 证照有效期开始日期
    LicenseEndDate string `json:"license_end_date" structs:"license_end_date"` // 证照有效期截止日期格式：yyyyMMdd。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当license_validity_type&#x3D;0时必填；当license_validity_type&#x3D;1时为空；当use_head_info_flag&#x3D;Y时不填
    FoundDate string `json:"found_date" structs:"found_date"` // 成立时间
    RegCapital string `json:"reg_capital" structs:"reg_capital"` // 注册资本保留两位小数；条件选填，国营企业、私营企业、外资企业、事业单位、其他、集体经济必填，政府机构、个体工商户可为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：100.00&lt;/font&gt;
    RegDistrictId string `json:"reg_district_id" structs:"reg_district_id"` // 注册区
    RegDetail string `json:"reg_detail" structs:"reg_detail"` // 注册详细地址
    DistrictId string `json:"district_id" structs:"district_id"` // 经营区
    DetailAddr string `json:"detail_addr" structs:"detail_addr"` // 经营详细地址scene_type&#x3D;OFFLINE/ALL时必填；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：上海市徐汇区XX路XX号&lt;/font&gt;
    LegalName string `json:"legal_name" structs:"legal_name"` // 法人姓名
    LegalCertType string `json:"legal_cert_type" structs:"legal_cert_type"` // 法人证件类型
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 法人证件号码
    LegalCertValidityType string `json:"legal_cert_validity_type" structs:"legal_cert_validity_type"` // 法人证件有效期类型
    LegalCertBeginDate string `json:"legal_cert_begin_date" structs:"legal_cert_begin_date"` // 法人证件有效期开始日期
    LegalCertEndDate string `json:"legal_cert_end_date" structs:"legal_cert_end_date"` // 法人证件有效期截止日期日期格式：yyyyMMdd， &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;当legal_cert_validity_type&#x3D;0时必填；&lt;br/&gt;当legal_cert_validity_type&#x3D;1时为空；&lt;br/&gt;当use_head_info_flag&#x3D;Y时不填
    LegalAddr string `json:"legal_addr" structs:"legal_addr"` // 法人证件地址
    LegalCertBackPic string `json:"legal_cert_back_pic" structs:"legal_cert_back_pic"` // 法人身份证国徽面
    LegalCertFrontPic string `json:"legal_cert_front_pic" structs:"legal_cert_front_pic"` // 法人身份证人像面
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 联系人手机号
    ContactEmail string `json:"contact_email" structs:"contact_email"` // 联系人电子邮箱
    LoginName string `json:"login_name" structs:"login_name"` // 管理员账号
    RegAcctPic string `json:"reg_acct_pic" structs:"reg_acct_pic"` // 开户许可证企业商户需要，结算账号为对公账户必填；通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F08；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    OpenLicenceNo string `json:"open_licence_no" structs:"open_licence_no"` // 基本存款账户编号或核准号条件选填；当use_head_info_flag&#x3D;Y时不填 ；&lt;br/&gt;基本存款账户信息请填写基本存款账户编号；开户许可证请填写核准号。&lt;br/&gt;当注册地址或经营地址为如下地区时必填：江苏省、浙江省、湖南省、湖北省、云南省、贵州省、陕西省、河南省、吉林省、黑龙江省、福建省、海南省、重庆市、青海省、宁夏回族自治区；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：J2900123456789&lt;/font&gt;
    CardInfo string `json:"card_info" structs:"card_info"` // 银行卡信息配置
    SettleCardFrontPic string `json:"settle_card_front_pic" structs:"settle_card_front_pic"` // 卡正面**对私必填**。通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F13；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    SettleCertBackPic string `json:"settle_cert_back_pic" structs:"settle_cert_back_pic"` // 持卡人身份证国徽面**对私必填**。通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F56；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    SettleCertFrontPic string `json:"settle_cert_front_pic" structs:"settle_cert_front_pic"` // 持卡人身份证人像面**对私必填**。通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F55；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    AuthEnturstPic string `json:"auth_enturst_pic" structs:"auth_enturst_pic"` // 授权委托书**对私非法人、对公非同名结算必填**；通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F15；开通银行电子账户（中信E管家）需提供F520；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    HeadHuifuId string `json:"head_huifu_id" structs:"head_huifu_id"` // 上级汇付Id如果head_office_flag&#x3D;0，则字段必填，如果head_office_flag&#x3D;1，上级汇付Id不可传&lt;br/&gt;如果headOfficeFlag&#x3D;0，useHeadInfoFlag&#x3D;Y,且head_huifu_id不为空则基本信息部分复用上级的基本信息。&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    MerIcp string `json:"mer_icp" structs:"mer_icp"` // 商户ICP备案编号商户ICP备案编号或网站许可证号；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：沪ICP备06046402号-28 &lt;/font&gt;&lt;br/&gt;类型为PC网站时，且为企业商户，且开通快捷或网银，或大额转账，或余额支付或分账业务（20%（不含）-100%），或为个人商户开通分账业务（10%（不含）-100%），必填
    StoreHeaderPic string `json:"store_header_pic" structs:"store_header_pic"` // 店铺门头照
    StoreIndoorPic string `json:"store_indoor_pic" structs:"store_indoor_pic"` // 店铺内景/工作区域照
    StoreCashierDeskPic string `json:"store_cashier_desk_pic" structs:"store_cashier_desk_pic"` // 店铺收银台/公司前台照

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataEntRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataEntRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataEntRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataEntRequest(reqParam V2MerchantBasicdataEntRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_ENT
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
