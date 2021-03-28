package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)


const MaxShortLinkLength = 8


type Query struct {
	ShortLink string  `json:"shortLink"`

	Author string  `json:"author"`
	Keyword string `json:"keyword"`
}


func (q *Query) Hash() string {
	if len(q.ShortLink) > 0 {
		return q.ShortLink
	}

	hs := fmt.Sprintf("Author:%s,Keyword:%s", q.Author, q.Keyword)
	md5Str := hex.EncodeToString(md5.New().Sum([]byte(hs)))

	return md5Str
}


type Document struct {
	Title string `json:"title"`
	Content string  `json:"content"`
	Author string  `json:"author"`

	ShortLink string  `json:"shortLink"`
	Created int64  `json:"created"`
}


func (d *Document) GenShortLink() {
	if d != nil {
		hs := fmt.Sprintf("time:%d,author:%s,content:%s,title:%s",
			d.Created, d.Author, d.Content, d.Title)
		md5Str := hex.EncodeToString(md5.New().Sum([]byte(hs)))
		if len(md5Str) > MaxShortLinkLength {
			md5Str = md5Str[:MaxShortLinkLength]
		}
		d.ShortLink = md5Str
	}
}


func NewDocument(content string, author string) *Document {
	nowTimeStamp := time.Now().UnixNano()
	doc := Document{
		Content:      content,
		Author:       author,
		ShortLink:    "",
		Created:      nowTimeStamp,
	}
	doc.GenShortLink()

	return &doc
}
