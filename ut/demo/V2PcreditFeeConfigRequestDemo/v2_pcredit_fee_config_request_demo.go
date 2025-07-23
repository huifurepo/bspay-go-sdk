/**
 * 商户分期配置 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2PcreditFeeConfigRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2PcreditFeeConfigRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2PcreditFeeConfigRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2PcreditFeeConfigRequest(dgReq)
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
    // 异步通知地址
    extendInfoMap["async_return_url"] = "http://192.168.85.157:30031/sspm/testVirgo"
    // 银行分期费率
    extendInfoMap["bank_fq_list"] = get82eb6e11C8b84028A9b7999d1e243635()
    // 花呗分期费率
    extendInfoMap["hb_fq_fee_list"] = getF3f2d9dbB869474d9433D4cf46bfc278()
    // 白条分期配置对象
    // extendInfoMap["jdbt_data"] = getA7c308118f864036B79d041ceb8deee1()
    // 银联聚分期配置对象
    // extendInfoMap["yljfq_data"] = getB2399be5E5004e3a87f2661458655bde()
    return extendInfoMap
}

func get7014a3deD77d423689daEcda59e2f3e9() interface{} {
    dto := make(map[string]interface{})
    // 银行编号
    dto["bank_code"] = "01040000"
    // 银行名称
    dto["bank_name"] = ""
    // 银联收单分期费率（%）
    dto["bank_acq_period"] = "6"
    // 用户付息费率
    dto["bank_fee_rate"] = "2"
    // 交易手续费外扣标记
    dto["out_fee_flag"] = ""
    // 手续费外扣的汇付商户号
    dto["out_fee_huifu_id"] = ""
    // 银联分期3期开关
    dto["three_period_switch"] = "Y"
    // 银联分期3期总费率（%）
    dto["three_period"] = "10"
    // 银联分期6期开关
    dto["six_period_switch"] = "Y"
    // 银联分期6期总费率（%）
    dto["six_period"] = "16"
    // 银联分期12期开关
    dto["twelve_period_switch"] = "Y"
    // 银联分期12期总费率（%）
    dto["twelve_period"] = "0.0001"

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get82eb6e11C8b84028A9b7999d1e243635() string {
    dto := make(map[string]interface{})
    // 银联入网模式
    dto["ent_way"] = "1"
    // 商户汇付Id
    dto["huifu_id"] = "6666000003156435"
    // 银行卡分期状态
    dto["bank_card_fq_status"] = "1"
    // 银行卡分期费率
    dto["bank_fq_fee_list"] = get7014a3deD77d423689daEcda59e2f3e9()
    // 贴息模式
    dto["fee_model"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getF3f2d9dbB869474d9433D4cf46bfc278() string {
    dto := make(map[string]interface{})
    // 商户汇付Id
    dto["huifu_id"] = "6666000003156435"
    // 花呗分期状态
    // dto["hb_fq_status"] = ""
    // 花呗分期3期开关
    dto["hb_three_period_switch"] = "Y"
    // 花呗收单分期3期费率（%）
    dto["hb_three_acq_period"] = "5"
    // 花呗分期3期利率（%）
    dto["hb_three_period"] = "10"
    // 花呗分期6期开关
    dto["hb_six_period_switch"] = "Y"
    // 花呗收单分期6期费率（%）
    dto["hb_six_acq_period"] = "5"
    // 花呗分期6期利率（%）
    dto["hb_six_period"] = "10"
    // 花呗分期12期开关
    dto["hb_twelve_period_switch"] = "Y"
    // 花呗收单分期12期费率（%）
    dto["hb_twelve_acq_period"] = "15"
    // 花呗分期12期利率（%）
    dto["hb_twelve_period"] = "11"
    // 交易手续费外扣标记
    dto["out_fee_flag"] = ""
    // 手续费外扣的汇付商户号
    dto["out_fee_huifu_id"] = ""
    // 花呗分期24期开关
    // dto["hb_twentyfour_period_switch"] = ""
    // 花呗收单分期24期费率（%）
    // dto["hb_twentyfour_acq_period"] = ""
    // 花呗分期24期利率（%）
    // dto["hb_twentyfour_period"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get6f420ef657224a4a831dDd63d6e7e4be() interface{} {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 业务开通标识
    // dto["open_flag"] = "test"
    // 手续费率(%)
    // dto["fee_rate"] = ""
    // 手续费扣取方式
    // dto["fee_rec_type"] = ""
    // 交易手续费扣款标记
    // dto["fee_flag"] = ""
    // 手续费外扣的汇付商户号
    // dto["out_fee_huifu_id"] = ""
    // 手续费外扣的汇付账户号
    // dto["out_fee_acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getA7c308118f864036B79d041ceb8deee1() string {
    dto := make(map[string]interface{})
    // 商户汇付Id
    // dto["huifu_id"] = "test"
    // 签约人类型
    // dto["sign_user_type"] = "test"
    // 签约人手机号
    // dto["mobile_no"] = "test"
    // 挂网协议地址3-挂网协议必填；示例值：https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/dg_gwxy/PaymentServiceAgreement_xxxx.html
    // dto["agreement_url"] = "test"
    // 京东白条费率数据
    // dto["jdbt_fee_data"] = get6f420ef657224a4a831dDd63d6e7e4be()
    // 签约人姓名
    // dto["name"] = ""
    // 签约人身份证号
    // dto["cert_no"] = ""
    // 协议类型
    // dto["agreement_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF8d3eb05Cd5b453bB2f5Fafdcdb1b4e3() interface{} {
    dto := make(map[string]interface{})
    // 支付场景
    // dto["pay_scene"] = "test"
    // 业务开通标识
    // dto["open_flag"] = "test"
    // 贴息模式
    // dto["discount_model"] = ""
    // 手续费率(%)
    // dto["fee_rate"] = ""
    // 手续费扣取方式
    // dto["fee_rec_type"] = ""
    // 交易手续费扣款标记
    // dto["fee_flag"] = ""
    // 手续费外扣的汇付商户号
    // dto["out_fee_huifu_id"] = ""
    // 手续费外扣的汇付账户号
    // dto["out_fee_acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getEc6251de3daf467e83828ace2ce369ae() interface{} {
    dto := make(map[string]interface{})
    // 文件id
    // dto["file_id"] = "test"
    // 文件类型
    // dto["file_type"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getB2399be5E5004e3a87f2661458655bde() string {
    dto := make(map[string]interface{})
    // 商户汇付Id
    // dto["huifu_id"] = "test"
    // 签约人类型
    // dto["sign_user_type"] = "test"
    // 签约人手机号
    // dto["mobile_no"] = "test"
    // 银联聚分期费率数据
    // dto["yljfq_fee_data"] = getF8d3eb05Cd5b453bB2f5Fafdcdb1b4e3()
    // 签约人姓名
    // dto["name"] = ""
    // 签约人身份证号
    // dto["cert_no"] = ""
    // 补充业务信息
    // dto["file_list"] = getEc6251de3daf467e83828ace2ce369ae()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

