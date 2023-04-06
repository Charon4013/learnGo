package main

import (
	"shendai/get_bilibili/downloader"
	myfmt "shendai/get_bilibili/fmt"
)

func main() {

	request := downloader.InfoRequest{
		Bvids: []string{"BV1JG4y1r7Eq", "BV1i94y1R7yN"},
	}

	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	for _, info := range response.Infos {
		myfmt.Logger.Printf("Title: %s \n desc: %s\n", info.Data.Title, info.Data.Desc)

	}

}
