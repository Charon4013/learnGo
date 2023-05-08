package main

import (
	"bytes"
	"encoding/base64"
	"log"
	"time"

	bili "github.com/JimmyZhangJW/biliStreamClient"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const (
	SECRET_ID = "114514"
	SecretKey = "1919810"
)

func main() {
	biliClient := bili.New()
	biliClient.Connect(923833)
	defer biliClient.Disconnect()
	for {
		packBody := <-biliClient.Ch
		switch packBody.Cmd {
		case "DAMMU_MSG":
			// log.Println(packBody.ParseDanmu())
			danmu, err := packBody.ParseDanmu()
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(danmu)
			sanitizedMsg := bili.Sanitize(danmu.Message)
			if bili.IsContainChineseWord(sanitizedMsg) {
				encodedVoice, err := bili.GetVoiceFromTencentCloud(SECRET_ID, SecretKey, bili.DefaultGirlVoice, sanitizedMsg)
				if err != nil {
					log.Fatalln(err)
				}
				data, err := base64.StdEncoding.DecodeString(encodedVoice)

				streamer, format, err := wav.Decode(bytes.NewReader(data))
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(streamer)
			}

		case "SEND_GIFT":
			log.Println(packBody.ParseGift())
		case "COMBO_SEND":
			log.Println(packBody.ParseGiftCombo())
		default:
			log.Println(packBody.Cmd)
		}
	}
}
