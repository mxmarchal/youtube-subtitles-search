# YouTube Subtitles Search

This project is a web application that allows users to fetch and display transcripts from YouTube videos. It's built with Go for the backend and vanilla JavaScript for the frontend.

## Features

- Fetch transcripts from YouTube videos using their URL
- Display transcripts with timecodes
- Simple and responsive user interface

## Tech Stack

- Backend: Go (Gin framework)
- Frontend: HTML, CSS, JavaScript
- Docker for containerization
- Weaviate for vector search (not implemented yet)

## Project Structure

```
.
├── cmd
│   └── app
│       └── main.go
├── internal
│   ├── api
│   │   ├── handler.go
│   │   └── router.go
│   ├── config
│   │   └── config.go
│   └── service
│       └── transcript_service.go
├── static
│   ├── css
│   │   └── style.css
│   ├── js
│   │   └── app.js
│   └── index.html
├── tests
│   └── youtube_test.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Setup and Running

1. Clone the repository
2. Make sure you have Docker and Docker Compose installed
3. Run the following command in the project root:

```
docker-compose up --build
```

4. Open your browser and navigate to `http://localhost:8080`

## API Endpoints

- `GET /health`: Health check endpoint
- `GET /transcript/:videoId`: Fetch transcript for a given YouTube video ID

## Frontend

The frontend is a simple HTML page with JavaScript for interactivity. It allows users to input a YouTube URL and displays the fetched transcript.

## Backend

The backend is written in Go using the Gin framework. It handles API requests and interacts with the YouTube API to fetch transcripts.

## Testing

There's a basic test for the transcript fetching functionality. To run the tests, use the following command:

```
go test ./tests
```

## Future Improvements

- Implement vector search using Weaviate for more advanced transcript searching
- Add functionality to download and run an audio-to-text model when YouTube transcripts are not available
- Improve error handling and user feedback
- Implement caching for frequently accessed transcripts

## Note

This project is for learning purposes and experimentation with Go and AI/ML concepts. It should not be used in a production environment.
