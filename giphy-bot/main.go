package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/crispgm/go-g/string/tokenizer"
	"github.com/gin-gonic/gin"
	"github.com/go-lark/examples/giphy-bot/downloader"
	"github.com/go-lark/examples/giphy-bot/model"
	"github.com/go-lark/lark"
	larkgin "github.com/go-lark/lark-gin"
	"github.com/peterhellberg/giphy"
)

const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"

	imageSizeLimit = 1024 * 1024 * 20
)

var (
	giphyClient = giphy.DefaultClient
	imageStore  = &model.Store{}
)

func main() {
	// init DB
	imageStore.Connect()
	// init bot
	bot := lark.NewChatBot(appID, appSecret)
	bot.GetTenantAccessTokenInternal(true)
	bot.StartHeartbeat()
	defer bot.StopHeartbeat()

	// handle events
	r := gin.Default()
	middleware := larkgin.NewLarkMiddleware()
	r.Use(middleware.LarkChallengeHandler())
	r.Use(middleware.LarkMessageHandler())
	r.POST("/lark", func(c *gin.Context) {
		if msg, ok := middleware.GetMessage(c); ok {
			go handleRequest(bot, msg)
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.Run(":8044")
}

func handleRequest(bot *lark.Bot, eventMsg *lark.EventMessage) error {
	query := eventMsg.Event.RealText
	parser := tokenizer.NewTokenizer(query, tokenizer.DefaultDelimiters, " ")
	tokens := parser.Parse()
	cmdToken := tokens[0]

	fmt.Printf("OpenID: %s OpenChatID: %s RawQuery [%s] CMD [%s] Tokens [%v]\n", eventMsg.Event.OpenID, eventMsg.Event.OpenChatID, query, cmdToken, tokens)

	var imageKey, gifID string
	if len(tokens) == 0 {
		gifID = trendingGiphy(0)
	} else {
		argTokens := tokens[1:]
		if cmdToken == "/i" {
			if len(argTokens) > 0 {
				gifID = argTokens[0]
			}
		} else if cmdToken == "/s" {
			gifID = searchGiphy(argTokens)
		} else if cmdToken == "/r" {
			gifID = randomGiphy(argTokens)
		} else if cmdToken == "/t" {
			index := 0
			var err error
			if len(argTokens) > 0 {
				index, err = strconv.Atoi(argTokens[0])
				if err != nil {
					index = 0
				}
			}
			gifID = trendingGiphy(index)
		} else if cmdToken == "/h" {
			usage := "/i id\n/s search ...keywords\n/r random ...keywords\n/t trending [sequence]\nFor more info, please contact zhangwanlong@bytedance.com.\n"
			_, err := bot.PostText(usage, lark.WithChatID(eventMsg.Event.OpenChatID))
			return err
		} else {
			gifID = searchGiphy(tokens)
		}
	}

	if gifID == "" {
		fmt.Println("Call Giphy API failed")
		msgbuf := lark.NewMsgBuffer(lark.MsgText)
		msgbuf.BindOpenChatID(eventMsg.Event.OpenChatID).BindReply(eventMsg.Event.OpenMessageID).Text("Call Giphy API failed")
		_, err := bot.PostMessage(msgbuf.Build())
		return err
	}

	var err error
	timeStart := time.Now().Unix()
	imageKey, err = fetchImage(*bot, gifID)
	timeEnd := time.Now().Unix()
	fmt.Println("Download", gifID, "elapsed:", (timeEnd - timeStart))
	if err != nil {
		prompt := fmt.Sprintf("Fetch image failed: %v", err)
		msgbuf := lark.NewMsgBuffer(lark.MsgText)
		msgbuf.BindOpenChatID(eventMsg.Event.OpenChatID).BindReply(eventMsg.Event.OpenMessageID).Text(prompt)
		_, err = bot.PostMessage(msgbuf.Build())
		return err
	}
	var resp *lark.PostMessageResponse
	if !checkImageSize(fmt.Sprintf("./images/%s.gif", gifID)) {
		msgbuf := lark.NewMsgBuffer(lark.MsgText)
		msgbuf.BindOpenChatID(eventMsg.Event.OpenChatID).BindReply(eventMsg.Event.OpenMessageID).Text("`").Text(gifID).Text("`").Text(" is over size")
		resp, err = bot.PostMessage(msgbuf.Build())
	} else {
		msgbuf := lark.NewMsgBuffer(lark.MsgImage)
		msgbuf.BindOpenChatID(eventMsg.Event.OpenChatID).BindReply(eventMsg.Event.OpenMessageID).Image(imageKey)
		resp, err = bot.PostMessage(msgbuf.Build())
	}

	if err != nil {
		fmt.Println("PostImage failed", err)
	}
	if resp.Code != 0 {
		fmt.Println("PostImage returns error with:", resp.Code)
	}
	return err
}

func searchGiphy(tokens []string) string {
	r, err := giphyClient.Search(tokens)
	var result string
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, item := range r.Data {
		result = item.ID
		break
	}
	return result
}

func randomGiphy(tokens []string) string {
	r, err := giphyClient.Random(tokens)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return r.Data.ID
}

func trendingGiphy(index int) string {
	r, err := giphyClient.Trending()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if len(r.Data) <= index-1 {
		return r.Data[0].ID
	}
	return r.Data[index].ID
}

func fetchImage(bot lark.Bot, gifID string) (string, error) {
	image, db := imageStore.Get(gifID)
	if db.Error == nil {
		if len(image.ImageKey) > 0 {
			fmt.Println(image.ImageKey, "cached")
			return image.ImageKey, nil
		}
	}
	useProxy := os.Getenv("GIPHY_PROXY")
	parellelDownload := os.Getenv("GIPHY_PARELLEL_DOWNLOAD")
	imagePath := fmt.Sprintf("./images/%s.gif", gifID)
	if stat, err := os.Stat(imagePath); os.IsNotExist(err) || stat.Size() == 0 {
		url := downloadableURL(gifID)
		if parellelDownload == "true" {
			err = downloader.DownloadParallelly(imagePath, url)
			if err != nil {
				fmt.Println(image.ImageKey, "Download parellelly failed")
			}
		} else if len(useProxy) > 0 {
			err = downloader.DownloadFileWithProxy(imagePath, useProxy, url)
			if err != nil {
				fmt.Println(image.ImageKey, "Download with proxy failed")
			}
		} else {
			err = lark.DownloadFile(imagePath, url)
			if err != nil {
				fmt.Println(image.ImageKey, "Download finally failed")
				return "", err
			}
		}
	}
	resp, err := bot.UploadImage(imagePath)
	if err != nil {
		fmt.Println(image.ImageKey, "Upload failed")
		return "", err
	}
	imageStore.Create(gifID, resp.Data.ImageKey)
	return resp.Data.ImageKey, nil
}

func downloadableURL(id string) string {
	return fmt.Sprintf("https://i.giphy.com/media/%s/giphy.gif", id)
}

func checkImageSize(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	if f.Size() > imageSizeLimit {
		return false
	}
	return true
}
