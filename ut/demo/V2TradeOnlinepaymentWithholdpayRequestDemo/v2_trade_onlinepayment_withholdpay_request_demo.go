/**
 * 代扣支付 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2TradeOnlinepaymentWithholdpayRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2TradeOnlinepaymentWithholdpayRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2TradeOnlinepaymentWithholdpayRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 商户号
        HuifuId:"6666000109812884",
        // 用户客户号
        UserHuifuId:"6666000109818115",
        // 绑卡id
        CardBindId:"10024597199",
        // 订单金额
        TransAmt:"0.01",
        // 商品描述
        GoodsDesc:"代扣test",
        // 代扣类型
        WithholdType:"2",
        // 异步通知地址
        NotifyUrl:"http://www.chinapnr.com/",
        // 银行扩展数据
        ExtendPayData:getF459eb2755d4496dA22633b514331e0c(),
        // 风控信息
        RiskCheckData:get58632caaC6844b46Aa0e1f192cc7c1f1(),
        // 设备信息数据
        TerminalDeviceData:get1c595ea4332b42cc8ac8Aa44eb5fc579(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2TradeOnlinepaymentWithholdpayRequest(dgReq)
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
    // 备注
    extendInfoMap["remark"] = "reamrk123"
    // 账户号
    // extendInfoMap["acct_id"] = ""
    // 订单失效时间
    extendInfoMap["time_expire"] = "20221212121212"
    // 分账对象
    // extendInfoMap["acct_split_bunch"] = get42e47512E9a6418680173509165b2100()
    // 补贴支付信息
    // extendInfoMap["combinedpay_data"] = get68d1322323b04623Baf08462a64b731d()
    return extendInfoMap
}

func get1d06be63D3244552Bd188fa04a8e6ee4() interface{} {
    dto := make(map[string]interface{})
    // 支付金额
    // dto["div_amt"] = ""
    // 商户号
    // dto["huifu_id"] = ""
    // 分账百分比%
    // dto["percentage_div"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get42e47512E9a6418680173509165b2100() string {
    dto := make(map[string]interface{})
    // 分账信息列表
    // dto["acct_infos"] = get1d06be63D3244552Bd188fa04a8e6ee4()
    // 百分比分账标志
    // dto["percentage_flag"] = ""
    // 是否净值分账
    // dto["is_clean_split"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getF459eb2755d4496dA22633b514331e0c() string {
    dto := make(map[string]interface{})
    // 业务种类
    dto["biz_tp"] = "012345"
    // 商品简称
    dto["goods_short_name"] = "看看"
    // 网关支付受理渠道
    // dto["gw_chnnl_tp"] = "test"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get58632caaC6844b46Aa0e1f192cc7c1f1() string {
    dto := make(map[string]interface{})
    // 基站地址经纬度、基站地址、IP地址三组信息至少填写一组；&lt;br/&gt;【mcc】+【mnc】+【location_cd】+【lbs_num】&lt;br/&gt;- mcc:移动国家代码，460代表中国；3位长&lt;br/&gt;- mnc：移动网络号码；2位长；&lt;br/&gt;- location_cd：位置区域码，16进制，5位长&lt;br/&gt;- lbs_num：基站编号，16进制，5位长&lt;br/&gt;- 注意若位数不足用空格补足；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：460001039217563&lt;/font&gt;，460（mcc)， 00(mnc)，10392(location_cd)， 17563(lbs_num)
    dto["base_station"] = ""
    // ip地址经纬度、基站地址、IP地址三组信息至少填写一组；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：172.28.52.52&lt;/font&gt;
    dto["ip_addr"] = "192.168.1.1"
    // 纬度纬度整数位不超过2位，小数位不超过6位。格式为：+表示北纬，-表示南纬。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：+37.12&lt;/font&gt;；&lt;br/&gt;经纬度、基站地址、IP地址三组信息至少填写一组
    dto["latitude"] = ""
    // 经度经度整数位不超过3位，小数位不超过5位；格式为:+表示东经，-表示西经。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：-121.213&lt;/font&gt;；&lt;br/&gt;经纬度、基站地址、IP地址三组信息至少填写一组
    dto["longitude"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get1c595ea4332b42cc8ac8Aa44eb5fc579() string {
    dto := make(map[string]interface{})
    // 交易设备ip
    dto["device_ip"] = "172.31.31.145"
    // 设备类型
    dto["device_type"] = "1"
    // 交易设备gps
    // dto["device_gps"] = ""
    // 交易设备iccid
    // dto["device_icc_id"] = ""
    // 交易设备imei
    // dto["device_imei"] = ""
    // 交易设备imsi
    // dto["device_imsi"] = ""
    // 交易设备mac
    // dto["device_mac"] = ""
    // 交易设备wifimac
    // dto["device_wifi_mac"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get68d1322323b04623Baf08462a64b731d() string {
    dto := make(map[string]interface{})
    // 补贴方汇付编号
    // dto["huifu_id"] = "test"
    // 补贴方类型
    // dto["user_type"] = "test"
    // 补贴方账户号
    // dto["acct_id"] = "test"
    // 补贴金额
    // dto["amount"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

