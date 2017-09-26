package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
//	"time"
  iconv "github.com/djimenez/iconv-go"
    "database/sql"
  _ "github.com/go-sql-driver/mysql"
)


func getNewsContent(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}
	
	newsCont := ""
	doc.Find("div#mainContent").Each(func(_ int, s *goquery.Selection) {
		content,_ := iconv.ConvertString(s.Text(), "gb2312", "utf-8")
		newsCont = newsCont + content
	})
	
	return newsCont, nil
}
const (
    targetfqdn = "github-ranking.com"
)

func IsRelativePath(url string) bool {
    if strings.Index(url, "//") == 0 {
        return false
    } else if strings.Index(url, "/") == 0 {
        return true
    } else {
        return false
    }
}

func GetAbsoluteURLFromRelativePath(scheme string, fqdn string, relativePath string) string {
    return scheme + "://" + fqdn + relativePath
}


func hasNextPageURL(doc *goquery.Document) (string, bool) {
    nexturl, exists := doc.Find("div > section.item > a").First().Attr("href")
    if exists == true && IsRelativePath(nexturl) {
        return GetAbsoluteURLFromRelativePath("https", targetfqdn, nexturl), true
    }
    return nexturl, false
}

func main() {
	//メイン処理のループ間隔(秒)
	const sleepInterval = 1000

	const news6parkUrl = "http://www.6park.com/" 

	// DB接続
	db, err := sql.Open("mysql", "wpuser:asdffdsa1234@/wp_myblog")
	if err != nil {
		panic(err.Error())
  	}
	
	defer db.Close() // 関数がリターンする直前に呼び出される
	
	// 最大PostIDを取得
	var maxCount int64
	rows, err := db.Query(`SELECT MAX(id) FROM wp_posts`,)
	
	if err != nil {
		panic(err.Error())
	}
	
	for rows.Next() {
		if err := rows.Scan(&maxCount); err != nil {
				panic(err.Error())
			}
	}
	rows.Close()
	
///	c := time.Tick(sleepInterval * time.Second)
///	for now := range c {
		doc, err := goquery.NewDocument(news6parkUrl)
		if err != nil {
			fmt.Println(err)
		}
//		if(exist == true ) {
//			fmt.Println(newurl)
//		} else {
//			fmt.Println("not find")
//		}

		//ランダムに取得したwikipediaのタイトルと取得日時を表示する
///		doc.Find("head").Each(func(i int, s *goquery.Selection) {
///			title := s.Find("title").Text()
///			//取得したタイトルから"- Wikipedia"という文字列を削除
///			replacedTitle := strings.Replace(title, "- Wikipedia", "", -1)
///
///			fmt.Println(strings.Repeat("☆", 50))
///			fmt.Println(now.Format("2006-01-02 15:4:5"), replacedTitle)
///			fmt.Println(replacedTitle)
///			fmt.Println(strings.Repeat("☆", 50))
///		})


		//divのidを目印に説明文を取得して表示する
		doc.Find("div#parknews > a").Each(func(i int, s *goquery.Selection) {
///			fmt.Println(s)


			url, _ := s.Attr("href")

			if strings.HasPrefix(url, "./index.php") {
				url = news6parkUrl+ fmt.Sprint(url[2:len(url)])
			}
			
			fmt.Println(url)
//			s.Find("p").Each(func(i int, s *goquery.Selection) {
//				
//				utf8Txt := make([]byte, len(gb2312Txt))
//				utf8Txt = utf8Txt[:]	
//				iconv.Convert(gb2312Txt, utf8Txt, "gb2312", "utf-8")	

				utf8Title,_ := iconv.ConvertString(s.Text(), "gb2312", "utf-8")	
				fmt.Println(utf8Title)

				utf8Content,_:= getNewsContent(url)
				maxCount += 1

				
				fmt.Println(utf8Content)
/*
				// WP_POSTSへ登録
				_, err := db.Exec(`INSERT INTO wp_posts VALUES (?, ?, current_timestamp, (current_timestamp - interval 9 hour), ?, ?, '', 'publish', 'closed', 'closed', '', ?, '', '', current_timestamp, (current_timestamp - interval 9 hour), '', 0, ?, 0, 'post', '' ,0)`,
							maxCount,
							2,
							utf8Content, 
							utf8Title, 
							maxCount, 
							"http://104.198.63.85/?p="+fmt.Sprint(maxCount),
						)
				if err != nil {
					panic(err.Error())
				}
*/
///			})

			fmt.Println()
			
		})
///	}
}


