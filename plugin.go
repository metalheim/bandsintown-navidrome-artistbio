//go:build wasip1

package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/navidrome/navidrome/plugins/api"
	"github.com/navidrome/navidrome/plugins/host/http"
)

var (
	ErrNotFound       = api.ErrNotFound
	ErrNotImplemented = api.ErrNotImplemented

	client = http.NewHttpService()
)

func (BandsInTownAgent) GetArtistBiography(ctx context.Context, req *api.ArtistBiographyRequest) (*api.ArtistBiographyResponse, error) {
	var bio string
	var err error

	if req.Name != "" {
		log.Printf("[BandsInTown Bio] Fetching for Name: %s", req.Name)
		//TODO: get appid from plugin config storage
		url := "https://rest.bandsintown.com/artists/" + req.Name + "/events?date=upcoming&app_id="+appid;
		resp, err := client.Get(ctx, &http.HttpRequest{Url: url.QueryEscape(url), TimeoutMs: 2000})
		if err != nil || resp.Status != 200 {
			log.Printf("[BandsInTown Bio] Error getting next Events from BandsInTown (status: %d): %v", resp.Status, err)
			return nil, ErrNotFound
		}
	}

	tourDateString, err := extractNextTourDates(resp.Body)
	
	if err != nil || len(tourDateString) == 0 {
		return nil, ErrNotFound
	}
	return &api.ArtistBiographyResponse{Biography: tourDateString}, nil
}

func extractNextTourDates(body []byte) (bio string) {
	var bio string;
	
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	
	for _, event := range data{
		var date	datetime	=	event.datetime
		var title	string		=	event.title
		var city	string		=	event.venue.city
		var venue	string		=	event.venue.name
		
		bio += date + "" + title + "" + "("+ venue + ", " + city + ")"
	}
	
	return bio, nil
}

// Not implemented methods
func (BandsInTownAgent) GetArtistImages(ctx context.Context, req *api.ArtistImageRequest) (*api.ArtistImageResponse, error) {
	return nil, ErrNotImplemented
}
func (BandsInTownAgent) GetArtistMBID(context.Context, *api.ArtistMBIDRequest) (*api.ArtistMBIDResponse, error) {
	return nil, ErrNotImplemented
}
func (BandsInTownAgent) GetSimilarArtists(context.Context, *api.ArtistSimilarRequest) (*api.ArtistSimilarResponse, error) {
	return nil, ErrNotImplemented
}
func (BandsInTownAgent) GetArtistTopSongs(context.Context, *api.ArtistTopSongsRequest) (*api.ArtistTopSongsResponse, error) {
	return nil, ErrNotImplemented
}
func (BandsInTownAgent) GetAlbumInfo(context.Context, *api.AlbumInfoRequest) (*api.AlbumInfoResponse, error) {
	return nil, ErrNotImplemented
}
func (BandsInTownAgent) GetAlbumImages(context.Context, *api.AlbumImagesRequest) (*api.AlbumImagesResponse, error) {
	return nil, ErrNotImplemented
}
func main() {}
func (BandsInTownAgent) OnInit(ctx context.Context, req *api.InitRequest) (*api.InitResponse, error) {
    // Access plugin configuration
    apiKey := req.Config["api_key"]
    appId  := req.Config["app_id"]
    if apiKey == "" {
        return &api.InitResponse{Error: "Missing API key in configuration"}, nil
    }
	if appId == "" {
        return &api.InitResponse{Error: "Missing appId in configuration"}, nil
    }


    // validate API key
    resp, err := httpClient.Get(ctx, &http.HttpRequest{
	//TODO: add apiKey to header (? read api documentation)
        Url: "https://rest.bandsintown.com/?app_id="+appid;
    })
    if err != nil {
        return &api.InitResponse{Error: "Failed to validate API key: " + err.Error()}, nil
    }

    if resp.StatusCode != 200 {
        return &api.InitResponse{Error: "Invalid API key"}, nil
    }

    return &api.InitResponse{}, nil
}
func init() {
	api.RegisterMetadataAgent(BandsInTownAgent{})
}
