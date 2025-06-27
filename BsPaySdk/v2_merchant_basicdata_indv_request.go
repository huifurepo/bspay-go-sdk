/**
 * 个人商户进件
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2MerchantBasicdataIndvRequest struct {
    ReqSeqId string `json:"req_seq_id" structs:"req_seq_id"` // 请求流水号
    ReqDate string `json:"req_date" structs:"req_date"` // 请求日期
    UpperHuifuId string `json:"upper_huifu_id" structs:"upper_huifu_id"` // 直属渠道号
    RegName string `json:"reg_name" structs:"reg_name"` // 商户名
    Mcc string `json:"mcc" structs:"mcc"` // 所属行业
    SceneType string `json:"scene_type" structs:"scene_type"` // 场景类型
    DistrictId string `json:"district_id" structs:"district_id"` // 经营区
    DetailAddr string `json:"detail_addr" structs:"detail_addr"` // 经营详细地址scene_type字段含有线下场景时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：上海市徐汇区XX路XX号&lt;/font&gt;
    LegalCertNo string `json:"legal_cert_no" structs:"legal_cert_no"` // 负责人证件号码
    LegalCertBeginDate string `json:"legal_cert_begin_date" structs:"legal_cert_begin_date"` // 负责人证件有效期开始日期
    LegalCertEndDate string `json:"legal_cert_end_date" structs:"legal_cert_end_date"` // 负责人证件有效期截止日期
    LegalAddr string `json:"legal_addr" structs:"legal_addr"` // 负责人身份证地址
    LegalCertBackPic string `json:"legal_cert_back_pic" structs:"legal_cert_back_pic"` // 负责人身份证国徽面
    LegalCertFrontPic string `json:"legal_cert_front_pic" structs:"legal_cert_front_pic"` // 负责人身份证人像面
    ContactMobileNo string `json:"contact_mobile_no" structs:"contact_mobile_no"` // 负责人手机号
    ContactEmail string `json:"contact_email" structs:"contact_email"` // 负责人电子邮箱
    CardInfo string `json:"card_info" structs:"card_info"` // 结算卡信息配置
    SettleCardFrontPic string `json:"settle_card_front_pic" structs:"settle_card_front_pic"` // 结算卡正面
    MerIcp string `json:"mer_icp" structs:"mer_icp"` // 商户ICP备案编号商户ICP备案编号或网站许可证号；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：沪ICP备06046402号-28 &lt;/font&gt;&lt;br/&gt;类型为PC网站时，且为企业商户，且开通快捷或网银，或大额转账，或余额支付或分账业务（20%（不含）-100%），或为个人商户开通分账业务（10%（不含）-100%），必填
    StoreHeaderPic string `json:"store_header_pic" structs:"store_header_pic"` // 店铺门头照文件类型：F22；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;&lt;br/&gt;微信/支付宝实名认证个人商户，门头照也使用此字段； &lt;br/&gt;门店场所：提交门店门口照片，要求招牌清晰可见; &lt;br/&gt;小微商户流动经营/便民服务：提交经营/服务现场照片
    StoreIndoorPic string `json:"store_indoor_pic" structs:"store_indoor_pic"` // 店铺内景/工作区域照文件类型：F24；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;&lt;br/&gt;微信/支付宝实名认证个人商户，内景照也使用此字段； &lt;br/&gt;门店场所：提交店内环境照片 &lt;br/&gt;小微商户流动经营/便民服务：可提交另一张经营/服务现场照片
    StoreCashierDeskPic string `json:"store_cashier_desk_pic" structs:"store_cashier_desk_pic"` // 店铺收银台/公司前台照商户线下场景需要提供；文件类型：F105；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
    HeadHuifuId string `json:"head_huifu_id" structs:"head_huifu_id"` // 上级商户汇付ID如果head_office_flag&#x3D;0，则字段必填&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2MerchantBasicdataIndvRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2MerchantBasicdataIndvRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2MerchantBasicdataIndvRequest(reqParam)
}

func (bp *BsPay) V2MerchantBasicdataIndvRequest(reqParam V2MerchantBasicdataIndvRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_MERCHANT_BASICDATA_INDV
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
