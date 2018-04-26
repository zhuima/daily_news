package robot

import (
	"github.com/sirupsen/logrus"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"fmt"
	"errors"
	"strings"
	"net/http"
)

func Bot(url string) (string, error) {
	var content string

	res, err := http.Get(url)
	if err != nil {
		logrus.Fatal("parst html fail", err.Error())
	} else {
		defer res.Body.Close()
		utfBody, err := iconv.NewReader(res.Body, "gb2312", "utf-8")
		if err != nil {
			logrus.Fatal("get info for html body ", err.Error())
			return "Faild", errors.New("get info for html body ")
		} else {
			doc, err := goquery.NewDocumentFromReader(utfBody)
			if err != nil {
				logrus.Fatal("goquery html body ", err.Error())
			}
			doc.Find("ol div.g").Each(func(i int, s *goquery.Selection) {
				a := s.Find("h3.r a")
				title := a.Text()
				cite := s.Find("div.s cite")
				link := cite.Text()
				if strings.HasPrefix(link, "http") {
					content = content + " \n" + fmt.Sprintf("- [%s](%s) \n", title, link)
					//logrus.Info("content is", content)
				} else {
					content = content + " \n" + fmt.Sprintf("- [%s](http://%s) \n", title, link)
				}
			})
		}
	}
	return content, nil

}
