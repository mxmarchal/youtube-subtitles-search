package tests

import (
	"testing"
	"youtube-subtitles-search/internal/service"
)

func TestFetchTranscript(t *testing.T) {
    transcript, err := service.GetTranscript("xG-Ef4ig07I")
    if err != nil {
        t.Fatalf("Failed to fetch transcript: %v", err)
    }
    if len(transcript) == 0 {
        t.Fatalf("Transcript is empty")
    }
}