package serve

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/waro163/wechat-message/mp"
)

/*
doc: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Message_encryption_and_decryption_instructions.html
dev test tool: https://developers.weixin.qq.com/apiExplorer?type=messagePush

api: path/api?signature=7f4c9e837603ad0814f9ede695e44c8bf82399b9
			&timestamp=1689578973&nonce=1553478527
			&encrypt_type=aes
			&msg_signature=6f8ff384b42497776b25301f04cd8a56155318a4

body:
	<xml>
    <ToUserName><![CDATA[gh_ae4671e1af0a]]></ToUserName>
    <Encrypt><![CDATA[tsA12LAbVNAenzN4ZSdNFJlr+Gw7XLZ7csLQTU7KT7BiPOERRvFssDWH9XSLko404kM5sHIxVdKrrL1lUm6la+lz2NqUspmbgXNGVSN39M30u00JU7CQWWIjfqBlu8uUVfhtDrZooCUNY/29WW8W/YwRuGraVNJGW+9aMjuJJ2okND4bChOAKHLdduN+8ljvDazIvMJqmrArbELp1PcPR4EH+B9IKZuvibRpWWN+cA2wbMmM4gwBmaadpgFTmivSKLoAZ/NLx/er2q+T+d/SMxmOx0t2TT3iuhTTWMc89JWohSImd47Npm8zNdXKF0jzCsFcz83yfTjsgt2iHYYZe9I1fU9ERa1Pa5smMUHdxqDnwz2KFrh8enwXozKkYW9WTNepOakmy2IK+BrBnsVVykqBQRQdoEi2cmLa7y0xjv9cS7WIAUdaE6QS5VTvJKmFanqaXqXcfrYIfpDUt8DmbmSsOGeQZPrCATGIYHEigB9DIO1NhER8V6c6I/6vbRWq]]></Encrypt>
	</xml>

signature:
	signature=sha1(sort(Token, timestamp, nonce))
msg_signature:
	msg_signature=sha1(sort(Token, timestamp, nonce, Encrypt))

encrypt:
	AESKey = Base64_Decode(EncodingAESKey + "=");
	FullStr = random(16B) + msg_len(4B) + msg + appid;
	msg_encrypt = Base64_Encode( AES_Encrypt( FullStr, AESKey ) );

decrypt:
	AESKey = Base64_Decode(EncodingAESKey + "=");
	TmpMsg = Base64_Decode(Encrypt)
	FullStr = AES_Decrypt(TmpMsg, AESKey);  FullStr 如前所述由4部分组成（random, msg_len, msg, appid）
	验证尾部的appid 是否正确（可选）
	去掉FullStr头部16字节的random、4字节的msg_len、和尾部的appid，即得到明文内容
*/

type SafeServer struct {
	Token        string
	AppID        string
	aesKey       []byte
	SkipValidata bool
	LogMsg       bool
	Handler      IMessageHandler
}

func (s *SafeServer) SetAESKey(key string) error {
	aesKey, err := base64.StdEncoding.DecodeString(key + "=")
	if err != nil {
		return err
	}
	s.aesKey = aesKey
	return nil
}

func (s *SafeServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// check signature
	if !mp.Validate(req, s.Token, s.SkipValidata) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'err_msg':'error signature'}"))
		return
	}
	query := req.URL.Query()
	// this is a bind url reqeust, just return echostr
	echostr := query.Get("echostr")
	if echostr != "" {
		w.Write([]byte(echostr))
		return
	}

	//handle wechat event, just for SAFE mode!!!
	var eventMsg mp.EncryptedXMLMsg
	if err := xml.NewDecoder(req.Body).Decode(&eventMsg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"errmsg":%s}`, err)))
		return
	}

	// check msg signature
	if !mp.MsgValidate(req, s.Token, eventMsg.EncryptedMsg, s.SkipValidata) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'err_msg':'error msg signature'}"))
		return
	}

	msgData, err := mp.DecryptMsg(eventMsg.EncryptedMsg, s.aesKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'err_msg':'error decrypt msg'}"))
		return
	}
	// check app_id
	if !s.SkipValidata {
		if msgData.AppID != s.AppID {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{'err_msg':'error app_id in msg'}"))
			return
		}
	}

	if s.Handler == nil {
		w.Write(nil)
		return
	}

	var mixMsg mp.MixMessage
	if err := xml.Unmarshal(msgData.Msg, &mixMsg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"errmsg":%s}`, err)))
		return
	}

	// process business logic by your handle function
	msgResp, err := s.Handler.HandleMsg(mixMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"errmsg":%s}`, err)))
		return
	}
	// generate whole xml format message to response
	data, err := mp.ReplyMsg(mixMsg, msgResp)
	if err != nil {
		w.Write(nil)
		return
	}
	encryptData, err := mp.EncryptMsg(data, s.aesKey, s.AppID, s.Token)
	if err != nil {
		return
	}
	w.Write(encryptData)
	w.WriteHeader(http.StatusOK)
}
