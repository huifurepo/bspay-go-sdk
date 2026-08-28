/**
 * @description: SM2国密算法签名、验签及密钥管理
 * @author: ivan.mao
 * @date: 2026/7/31 14:31
 */
package BsPaySdk

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

// Sm2Sign 使用SM2私钥对内容进行签名
//
// 签名流程:
//  1. 从商户配置中解析SM2私钥
//  2. 使用私钥对原文进行SM2签名
//  3. 将签名结果编码为十六进制字符串返回
//
// 参数:
//   - content: 待签名的原文（通常为hex编码后的字符串，与验签端保持一致）
//   - msc:     商户系统配置，需包含Sm2MerchPrivateKey字段
//
// 返回:
//   - string: 签名值的十六进制字符串
//   - error:  签名过程中的错误信息
func Sm2Sign(content string, msc *MerchSysConfig) (string, error) {
	privateKey, err := getSm2PrivateKey(msc)
	if err != nil {
		BspayPrintln("error in get sm2 private key: " + err.Error())
		return "", err
	}
	if privateKey == nil {
		BspayPrintln("error in get sm2 private key: key is nil")
		return "", errors.New("sm2 private key is nil")
	}

	// 使用SM2私钥对内容进行签名
	// tjfoc/gmsm库中 *sm2.PrivateKey 的 Sign 方法始终使用SM2算法，
	// opts 传 nil 时使用默认UID "1234567812345678"，与验签端保持一致
	signature, err := privateKey.Sign(rand.Reader, []byte(content), nil)
	if err != nil {
		BspayPrintln("sm2 sign error: " + err.Error())
		return "", err
	}

	// 返回签名的十六进制字符串
	return hex.EncodeToString(signature), nil
}

// Sm2SignVerify 使用SM2公钥验证签名
//
// 验签流程:
//  1. 从商户配置中解析SM2公钥
//  2. 将hex编码的签名值解码为原始字节
//  3. 使用公钥验证签名
//
// 参数:
//   - content: 原文内容（与签名时传入的content一致）
//   - sign:    签名值的十六进制字符串
//   - msc:     商户系统配置，需包含Sm2HuifuPublicKey字段
//
// 返回:
//   - bool:  验签结果，true表示验签通过
//   - error: 验签过程中的错误信息（注意：验签失败不返回error，而是返回false）
func Sm2SignVerify(content string, sign string, msc *MerchSysConfig) (bool, error) {
	publicKey, err := getSm2PublicKey(msc)
	if err != nil {
		BspayPrintln("error in get sm2 public key: " + err.Error())
		return false, err
	}
	if publicKey == nil {
		BspayPrintln("error in get sm2 public key: key is nil")
		return false, errors.New("sm2 public key is nil")
	}

	// 将hex编码的签名值解码为原始字节，与Sm2Sign中的hex编码对应
	signBytes, err := hex.DecodeString(sign)
	if err != nil {
		BspayPrintln("sm2 verify sign decode error: " + err.Error())
		return false, err
	}

	// 使用SM2公钥验证签名
	return publicKey.Verify([]byte(content), signBytes), nil
}

// getSm2PrivateKey 从商户配置中解析SM2私钥
//
// 私钥以Base64编码的PKCS8格式存储在配置中，
// 本方法负责解码并解析为 *sm2.PrivateKey 对象。
// 使用 defer+recover 拦截解析过程中可能的panic（如格式错误）。
//
// 参数:
//   - msc: 商户系统配置
//
// 返回:
//   - *sm2.PrivateKey: SM2私钥对象
//   - error: 解析过程中的错误信息
func getSm2PrivateKey(msc *MerchSysConfig) (key *sm2.PrivateKey, err error) {
	priKeyString := msc.Sm2MerchPrivateKey
	if priKeyString == "" {
		BspayPrintln("unknow sm2 private key data")
		return nil, errors.New("unknow sm2 private key data")
	}

	// 拦截解析过程中可能的panic（如格式错误导致的底层panic）
	// 使用命名返回值，确保panic时也能返回错误信息
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error in get sm2 private key: %v", r)
			BspayPrintln(err.Error())
			key = nil
		}
	}()

	// 私钥配置使用base64编码存储，需要先解码为原始字节
	privateKeyData, err := base64.StdEncoding.DecodeString(priKeyString)
	if err != nil {
		BspayPrintln("error in decode sm2 private key: " + err.Error())
		return nil, err
	}

	// 解析PKCS8格式的私钥数据，返回SM2私钥对象
	key, err = x509.ParsePKCS8PrivateKey(privateKeyData, nil)
	if err != nil {
		BspayPrintln("error in get sm2 private key: " + err.Error())
		return nil, err
	}

	return key, nil
}

// getSm2PublicKey 从商户配置中解析SM2公钥
//
// 公钥以Base64编码的PKIX格式存储在配置中，
// 本方法负责解码并解析为 *sm2.PublicKey 对象。
// 由于 x509.ParsePKIXPublicKey 返回 *ecdsa.PublicKey，
// 需要手动转换为 *sm2.PublicKey（使用国密P256Sm2曲线）。
// 使用 defer+recover 拦截解析过程中可能的panic（如格式错误）。
//
// 参数:
//   - msc: 商户系统配置
//
// 返回:
//   - *sm2.PublicKey: SM2公钥对象
//   - error: 解析过程中的错误信息
func getSm2PublicKey(msc *MerchSysConfig) (key *sm2.PublicKey, err error) {
	pubKeyString := msc.Sm2HuifuPublicKey
	if pubKeyString == "" {
		BspayPrintln("unknow sm2 public key data")
		return nil, errors.New("unknow sm2 public key data")
	}

	// 拦截解析过程中可能的panic（如格式错误导致的底层panic）
	// 使用命名返回值，确保panic时也能返回错误信息
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error in get sm2 public key: %v", r)
			BspayPrintln(err.Error())
			key = nil
		}
	}()

	// 公钥配置使用base64编码存储，需要先解码为原始字节
	publicKeyData, err := base64.StdEncoding.DecodeString(pubKeyString)
	if err != nil {
		BspayPrintln("error in decode sm2 public key: " + err.Error())
		return nil, err
	}

	// 解析PKIX格式的公钥数据，返回的是 *ecdsa.PublicKey 类型
	parsedKey, err := x509.ParsePKIXPublicKey(publicKeyData)
	if err != nil {
		BspayPrintln("error in get sm2 public key: " + err.Error())
		return nil, err
	}

	// 类型断言：确认解析出的是 *ecdsa.PublicKey
	ecdsaKey, ok := parsedKey.(*ecdsa.PublicKey)
	if !ok {
		BspayPrintln("error: key is not ecdsa public key")
		return nil, errors.New("key is not ecdsa public key")
	}

	// 从 *ecdsa.PublicKey 构造 *sm2.PublicKey
	// SM2使用国密P256Sm2曲线，与标准ECDSA的曲线不同
	// X、Y为公钥的坐标点，使用 new(big.Int).Set 深拷贝以避免引用原始对象
	key = &sm2.PublicKey{
		Curve: sm2.P256Sm2(), // SM2国密曲线
		X:     new(big.Int).Set(ecdsaKey.X),
		Y:     new(big.Int).Set(ecdsaKey.Y),
	}

	return key, nil
}

// Sm2KeyPair SM2密钥对，包含Base64编码的私钥和公钥
type Sm2KeyPair struct {
	PrivateKey string // Base64编码的PKCS8格式私钥
	PublicKey  string // Base64编码的PKIX格式公钥
}

// GenerateSm2KeyPair 生成SM2密钥对
//
// 使用国密SM2算法随机生成一对公私钥，
// 私钥以PKCS8格式编码，公钥以PKIX格式编码，
// 最终均以Base64字符串形式返回，便于配置存储。
//
// 返回:
//   - *Sm2KeyPair: 包含Base64编码公私钥的密钥对
//   - error: 生成过程中的错误信息
func GenerateSm2KeyPair() (*Sm2KeyPair, error) {
	// 使用密码学安全的随机数源生成SM2私钥
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		BspayPrintln("generate sm2 key pair error: " + err.Error())
		return nil, err
	}

	// 将私钥序列化为PKCS8格式
	privateKeyBytes, err := x509.MarshalSm2PrivateKey(privateKey, nil)
	if err != nil {
		BspayPrintln("marshal sm2 private key error: " + err.Error())
		return nil, err
	}

	// 将公钥序列化为PKIX格式
	publicKeyBytes, err := x509.MarshalSm2PublicKey(&privateKey.PublicKey)
	if err != nil {
		BspayPrintln("marshal sm2 public key error: " + err.Error())
		return nil, err
	}

	// 返回Base64编码的密钥对
	return &Sm2KeyPair{
		PrivateKey: base64.StdEncoding.EncodeToString(privateKeyBytes),
		PublicKey:  base64.StdEncoding.EncodeToString(publicKeyBytes),
	}, nil
}
