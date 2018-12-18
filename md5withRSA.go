package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
)

func main() {
	srcStr := "POSID=029942036&BRANCHID=370000000&ORDERID=607261&PAYMENT=0.01&CURCODE=01&REMARK1=58779&REMARK2=&ACC_TYPE=12&SUCCESS=Y&TYPE=1&REFERER=&CLIENTIP=11.168.64.169&ACCDATE=20181210"
	signMD5WithRSA(data []byte, priv *rsa.PrivateKey)
}

func signMD5WithRSA(data []byte, priv *rsa.PrivateKey) ([]byte, error) {
	return rsa.SignPKCS1v15(nil, priv, crypto.MD5, md5.Sum(data))
}
func decryptMD5WithRSA() {

}
