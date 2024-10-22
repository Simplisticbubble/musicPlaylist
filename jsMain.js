/**
 * Function to convert a YouTube video to MP3 format.
 *
 * @param {string} youtubeUrl - The URL of the YouTube video.
 * @returns {string} The URL of the converted MP3 file.
 */
function convertYoutubeToMp3(youtubeUrl) {
    // Check if the YouTube URL is valid
    if (!isValidYoutubeUrl(youtubeUrl)) {
        throw new Error("Invalid YouTube URL");
    }

    // Extract the video ID from the YouTube URL
    const videoId = extractVideoId(youtubeUrl);

    // Generate the download URL for the MP3 file
    const mp3Url = generateMp3Url(videoId);

    // Return the URL of the converted MP3 file
    return mp3Url;
}

/**
 * Function to check if a YouTube URL is valid.
 *
 * @param {string} url - The YouTube URL to be validated.
 * @returns {boolean} True if the URL is valid, false otherwise.
 */
function isValidYoutubeUrl(url) {
    // Regular expression pattern to match YouTube URLs
    const youtubeUrlPattern = /^(https?:\/\/)?(www\.)?youtube\.com\/watch\?v=[\w-]{11}$/;

    // Check if the URL matches the pattern
    return youtubeUrlPattern.test(url);
}

/**
 * Function to extract the video ID from a YouTube URL.
 *
 * @param {string} url - The YouTube URL.
 * @returns {string} The video ID.
 */
function extractVideoId(url) {
    // Extract the video ID from the URL
    const videoId = url.split("v=")[1];

    // Return the video ID
    return videoId;
}

/**
 * Function to generate the download URL for the MP3 file.
 *
 * @param {string} videoId - The video ID of the YouTube video.
 * @returns {string} The download URL for the MP3 file.
 */
function generateMp3Url(videoId) {
    // Construct the download URL using the video ID
    const mp3Url = `https://www.youtube.com/convert/${videoId}`;

    // Return the download URL
    return mp3Url;
}

// Usage Example for convertYoutubeToMp3

const youtubeUrl = "https://www.youtube.com/watch?v=dTQ_Yd7MLEg";
try {
    const mp3Url = convertYoutubeToMp3(youtubeUrl);
    console.log(`The MP3 file can be downloaded from: ${mp3Url}`);
} catch (error) {
    console.error(error.message);
}