package webscraper

import (
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type ScraperService struct {
	mu sync.Mutex
}

func NewScraper() *ScraperService {
	return &ScraperService{}
}

// webページからテキストを抽出する
func (s *ScraperService) ExtractText(url string) (*[]string, error) {
	var text string

	// 新しいCollyコレクタを初期化
	c := colly.NewCollector()

	// HTML要素からテキストを抽出するハンドラ
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// body内のすべてのテキストノードを取得
		s.mu.Lock() // 排他制御
		text = e.Text
		s.mu.Unlock()
	})

	// エラーハンドリング
	c.OnError(func(r *colly.Response, err error) {
		log.Println("エラー:", err)
	})

	// リクエストを発行
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	// 改行情報を削除
	text = strings.ReplaceAll(text, "\n", "")
	// タブ情報を削除
	text = strings.ReplaceAll(text, "\t", "")
	// 句点で分割
	sentences := strings.Split(text, "。")

	// 抽出されたテキストを返す
	return &sentences, nil
}
