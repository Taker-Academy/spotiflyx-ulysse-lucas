package api

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetYoutubeClient() (*youtube.Service, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
    ctx := context.Background()
    service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        fmt.Println("Erreur lors de la création du service YouTube :", err)
        return nil, err
    }
	return service, err
}

func parseVideoUrl(url string) string {
	start := strings.Index(url, "youtu.be/") + 9
	end := strings.Index(url, "?")
	if (end == -1 || end > len(url)) {
		end = len(url)
	}
	if start < 0 || start > len(url) {
		return "";
	}
	videoID := url[start:end]
	if len(videoID) != 11 {
		return "";
	}
	return videoID
}

func GetVideoInfo(url string) (*youtube.VideoListResponse, error) {
	client, err := GetYoutubeClient()
	if err != nil {
		fmt.Println("Error getting Youtube client:", err)
		return nil, err
	}

	videoID := parseVideoUrl(url)

	// Retrieve video details
	call := client.Videos.List([]string{"snippet"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		fmt.Println("Error retrieving video details:", err)
		return nil, err
	}

	// Check if there are matching videos
	if len(response.Items) == 0 {
		fmt.Println("No video found with this ID.")
		return nil, errors.New("non existing id")
	}

	videoListResponse := &youtube.VideoListResponse{
		Kind:  response.Kind,
		Etag:  response.Etag,
		Items: []*youtube.Video{response.Items[0]},
		PageInfo: &youtube.PageInfo{
			TotalResults: response.PageInfo.TotalResults,
			ResultsPerPage: response.PageInfo.ResultsPerPage,
		},
	}

	return videoListResponse, nil
}
