/**
 * 分期支付 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeInstallmentPaymentRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeInstallmentPaymentRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeInstallmentPaymentRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000100000000",
        // 交易金额
        TransAmt:"100.00",
        // 分期数
        InstallmentNum:"3",
        // 商品描述
        GoodsDesc:"手机",
        // 安全信息
        RiskCheckData:getRiskCheckData(),
        // 京东白条分期信息trans_type&#x3D;JDBT时，必填jsonObject字符串，京东白条分期相关信息通过该参数集上送
        JdbtData:getJdbtData(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeInstallmentPaymentRequest(dgReq)
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
    // 账户号
    // extendInfoMap["acct_id"] = ""
    // 交易类型
    // extendInfoMap["trans_type"] = ""
    // 支付场景
    // extendInfoMap["pay_scene"] = ""
    // 交易有效期
    extendInfoMap["time_expire"] = "20241008235959"
    // 手续费扣款标志
    // extendInfoMap["fee_flag"] = ""
    // 延时标识
    // extendInfoMap["delay_acct_flag"] = ""
    // 备注
    extendInfoMap["remark"] = "备注"
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://www.baidu.com/onlineAsync"
    // 分账对象
    extendInfoMap["acct_split_bunch"] = getAcctSplitBunch()
    return extendInfoMap
}

func getAcctInfosRc() interface{} {
    dto := make(map[string]interface{})
    // 商户号
    dto["huifu_id"] = "6666000100000000"
    // 分账金额
    dto["div_amt"] = ""
    // 分账百分比
    dto["percentage_div"] = "70.00"
    // 账户号
    dto["acct_id"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func getAcctSplitBunch() string {
    dto := make(map[string]interface{})
    // 百分比分账标志
    dto["percentage_flag"] = "Y"
    // 是否净值分账
    dto["is_clean_split"] = "N"
    // 分账明细
    dto["acct_infos"] = getAcctInfosRc()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskCheckData() string {
    dto := make(map[string]interface{})
    // 经度
    dto["longitude"] = "126.630128"
    // 纬度
    dto["latitude"] = "126.630128"
    // 基站地址
    dto["base_station"] = "3"
    // IP地址
    dto["ip_addr"] = "182.33.21.4"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getJdbtData() string {
    dto := make(map[string]interface{})
    // 商品数量
    dto["goods_num"] = "3"
    // 下单来源
    dto["order_source"] = "微信APP扫一扫"
    // 请求来源类型
    dto["order_source_type"] = "H5"
    // 同步通知页面
    dto["callback_url"] = "https://www.baidu.com"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

