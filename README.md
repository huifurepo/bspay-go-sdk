# GO SDK 简介
    为了提高客户接入的便捷性，本系统提供 SDK 方式接入，使用本 SDK 将极大的简化开发者的工作，开发者将无需考虑通信、签名、验签等，只需要关注业务参数的输入。

## SDK 项目结构说明
- BsPaySdk -- SDK核心包, 包含通信, 加解签, 接口参数对象等
- config -- 商户配置
- ut -- 接口调用/参数赋值演示demo

## SDK 接入说明 
以下两种方式任选其一
1. 直接在go.mod中引用(require github.com/huifurepo/bspay-go-sdk/BsPaySdk [version])
2. 直接下载源码文件, 将BsPaySdk(SDK核心包)源码放入项目中


## SDK 使用说明
    接口命名直接根据接口URL来命名, 方便用户使用, 需要使用某接口时, 可直接使用接口中文名, 或接口URL(驼峰格式)进行搜索, 找到对应的struct, demo等

1. 配置初始化
- 初始化为全局配置(多商户模式下, 可初始化多份)
- 第一个参数true:生产环境, false:测试环境
- 第二个参数商户公私钥等信息配置
```
dgSDK, _ := BsPaySdk.NewBsPay(false, "./config/config.json")
```

2. 组装请求参数
- 为了接口使用更加方便, 我们将参数粗分为必填/非必填, 必填直接放在结构体内, 非必填放在结构体的map字段ExtendInfos中
```
	dgReq := BsPaySdk.V2MerchantBasicdataQueryRequest{
		// 请求流水号
		ReqSeqId: tool.GetReqSeqId(),
		// 请求日期
		ReqDate: tool.GetCurrentDate(),
		// 汇付客户Id
		HuifuId: "6666000108854952",
	}
	// 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()
```

3. 发起API调用
```
resp, err := dgSDK.V2MerchantBasicdataQueryRequest(dgReq)
```

