package api

import (
	"net/http"
	"youtube-subtitles-search/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTranscriptHandler(c *gin.Context) {
	videoId := c.Param("videoId")
	transcript, err := service.GetTranscript(videoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type TranscriptResponse struct {
		Timecode string `json:"timecode"`
		Text     string `json:"text"`
	}

	transcriptResponse := make([]TranscriptResponse, len(transcript))
	for i, segment := range transcript {
		transcriptResponse[i] = TranscriptResponse{
			Timecode: segment.OffsetText,
			Text:     segment.Text,
		}
	}
	c.JSON(http.StatusOK, transcriptResponse)
}

func GetHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}