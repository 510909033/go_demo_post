package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Content struct {
	Path    string
	Content string
}

func main() {
	logfile := `E:\logs\tmp\log1.log`
	f, _ := os.Create(logfile)

	dir := `C:\phpStudy\PHPTutorial\WWW\remote-pregenancy\scripts`

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			bl, _ := ioutil.ReadFile(path)
			content := Content{
				Path:    path,
				Content: string(bl),
			}
			encodeBl, _ := json.Marshal(&content)
			f.WriteString(string(encodeBl))
			f.WriteString("\n")

		}

		return nil
	})
}
