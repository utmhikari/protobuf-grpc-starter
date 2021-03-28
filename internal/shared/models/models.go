package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	pb "github.com/utmhikari/protobuf-grpc-starter/api/pb/base"
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

func (q *Query) IsEmpty() bool {
	if nil == q {
		return true
	}
	return "" == q.ShortLink && "" == q.Author && "" == q.Keyword
}


func (q *Query) IsShortLinkQuery() bool {
	return nil != q && q.ShortLink != ""
}


type Document struct {
	Title string `json:"title"`
	Content string  `json:"content" binding:"required"`
	Author string  `json:"author"`

	ShortLink string  `json:"shortLink"`
	Created int64  `json:"created"`
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


func NewDocumentPromProto(pbDoc *pb.Document) *Document {
	if nil == pbDoc {
		return nil
	}

	return &Document{
		Title: pbDoc.GetTitle(),
		Content: pbDoc.GetContent(),
		Author: pbDoc.GetAuthor(),
		ShortLink: pbDoc.GetShortLink(),
		Created: pbDoc.GetCreated(),
	}
}


func (d *Document) ToProto() *pb.Document{
	if nil == d {
		return nil
	}

	return &pb.Document{
		Title: d.Title,
		Content: d.Content,
		Author: d.Author,
		ShortLink: d.ShortLink,
		Created: d.Created,
	}
}
