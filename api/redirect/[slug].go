package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type urlRedirect struct {
	Slug      string `bson:"slug"`
	TargetURL string `bson:"targetURL"`
}

var httpPrefix = regexp.MustCompile("^http")

// RedirectHandler redirects to the URL previously associated to the given slug, or to https://casca.dev/goto if no URL is available.
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	origin := "https://casca.dev"
	if r.Header.Get("Origin") == "https://carmeloscandaliato.com" {
		origin = "https://carmeloscandaliato.com"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		return
	}

	targetURL, err := handlerWithError(w, r)
	fmt.Printf(targetURL)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Redirect(w, r, "https://casca.dev/goto", http.StatusPermanentRedirect)
		return
	}
	http.Redirect(w, r, targetURL, http.StatusPermanentRedirect)
}

func handlerWithError(w http.ResponseWriter, r *http.Request) (string, error) {
	slug := r.URL.Query().Get("slug")
	if slug == "" {
		return "", e("could not extract slug from request", nil)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB")))
	if err != nil {
		return "", e("error while connecting to database", err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("goto").Collection("redirects")

	var redir urlRedirect
	err = collection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&redir)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", e("not found", err)
		}
		return "", e("error while looking for existing redirect", err)
	}

	if !httpPrefix.MatchString(redir.TargetURL) {
		return "https://" + redir.TargetURL, nil
	}
	return redir.TargetURL, nil
}

func e(message string, err error) error {
	if err == nil {
		return fmt.Errorf("%s", message)
	}
	return fmt.Errorf("%s: %v", message, err)
}
