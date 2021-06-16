package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

type TitleTag struct {
	Level1 string `json:"level1"`
	Level2 string `json:"level2"`
	Level3 string `json:"level3"`
	Level4 string `json:"level4"`
}

var titleUserId1 = map[string]string{}

type IdTitle struct {
	Id       string
	OldTitle string
	NewTitle string
}

func modifyTitle() {
	//var m = map[string]string{
	//	"891358": "白雪，贝儿，给孩子装饰房间！ #小猪佩奇  #儿童玩具  #白雪公主玩具", "891360": "#骗你生儿子系列 #二胎孕妈日常 #服务号10亿补贴 孕晚期不知道是不是能辨别人声啦，每次哥哥说话就动，到了晚期每天都睡不好，腰特别痛，屁股也疼，有同感的吗？", "891361": "#孕期知识 #母婴知识 #胎儿发育", "891362": "我姓赵，农村人，目前在广东有自己的工厂，没靠男人，没啃老，刷到是缘分，交个朋友，认识我不会让你吃亏 #儿童玩具", "891363": "#分享宝藏男孩 #无忧无虑的童年 #动画片",
	//}

	filename := `C:\Users\Administrator\Desktop\8.33.0\8.51\title-lskajlkjfalf`
	bList, _ := ioutil.ReadFile(filename)
	var m map[string]string
	json.Unmarshal(bList, &m)

	f, err := os.Create(`C:\Users\Administrator\Desktop\8.33.0\8.51\title-compare1.txt`)
	if err != nil {
		panic(err)
	}
	_ = f

	var idTitleList = make([]IdTitle, 0)

	count := 0
	for k, v := range m {
		_ = k
		newTitle := demoRegx(v)
		if v != newTitle {
			//fmt.Println(v)
			//fmt.Println(newTitle)
			//fmt.Println("")

			v1 := strings.Replace(v, " ", "", -1)
			newTitle1 := strings.Replace(newTitle, " ", "", -1)
			if newTitle1 == "" {
				count++
			}
			if v1 != newTitle1 {
				idTitleList = append(idTitleList, IdTitle{
					Id:       k,
					OldTitle: v,
					NewTitle: newTitle,
				})
				//f.WriteString("before：" + v + "\n")
				//f.WriteString("af ter：" + newTitle + "\n\n")
			}
		}
	}
	fmt.Println(count)
	bList, _ = json.Marshal(idTitleList)
	os.WriteFile(`C:\Users\Administrator\Desktop\8.33.0\8.51\id_title.json`, bList, 0755)
}

func demoRegx(title string) string {
	//快手、创作者、补贴 、流量 、话题、活动
	ruleList := []string{"快手", "创作者", "补贴", "流量", "话题", "活动"}

	titleSplit := strings.Split(title, " ")
	newTitleList := make([]string, 0)
	for _, shortTitle := range titleSplit {
		shortTitle = strings.TrimSpace(shortTitle)

		for _, rule := range ruleList {
			for _, fuhao := range []string{"@", "#"} {
				var ruleStr string
				if fuhao == "@" {
					ruleStr = fmt.Sprintf(`%s.*`, fuhao)
				} else {
					ruleStr = fmt.Sprintf(`%s.*%s.*`, fuhao, rule)
				}

				re := regexp.MustCompile(ruleStr)
				if re.Match([]byte(shortTitle)) {
					//log.Println("标题命中 ", shortTitle, "|", ruleStr)

					shortTitle = re.ReplaceAllString(shortTitle, "")
					//log.Println("替换后的结果" + shortTitle)
					shortTitle = strings.TrimSpace(shortTitle)
				}
			}
		} //end titleSplit
		shortTitle = strings.TrimSpace(shortTitle)
		if shortTitle != "" {
			newTitleList = append(newTitleList, shortTitle)
		}
	}
	log.Println(strings.Join(newTitleList, " "))
	return strings.Join(newTitleList, " ")
}

func demokuaishou() {
	filename := `C:\Users\Administrator\Desktop\8.33.0\8.46\内容-标签四级目录-对应快手马甲号0510.xlsx`
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("现有三级目录马甲号")

	rows = rows[1:]
	for _, row := range rows {
		for _, colCell := range row {
			if len(row) < 7 {
				//fmt.Println(row)
				continue
			}
			_ = colCell
			titleUserId1[row[0]] = row[6]
			titleUserId1[row[1]] = row[6]
			titleUserId1[row[2]] = row[6]
			titleUserId1[row[3]] = row[6]
			//fmt.Print(k, colCell, "\t")
		}
	}

	delete(titleUserId1, "")
	fmt.Println(titleUserId1)

	b, _ := json.Marshal(titleUserId1)
	fmt.Println(string(b))
	//fmt.Println("package app_index\n var IosConfig = []byte(`" + string(b) + "`)")
	//ioutil.WriteFile("C:/Users/Administrator/Desktop/8.33.0/8.46/ios-app-search/ios_inner_search_data.go", []byte("package app_index\n var IosConfig = []byte(`"+string(b)+"`)"), 0755)

}

func aaaaa() {

}

func DemoFushi() {
	_, fullFilename, _, _ := runtime.Caller(0)
	fmt.Println(fullFilename)

	var filename string
	filename = path.Base(fullFilename)
	fmt.Println("filename=", filename)

	return

	var data = make(map[string]map[string][]string)
	data["备孕"] = make(map[string][]string)
	data["孕期"] = make(map[string][]string)
	data["辅食"] = make(map[string][]string)

	filename = `C:\Users\Administrator\Desktop\8.33.0\8.52\功效映射(1).xlsx`
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("Sheet1")

	//rows = rows[:]
	for _, row := range rows {

		for k, colCell := range row {
			if k == 0 || k == 1 {
				continue
			}
			if row[0] == "" {
				//log.Println(row[1])
				for a, _ := range data {
					log.Println(a, row[1], colCell)
					data[a][row[1]] = append(data[a][row[1]], colCell)
				}
			} else {
				if row[0] == "产后" {
					row[0] = "辅食"
				}

				data[row[0]][row[1]] = append(data[row[0]][row[1]], colCell)
			}

		}

	}

	//b, _ := json.Marshal(data["备孕"])
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
	//fmt.Println("package app_index\n var IosConfig = []byte(`" + string(b) + "`)")
	//ioutil.WriteFile("C:/Users/Administrator/Desktop/8.33.0/8.46/ios-app-search/ios_inner_search_data.go", []byte("package app_index\n var IosConfig = []byte(`"+string(b)+"`)"), 0755)

}
