package check

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/waro163/wechat-message/mp"
)

func TestSubscribeStruct(t *testing.T) {
	var res1, res2, res3 mp.MixMessage
	err := xml.Unmarshal([]byte(SubscribeMsgDemo1), &res1)
	if err != nil {
		t.Error("not suit for SubscribeMsgDemo1 type", err)
	}
	fmt.Println(res1)
	err = xml.Unmarshal([]byte(SubscribeMsgDemo2), &res2)
	if err != nil {
		t.Error("not suit for SubscribeMsgDemo2 type", err)
	}
	fmt.Println(res2)
	err = xml.Unmarshal([]byte(SubscribeMsgDemo3), &res3)
	if err != nil {
		t.Error("not suit for SubscribeMsgDemo3 type", err)
	}
	fmt.Println(res3)
	t.Log("MixMessage satisfy all Subscribe message struct")
}
