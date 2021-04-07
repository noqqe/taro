package cmd

import (
	"fmt"
	"os"

	"gopkg.in/masci/flickr.v2"
)

func UploadToFlickr(name string) {

	// fetching photo from database
	photo := getPhoto(name)
	body := GetPhotoFromS3(photo.Name)
	// retrieve Flickr credentials from env vars
	apik := os.Getenv("FLICKR_API_KEY")
	apisec := os.Getenv("FLICKR_API_SECRET")
	token := os.Getenv("FLICKR_API_OAUTH_TOKEN")
	tokenSecret := os.Getenv("FLICKR_API_OAUTH_TOKEN_SECRET")

	// do not proceed if credentials were not provided
	if apik == "" || apisec == "" || token == "" || tokenSecret == "" {
		fmt.Fprintln(os.Stderr, "Please set FLICKRGO_API_KEY, FLICKRGO_API_SECRET, "+
			"FLICKRGO_OAUTH_TOKEN and FLICKRGO_OAUTH_TOKEN_SECRET env vars")
		os.Exit(1)
	}

	// create an API client with credentials
	client := flickr.NewFlickrClient(apik, apisec)
	client.OAuthToken = token
	client.OAuthTokenSecret = tokenSecret

	// upload a photo
	params := flickr.NewUploadParams()
	params.Title = photo.Name
	params.Tags = photo.Tags
	params.Description = photo.Alt
	resp, err := flickr.UploadReader(client, body, photo.Name, params)
	if err != nil {
		fmt.Println("Failed uploading:", err)
		if resp != nil {
			fmt.Println(resp.ErrorMsg)
		}
		os.Exit(1)
	} else {
		fmt.Println("Photo uploaded, id:", resp.ID)
	}

}
