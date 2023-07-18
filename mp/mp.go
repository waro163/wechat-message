package mp

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"

	msg "github.com/waro163/wechat-message/mp/message"
	"github.com/waro163/wechat-message/utils"
)

func Validate(req *http.Request, token string, skip bool) bool {
	if skip {
		return true
	}
	query := req.URL.Query()
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	signature := query.Get("signature")
	return signature == Signature(token, timestamp, nonce)
}

// MsgValidate 加密数据的签名校验
func MsgValidate(req *http.Request, token, msgEncrypt string, skip bool) bool {
	if skip {
		return true
	}
	query := req.URL.Query()
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	msgSignature := query.Get("msg_signature")
	return msgSignature == Signature(token, timestamp, nonce, msgEncrypt)
}

func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

type DecryptMsgData struct {
	Msg   []byte
	AppID string
}

// DecryptMsg 解密安全模式的加密数据
func DecryptMsg(encryptMsg string, aesKey []byte) (*DecryptMsgData, error) {
	aesMsg, err := base64.StdEncoding.DecodeString(encryptMsg)
	if err != nil {
		return nil, err
	}
	buf, err := utils.AesDecrypt(aesMsg, aesKey)
	if err != nil {
		return nil, err
	}
	var msgLen int32
	binary.Read(bytes.NewBuffer(buf[16:20]), binary.BigEndian, &msgLen)
	// TODO: check msgLen
	msgData := buf[20 : 20+msgLen]
	appID := buf[20+msgLen:]
	return &DecryptMsgData{
		Msg:   msgData,
		AppID: string(appID),
	}, nil
}

// EncryptMsg 对被动回复的消息进行加密并签名返回
func EncryptMsg(rawMsg, aesKey []byte, appID, token string) ([]byte, error) {
	random := make([]byte, 16)
	if _, err := rand.Read(random); err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, uint32(len(rawMsg))); err != nil {
		return nil, err
	}
	len := buf.Bytes()

	plainText := bytes.Join([][]byte{random, len, rawMsg, []byte(appID)}, nil)
	encryptMsg, err := utils.AesEncrypt(plainText, aesKey)
	if err != nil {
		return nil, err
	}
	encryptMsgData := base64.StdEncoding.EncodeToString(encryptMsg)
	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	resp := &EncryptedRespXMLMsg{
		Encrypt:      msg.CDATA(encryptMsgData),
		TimeStamp:    timeStamp,
		Nonce:        msg.CDATA(random),
		MsgSignature: msg.CDATA(Signature(token, timeStamp, string(random), encryptMsgData)),
	}
	return xml.Marshal(resp)
}
