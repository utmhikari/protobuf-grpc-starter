package document

import (
	"errors"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/models"
	"log"
)

func GetDocumentByShortLink(shortLink string) (*models.Document, error) {
	log.Printf("get doc by short-link -> %s\n", shortLink)
	if "" == shortLink {
		return nil,  errors.New("empty shortLink")
	}

	// TODO: find from cache & db, then save to cache
	return nil, errors.New("not implemented")
}


func GetDocumentsByQuery(query *models.Query) ([]*models.Document, error) {
	if nil == query || query.IsEmpty() {
		return nil, errors.New("empty query")
	}

	log.Printf("get docs by query -> %+v\n", query)

	if query.IsShortLinkQuery() {
		doc, err := GetDocumentByShortLink(query.ShortLink)
		if err != nil {
			return nil, err
		}
		return []*models.Document{doc}, nil
	}

	// TODO: query doc in db
	return nil, errors.New("not implemented")
}


func CreateDocument(doc *models.Document) error {
	if nil == doc {
		return errors.New("empty doc")
	}

	// generate shortlink before save it
	doc.GenShortLink()
	log.Printf("create doc -> %+v\n", doc)

	// TODO: save doc
	return nil
}
