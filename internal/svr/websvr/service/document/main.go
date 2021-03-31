package document

import (
	"context"
	"errors"
	"github.com/utmhikari/protobuf-grpc-starter/api/pb/cache"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/models"
	cacheService "github.com/utmhikari/protobuf-grpc-starter/internal/svr/websvr/service/cache"
	"github.com/utmhikari/protobuf-grpc-starter/internal/svr/websvr/service/db"
	"log"
	"strings"
)

func GetDocumentByShortLink(shortLink string) (*models.Document, error) {
	log.Printf("get doc by short-link -> %s\n", shortLink)
	if "" == shortLink {
		return nil,  errors.New("empty shortLink")
	}

	// find by cache
	resp, err := cacheService.GetClient().GetDocument(
		context.Background(), &cache.GetDocumentRequest{ShortLink: shortLink})
	switch {
	case err != nil:
		log.Printf("%s get cache failed, internal error: %s\n", shortLink, err.Error())
		break
	case resp == nil:
		log.Printf("%s get cache failed, nil response\n", shortLink)
		break
	case !resp.Status.Success:
		log.Printf("%s get cache failed, error: %s\n", shortLink, resp.Status.Message)
		break
	default:
		log.Printf("%s get cache success\n", shortLink)
		return models.NewDocumentPromProto(resp.Document), nil
	}

	// find from db
	log.Printf("failed to get document %s from cache, try find from db...\n", shortLink)
	docs, numDocs, err := db.GetAllDocs()
	if err != nil {
		return nil, err
	}

	for i := 0; i < numDocs; i++ {
		doc := &(*docs)[i]
		if shortLink == doc.ShortLink {
			// save to cache
			setResp, setErr := cacheService.GetClient().SetDocument(
				context.Background(), &cache.SetDocumentRequest{Document: doc.ToProto()})
			switch {
			case setErr != nil:
				log.Printf("%s set cache failed, internal error: %s\n", shortLink, setErr.Error())
				break
			case setResp == nil:
				log.Printf("%s set cache failed, nil response\n", shortLink)
				break
			case !setResp.Status.Success:
				log.Printf("%s set cache failed, error: %s\n", shortLink, setResp.Status.Message)
				break
			default:
				log.Printf("%s set cache success\n", shortLink)
			}
			return doc, nil
		}
	}

	return nil, errors.New("cannot find doc with short-link " + shortLink)
}


func GetDocumentsByQuery(query *models.Query) ([]*models.Document, error) {
	log.Printf("get docs by query -> %+v\n", query)
	if nil == query || query.IsEmpty() {
		return nil, errors.New("empty query")
	}

	if query.IsShortLinkQuery() {
		doc, err := GetDocumentByShortLink(query.ShortLink)
		if err != nil {
			return nil, err
		}
		return []*models.Document{doc}, nil
	}

	docs, numDocs, err := db.GetAllDocs()
	if err != nil {
		return nil, err
	}

	documents := make([]*models.Document, 0)
	for i := 0; i < numDocs; i++ {
		doc := &(*docs)[i]

		if "" != query.Author {
			if doc.Author != query.Author {
				continue
			}
		}

		if "" != query.Keyword {
			if !strings.Contains(doc.Title, query.Keyword) &&
				!strings.Contains(doc.Content, query.Keyword) &&
				!strings.Contains(doc.Author, query.Keyword) {
				continue
			}
		}

		documents = append(documents, doc)
	}

	return documents, nil
}


func CreateDocument(doc *models.Document) error {
	if nil == doc {
		return errors.New("empty doc")
	}

	doc.GenMeta()
	log.Printf("create doc -> %+v\n", doc)

	return db.Save(doc)
}
