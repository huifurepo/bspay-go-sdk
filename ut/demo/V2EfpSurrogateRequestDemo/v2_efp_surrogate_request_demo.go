/**
 * 全渠道资金付款申请 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2EfpSurrogateRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2EfpSurrogateRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2EfpSurrogateRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户汇付id
        HuifuId:"6666000108422302",
        // 交易金额.单位:元，2位小数
        CashAmt:"0.01",
        // 银行账号
        CardNo:"cDH2Gq/a7PnrH5tvA6JNEFEcLewpEW3x5cRiyJRUEwpoqTMmp74ObRCJCarqKPRnMrnHbXfa1WGAXW24f6SLiDKqCdI0zc5+tQtKBXS5Kh/mfmJIDNU710i5IDs+7luEIpRrsppg6YJejRhGY0TvPVY19dRaJ3KxIeyTkUDv/9KEb8TELxm2GBgfiFlKVPKxf95WpaZWV2kT3rh0ddJXA9YgUvHcTcEEY7GeCv5OHOaquIcP38kv27ZL2rScqgGpmluhyevPtDmvXRkdGK68IfNnWeqfCRjDAdVqcMskTb5Ajq8dtnNlx7uuSwYYKBeJKCzoPX8SE5X+f/9d62Cutw==",
        // 银行编号
        BankCode:"01050000",
        // 银行卡用户名
        CardName:"交通银行股份有限公司",
        // 对公对私标识
        CardAcctType:"E",
        // 省份
        ProvId:"310000",
        // 地区
        AreaId:"310100",
        // 手机号对私必填，使用斗拱系统的公钥对手机号进行RSA加密得到秘文；  示例值：b9LE5RccVVLChrHgo9lvp……PhWhjKrWg2NPfbe0mkUDd
        MobileNo:"AJnlbnjQcbTgyDv2NSNdVpMlpE5PkMqtppZj1AQ7yxAbvPhWHwHUzq7J+6C8PIrsHWwI6iwAo07N77zUIbMmORzRY1eENJ9intq0/nGEbRDQ3s6EtV/AXVUR9Pv+GOqetpX5Yi+htEbpKObW8V+jEUngz4L08E5VsPLSjmLKeLkVXGKiMr8jeZf/+QAhDiJFyi533dxHL+KPT0qCa3iebau1NXy17sZm4izmeYf35LxTlgZbQdxhC50z3zlkhZvMsArtod1CmlzI+SB5T3bwqpVkR22o6BkTbLrqBZp+zz5x99o6sqIEKMrwKYjDOJ0UjYsjn+KFTa+PFvJzstmqhg==",
        // 证件类型
        CertType:"11",
        // 证件号
        CertNo:"KbQ+WwhycbCOeIbrB+pH+eEsJPcYo2Q1IhMUQosshs00qy7hor+CA71bZLMazVOuFkeJxex9BfhR9W2hQNbRaqdWI4yxkDOTw9Qkx1PDTDl/n8CXpxWqQKhObCE5UEd5b+M/wWe+iKNYGcJkcoyswHdMA8kZoezxqwVUi0tbq//1Ov+kTyMVhmIwNbWJpahDvS+f780opCAtlMbz9hl25EcPpeTtNgbruKY+jeO4j6oejFK0epg616uC9jQalryERsX4EjaLqQrtd5nwZBkASc5Up56xkVqvaOo+6hFQP/KbCymxWbM3J0/PFsJtv/CPM4+9JkWusX/Q1ZEH8wdZ+A==",
        // 统一社会信用代码对公必填
        LicenceCode:"9131000010000595XD",
        // 挂网协议地址
        AgreementUrl:"https://cloudpnrcdn.oss-cn-shanghai.aliyuncs.com/opps/api/prod/dg_gwxy/PaymentServiceAgreement_surrogate_20241113.html",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2EfpSurrogateRequest(dgReq)
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
    // 原请求流水号
    extendInfoMap["efp_req_seq_id"] = "20241128164919defaultit687891821"
    // 原请求日期
    extendInfoMap["efp_req_date"] = "20241128"
    // 联行号
    extendInfoMap["branch_code"] = ""
    // 备注
    extendInfoMap["remark"] = "saas申请付款"
    // 异步通知地址
    extendInfoMap["notify_url"] = "http://www.service.com/getresp"
    return extendInfoMap
}

