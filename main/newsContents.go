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

func getNewsContent(url string, selector string, encoding string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}
	
	newsCont := ""
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		content,_ := iconv.ConvertString(s.Text(), encoding, "utf-8")
		newsCont = newsCont + content
	})
	
	return newsCont, nil
}

type TopNews struct {
	title string
	url string
}

func getNewsTop (topUrl string, urlSelector string, titleSelector string, encoding string) ([]TopNews) {
	var rtnArr []TopNews
//	fmt.Println(topUrl + urlSelector + encoding)
	doc, err := goquery.NewDocument(topUrl)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(urlSelector).Each(func(i int, s *goquery.Selection) {
///		fmt.Println(s.Html())
		url, _ := s.Attr("href")

		if strings.HasPrefix(url, "./index.php") {
			url = topUrl + fmt.Sprint(url[2:len(url)])
		}

///		if ! strings.HasPrefix(url, "http://news.6park.com/newspark") {
///			return
///		}
//		fmt.Println(url)
		var tempTitle string
		if titleSelector == "" {
			tempTitle = s.Text()
		} else {
			tempTitle = s.Find(titleSelector).First().Text()
		}
		utf8Title,_ := iconv.ConvertString(tempTitle, encoding, "utf-8")	
			rtnArr = append(rtnArr, TopNews{ utf8Title, url } )
	})
	
	return rtnArr
}

func getNewsTop__(topUrl string, selector1 string, selector2 string, encoding string) () {
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


	doc, err := goquery.NewDocument(topUrl)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(selector1).Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")

		if strings.HasPrefix(url, "./index.php") {
			url = topUrl + fmt.Sprint(url[2:len(url)])
		}

		if ! strings.HasPrefix(url, "http://news.6park.com/newspark") {
			return
		}
		fmt.Println(url)

		utf8Title,_ := iconv.ConvertString(s.Text(), encoding, "utf-8")	
		utf8Content,_:= getNewsContent(url, selector2, encoding)

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

	})
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




