/**
 * 快捷绑卡确认接口 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2QuickbuckleConfirmRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2QuickbuckleConfirmRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2QuickbuckleConfirmRequest{
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 汇付商户Id
        HuifuId:"6666000109133323",
        // 用户id
        OutCustId:"6666000150648142",
        // 订单号
        OrderId:"20220408105303542461831",
        // 订单日期
        OrderDate:"20220408",
        // 银行卡号
        CardId:"VDFveDULoK6qhWuN34P8XF7tuZow4eg74zEPKjNVwSjQTW572jqmYjpRDbEtX0f56hMQUfFv5wFRjvDIY+yTl+SFKgBmoNlPDm9B3mDjOa8er5FEnI5QWgvyuqSxHWJf2eIjU7OHt3q3/0ZsbetNzIAiaAZdicn9Sawsr9yZ8ZOdBmhIPO5tqJyDkoKO5Tpj0VyZMJ5ugSDx/2XGSmX3dHQ1JKlZ/7rovB+WchA9BzZQzSTcvmmdV6mzuyTPWRzxfhyVPAVLzh5XhKyl6SMKmzOM1zSNMPtzDhTaUG4XLSM5A0+Tqt3O4tSi16f5zn2csgwbN6TMd5jrXNzC8UTpdQ==",
        // 银行卡开户姓名
        CardName:"dj1TYq2WxbX8ti/7QjuSCqYKzCusdDe/+1wIVx23iFvxMcxkiV0rUK8Hc1PTw4H+qy+6RkgDh06ZNH4EXmYpl/AhfSjyMlSgF1O1YR4/WvKzRciwo1FvEOtRU6X0isOjA+NA+lv4A7ejGTtJP3dyam0j/IsOYlOriT8aGtBgfsTw+Dc+e3coZ3iCTP0Iwk6fC/BSs/GpM7H21qcXR9OyewIgSQnU5PaV/TKTaCxtLZM+6xf8Cuulg/LK5Jth3pzEkLrY+eziftxd2jCn5E7a3pyRHVD4d5VeQZQ9kF87JFjWKhMTAOVV7znXE0hutZP124UNN48FgAyAEZa+k4fWfg==",
        // 银行卡绑定证件类型
        CertType:"00",
        // 银行卡绑定身份证
        CertId:"OLOxrl809XmVIMmPRTIyIpJHxi4HFMJNmxGz1nhZAtps9xPEVDysU3UZPBZfsNFLFcXEMENPsJ+tymbYt42dNQ+76hyEtx+qz0BQhU8SKV8H5itrI4kaXpaJf+sV8OewrJkhDidPdz01vqMEDQRhaMnNwl8snHZxDdpu7HVUz5JwsLYfBbZP2Q4CZpVWS3NunQahZ8zHQ+8EhvYVH1vE7f/M6nJt1/4GoHz9C/UnuYudV0S2uBhlywbX+YkRGNMl/oz5+UNeMDA2jqhtTreSD4+w7S82L53rqKsAbosodOPo4OAMZk4/0QOH5LDbqFByVM97mzHfw7z+Tx/dsXJ8QQ==",
        // 银行卡绑定手机号
        CardMp:"fD4awMK+JdxQi+Qzl1xgJbgq4jlNTKFSKlts2C9hJhFbu0J6K7mHmViRh5wG3bDdYAKbKEAz+Uzd1xyEeYZsFNi3jQgu8gRv5sTsjHZHYIbvvq1ju2nLXrzq8kTzVXnWRXB0oHxy6EnN2xuvaC3mT89sW5BND09J7qy+Va3Y/t1nTqz4oEE5qL+TTjm6Fv6BY8ac/T2mKaeHtN27Ufj4hmJnGTtcTuoS0uQ6bEksQiTHwK2QG7EBMInnoYiJD15cAPwQeR0xmZNAwEXehtxvIAAfFpAiLqe/2G9whSyoT/BlsrhYf+958bis856dTm6Lf6nAVjQbNvh6Iyhdu7H1Rw==",
        // 验证码
        VerifyCode:"111111",
        // CVV2信用卡代扣专用需要密文传输，使用汇付RSA公钥加密(加密前3位，加密后最长2048位），[参见参考文档](https://paas.huifu.com/partners/guide/#/api_jiami_jiemi)；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Ly+fnExeyPOTzf……rGq0l0NTebfc6XJXZss&#x3D;&lt;/font&gt;
        // VipCode:"test",
        // 卡有效期信用卡代扣专用，格式：MMYY，需要密文传输。&lt;br/&gt;使用汇付RSA公钥加密(加密前4位，加密后最长2048位），[参见参考文档](https://paas.huifu.com/partners/guide/#/api_jiami_jiemi)；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：Ly+fnExeyPO……ebfc6XJXZss&#x3D;&lt;/font&gt;
        // Expiration:"test",
        // 设备信息域 
        // TrxDeviceInf:getTrxDeviceInf(),
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2QuickbuckleConfirmRequest(dgReq)
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
    // 商户名称
    extendInfoMap["merch_name"] = "测试"
    // 验卡流水号
    extendInfoMap["trans_id"] = "20220408105303542461831"
    // 卡的借贷类型
    // extendInfoMap["dc_type"] = ""
    // 风控信息
    // extendInfoMap["risk_info"] = getRiskInfo()
    return extendInfoMap
}

func getTrxDeviceInf() string {
    dto := make(map[string]interface{})
    // 银行预留手机号
    // dto["trx_mobile_num"] = "test"
    // 设备类型
    // dto["trx_device_type"] = "test"
    // 交易设备IP
    // dto["trx_device_ip"] = "test"
    // 交易设备MAC
    // dto["trx_device_mac"] = ""
    // 交易设备IMEI
    // dto["trx_device_imei"] = ""
    // 交易设备IMSI
    // dto["trx_device_imsi"] = ""
    // 交易设备ICCID
    // dto["trx_device_icc_id"] = ""
    // 交易设备WIFIMAC
    // dto["trx_device_wfifi_mac"] = ""
    // 交易设备GPS
    // dto["trx_device_gps"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getRiskInfo() string {
    dto := make(map[string]interface{})
    // IP类型
    // dto["ip_type"] = "test"
    // IP值
    // dto["source_ip"] = "test"
    // 设备标识
    // dto["device_id"] = ""
    // 设备类型
    // dto["device_type"] = ""
    // 银行预留手机号
    // dto["mobile"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

