package check

import (
	"encoding/xml"
	"testing"

	"github.com/waro163/wechat-message/mp"
)

func TestMassJobEvent(t *testing.T) {
	var res mp.MixMessage
	err := xml.Unmarshal([]byte(MassJobEventDemo), &res)
	if err != nil {
		t.Error("not suit for MassJobEventDemo type", err)
	}
	t.Log(res)
}
