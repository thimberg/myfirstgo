package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
  iconv "github.com/djimenez/iconv-go"
)


func main() {
	//メイン処理のループ間隔(秒)
	const sleepInterval = 10

	const news6parkUrl = "http://m.6park.com/index.php" 

	c := time.Tick(sleepInterval * time.Second)
	for now := range c {
		doc, err := goquery.NewDocument(news6parkUrl)
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
		doc.Find(".titleList").Each(func(i int, s *goquery.Selection) {
			s.Find("p").Each(func(i int, s *goquery.Selection) {
				
//				utf8Txt := make([]byte, len(gb2312Txt))
//				utf8Txt = utf8Txt[:]	
//				iconv.Convert(gb2312Txt, utf8Txt, "gb2312", "utf-8")	

				utf8Txt,_ := iconv.ConvertString(s.Text(), "gb2312", "utf-8")	
				fmt.Println(utf8Txt)
			})

			fmt.Println()
		})
	}
}
