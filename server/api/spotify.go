package api

import (
	"context"
	"fmt"
	"os"
	"strings"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"
)

func GetSpotifyClient() (*spotify.Client, context.Context, error) {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		fmt.Printf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	return spotify.New(httpClient), ctx, err;
}

func parseMusicUrl(url string) string {
	start := strings.Index(url, "track/") + 6
	end := strings.Index(url, "?")
	if (end == -1 || end > len(url)) {
		end = len(url)
	}
	if start < 0 || start > len(url) {
		return "";
	}
	trackID := url[start:end]
	if len(trackID) != 22 {
		return "";
	}
	return trackID
}

func GetTrackInfo(url string, ctx context.Context, client *spotify.Client) (*spotify.FullTrack, error) {
	trackID := parseMusicUrl(url)
	track, err := client.GetTrack(ctx, spotify.ID(trackID))
	if err != nil {
		fmt.Println(err)
	}
	return track, err
}
