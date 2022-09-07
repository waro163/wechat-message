### 概述
该仓库主要是用来解析微信公众号和微信小程序事件推送消息的解析

其中微信公众号事件推送消息主要来自于
- [用户普通消息](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html)
- [事件推送](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_event_pushes.html)
- [菜单事件推送](https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Custom_Menu_Push_Events.html)

小程序事件推送主要来自于
- [订阅消息事件推送](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/subscribe-message.html)
- [客服消息事件](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/customer-message/receive.html)

### 必要性

本仓库部分代码来源于[wechat](https://github.com/silenceper/wechat).

但是上述仓库中事件推送消息解析部分不对外暴漏，且耦合过于紧密，难以单独拆出解析部分的代码进行使用。对于内部服务来说，只需要解析即可，为了防止重复开发因此有此仓库。

### 代码结构
```
|-mp/ 微信公众号事件消息解析
|  |-
|
|-mini/ 微信小程序事件消息解析
|  |-
```