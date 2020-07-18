package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type urlRedirect struct {
	ShortURL  string `bson:"-" json:"shortURL"`
	Slug      string `bson:"slug" json:"-"`
	TargetURL string `bson:"targetURL" json:"targetURL"`
}

type requestBody struct {
	TargetURL string
}

type handlerError struct {
	message string
	code    int
}

type errorResponse struct {
	Message string `json:"error"`
}

const timeoutSec = 1

var letterRunes = []rune("abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ShortenHandler creates a shortURL from a given targetURL.
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://casca.dev, https://carmeloscandaliato.com, ")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		return
	}

	urlRedirect, err := handlerWithError(w, r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(&errorResponse{Message: err.Error()})
		return
	}
	urlRedirect.ShortURL = toURL(urlRedirect.Slug)
	json.NewEncoder(w).Encode(urlRedirect)
}

func handlerWithError(w http.ResponseWriter, r *http.Request) (*urlRedirect, error) {
	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return nil, e("error while parsing request", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB")))
	if err != nil {
		return nil, e("error while connecting to database", err)
	}
	// TODO: warn if connection leaks?
	defer client.Disconnect(context.TODO())

	collection := client.Database("goto").Collection("redirects")

	var redir urlRedirect
	err = collection.FindOne(context.TODO(), bson.M{"targetURL": reqBody.TargetURL}).Decode(&redir)
	if err == nil {
		return &redir, nil
	}
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, e("error while looking for existing redirect", err)
	}

	slug, err := unusedSlug(collection)
	if err != nil {
		return nil, e("error while generating short url", err)
	}

	redir = urlRedirect{TargetURL: reqBody.TargetURL, Slug: slug}
	_, err = collection.InsertOne(context.TODO(), redir)
	if err != nil {
		return nil, e("error while adding record to database", err)
	}

	return &redir, nil
}

func e(message string, err error) error {
	return fmt.Errorf("%s: %v", message, err)
}

func unusedSlug(c *mongo.Collection) (string, error) {
	for {
		slug := randStringRunes(4)
		err := c.FindOne(context.TODO(), bson.M{"slug": slug}).Err()
		if err == nil {
			continue
		}
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return "", err
		}
		return slug, nil
	}
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func toURL(slug string) string {
	return fmt.Sprintf("https://casca.dev/goto/%s", slug)
}
