/**
 * 微信支付宝入驻信息修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBusiAtModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBusiAtModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBusiAtModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // AT信息修改列表
        AtRegList:getAtRegList(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBusiAtModifyRequest(dgReq)
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
    // 业务开通结果异步消息接收地址
    extendInfoMap["busi_async_return_url"] = "http://service.example.com/to/path"
    // 支付成功页商户LOGO图片
    // extendInfoMap["ali_mer_logo"] = ""
    return extendInfoMap
}

func getAtRegList() string {
    dto := make(map[string]interface{})
    // 商户汇付ID
    dto["huifu_id"] = "6666000***456098"
    // 产品号
    dto["product_id"] = "ZDTEST"
    // 业务开通类型
    dto["fee_type"] = "03"
    // 支付通道
    dto["pay_way"] = "W"
    // 子渠道号
    dto["pay_channel_id"] = "JP00001"
    // 经营简称
    dto["short_name"] = "盈盈超市3.0"
    // 客服电话
    dto["service_phone"] = "1752***5001"
    // 商户名称
    // dto["mer_name"] = ""
    // 营业执照类型
    // dto["business_license_type"] = ""
    // 商户营业执照号
    // dto["license_code"] = ""
    // 法人身份证号
    // dto["legal_cert_no"] = ""
    // 行业分类
    // dto["cls_id"] = ""
    // 申请服务
    // dto["service_codes"] = ""
    // 结算卡
    // dto["settle_card_no"] = ""
    // 结算卡户名
    // dto["settle_card_name"] = ""
    // 商户结算卡开卡行支行名称
    // dto["mer_card_bank_branch_name"] = ""
    // 支付宝登录账号
    // dto["alipay_account"] = ""
    // 联系人类型
    // dto["contact_type"] = ""
    // 联系人姓名
    // dto["contact_name"] = ""
    // 联系人手机号
    // dto["contact_mobile"] = ""
    // 联系人邮箱
    // dto["contact_email"] = ""
    // 商户地址
    // dto["mer_addr"] = ""
    // 省份编码
    // dto["province_code"] = ""
    // 城市编码
    // dto["city_code"] = ""
    // 区县编码
    // dto["district_code"] = ""
    // 拟申请的间联商户等级
    // dto["indirect_level"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

