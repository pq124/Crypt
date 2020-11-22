package base

import "encoding/base64"

//使用base64进行编码
func Base64Encode(data []byte)(string){

    encoding:=base64.StdEncoding
    dst:=make([]byte,0)
    encoding.Decode(dst,data)
    return string(dst)
	//encoding:=base64.StdEncoding
	//dst:=make([]byte,0)
	//encoding.Encode(dst,data)
	////encoding.EncodeToString()
	//return string(dst)
}
func Base64Decode(data string)string  {
	//encoding:=base64.StdEncoding
	//dst:=make([]byte,0)
	//encoding.Decode(dst,[]byte(data))
	//return string(dst)

	encoding:=base64.StdEncoding
	dst:=make([]byte,0)
	encoding.Decode(dst,[]byte(data))
	return string(dst)

}