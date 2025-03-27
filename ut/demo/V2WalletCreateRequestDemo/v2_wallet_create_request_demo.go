/**
 * 钱包开户 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2WalletCreateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2WalletCreateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2WalletCreateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000107309462",
        // 个人姓名钱包账户开户人的本人真实姓名；wallet_type&#x3D;1时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：张三&lt;/font&gt;
        Name:"张三",
        // 钱包绑定手机号
        MobileNo:"DezMPeP4zos5Y5ltOuuXBaJ7WshobdsrWiNDaWRdIcrUWoF4S8HXlm7bXkuC92nxWMHS2iw2qs3bBbKibS3BcVYFAwPZXyFr+LRFVcfskWSVnBU97JA3ARUbMmGKPqAOwp0I1S0ybtDdhQUodgjtUrAGWHUObzx0Qw0hhU/0ZEnfOV4HNrXGTmI1+z4JDubi07wJ1NsB5XDb62U3ops2aJUVNSWwzzFQnHu6YJG9wgc40PeJYxZNJXzSh0WaLOhxWFeAzpz4Fe+D0xYB9cigB6DKM51YkO0oTk28Xz5TuPJzmSz3Nl36wWVruekAamrv0W7o0PqmPOIOQIyqc8bfOw==",
        // 手机短信验证码
        VerifyCode:"244372",
        // 短信验证流水号
        VerifySeqId:"WALLET0000000054024905",
        // 跳转地址
        FrontUrl:"https://www.huifu.com/products-services/",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2WalletCreateRequest(dgReq)
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
    // 请求失效时间
    extendInfoMap["time_expired"] = "30"
    // 钱包类型
    extendInfoMap["wallet_type"] = "1"
    // 电子邮箱
    extendInfoMap["email"] = "123@huifu.com"
    // 实名验证信息
    extendInfoMap["real_name_veri_info"] = getDd8b4ee218634d47Ba4c37c34ccd5b31()
    // 风险验证方式
    extendInfoMap["risk_check_type"] = "2"
    // 是否开通免密支付
    extendInfoMap["encryption_free_flag"] = "Y"
    // 单笔支付免密额度
    extendInfoMap["encryption_free_quota"] = "100.00"
    // 单日支付免密次数
    extendInfoMap["encryption_free_times"] = "5"
    // 异步通知地址
    extendInfoMap["notify_url"] = "https://bspay-stag.cloudpnr.com/sw/manager/callback/store"
    return extendInfoMap
}

func get2168fcfd89724371B0c2Ac35d6c5e564() string {
    dto := make(map[string]interface{})
    // 银行预留手机号
    dto["trx_mobile_num"] = "13771817106"
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

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getD8d893a20b24463c80ecC5a6f6543239() string {
    dto := make(map[string]interface{})
    // IP类型
    dto["ip_type"] = "04"
    // IP值
    dto["source_ip"] = "10.10.0.1"
    // 设备标识
    dto["device_id"] = "030147441006000182623"
    // 设备类型
    dto["device_type"] = ""
    // 银行预留手机号
    dto["mobile"] = "13771817106"

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get9e6d80e68292401a810e1570231010f6() string {
    dto := make(map[string]interface{})
    // 银行卡号
    dto["card_no"] = "Fihn3upMcPZThzo60j6zj8NJmmqSTc/dWUQGo5VaF/+BZHUNKUDPqd++rVhS18on4bNMKv7k8tUBlWUS8caZLdKhrouE5WvYlYGkWZZfArol0XUOftwryGdBL/YY7q1DyDBCe6VV9ZZTRb17BTTQrV8whfiVXSo1LKjLS4jesm182OJSmz5fZ3RB6MlpT1PmQWQjh/GEaOAV3isF0N314Y2Sp5WNanekXd4uaOXVX8MIecL/ykAZjp4gOgzOMm5gQo/JR3mrbcv7+ifL2+4SXXeSkH2fkuplGGZjEKUZOAvTGL9WN+VlZspVp+H5QJ+LiBw/4Hdti1eiNVCf2U3EqA=="
    // 银行卡绑定手机号
    dto["card_mobile"] = "AWWRBrehFsnxJgaxHH3WMH9wMpwT6rIVIRyenCU5VAe3l2oKI5u5VUXGChUu2XYNaI5chBzLatH1hepwkKjLUXccaNZ1v2SSttOhVBX3J8eLGQ8aHtrbeQMetIglC+6yH+dg7MBRYtSJQT/hB8x/EB68394BHe6vDHDbJly4HFc3jy2ScJVwfNtrtEKeckSK3id2x/qSLtMNAl/QYc/CEQ4QBVFJzKvIPAUooaXMOmtbQrP9QNyDeXDfe9w1+Q3q0No6+m6hgRNkPTP92RnYVzqXf8TaXm7KCQg+Gj3B6dX0bxqoJ6gMvFZOR/Aq5MXvvbYZMyRYgOWO/YFsfdJdUg=="
    // 身份证照片人像面首次绑定银行卡时，需上传身份证照片。文件类型：F40  &lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    dto["cert_img_portrait"] = "57cc7f00-600a-33ab-b614-6221bbf2e529"
    // 身份证照片国徽面首次绑定银行卡时，需上传身份证照片。文件类型：F41  &lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e529&lt;/font&gt;
    dto["cert_img_emblem"] = "57cc7f00-600a-33ab-b614-6221bbf2e530"
    // 设备信息域
    dto["trx_device_info"] = get2168fcfd89724371B0c2Ac35d6c5e564()
    // 银行卡类型
    dto["dc_type"] = "D"
    // 风控信息
    dto["risk_info"] = getD8d893a20b24463c80ecC5a6f6543239()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getDd8b4ee218634d47Ba4c37c34ccd5b31() string {
    dto := make(map[string]interface{})
    // 个人证件类型
    dto["cert_type"] = "00"
    // 个人证件号码
    dto["cert_no"] = "Jv9/zMKdL3r50MQsQxvsq1rhm45P1Fv/sS8Hvp5XsMV+/lSI10onYqUFGbyv5AWrktbG+EUa6oe6/+0gbDQrFHCuGR2VAUpFPxWbQmQqQbT2x3aWcqpXS/9szoFILe7y7QC2ggwzwxE5sY7T/sCGDvpbbNo7ROZkosKUrGN3WWG81hcOLuAXhmB9qezUCuXzlejbt2RGWYsKz4Kdr+15zPakPr6jaF8ay8VXhFhEmdTPp51Z/nuuXXk6qbsmTU/e4+HUCY2v8By/XSybdie299fD/vHUK/gyvZLiG1t/WYsnvPLDnbhmxkRWEYXYMxijaQxOF7ZAqM7NyfVythO0ag=="
    // 证件有效期开始日期
    dto["cert_begin_date"] = "20230426"
    // 证件有效期类型
    dto["cert_validity_type"] = "0"
    // 开户人职业
    dto["occupation"] = "1A"
    // 开户人地址
    dto["address"] = "上海市徐汇区桂林路"
    // 证件有效期结束日期
    dto["cert_end_date"] = "20230626"
    // 钱包绑定卡信息
    dto["bind_card_info"] = get9e6d80e68292401a810e1570231010f6()

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

