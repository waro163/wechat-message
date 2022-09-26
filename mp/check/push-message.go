package check

// 普通消息推送实例

var TextDemo = `<xml>
<ToUserName><![CDATA[gh_d049d0472d74]]></ToUserName>
<FromUserName><![CDATA[oJlus55n0z6moe2LXDnZWMJkaq7I]]></FromUserName>
<CreateTime>1662522877</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[你好]]></Content>
<MsgId>23801603303036899</MsgId>
</xml> `

var ImageDemo = `<xml>
<ToUserName><![CDATA[gh_d049d0472d74]]></ToUserName>
<FromUserName><![CDATA[oJlus55n0z6moe2LXDnZWMJkaq7I]]></FromUserName>
<CreateTime>1662522888</CreateTime>
<MsgType><![CDATA[image]]></MsgType>
<PicUrl><![CDATA[http://mmbiz.qpic.cn/mmbiz_jpg/mBp0rAXmfrDWxRPzUSBSjgKwEicKIibicqKxzI4f6G9ibUhrPKnxvjJkiaTZXicepGlDXGvmN3hAku9VJHERCGiad7ZkA/0]]></PicUrl>
<MsgId>23801602667960890</MsgId>
<MediaId><![CDATA[AQrb7Sk-JYMgDd0WK2IHJ6eMp2AgOyGwji3MQKDe4axZDcGDxGxtA4PzIsvtrMha]]></MediaId>
</xml> `

var VoiceDemo = `<xml>
<ToUserName><![CDATA[gh_d049d0472d74]]></ToUserName>
<FromUserName><![CDATA[oJlus55n0z6moe2LXDnZWMJkaq7I]]></FromUserName>
<CreateTime>1662523726</CreateTime>
<MsgType><![CDATA[voice]]></MsgType>
<MediaId><![CDATA[AQrb7Sk-JYMgDd0WK2IHJwlpwqt32H0dk2oY_PGq-H-PiYf57hYdZgDUQGH3j2Ff]]></MediaId>
<Format><![CDATA[amr]]></Format>
<MsgId>23801614092693730</MsgId>
<Recognition><![CDATA[]]></Recognition>
</xml> `

var VideoDemo = `<xml>
<ToUserName><![CDATA[gh_d049d0472d74]]></ToUserName>
<FromUserName><![CDATA[oJlus55n0z6moe2LXDnZWMJkaq7I]]></FromUserName>
<CreateTime>1662523938</CreateTime>
<MsgType><![CDATA[video]]></MsgType>
<MediaId><![CDATA[AQrb7Sk-JYMgDd0WK2IHJ87BCZkq8fPUzx5qObTYnPESTWmyu-V_ALkNIMSly-AWEUV-x9AQmrPoPoeeUgY41w]]></MediaId>
<ThumbMediaId><![CDATA[AQrb7Sk-JYMgDd0WK2IHJwL-365_uSgadp8JEhwn6Cy0spFV_2rCPV1otIa68f6y]]></ThumbMediaId>
<MsgId>23801618750853655</MsgId>
</xml> `

var LocationDemo = `<xml>
<ToUserName><![CDATA[gh_d049d0472d74]]></ToUserName>
<FromUserName><![CDATA[oJlus55n0z6moe2LXDnZWMJkaq7I]]></FromUserName>
<CreateTime>1662524003</CreateTime>
<MsgType><![CDATA[location]]></MsgType>
<Location_X>39.909092</Location_X>
<Location_Y>116.481995</Location_Y>
<Scale>15</Scale>
<Label><![CDATA[朝阳区建国路77号]]></Label>
<MsgId>23801619415690638</MsgId>
</xml> `

var LinkDemo = `<xml>
<ToUserName><![CDATA[gh_d049d0472d74]]></ToUserName>
<FromUserName><![CDATA[oJlus55n0z6moe2LXDnZWMJkaq7I]]></FromUserName>
<CreateTime>1662538701</CreateTime>
<MsgType><![CDATA[link]]></MsgType>
<Title><![CDATA[B站云原生混部技术实践]]></Title>
<Description><![CDATA[本文介绍B站云原生混部技术实践]]></Description>
<Url><![CDATA[http://mp.weixin.qq.com/s?__biz=MzAwMDU1MTE1OQ==&mid=2653560636&idx=1&sn=eb6db387f0c791da9f66561297fb85a9&chksm=81398ea4b64e07b260344be83b19cfae56d9e135b31c7277a331fec340806524a911ee4da376&mpshare=1&scene=24&srcid=0906x71DMWDjCKHG6iwd2AGp&sharer_sharetime=1662442429205&sharer_shareid=7e6b6bcdfcdea5264ec06f9d884dc099#rd]]></Url>
<MsgId>23801831802242706</MsgId>
</xml> `
