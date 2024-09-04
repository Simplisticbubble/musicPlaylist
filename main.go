package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
)

func downloadYouTubeVideoAsAudio(url string) error {
	// Extract video ID from the URL
	videoID := url[32:]
	fmt.Printf("Extracted video ID: %s\n", videoID)

	// Create a new YouTube client
	client := youtube.Client{}

	// Get the video information
	video, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Printf("Error getting video info: %v\n", err)
		return err
	}

	// Filter formats to get only audio channels
	formats := video.Formats.WithAudioChannels()

	// Get the stream for the highest quality audio format
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		fmt.Printf("Error getting stream: %v\n", err)
		return err
	}
	defer stream.Close()

	// Create the output file
	file, err := os.Create(filepath.Base(videoID) + ".mp3")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return err
	}
	defer file.Close()

	// Copy the stream to the file
	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Printf("Error copying stream to file: %v\n", err)
		return err
	}

	fmt.Println("Download completed successfully!")
	return nil
}

func main() {
	url := "https://www.youtube.com/watch?v=RkQNm99y8fg"
	err := downloadYouTubeVideoAsAudio(url)
	if err != nil {
		fmt.Println(err)
		return
	}
}
