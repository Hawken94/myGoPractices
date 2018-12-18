package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	md5decrypt()

}

func add(a, b int32) (int32, bool) {
	return a + b, true
}

func md5decrypt() {
	pubKey := "30819c300d06092a864886f70d010101050003818a00" +
		"3081860281805fa728db4d2f0487ad96405d898d022ec2e71fc21c7c8b46cdd16edcbb2474709ed5355" +
		"d89272c505ea3c2628b5e69aa66b699c6afeb9ec7aeb0445851809d9511d66e55f8c9283df8359db" +
		"1cd069640be572177793f08e9a4cedace006a533a51c8119059d465cbc28619d5b0be6df3eaab" +
		"161223cdfad14c31653b991606bd020111"
	srcStr := "POSID=029942036&BRANCHID=370000000&ORDERID=607261&PAYMENT=0.01&CURCODE=01&REMARK1=58779&REMARK2=&ACC_TYPE=12&SUCCESS=Y&TYPE=1&REFERER=&CLIENTIP=11.168.64.169&ACCDATE=20181210"
	signStr := "2d7c3f1d487ec1331e0846e011dce607e15030cf7f98a2c36ae2fb62ae9a42a34a009ca62796f3983ccd68b053ccec9ff26b7ddb8c08c993c68f08f1653440762669f5350b0e3075c804fb98241575180613d69bee7142bdbb7d2d2c4921ba25d4c4e2c968485a9fa694fe0ae38e3fd0847d1140664f02fe9f1efeb5102c18fc"
	sign, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		fmt.Println("延签格式转换错误")
	}
	fmt.Println(string(sign))
	public, err := hex.DecodeString(pubKey)
	if err != nil {
		fmt.Println("证书格式错误,err=", err, public)
	}
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		fmt.Println("证书错误,err=", err)
	}
	t := md5.New()
	io.WriteString(t, srcStr)
	digest := t.Sum(nil)
	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.MD5, digest, sign)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
	}
}
