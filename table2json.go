package table2json

import (
	"bytes"
	"encoding/json"
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ConvertMap(html string) ([]map[string]string, error) {
	result, err := ConvertList(html)
	if err != nil {
		return nil, err
	}
	list := []map[string]string{}
	for i := 0; i < len(result[0]); i++ {
		if i == 0 {
			continue
		}
		kvData := map[string]string{}
		for _, n := range result {
			kvData[n[0]] = n[i]
		}
		list = append(list, kvData)
	}
	return list, nil
}

func ConvertJsonString(html string) (string, error) {
	result, err := ConvertMap(html)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	var str bytes.Buffer
	_ = json.Indent(&str, b, "", "    ")
	return str.String(), err
}

func ConvertList(html string) ([][]string, error) {
	b := bytes.NewBufferString(html)
	dom, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		return nil, err
	}
	result := [][]string{}
	if dom.Find("tr").First().Find("th").Size() == 0 {
		return nil, errors.New("table format error")
	}
	dom.Find("tr").First().Find("th").Each(func(i int, s *goquery.Selection) {
		result = append(result, []string{strings.TrimSpace(s.Text())})
	})
	dom.Find("tr").Next().Each(func(tri int, trs *goquery.Selection) {
		trs.Find("td").Each(func(tdi int, tds *goquery.Selection) {
			for i := range result {
				if len(result[i])-1 > tri {
					continue
				}
				tdLenth := getRowspanLenth(tds)
				tdWidth := getRowspanWidth(tds)
				tds.Find("br").ReplaceWithHtml(" \n ")
				tdContent := strings.TrimSpace(tds.Text())
				for count := 0; count < tdWidth; count++ {
					result[i+count] = append(result[i+count], slices.Repeat([]string{tdContent}, tdLenth)...)
				}
				break
			}
		})
	})
	return result, nil
}

func getRowspanLenth(e *goquery.Selection) int {
	str, ok := e.Attr("rowspan")
	if !ok {
		return 1
	}
	result, _ := strconv.Atoi(str)
	return result
}

func getRowspanWidth(e *goquery.Selection) int {
	str, ok := e.Attr("colspan")
	if !ok {
		return 1
	}
	result, _ := strconv.Atoi(str)
	return result
}
