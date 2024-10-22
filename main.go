package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/kkdai/youtube/v2"
)

const downloadsFolder = "downloads"

func createDownloadsFolder() error {
	if _, err := os.Stat(downloadsFolder); os.IsNotExist(err) {
		err := os.Mkdir(downloadsFolder, 0755)
		if err != nil {
			return fmt.Errorf("failed to create downloads folder: %v", err)
		}
		fmt.Println("Downloads folder created.")
	} else if err != nil {
		return fmt.Errorf("error checking downloads folder: %v", err)
	}
	return nil
}

func downloadYouTubeVideoAsAudio(url string) error {
	// Create downloads folder if it doesn't exist
	if err := createDownloadsFolder(); err != nil {
		return err
	}

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

	// Create the output file path
	var vidTitle = keepLettersAndReplaceSpaces(video.Title)
	outputPath := filepath.Join(downloadsFolder, vidTitle+".mp3")
	file, err := os.Create(outputPath)
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

	fmt.Printf("Download completed successfully! Saved to: %s\n", outputPath)
	return nil
}
func keepLettersAndReplaceSpaces(input string) string {
	var builder strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) {
			builder.WriteRune(char)
		} else if char == ' ' {
			builder.WriteRune('-')
		}
	}
	return builder.String()
}

func getPlaylistURLs(playlistID string) ([]string, error) {
	client := youtube.Client{}
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		return nil, fmt.Errorf("failed to get playlist: %v", err)
	}

	var urls []string
	for _, entry := range playlist.Videos {
		urls = append(urls, fmt.Sprintf("https://www.youtube.com/watch?v=%s", entry.ID))
	}

	return urls, nil
}

func main() {
	playlistID := "PL9_XszuQGuzhuuums8Bkc0dj_FRLo9puN"
	urls, err := getPlaylistURLs(playlistID)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("Found %d video URLs:\n", len(urls))
	for _, url := range urls {
		err := downloadYouTubeVideoAsAudio(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
