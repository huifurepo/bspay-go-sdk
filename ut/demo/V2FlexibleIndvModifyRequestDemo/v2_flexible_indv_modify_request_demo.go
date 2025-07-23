/**
 * 灵工个人用户修改 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2FlexibleIndvModifyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2FlexibleIndvModifyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2FlexibleIndvModifyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商/商户汇付Id
        UpperHuifuId:"6666000108329682",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2FlexibleIndvModifyRequest(dgReq)
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
    // 基本信息
    // extendInfoMap["basic_info"] = get34a3e4b95be148d7B2547819e1b99346()
    // 取现配置列表
    extendInfoMap["cash_config"] = getFe3a0d1d38ec4885B59603888696de3f()
    // 卡信息
    extendInfoMap["card_info"] = getE8224b0485914e82A6f679fce70d1655()
    return extendInfoMap
}

func get34a3e4b95be148d7B2547819e1b99346() string {
    dto := make(map[string]interface{})
    // 个人证件有效期类型
    // dto["cert_validity_type"] = ""
    // 个人证件有效期开始日期
    // dto["cert_begin_date"] = ""
    // 个人证件有效期截止日期
    // dto["cert_end_date"] = ""
    // 手机号
    // dto["mobile_no"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getFe3a0d1d38ec4885B59603888696de3f() string {
    dto := make(map[string]interface{})
    // 提现手续费（固定/元）fix_amt与fee_rate至少填写一项， 需保留小数点后两位，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;注：当cash_type&#x3D;D1时为节假日取现手续费
    dto["fix_amt"] = ""
    // 提现手续费率（%）fix_amt与fee_rate至少填写一项，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;注：1、如果fix_amt与fee_rate都填写了则手续费&#x3D;fix_amt+支付金额\*fee_rate2、当cash_type&#x3D;D1时为节假日取现手续费
    dto["fee_rate"] = "10.00"
    // D1工作日取现手续费固定金额单位元，需保留小数点后两位。不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;D1取现配置时选填，其他取现配置无效；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
    dto["weekday_fix_amt"] = ""
    // D1工作日取现手续费率单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;D1取现配置时选填，其他取现配置无效；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
    dto["weekday_fee_rate"] = ""
    // 业务类型
    dto["cash_type"] = "D0"
    // 是否交易手续费外扣
    dto["out_fee_flag"] = ""
    // 手续费承担方
    dto["out_fee_huifu_id"] = ""
    // 交易手续费外扣的账户类型
    dto["out_fee_acct_type"] = ""
    // 是否优先到账
    dto["is_priority_receipt"] = ""
    // 开通状态
    dto["switch_state"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getE8224b0485914e82A6f679fce70d1655() string {
    dto := make(map[string]interface{})
    // 卡号
    dto["card_no"] = "6228481269040908115"
    // 银行所在省
    dto["prov_id"] = "310000"
    // 银行所在市
    dto["area_id"] = "310100"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

