package database

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/ihummer/go-firebase/models"
)

// CreateProduct creates a new product in the database
func CreateProduct(db *firestore.Client, p *models.Product) (string, *firestore.DocumentRef, error) {
	ref, _, err := db.Collection("products").Add(context.Background(), p)
	if err != nil {
		return "", nil, err
	}

	return ref.ID, ref, nil
}

func GetAllProducts(db *firestore.Client) ([]*models.Product, error) {
	// iter := db.Collection("products").Documents(context.Background())
	var products []*models.Product

	// for {
	// 	doc, err := iter.Next()
	// 	if err == iter. {
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
