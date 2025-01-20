/**
 * 银行大额支付绑卡 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2LargeamtBindcardBindRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2LargeamtBindcardBindRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2LargeamtBindcardBindRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 商户号
        HuifuId:"6666000108312997",
        // 卡类型
        CardType:"1",
        // 银行账户名
        CardName:"许溪证",
        // 银行卡号密文
        CardNo:"GCMghaHLsWffNmBl/uuvVnv+kzwvBSLaZR+AsldnabAMzjPUzw4SMe2DX8IvVTM/Qb/tbiQwayeQ+TwkeSyQ0IB6oy/BNgM3rl7wZsdTzKbyigyGQvtOYsauk3IUuiJ8ptJ1k0C4Ysd5Z4+6ApLmOZhAem1pqu+DUk8EpKMj37RDgk3zWgVIf1wX9nBaSN1IGIoVjmweg8/r/UVWqCKoYrEWHxO1R0elZM9+hXTwXEKHFc2L2yossgDGjJDKuykaN0DzVunz1uQbxuvg4lMCmycSRjlQ1MCsIzqs4oiVNW3PCqAwoFkdRKL879e5/EsvohJJNVuX6YOeefFdJOC8Ug==",
        // 银行编码
        BankCode:"",
        // 手机号
        MobileNo:"dFw39mqjcPyZJk5ukKiH5oL+LyJLdJ2DfPgXcOCCgYfsUuCcOJLPnBc6f0nybPDBnfgLCcK31wG5TLFi97EttpBsrQVI6kEWMrxUAAcIehSMuWEBBuGG8QnaayE0tZa2gSgQZgFltCrkgfQ08N6TwLmvZEJ3z+gudsIPRaMXAgxMgnyH6xjuFbdOWJKfgcTQxpIirIQg0bWPpBPnO6HizB3z435qeep7WVCRK7c+tvYxjLRm7jDEeUCd9c0yZ4eKWOt1vLini6kqAwXuCTXa10z1NEnGbFlBrOK5/R5ZK977BmuAD7ZLuHU6T/j2Ca1nG6JOJwXT827CVo/sU7osjQ==",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2LargeamtBindcardBindRequest(dgReq)
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
    extendInfoMap["acct_id"] = "C02730634"
    // 银行卡绑定身份证
    extendInfoMap["cert_no"] = "k05wtsAi+WSv8rdRaj24nOGQetL4L8k5VFRGPdljb1O/5pOJYe4o3ofwiKNjaVyAwvGFWIqMNEu0GU1gcq+UDmnabOROcneJVNGu+XMy5J9I55OqBDOC5lIiiuSWQux7TlaDCZT7ACpYHjRI2r3bzOASgzPXebjYLllnuEg2kxYpGqJBe8jsjaTpAzoEB1Yoy6I0sAn4xxl8IjmGu5AHEA/drWyrAIT0GsEhmeR6wkJK3iCjShqIQ317BkNBzXsdt8dGZLF4M/7iwiQXaVP2KLWKtX+gn2oI19ckTTiFXnvuNtNPJEUEayTbsAODHKvo5wsYLUdbnO2UFJ6wlE3rOQ=="
    // 联行号
    extendInfoMap["branch_code"] = "105290071051"
    // 银行所在省
    extendInfoMap["prov_id"] = "130000"
    // 银行所在市
    extendInfoMap["area_id"] = "130300"
    // 补充资质材料列表
    extendInfoMap["file_list"] = getFileList()
    return extendInfoMap
}

func getFileList() string {
    dto := make(map[string]interface{})
    // 文件jfileID
    // dto["file_id"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

