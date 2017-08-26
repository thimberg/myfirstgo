package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

func main() {
	//メイン処理のループ間隔(秒)
	const sleepInterval = 10
	//wikipediaのおまかせ表示URL
//	const wikipediaRandomUrl = "https://ja.wikipedia.org/wiki/特別:おまかせ表示"

	const wikipediaRandomUrl = "http://m.6park.com/index.php" 

	c := time.Tick(sleepInterval * time.Second)
	for now := range c {
		doc, err := goquery.NewDocument(wikipediaRandomUrl)
		if err != nil {
			fmt.Println(err)
		}

		//ランダムに取得したwikipediaのタイトルと取得日時を表示する
		doc.Find("head").Each(func(i int, s *goquery.Selection) {
			title := s.Find("title").Text()
			//取得したタイトルから"- Wikipedia"という文字列を削除
			replacedTitle := strings.Replace(title, "- Wikipedia", "", -1)

			fmt.Println(strings.Repeat("☆", 50))
			fmt.Println(now.Format("2006-01-02 15:4:5"), replacedTitle)
			fmt.Println(strings.Repeat("☆", 50))
		})

		//divのidを目印に説明文を取得して表示する
		doc.Find("#bodyContent #mw-content-text").Each(func(i int, s *goquery.Selection) {
			s.Find("p").Each(func(i int, s *goquery.Selection) {
				fmt.Println(s.Text())
			})

			fmt.Println()
		})
	}
}
