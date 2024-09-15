package service

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

func getLanguageTranscript(video *youtube.Video) (string, error) {
	priorityLanguages := []string{"en", "es", "fr"}
	captions := video.CaptionTracks
	if len(captions) == 0 {
		return "", fmt.Errorf("no captions found")
	}

	for _, language := range priorityLanguages {
		for _, caption := range captions {
			if caption.LanguageCode == language {
				return caption.LanguageCode, nil
			}
		}
	}

	return captions[0].LanguageCode, nil
}

func GetTranscript(videoId string) (youtube.VideoTranscript, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(videoId)
	if err != nil {
		fmt.Println(err)
		return youtube.VideoTranscript{}, err
	}

	language, err := getLanguageTranscript(video)
	if err != nil {
		fmt.Println(err)
		return youtube.VideoTranscript{}, err
	}

	transcript, err := client.GetTranscript(video, language)
	if err != nil {
		fmt.Println(err)
		return youtube.VideoTranscript{}, err
	}

	return transcript, nil
}