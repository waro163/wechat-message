### 概述

该仓库主要是用来解析微信公众号和微信小程序事件推送消息的解析

其中微信公众号事件推送消息主要来自于

- [用户普通消息](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html)
- [事件推送](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_event_pushes.html)
- [菜单事件推送](https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Custom_Menu_Push_Events.html)
- [订阅通知事件推送](https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#%E4%BA%8B%E4%BB%B6%E6%8E%A8%E9%80%81)
- [模版消息事件推送](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#%E4%BA%8B%E4%BB%B6%E6%8E%A8%E9%80%81)
- [发布能力事件推送](https://developers.weixin.qq.com/doc/offiaccount/Publish/Callback_on_finish.html)
- [群发消息事件推送](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#7)
- [卡券事件推送](https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons_Vouchers_and_Cards_Event_Push_Messages.html)
- [用户授权信息变更推送](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/authorization_change.html)

小程序事件推送主要来自于

- [订阅消息事件推送](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/subscribe-message.html)
- [客服消息事件](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/customer-message/receive.html)

### 必要性

本仓库部分代码来源于[wechat](https://github.com/silenceper/wechat).

但是上述仓库中事件推送消息解析部分不对外暴漏，且耦合过于紧密，难以单独拆出解析部分的代码进行使用。对于内部服务来说，只需要解析即可，为了防止重复开发因此有此仓库。

### 代码结构

```
|-mp/ 微信公众号事件消息解析
|  |-check/ 各种消息格式示例
|  |-message/ 回复消息结构
|  |-serve/ ServeHTTP方法
|  |  |-serve.go 明文数据
|  |  |-safe-serve.go 密文数据
|  |-event.go 微信事件格式
|  |-message.go 微信消息格式
|  |-mp.go 签名方法和加解密方法
|  |-reply.go 回复方法
|
|-mini/ 微信小程序事件消息解析
|  |-
|
|-utils/ 工具方法
|  |-crypto.go aes cbc加解密
```
