package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"strconv"
)

const (
	cacheSize = 1000
)

var (
	token  string
	times  int
	client *resty.Client
)

func init() {
	flag.StringVar(&token, "token", "", "token")
	flag.IntVar(&times, "times", 100, "times")
	client = resty.New().
		SetBaseURL("https://cat-match.easygame2021.com").
		SetHeader("Host", "cat-match.easygame2021.com").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip,compress,br,deflate").
		SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c25) NetType/WIFI Language/zh_CN").
		SetHeader("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/15/page-frame.html").
		SetQueryParam("rank_score", "1").
		SetQueryParam("rank_state", "1").
		SetQueryParam("rank_time", strconv.FormatInt(1000, 10)).
		SetQueryParam("rank_role", "1").
		SetQueryParam("skin", "1")
}

func main() {
	var (
		ch           = make(chan bool, cacheSize)
		successTimes int
	)
	flag.Parse()
	if len(token) == 0 {
		log.Println("token 为空, 请继续输入token, 或者直接使用 -token 参数")
		fmt.Scanf("%s", &token)
	}
	for {
		go func() {
			ch <- true
			ctx := context.Background()
			if err := Send(ctx, token, ch); err != nil {
				log.Println(err)
				return
			}
			successTimes++
			fmt.Println("成功次数共", successTimes)
		}()
	}
}

func Send(ctx context.Context, theToken string, ch chan bool) error {
	var (
		resp   *resty.Response
		err    error
		result gjson.Result
	)
	resp, err = client.R().SetContext(ctx).SetHeader("t", theToken).Get("/sheep/v1/game/game_over")
	if err != nil {
		<-ch
		return fmt.Errorf("请求失败: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		<-ch
		return fmt.Errorf("[%d] 请求错误: %s", resp.StatusCode(), resp.String())
	}
	if result = gjson.Parse(resp.String()); result.Get("err_code").Int() != 0 {
		<-ch
		return fmt.Errorf("请求错误: %s", resp.String())
	}
	<-ch
	return nil
}
