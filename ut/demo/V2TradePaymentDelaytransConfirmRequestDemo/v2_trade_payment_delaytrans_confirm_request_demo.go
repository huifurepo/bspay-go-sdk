/**
 * 交易确认 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradePaymentDelaytransConfirmRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradePaymentDelaytransConfirmRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradePaymentDelaytransConfirmRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109133323",
        // 交易类型**原交易为快捷支付必填：QUICK_PAY**；&lt;br/&gt;**原交易为余额支付必填：ACCT_PAYMENT**；&lt;br/&gt;原交易为全域资金必填：REMITTANCE_PAY；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：ACCT_PAYMENT&lt;/font&gt;
        // PayType:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradePaymentDelaytransConfirmRequest(dgReq)
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
    // 原交易请求日期
    extendInfoMap["org_req_date"] = "20220512"
    // 原交易请求流水号
    extendInfoMap["org_req_seq_id"] = "20220512195832E06521"
    // 原交易商户订单号
    // extendInfoMap["org_mer_ord_id"] = ""
    // 原交易全局流水号
    extendInfoMap["org_hf_seq_id"] = ""
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getA7256caaCe6b490680ba86babafe4967()
    // 安全信息
    extendInfoMap["risk_check_data"] = get5dc1f4fb2ca7478382dc4811aa4903f8()
    // 备注
    extendInfoMap["remark"] = "remark123"
    // 灵活用工标志
    // extendInfoMap["hyc_flag"] = ""
    // 灵活用工平台
    // extendInfoMap["lg_platform_type"] = ""
    // 代发模式
    // extendInfoMap["salary_modle_type"] = ""
    // 落地公司商户号
    // extendInfoMap["bmember_id"] = ""
    // 乐接活请求参数集合
    // extendInfoMap["ljh_data"] = get4da1733e10a041c7Aa241338ac87fac0()
    // 异步通知地址
    // extendInfoMap["notify_url"] = ""
    return extendInfoMap
}

func getF77aaa99F28d44d5820628512e341eac() interface{} {
    dto := make(map[string]interface{})
    // 分账金额(元)单位元，需保留小数点后两位，最低传入0.01 ，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt; ，percentage_flag非Y时必填；&lt;br/&gt;percentage_flag&#x3D;Y时div_amt不填，div_amt&#x3D;total_div_amt*percentage_div
    dto["div_amt"] = "0.01"
    // 分账接收方ID
    dto["huifu_id"] = "6666000109133323"
    // 分账百分比%
    // dto["percentage_div"] = ""
    // 分账接收方账户号
    // dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getA7256caaCe6b490680ba86babafe4967() string {
    dto := make(map[string]interface{})
    // 分账总金额（元）本次交易确认总额。需保留小数点后两位&lt;br/&gt;percentage_flag&#x3D;Y时必填。该金额与分账百分比用来计算分账金额。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：10.00&lt;/font&gt;；
    // dto["total_div_amt"] = "test"
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 分账明细
    dto["acct_infos"] = getF77aaa99F28d44d5820628512e341eac()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get5dc1f4fb2ca7478382dc4811aa4903f8() string {
    dto := make(map[string]interface{})
    // ip地址
    // dto["ip_addr"] = ""
    // 基站地址
    dto["base_station"] = "3"
    // 纬度
    dto["latitude"] = "2"
    // 经度
    dto["longitude"] = "1"
    // 产品子类
    // dto["sub_product"] = ""
    // 转账原因
    // dto["transfer_type"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get4da1733e10a041c7Aa241338ac87fac0() string {
    dto := make(map[string]interface{})
    // 税源地ID
    // dto["tax_area_id"] = ""
    // 任务模板ID
    // dto["template_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

