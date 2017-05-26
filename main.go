package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	//"theTwinFiles/talkFetch"
	// "./talkFetch/transcriptsPage"
	// "./talkFetch/videoPage"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/imdario/mergo"
)

type TedTalk struct {
	TalkVideoPage      VideoPage      `json:"TalkVideoPage"`
	TalkTranscriptPage TranscriptPage `json:"TalkTranscriptPage"`
}

type VideoPage struct {
	TalkURL                 string   `json:"VideoURL"`
	AvailableSubtitlesCount string   `json:"AvailableSubtitlesCount"`
	Speaker                 string   `json:"Speaker"`
	Duration                string   `json:"Duration"`
	TimeFilmed              string   `json:"TimeFilmed"`
	TalkViewsCount          string   `json:"TalkViewsCount"`
	TalkTopicsList          []string `json:"TalkTopicsList"`
	TalkCommentsCount       string   `json:"TalkCommentsCount"`
}

type talkTranscript struct {
	LocalTalkTitle              string   `json:"LocalTalkTitle"`
	Paragraphs                  []string `json:"Paragraphs"`
	TimeStamps                  []string `json:"TimeStamps"`
	TalkTranscriptAndTimeStamps []string `json:"TalkTranscriptAndTimeStamps"`
}

type TranscriptPage struct {
	AvailableTranscripts []string                  `json:"AvailableTranscripts"`
	DatePosted           string                    `json:"DatePosted"`
	Rated                string                    `json:"Rated"`
	ImageURL             string                    `json:"ImageURL"`
	TalkTranscript       map[string]talkTranscript `json:"TalkTranscript"`
}

func main() {

	// Add logger and stubs for better debugging
	checkInternet()
	videoURL := os.Args[1]
	//videoURL := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"
	//videoURL := "https://www.ted.com/talks/elon_musk_the_future_we_re_building_and_boring"

	// We are knowingly making sync. calls to the main Video page and
	// in case we find there are One or more subtitle lanuguages we make
	// more async. requests
	var videoPageInfo VideoPage
	videoPageInfo = videoFetchInfo(videoURL)

	// Checking if there are any subtitles at all
	// In case there are, we send a default query to fetch the list of available languages
	numOfSubtitles, _ := strconv.ParseInt(videoPageInfo.AvailableSubtitlesCount, 10, 64)

	// This function will cause the program to EXIT if there are no subtitles
	// Else we continue to fill the basic page
	exitIfNoSubtitlesExist(numOfSubtitles)

	transcriptEnURL := videoURL + "/transcript?language=en"

	// Since we've already made the request to default lang transcript
	// we fill in the common details into a transcript info struct
	transcriptPageCommonInfo := transcriptFetchCommonInfo(transcriptEnURL)

	urls := genTranscriptURLs(langCodes, transcriptPageCommonInfo.AvailableTranscripts, videoURL)
	//fmt.Println(transcriptCommonInfo.AvailableTranscripts)

	// @@@@@@@@@@
	// Page UnCommon

	//var transcriptS []talkTranscript

	langSpecificMap := make(map[string]talkTranscript)

	var wg sync.WaitGroup

	numOfURLs := len(urls)
	//fmt.Println(numOfURLs)
	wg.Add(numOfURLs)

	for _, url := range urls {

		go func(url string) {
			defer wg.Done()
			//color.Green(url)
			x, langName := transcriptFetchUncommonInfo(url)
			langSpecificMap[langName] = x
			//transcriptS = append(transcriptS, x)
		}(url)

	}

	wg.Wait()

	//writeJSON(videoPageInfo)

	//fmt.Println(transcriptS)

	// STUB for the actual construction of the complete talk struct

	var transcriptPageUnCommonInfo TranscriptPage
	transcriptPageUnCommonInfo.TalkTranscript = langSpecificMap

	transcriptPageCompleteInfo := transcriptPageCommonInfo
	mergo.Merge(&transcriptPageCompleteInfo, transcriptPageUnCommonInfo)
	//writeJSON(transcriptPageCompleteInfo)

	//temp1, _ := json.Marshal(transcriptPageCompleteInfo)
	//fmt.Println(string(temp1))

	var tedTalk TedTalk
	tedTalk.TalkVideoPage = videoPageInfo
	tedTalk.TalkTranscriptPage = transcriptPageCompleteInfo
	//mergo.Merge(&tedTalk, transcriptPageCompleteInfo)
	//fmt.Println(tedTalk)
	//temp2, _ := json.Marshal(tedTalk)
	//fmt.Println(string(temp2))
	writeJSON(tedTalk)
} // end of main()

func writeJSON(aStruct TedTalk) {

	temp1, _ := json.Marshal(aStruct)
	//fmt.Println(string(temp1))
	htmlSplit := strings.Split(aStruct.TalkVideoPage.TalkURL, "/")
	talkName := htmlSplit[len(htmlSplit)-1]

	fileName := "./" + talkName + ".json"

	f, err := os.Create(fileName)
	checkErr(err)

	f.Write(temp1)
	defer f.Close()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func checkInternet() {
	// Make a GET request
	rs, err := http.Get("https://google.com")
	// Process response
	if err != nil {
		color.Red("We're OFF-Line!")
		//panic("Not connected to the net") // More idiomatic way would be to print the error and die unless it's a serious error

		// Learn about exit status in Golang
		os.Exit(1)
	}

	defer rs.Body.Close()

}

func exitIfNoSubtitlesExist(numOfSubtitles int64) {
	if numOfSubtitles < 1 {
		color.Red("No subtitles available yet")
		os.Exit(1)
	}
}

func videoFetchInfo(url string) VideoPage {

	videoPage, _ := goquery.NewDocument(url)

	videoPageInstance := VideoPage{
		TalkURL:                 videoTalkURL(url),
		AvailableSubtitlesCount: videoAvailableSubtitlesCount(videoPage),
		Speaker:                 videoSpeaker(videoPage),
		Duration:                videoDuration(videoPage),
		TimeFilmed:              videoTimeFilmed(videoPage),
		TalkViewsCount:          videoTalkViewsCount(videoPage),
		TalkTopicsList:          videoTalkTopicsList(videoPage),
		TalkCommentsCount:       videoTalkCommentsCount(videoPage),
	}

	return videoPageInstance
}

func transcriptFetchCommonInfo(url string) TranscriptPage {
	transcriptPage, _ := goquery.NewDocument(url)

	transcriptPageInstance := TranscriptPage{

		AvailableTranscripts: transcriptAvailableTranscripts(transcriptPage),
		DatePosted:           transcriptDatePosted(transcriptPage),
		Rated:                transcriptRated(transcriptPage),
		ImageURL:             transcriptGetImage(transcriptPage, url),
	}
	return transcriptPageInstance
}

func transcriptFetchUncommonInfo(url string) (talkTranscript, string) {

	//fmt.Println(url)
	transcriptPage, _ := goquery.NewDocument(url)
	//fmt.Println(transcriptLocalTalkTitle(transcriptPage))

	transcript := talkTranscript{

		LocalTalkTitle:              transcriptLocalTalkTitle(transcriptPage),
		Paragraphs:                  transcriptTalkTranscript(transcriptPage),
		TimeStamps:                  transcriptTimeStamps(transcriptPage),
		TalkTranscriptAndTimeStamps: transcriptTalkTranscriptAndTimeStamps(transcriptPage),
	}

	langName := strings.Split(url, "=")[1]
	return transcript, langName
}
