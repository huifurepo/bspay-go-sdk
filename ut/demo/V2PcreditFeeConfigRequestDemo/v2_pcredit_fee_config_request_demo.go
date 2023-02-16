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
    extendInfoMap["bank_fq_list"] = getBankFqList()
    // 花呗分期费率
    extendInfoMap["hb_fq_fee_list"] = getHbFqFeeList()
    return extendInfoMap
}

func getBankFqFeeList() interface{} {
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

func getBankFqList() string {
    dto := make(map[string]interface{})
    // 银联入网模式
    dto["ent_way"] = "1"
    // 商户汇付Id
    dto["huifu_id"] = "6666000003156435"
    // 银行卡分期状态
    dto["bank_card_fq_status"] = "1"
    // 银行卡分期费率
    dto["bank_fq_fee_list"] = getBankFqFeeList()
    // 贴息模式
    dto["fee_model"] = "1"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getHbFqFeeList() string {
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

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

