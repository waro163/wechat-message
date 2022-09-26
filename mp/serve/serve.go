package serve

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/waro163/wechat-message/mp"
)

type Server struct {
	Token        string
	SkipValidata bool
	Handler      func(mp.MixMessage) (interface{}, error)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

	//handle wechat event, just for not SAFT mode!!!
	var eventMsg mp.MixMessage
	rawXMLMsgBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = xml.Unmarshal(rawXMLMsgBytes, &eventMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if s.Handler == nil {
		w.Write([]byte("success"))
	}

	// process business logic by your handle function
	msgResp, err := s.Handler(eventMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// generate whole xml format message to response
	data, err := mp.ReplyMsg(eventMsg, msgResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
