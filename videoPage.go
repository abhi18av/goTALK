package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// @@@@@@@@@@@@@@@@@@
// VIDEO PAGE

func videoAvailableSubtitlesCount(doc *goquery.Document) string {

	subtitles := doc.Find(".player-hero__meta__link").Contents().Text()
	//fmt.Println(subtitles)

	//for _, x := range strings.Split(subtitles, "\n") {
	//fmt.Println(x)
	//println("~~~~~~")
	//}

	y := strings.Split(subtitles, "\n")
	z := strings.Split(y[3], " ")[0]
	// In case I need an INT
	//numOfSubtitles, _ := strconv.ParseInt(z, 10, 32)
	numOfSubtitles := z
	return numOfSubtitles
}

func videoSpeaker(doc *goquery.Document) string {
	speaker := doc.Find(".talk-speaker__name").Contents().Text()
	//fmt.Println(speaker)
	speaker = strings.Trim(speaker, "\n")
	return speaker
}

/*
// This is now taken from the transcripts page
func title(doc *goquery.Document) {
	title := doc.Find(".player-hero__title__content").Contents().Text()
	fmt.Println(title)
}
*/

func videoDuration(doc *goquery.Document) string {

	duration := doc.Find(".player-hero__meta").Contents().Text()
	//fmt.Println(duration)

	//for _, x := range strings.Split(duration, "\n") {
	//	fmt.Println(x)
	//	println("~~~~~~")
	//}

	x := strings.Split(duration, "\n")
	//fmt.Println(x[6])
	return x[6]

}

// TimeFilmed : Time at which the talk was filmed
func videoTimeFilmed(doc *goquery.Document) string {

	talkFilmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(talkFilmed)

	y := strings.Split(talkFilmed, "\n")
	//fmt.Println(y[11])
	return y[11]
}

func videoTalkViewsCount(doc *goquery.Document) string {

	talkViewsCount := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talkViewsCount)

	a := strings.Split(talkViewsCount, "\n")
	b := strings.TrimSpace(a[2])
	//fmt.Println(b)
	return b

}

func videoTalkTopicsList(doc *goquery.Document) []string {

	talkTopics := doc.Find(".talk-topics__list").Contents().Text()

	c := strings.Split(talkTopics, "\n")
	var topics []string
	for i := 3; i < len(c); i++ {
		//fmt.Println(c[i])
		if c[i] == "" {

		} else {
			topics = append(topics, c[i])
		}
	}
	return topics
}

func videoTalkCommentsCount(doc *goquery.Document) string {

	talkCommentsCountt := doc.Find(".h11").Contents().Text()
	//fmt.Println(talkCommentsCountt)
	d := strings.Split(talkCommentsCountt, " ")
	//fmt.Println(d[0])
	return strings.TrimLeft(d[0], "\n")
}

func videoTalkURL(videoURL string) string {
	return videoURL
}
