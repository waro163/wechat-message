package message

import "github.com/waro163/wechat-message/utils"

// News 图文消息
type News struct {
	CommonMsg

	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles>item,omitempty"`
}

// NewNews 初始化图文消息
func NewNews(articles []*Article) *News {
	news := new(News)
	news.ArticleCount = len(articles)
	news.Articles = articles
	news.SetMsgType(MsgTypeNews)
	news.SetCreateTime(utils.GetCurrTS())
	return news
}

// Article 单篇文章
type Article struct {
	Title       CDATA `xml:"Title"`
	Description CDATA `xml:"Description"`
	PicURL      CDATA `xml:"PicUrl"`
	URL         CDATA `xml:"Url"`
}

// NewArticle 初始化文章
func NewArticle(title, description, picURL, url string) *Article {
	article := new(Article)
	article.Title = CDATA(title)
	article.Description = CDATA(description)
	article.PicURL = CDATA(picURL)
	article.URL = CDATA(url)
	return article
}
