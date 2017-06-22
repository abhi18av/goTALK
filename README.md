# goTALK
This little program fetches all the basic stats and the transcripts of a given TedTalk

## Build 

- Make sure you have `go-lang` installed.
- On mac it's `brew install go` 
- Now `got get` the following packages 
  -   `go get github.com/PuerkitoBio/goquery`
  -   `go get github.com/fatih/color`
  -   `go get github.com/imdario/mergo`
- Clone the repo with `git clone https://github.com/abhi18av/goTALK`
- Move inside the folder `$ cd goTALK`
- Build the package with `go build .`

Awesome! Now you have the `goTALK` binary!


### Usage

Use the `goTALK` binary to fetch the stats for any talk as 

```sh

$ ./goTALK https://www.ted.com/talks/elon_musk_the_future_we_re_building_and_boring

```



This package is a fundamental part of an umbrella pet project.
