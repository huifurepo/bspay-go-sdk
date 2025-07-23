/**
 * 商户活动报名 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantActivityAddRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantActivityAddRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantActivityAddRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 汇付客户Id
        HuifuId:"6666000103627938",
        // 营业执照图片调用[图片上传接口](http://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)获取jfile文件id；[示例图片](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/spin/imgs/%E8%90%A5%E4%B8%9A%E6%89%A7%E7%85%A7%E7%A4%BA%E4%BE%8B.png)参考&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;&lt;br/&gt;活动类型为支付宝谷雨活动时无需填写任何资料
        BlPhoto:"42204258-967e-373c-88d2-1afa4c7bb8ef",
        // 店内环境图片参加教育食堂、非校园餐饮、非盈利、停车缴费行业时必传；调用[图片上传接口](http://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)获取jfile文件id；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;&lt;br/&gt;活动类型为支付宝谷雨活动时无需填写任何资料
        DhPhoto:"42204258-967e-373c-88d2-1afa4c7bb8ef",
        // 手续费类型
        FeeType:"7",
        // 整体门面图片（门头照）参加教育食堂行业、非校园餐饮、非盈利、线下教培、公办医院、商业医疗时必传；调用[图片上传接口](http://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)获取jfile文件id；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;&lt;br/&gt;活动类型为支付宝谷雨活动时无需填写任何资料&lt;br/&gt;若为线下教培活动,[示例图片](https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/spin/imgs/%E9%97%A8%E5%A4%B4%E7%85%A7%E7%A4%BA%E4%BE%8B.png)参考
        MmPhoto:"42204258-967e-373c-88d2-1afa4c7bb8ef",
        // 收银台照片参加教育食堂行业、线下教培、公办医院时必传；调用[图片上传接口](http://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)获取jfile文件id；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;&lt;br/&gt;活动类型为支付宝谷雨活动时无需填写任何资料
        SytPhoto:"42204258-967e-373c-88d2-1afa4c7bb8ef",
        // 支付通道
        PayWay:"W",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantActivityAddRequest(dgReq)
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
    // 活动类型
    extendInfoMap["activity_type"] = "BLUE_SEA"
    // 二级商户号
    extendInfoMap["sub_mer_id"] = "W5503418657189757903"
    // 二级商户名称
    extendInfoMap["sub_mer_name"] = ""
    // 异步通知地址
    extendInfoMap["async_return_url"] = "http://192.168.85.157:30031/sspm/testVirgo"
    // 证明文件图片
    extendInfoMap["certificate_file_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 收费样本
    extendInfoMap["charge_sample_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 照会
    extendInfoMap["diplomatic_note_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 事业单位法人证书图片
    extendInfoMap["inst_org_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 法人身份证图片
    extendInfoMap["legal_person_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 法人登记证书图片
    extendInfoMap["legal_person_reg_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 医疗执业许可证图片
    extendInfoMap["medical_license_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 民办非企业单位登记证书图片
    extendInfoMap["nonenterprise_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 组织机构代码证图片
    extendInfoMap["org_cert_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 机构资质证明照片
    extendInfoMap["org_qualifi_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 办学资质图片
    extendInfoMap["school_license_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 门店省市区编码
    extendInfoMap["shop_add_code"] = "110101"
    // 门店街道名称
    extendInfoMap["shop_street"] = "门店街道名称"
    // 门店租赁证明
    extendInfoMap["store_tenancy_proof_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 合作资质证明
    extendInfoMap["cooper_certi_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 优惠费率承诺函
    extendInfoMap["activity_rate_commit_photo"] = "42204258-967e-373c-88d2-1afa4c7bb8ef"
    // 商户同名银行账户信息
    extendInfoMap["bank_account"] = getA9362d330e0241b090a8D3be2713bcbf()
    // 银行开户证明图片
    extendInfoMap["bank_account_prove_photo"] = ""
    // 机构银行合作授权函图
    extendInfoMap["bank_agreement_photo"] = ""
    // 行业编码
    extendInfoMap["industry_code"] = ""
    // 商户行业资质图片
    extendInfoMap["industry_photo"] = ""
    // 负责人授权函图片
    extendInfoMap["legal_person_auth_photo"] = ""
    // 食堂经营相关资质
    // extendInfoMap["food_qualification_proof"] = ""
    // 活动费率%
    // extendInfoMap["preferential_rate"] = ""
    // 活动到期后费率%
    // extendInfoMap["expiration_rate"] = ""
    // 学校号
    // extendInfoMap["school_id"] = ""
    return extendInfoMap
}

func getA9362d330e0241b090a8D3be2713bcbf() string {
    dto := make(map[string]interface{})
    // 账户名
    dto["card_name"] = "张三"
    // 银行账号
    dto["card_no"] = "6228480402637874213"
    // 开户行名称
    dto["bank_branch_name"] = "招商银行"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

