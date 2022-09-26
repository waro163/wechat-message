package check

var SubscribeMsgDemo1 = `
<xml>
    <ToUserName><![CDATA[gh_123456789abc]]></ToUserName>
    <FromUserName><![CDATA[otFpruAK8D-E6EfStSYonYSBZ8_4]]></FromUserName>
    <CreateTime>1610969440</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe_msg_popup_event]]></Event>
    <SubscribeMsgPopupEvent>
        <List>
            <TemplateId><![CDATA[VRR0UEO9VJOLs0MHlU0OilqX6MVFDwH3_3gz3Oc0NIc]]></TemplateId>
            <SubscribeStatusString><![CDATA[accept]]></SubscribeStatusString>
            <PopupScene>2</PopupScene>
        </List>
        <List>
            <TemplateId><![CDATA[9nLIlbOQZC5Y89AZteFEux3WCXRRRG5Wfzkpssu4bLI]]></TemplateId>
            <SubscribeStatusString><![CDATA[reject]]></SubscribeStatusString>
            <PopupScene>2</PopupScene>
        </List>
    </SubscribeMsgPopupEvent>
</xml>
`
var SubscribeMsgDemo2 = `
<xml>
    <ToUserName><![CDATA[gh_123456789abc]]></ToUserName>
    <FromUserName><![CDATA[otFpruAK8D-E6EfStSYonYSBZ8_4]]></FromUserName>
    <CreateTime>1610969440</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe_msg_change_event]]></Event>
    <SubscribeMsgChangeEvent>
        <List>
            <TemplateId><![CDATA[VRR0UEO9VJOLs0MHlU0OilqX6MVFDwH3_3gz3Oc0NIc]]></TemplateId>
            <SubscribeStatusString><![CDATA[reject]]></SubscribeStatusString>
        </List>
    </SubscribeMsgChangeEvent>
</xml>
`

var SubscribeMsgDemo3 = `
<xml>
    <ToUserName><![CDATA[gh_123456789abc]]></ToUserName>
    <FromUserName><![CDATA[otFpruAK8D-E6EfStSYonYSBZ8_4]]></FromUserName>
    <CreateTime>1610969468</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe_msg_sent_event]]></Event>
    <SubscribeMsgSentEvent>
        <List>
            <TemplateId><![CDATA[VRR0UEO9VJOLs0MHlU0OilqX6MVFDwH3_3gz3Oc0NIc]]></TemplateId>
            <MsgID>1700827132819554304</MsgID>
            <ErrorCode>0</ErrorCode>
            <ErrorStatus><![CDATA[success]]></ErrorStatus>
        </List>
    </SubscribeMsgSentEvent>
</xml>`
