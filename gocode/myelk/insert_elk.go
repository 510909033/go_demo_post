package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type User struct {
	Id   interface{}
	Name string
	Desc string
}

func insert() {
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 5)

	var demo *Demo
	for i := 0; i < 1; i++ {
		wg.Add(1)
		ch <- struct{}{}
		title := demo.GetOneLineText()

		go func(i int, title string) {
			defer func() {
				<-ch
			}()
			defer wg.Done()

			u := User{}
			//			u.Id = i
			u.Id = "lsdkflsfjlsf"
			u.Name = fmt.Sprintf("%d--name", i)
			u.Desc = title

			var body string
			if b1, err := json.Marshal(&u); err == nil {
				body = string(b1)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index: "test",
				//DocumentID: strconv.Itoa(i + 1),
				Body:    strings.NewReader(body),
				Refresh: "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), esclient)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d, req=%#v\n", res.Status(),
						r["result"],
						int(r["_version"].(float64)),
						req)
					log.Printf("res=%+v, a=%+v , r=%#v\n", res, res.String(), r)
					if a, err := ioutil.ReadAll(res.Body); true {
						fmt.Printf("\r\r--------\r%s, err=%+v\n", a, err)
					}
				}
			}
		}(i, title)
	}
	wg.Wait()

}
