package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var f *os.File
var err error
var allLink = make(map[string]string)

func main() {
	f, err = os.OpenFile("e:/myname.html", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}

	defer func() {
		for link, title := range allLink {
			f.WriteString(fmt.Sprintf(`<a href="%s" target="_blank">%s</a><br>`, link, title))
		}
	}()

	pg := 2
	for {
		one(pg)
		pg++
		if pg > 70 {
			return
		}
	}

}

func one(pg int) error {
	url := "http://www.52shijing.com/category/270/index_" + fmt.Sprintf("%d", pg) + ".html"
	fmt.Println("url:", url)
	fmt.Println("allLink len=", len(allLink))

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	//	resp.Request.

	list, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//	fmt.Println(string(list))

	//<span class="h5">语字起名女孩名字怎么样 女孩关于语字起名大全</span>

	//	r1 := regexp.MustCompile("p([a-z]+)ch")
	r1 := regexp.MustCompile(`<li><a href="(.*)" title="(.*)" target="_blank">`)

	//	titleList := r1.FindAllString(string(list), 5)
	titleList := r1.FindAllStringSubmatch(string(list), 30)
	//	fmt.Println(titleList)

	var chars = []string{"金", "水", "鼠"}

	for _, v := range titleList {
		fmt.Println(v[1], v[2])
		link := v[1]
		title := v[2]
		if strings.ContainsAny(title, strings.Join(chars, "")) {

			for _, char := range chars {
				color := "red"
				if char == "金" {
					color = "orange"
				} else if char == "水" {
					color = "green"
				}
				title = strings.ReplaceAll(title, char, fmt.Sprintf(`<b style="color:%s;">%s</b>`, color, char))
			}
			allLink["http://www.52shijing.com/"+link] = title
		}
	}

	return nil
}
