/*
 * @Author: huaze.fan
 * @LastEditTime: 2022-11-14 16:21:44
 * @LastEditors: huaze.fan
 * @Description: Good Good Study！Day Day Up
 */
package BsPaySdk

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

/**
私钥加签
*/
func RsaSign(content string, msc *MerchSysConfig) (result string, err error) {
	privateKey, err := getPrivateKey(msc)
	if err != nil {
		BspayPrintln("error in get private key: " + err.Error())
		return "", err
	} else if privateKey == nil {
		BspayPrintln("error in get private key")
		return "", errors.New("error in get private key")
	}

	h := sha256.New()
	h.Write([]byte([]byte(content)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash[:])
	if err != nil {
		BspayPrintln("sign error: " + err.Error())
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)

	return out, nil
}

/**
公钥验签
*/
func RsaSignVerify(base64Sign string, content string, msc *MerchSysConfig) (bool, error) {
	pubKey, err := getPublicKey(msc)
	if err != nil {
		BspayPrintln("error in get public key: " + err.Error())
		return false, err
	} else if pubKey == nil {
		BspayPrintln("error in get public key")
		return false, errors.New("error in get public key")
	}

	// 反解base64
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false, err
	}

	hash := sha256.New()
	hash.Write([]byte(content))
	bytes := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, bytes, sign)
	if err != nil {
		BspayPrintln("check sign error: " + err.Error())
		return false, err
	}

	return true, nil
}

/**
获取配置中的rsa私钥转成对象
*/
func getPrivateKey(msc *MerchSysConfig) (interface{}, error) {
	priKeyString := msc.RsaMerchPrivateKey
	if priKeyString == "" {
		BspayPrintln("unknow private key data")
		return nil, errors.New("unknow private key data")
	}

	// 加头尾
	if !strings.Contains(priKeyString, "-----BEGIN RSA PRIVATE KEY-----") {
		priKeyString = "-----BEGIN RSA PRIVATE KEY-----\n" + priKeyString
	}
	if !strings.Contains(priKeyString, "-----END RSA PRIVATE KEY-----") {
		priKeyString = priKeyString + "\n-----END RSA PRIVATE KEY-----"
	}

	// 拦截处理 key 时可能的错误
	// 拦截到以后，两个返回参数都是空了
	defer func() {
		err := recover()
		if err != nil {
			BspayPrintln("error in get private key")
		}
	}()
	keyByts, _ := pem.Decode([]byte(priKeyString))

	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts.Bytes)
	if err != nil {
		BspayPrintln("error in get private key: " + err.Error())
		return nil, err
	}

	// 正常返回
	return privateKey, nil
}

/**
获取配置中的rsa公钥转成对象
*/
func getPublicKey(msc *MerchSysConfig) (interface{}, error) {
	pubKeyString := msc.RsaHuifuPublicKey
	if pubKeyString == "" {
		BspayPrintln("unknow public key data")
		return nil, errors.New("unknow public key data")
	}

	// 加头尾
	if !strings.Contains(pubKeyString, "-----BEGIN PUBLIC KEY-----") {
		pubKeyString = "-----BEGIN PUBLIC KEY-----\n" + pubKeyString
	}
	if !strings.Contains(pubKeyString, "-----END PUBLIC KEY-----") {
		pubKeyString = pubKeyString + "\n-----END PUBLIC KEY-----"
	}

	// 拦截处理 key 时可能的错误
	// 拦截到以后，两个返回参数都是空了
	defer func() {
		err := recover()
		if err != nil {
			BspayPrintln("error in get public key")
		}
	}()
	keyByts, _ := pem.Decode([]byte(pubKeyString))

	pubKey, err := x509.ParsePKIXPublicKey(keyByts.Bytes)
	if err != nil {
		BspayPrintln("error in get public key: " + err.Error())
		return nil, err
	}

	// 正常返回
	return pubKey, nil
}
