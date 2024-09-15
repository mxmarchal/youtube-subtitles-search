const searchContainer = document.getElementById("search-container");
const youtubeUrlInput = document.getElementById("youtubeUrl");
const goButton = document.getElementById("goButton");

const transcriptContainer = document.getElementById("transcript-container");
const youtubeVideo = document.getElementById("youtube-video");
const transcript = document.getElementById("transcript");
const transcriptList = document.getElementById("transcript-list");

function toggleTranscriptFetching(isFetching) {
  goButton.disabled = isFetching;
  youtubeUrlInput.disabled = isFetching;
  goButton.innerHTML = isFetching ? "Loading..." : "Go";
}

function toggleTranscript(isVisible) {
  transcriptContainer.style.display = isVisible ? "block" : "none";
}

function toggleSearch(isVisible) {
  searchContainer.style.display = isVisible ? "block" : "none";
}

function generateTranscriptElement(transcript, dataId) {
  const transcriptElement = document.createElement("li");
  transcriptElement.setAttribute("data-id", dataId);

  transcriptElement.innerHTML = `<span class="ts-id">${dataId}</span> - <span class="timecode">${transcript.timecode}</span> - <span class="text">${transcript.text}</span>`;
  transcriptList.appendChild(transcriptElement);
}

function getTranscript() {
  const youtubeUrl = youtubeUrlInput.value;
  toggleTranscriptFetching(true);
  if (youtubeUrl.includes("youtube.com")) {
    const videoId = youtubeUrl.split("v=")[1];

    fetch(`/transcript/${videoId}`)
      .then((response) => {
        if (response.status !== 200) {
          alert("An error occurred while fetching the transcript");
          return;
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);
        // TODO: display transcript
        toggleTranscriptFetching(false);
        // data is an array of objects with timecode and text

        //empty transcript list
        transcriptList.innerHTML = "";
        data.forEach((transcript, index) => {
          generateTranscriptElement(transcript, index);
        });

        toggleTranscript(true);
        toggleSearch(false);
      })
      .catch((error) => {
        console.error("Error:", error);
        alert("An error occurred while fetching the transcript");
        toggleTranscriptFetching(false);
      });
  } else {
    alert("Please enter a valid Youtube URL");
    toggleTranscriptFetching(false);
  }
}

goButton.addEventListener("click", getTranscript);

document.addEventListener("DOMContentLoaded", () => {
  fetch("/health")
    .then((response) => response.json())
    .then((data) => {
      if (data.status === "ok") {
        goButton.disabled = false;
      } else {
        alert("Error: Server is down");
      }
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Error: Server is down");
    });
});
