package check

var PublishEventSuccess = `
<xml> 
  <ToUserName><![CDATA[gh_4d00ed8d6399]]></ToUserName>  
  <FromUserName><![CDATA[oV5CrjpxgaGXNHIQigzNlgLTnwic]]></FromUserName>  
  <CreateTime>1481013459</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[PUBLISHJOBFINISH]]></Event>
  <PublishEventInfo>
    <publish_id>2247503051</publish_id>
    <publish_status>0</publish_status>
    <article_id><![CDATA[b5O2OUs25HBxRceL7hfReg-U9QGeq9zQjiDvy
WP4Hq4]]></article_id>
    <article_detail>
      <count>1</count>
      <item>
        <idx>1</idx>
        <article_url><![CDATA[ARTICLE_URL]]></article_url>
      </item>
    </article_detail>
  </PublishEventInfo>
</xml>
`

var PublishEventReject = `
<xml> 
  <ToUserName><![CDATA[gh_4d00ed8d6399]]></ToUserName>  
  <FromUserName><![CDATA[oV5CrjpxgaGXNHIQigzNlgLTnwic]]></FromUserName>  
  <CreateTime>1481013459</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[PUBLISHJOBFINISH]]></Event>
  <PublishEventInfo>
    <publish_id>2247503051</publish_id>
    <publish_status>2</publish_status>
    <fail_idx>1</fail_idx>
    <fail_idx>2</fail_idx>
  </PublishEventInfo>
</xml>`
