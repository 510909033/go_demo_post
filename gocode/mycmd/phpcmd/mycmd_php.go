package phpcmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

/*
go执行php文件
*/

type Php struct {
}

func (s *Php) shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("php", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return err, stdout.String(), stderr.String()
}
func (s *Php) Start() {
	//	err, out, errout := Shellout("ls -ltr") //bash
	//	err, out, errout := Shellout(`C:/phpStudy/PHPTutorial/php/php-7.2.1-nts/php.exe E:\code_180622\all\aaa3-发私信\push-save_data.php`)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err, out, errout := s.shellout(`./demo_php.php`) //执行一个php
			//err, out, errout := s.shellout(`E:/push-save_data.php`) //执行一个php
			if err != nil {
				log.Printf("error: %v\n", err)
			}
			fmt.Println("--- stdout ---")
			fmt.Println(out)
			fmt.Println("--- stderr ---")
			fmt.Println(errout)

		}()
	}

	wg.Wait()
}

func (s *Php) Start2() {
	//	err, out, errout := Shellout("ls -ltr")
	//	err, out, errout := Shellout(`C:/phpStudy/PHPTutorial/php/php-7.2.1-nts/php.exe E:\code_180622\all\aaa3-发私信\push-save_data.php`)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err, out, errout := s.shellout(`./demo_php.php`) //执行一个php
			if err != nil {
				log.Printf("error: %v\n", err)
			}
			fmt.Println("--- stdout ---")
			fmt.Println(out)
			fmt.Println("--- stderr ---")
			fmt.Println(errout)

		}()
	}

	wg.Wait()
}

func tmp() {
	//	net.FileListener()

	//	net.DialUnix()

}
