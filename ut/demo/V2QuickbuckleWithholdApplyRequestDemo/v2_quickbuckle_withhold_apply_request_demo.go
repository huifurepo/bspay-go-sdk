/**
 * 代扣绑卡申请 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2QuickbuckleWithholdApplyRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2QuickbuckleWithholdApplyRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2QuickbuckleWithholdApplyRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求时间
        ReqDate:tool.GetCurrentDate(),
        // 汇付Id
        HuifuId:"6666000003078984",
        // 返回地址
        ReturnUrl:"http://www.huifu1234.com/",
        // 用户id
        OutCustId:"16666000106789536",
        // 绑卡订单号
        OrderId:"20230525081932677893621",
        // 绑卡订单日期
        OrderDate:"20230525",
        // 银行卡号
        CardId:"ZSSW+34A2soLbwLQ5SkZJO4Azy6BknTGkk6EYDTbGA+G0v+zcF3TnU4iYH171KB4ReLjJlY+hSy8MvgVbAx7dL9V7LvLFJd8RE+Lp6XKiIbVUCA1wd2Otp2jI2D32z5gUFqUbB4clRZyRyltXV3xmAWH4fLZDER3H+QwC0/UNF4=",
        // 银行卡开户姓名
        CardName:"H12ShtAyV4I4sOQqbISH4eMQUcmzpYOHggxRcXhxNoForh5qLyFgDrsSTn0nnepnPO8okfZYSWQlWIBzsRyyHYwAk94s2sO2Sz/6q4Jg2xDieeGDGrnrAphc8/OAN2OK8dMdbQzL12MvPQU/GX148MCxJzGvvdRFqTEPRLOLXTs=",
        // 银行卡绑定证件类型
        CertType:"00",
        // 银行卡绑定身份证
        CertId:"FviSPp2Xv6QYfRSYRZcouGAz4BvfZRS9nFKI/7daIUtn4JmBVMTDtrqKLCWeoY7WP4hQAz3rptjOe8WsuynRG3kQhBsXZB0v6e1X1+POD5FXVojquKQb1BF5tKlaOqTj/+G62URC3SWui26JzQQmjGhCORXXHFD7PPNJKusYhHI=",
        // 银行卡绑定手机号
        CardMp:"GmMLD+v2Mfc/vr9HOVFKOon3Dl4Q9cjze21X902G8Dnl2/2rpH8wpJUnufoYnI0nR9D2XkOm0ApOJL3ShiZxgLvnTaKrTDjRdrBJexhXbbhbfDx/2x+ZULvZHOEjzRI21tK2WKUzxDhX/lw/iXMjslKNVYlQ7as/aH5bLipf12g=",
        // 个人证件有效期类型
        CertValidityType:"0",
        // 个人证件有效期起始日
        CertBeginDate:"20140504",
        // 卡的借贷类型
        // DcType:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2QuickbuckleWithholdApplyRequest(dgReq)
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
    // 页面有效期
    extendInfoMap["expire_time"] = "15"
    // CVV2
    extendInfoMap["vip_code"] = "HOVFKOon3Dl4Q9cjze21X902G8Dnl2LvLFJd8RE+Lp6XKiIbVUCA1wd2Otp2jI2D32z5gUFqUbB4clRZyRyltXV3xmAWH4fLZDER3H+YwAk94s2sO2Sz/6q4Jg2xDieesO2Sz/6q4Jg2xDieeGDGbQzL12MvPQU/GX14xJzGvvd="
    // 卡有效期
    extendInfoMap["expiration"] = "IUtn4JmBVMTDtrqKLCWeoY7WP4hQAz3rptjOe8WsuySW+34SkZJO4Azy6BknTGkk6EA2soLbwLQ5SkZJO4Azy6BknTGkk6EX902G8Dnl2/2rpH8wpJUnufoYnI0nR9YDTbGA+G0v+ApOJL3ShiZxgLvnTaKrnU4iYH171KB4="
    // 个人证件有效期到期日
    extendInfoMap["cert_end_date"] = "20260504"
    // 设备信息域
    extendInfoMap["trx_device_info"] = getB56bfbeeE390414999542cc647b3ed9a()
    // 风控信息
    extendInfoMap["risk_info"] = getF494b2e9E5cf4fe592d8Fa81cc893071()
    // 代扣绑卡类型
    // extendInfoMap["binding_card_type"] = ""
    return extendInfoMap
}

func getB56bfbeeE390414999542cc647b3ed9a() interface{} {
    dto := make(map[string]interface{})
    // 银行预留手机号
    dto["trx_mobile_num"] = "15556622368"
    // 设备类型
    dto["trx_device_type"] = "1"
    // 交易设备IP
    dto["trx_device_ip"] = "10.10.0.1"
    // 交易设备MAC
    dto["trx_device_mac"] = "10.10.0.1"
    // 交易设备IMEI
    dto["trx_device_imei"] = "030147441006000182623"
    // 交易设备IMSI
    dto["trx_device_imsi"] = "030147441006000182623"
    // 交易设备ICCID
    dto["trx_device_icc_id"] = "030147441006000182623"
    // 交易设备WIFIMAC
    dto["trx_device_wfifi_mac"] = "030147441006000182623"
    // 交易设备GPS
    dto["trx_device_gps"] = "030147441006000182623"

    return dto;
}

func getF494b2e9E5cf4fe592d8Fa81cc893071() interface{} {
    dto := make(map[string]interface{})
    // IP类型
    dto["ip_type"] = "04"
    // IP值
    dto["source_ip"] = "192.168.1.2"
    // 设备标识
    dto["device_id"] = "123"
    // 设备类型
    dto["device_type"] = "1"
    // 银行预留手机号
    dto["mobile"] = "13778787106"
    // 协议编号
    // dto["agreement_no"] = ""
    // 协议地址
    // dto["agreement_url"] = ""

    return dto;
}

