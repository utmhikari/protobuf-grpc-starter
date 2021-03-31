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
	ShortLink string  `json:"shortLink" form:"shortLink"`

	Author string  `json:"author" form:"author"`
	Keyword string `json:"keyword" form:"keyword"`
}


func (q *Query) Hash() string {
	if len(q.ShortLink) > 0 {
		return q.ShortLink
	}

	hs := fmt.Sprintf("Author:%s,Keyword:%s", q.Author, q.Keyword)
	m := md5.New()
	m.Write([]byte(hs))
	md5Str := hex.EncodeToString(m.Sum(nil))

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


func (d *Document) GenMeta() {
	if d != nil {
		// gen created
		d.Created = time.Now().UnixNano()

		// gen shortlink
		hs := fmt.Sprintf("time:%d_author:%s_content:%s_title:%s",
			d.Created, d.Author, d.Content, d.Title)
		m := md5.New()
		m.Write([]byte(hs))
		md5Str := hex.EncodeToString(m.Sum(nil))
		// fmt.Printf("%s -> %s\n", hs, md5Str)

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
