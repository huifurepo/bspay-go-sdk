/**
 * 支付宝实名申请提交 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiAliRealnameApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiAliRealnameApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiAliRealnameApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 汇付ID
        HuifuId:"6666000105418240",
        // 主体信息
        AuthIdentityInfo:getAuthIdentityInfo(),
        // 联系人信息
        ContactPersonInfo:getContactPersonInfo(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiAliRealnameApplyRequest(dgReq)
  	if err != nil {
		fmt.Println("请求异常:", err)
	}

	fmt.Println("返回数据:", resp)
}

/**
 * 非必填字段
 */
func getExtendInfos() map[string]interface{} {
    // 设置非必填字段
    extendInfoMap := make(map[string]interface{})
    // 子渠道号
    extendInfoMap["pay_channel_id"] = "10000001"
    // 业务开通类型
    extendInfoMap["pay_scene"] = "1"
    // 法人身份信息
    extendInfoMap["legal_person_info"] = getLegalPersonInfo()
    // 受益人信息
    extendInfoMap["ubo_info"] = getUboInfo()
    return extendInfoMap
}

func getCertificateInfo() string {
    dto := make(map[string]interface{})
    // 登记证书类型**证照类型为登记证书时(certificate_type&#x3D;REGISTER_CERT)必填**。枚举：&lt;br/&gt;统一社会信用代码证书(CERTIFICATE_TYPE_2389)&lt;br/&gt;慈善组织公开募捐资格证书(CERTIFICATE_TYPE_2397)&lt;br/&gt;社会团体法人登记证书(CERTIFICATE_TYPE_2394)&lt;br/&gt;民办非企业单位登记证书(CERTIFICATE_TYPE_2395)&lt;br/&gt;基金会法人登记证书(CERTIFICATE_TYPE_2396)&lt;br/&gt;农民专业合作社法人营业执照(CERTIFICATE_TYPE_2398)&lt;br/&gt;宗教活动场所登记证(CERTIFICATE_TYPE_2399)&lt;br/&gt;其他证书/批文/证明(CERTIFICATE_TYPE_2400)&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：CERTIFICATE_TYPE_2389&lt;/font&gt;
    dto["cert_type"] = "CERTIFICATE_TYPE_2389"
    // 证照编号
    dto["cert_number"] = "9111000071093465XC"
    // 证照图片
    dto["cert_copy"] = "afce08c5-1548-30f8-bf70-1752c3012f66"
    // 证照商户名称
    dto["cert_merchant_name"] = "新新饭店"
    // 证照法人姓名
    dto["cert_legal_person"] = "李四"
    // 证照注册地址
    dto["cert_company_address"] = "浙江省杭州市西湖区1街道10号"
    // 证照生效时间
    dto["effect_time"] = "19990101"
    // 证照过期时间
    dto["expire_time"] = "长期"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getSupportCredentials() string {
    dto := make(map[string]interface{})
    // 小微经营类型
    dto["micro_biz_type"] = "MICRO_TYPE_STORE"
    // 门店名称
    dto["store_name"] = "张三"
    // 门店省市编码
    dto["province_code"] = "310000"
    // 门店省份
    dto["province"] = "上海"
    // 门店市行政区号
    dto["city_code"] = "310100"
    // 门店城市
    dto["city"] = "上海市"
    // 门店街道区号
    dto["district_code"] = "310107"
    // 门店街道
    dto["district"] = "普陀区"
    // 门店场所填写门店详细地址
    dto["store_address"] = "消息路"
    // 门店门头照信息或摊位照
    dto["store_door_img"] = "afce08c5-1548-30f8-bf70-1752c3012f66"
    // 门店店内照片或者摊位照侧面
    dto["store_inner_img"] = "51dd13bb-6268-36d0-ac84-c4cdc19eccba"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getQualificationInfoList() string {
    dto := make(map[string]interface{})
    // 行业类目id
    dto["mcc_code"] = "2015050700000000"
    // 行业经营许可证资质照片
    dto["image_list"] = "a5d373f6-3e79-405f-9993-fb7ea051c372"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getAuthIdentityInfo() string {
    dto := make(map[string]interface{})
    // 主体类型
    dto["business_type"] = "2"
    // 是否金融机构
    dto["finance_institution_flag"] = "N"
    // 金融机构类型
    dto["financial_type"] = "INST"
    // 金融机构许可证图片
    dto["finance_license_pics"] = "a5d373f6-3e79-405f-9993-fb7ea051c372"
    // 证照类型
    dto["certificate_type"] = "BUSINESS_CERT"
    // 登记证书信息
    dto["certificate_info"] = getCertificateInfo()
    // 单位证明函照片
    dto["company_prove_copy"] = "71da066c-5d15-3658-a86d-4e85ee67808a"
    // 辅助证明材料信息
    dto["support_credentials"] = getSupportCredentials()
    // 经营许可证
    dto["qualification_info_list"] = getQualificationInfoList()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getLegalPersonInfo() string {
    dto := make(map[string]interface{})
    // 证件持有人类型
    dto["legal_type"] = "SUPER"
    // 证件类型
    dto["card_type"] = "00"
    // 法人姓名
    dto["person_name"] = "李四"
    // 证件号码
    dto["card_no"] = "110101199909291419"
    // 证件生效时间
    dto["effect_time"] = "19990101"
    // 证件过期时间
    dto["expire_time"] = "长期"
    // 证件人像面
    dto["card_front_img"] = "afce08c5-1548-30f8-bf70-1752c3012f66"
    // 证件国徽面
    dto["card_back_img"] = "51dd13bb-6268-36d0-ac84-c4cdc19eccba"
    // 授权函照片
    dto["auth_letter_img"] = "51dd13bb-6268-36d0-ac84-c4cdc19eccba"
    // 是否为受益人
    dto["is_benefit_person"] = "N"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getContactPersonInfo() string {
    dto := make(map[string]interface{})
    // 联系人身份证号码
    dto["id_card_number"] = "120103198507275017"
    // 联系人姓名
    dto["name"] = "谢季"
    // 联系人手机号
    dto["mobile"] = "18900400032"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getUboInfo() string {
    dto := make(map[string]interface{})
    // 证件姓名
    dto["ubo_id_doc_name"] = "消化"
    // 证件类型
    dto["ubo_id_doc_type"] = "00"
    // 证件号码
    dto["ubo_id_doc_number"] = "110101199909291419"
    // 证件有效期开始时间
    dto["ubo_period_begin"] = "19990101"
    // 证件有效期结束时间
    dto["ubo_period_end"] = "20260606"
    // 证件人像面
    dto["ubo_id_doc_copy"] = "afce08c5-1548-30f8-bf70-1752c3012f66"
    // 证件国徽面
    dto["ubo_id_doc_copy_back"] = "51dd13bb-6268-36d0-ac84-c4cdc19eccba"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

