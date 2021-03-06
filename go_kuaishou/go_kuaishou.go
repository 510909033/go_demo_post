package go_kuaishou

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go_demo_post/parse_text"
	"go_demo_post/parse_text/rule"
	"log"
	"math"
	"reflect"
)

func DemoKuaishou() {
	client := resty.New()
	_ = client

	fmt.Printf("%d\n", uint64(math.MaxUint64)-1)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
		//SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		Post("http://172.16.7.242:12345/demo-json-post.php")

	fmt.Println(err)
	fmt.Println(resp)

	parseTextService := parse_text.NewParseText()
	filename := `C:\Users\Administrator\Desktop\8.33.0\8.46\kuaiphoto.json`
	filename = `C:\Users\Administrator\Desktop\8.33.0\8.46\kuaiphoto1.json`
	ruleService := &rule.RuleKuaiPhotoGetService{}
	if parseRes, err := parseTextService.Parse(filename, ruleService); true {
		log.Println(err)
		log.Println(parseRes.GetList())
		//if v,ok:=parseRes.GetList().([]*rule.RuleKuaiPhotoGetResultOne);ok{
		if v, ok := parseRes.GetList().([]interface{}); ok {
			log.Println(v)
		}
		val := reflect.ValueOf(parseRes.GetList())
		log.Println("kind", val.Kind())
		if val.Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				log.Println("i=", i, "value=", val.Index(i))
			}
		}
	}

}
