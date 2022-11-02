package check

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waro163/wechat-message/mp"
	msg "github.com/waro163/wechat-message/mp/message"
)

func TestMassJobEvent(t *testing.T) {
	var res mp.MixMessage
	err := xml.Unmarshal([]byte(MassJobEventDemo), &res)
	assert.Nil(t, err)
	assert.Equal(t, msg.MsgTypeEvent, res.MsgType)
	assert.Equal(t, mp.EventMassSendJobFinish, res.Event)
	assert.NotNil(t, res.CopyrightCheckResult)
	assert.NotNil(t, res.ArticleUrlResult)
	// if err != nil {
	// 	t.Error("not suit for MassJobEventDemo type", err)
	// }
	// t.Log(res)
}
