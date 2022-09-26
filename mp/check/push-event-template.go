package check

// 发送模版消息成功推送
var SendTemplateMsgSucc = `<xml> 
<ToUserName><![CDATA[gh_7f083739789a]]></ToUserName>  
<FromUserName><![CDATA[oia2TjuEGTNoeX76QEjQNrcURxG8]]></FromUserName>  
<CreateTime>1395658920</CreateTime>  
<MsgType><![CDATA[event]]></MsgType>  
<Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>  
<MsgID>200163836</MsgID>  
<Status><![CDATA[success]]></Status> 
</xml>`

// 用户拒绝接受模版消息，用户设置拒绝接收公众号消息
var UserRejectTemplate = `<xml> 
<ToUserName><![CDATA[gh_7f083739789a]]></ToUserName>  
<FromUserName><![CDATA[oia2TjuEGTNoeX76QEjQNrcURxG8]]></FromUserName>  
<CreateTime>1395658984</CreateTime>  
<MsgType><![CDATA[event]]></MsgType>  
<Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>  
<MsgID>200163840</MsgID>  
<Status><![CDATA[failed:user block]]></Status> 
</xml>`

// 其他原因失败
var SendTemplateFail = `<xml> 
<ToUserName><![CDATA[gh_7f083739789a]]></ToUserName>  
<FromUserName><![CDATA[oia2TjuEGTNoeX76QEjQNrcURxG8]]></FromUserName>  
<CreateTime>1395658984</CreateTime>  
<MsgType><![CDATA[event]]></MsgType>  
<Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>  
<MsgID>200163840</MsgID>  
<Status><![CDATA[failed: system failed]]></Status> 
</xml>`
