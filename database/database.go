package database

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/ihummer/go-firebase/models"
)

type Database interface {
	CreateProduct(db *firestore.Client, p *models.Product) (string, error)
	GetAllProducts(db *firestore.Client) ([]*models.Product, error)
}

type FirestoreDatabase struct{}

func (f FirestoreDatabase) CreateProduct(db *firestore.Client, p *models.Product) (string, error) {
	ref, _, err := db.Collection("products").Add(context.Background(), p)
	if err != nil {
		return "", err
	}

	return ref.ID, nil
}

func (f FirestoreDatabase) GetAllProducts(db *firestore.Client) ([]*models.Product, error) {
	// iter := db.Collection("products").Documents(context.Background())

	var products []*models.Product
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	p := new(models.Product)
	// 	doc.DataTo(p)
	// 	products = append(products, p)
	// }

	return products, nil
}
