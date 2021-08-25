package my_regexp

import (
	"log"
	"regexp"
)

func DemoMyRegexp() {
	demo := `<babytree_meta data="eyJsZW5ndGgiOjc0MSwibGFzdF9lZGl0X3RzIjoxNjI5MzU1NTA5LCJjX3R5cGUiOjEsInB1Yl90eXBlcyI6W10sIm5vX3B1Yl90eXBlcyI6W119"></babytree_meta>「数胎动时间」13:42~14:42 「胎动次数」5次<p><img src="http://pic08.babytreeimg.com/2021/0819/FjHFQBwr2kyjWvWBTo4fgz98aPRS_m.jpg" data="eyJ0YWciOiJpbWciLCJpZCI6IjE2MTc3ODk2NzYiLCJzX3R5cGUiOiJxaW5pdSIsImZpbGVfdHlwZSI6Ii5wbmciLCJmaWxlX3NpemUiOjc0MjU3NiwiYl93aXRkaCI6bnVsbCwiYl9oZWlnaHQiOjEyODAsImJfdXJsIjoiaHR0cDpcL1wvcGljMDcuYmFieXRyZWVpbWcuY29tXC8yMDIxXC8wODE5XC9GakhGUUJ3cjJreWpXdldCVG80Zmd6OThhUFJTX2IuanBnIiwibWJfd2l0ZGgiOm51bGwsIm1iX2hlaWdodCI6NDgwLCJtYl91cmwiOiJodHRwOlwvXC9waWMwNy5iYWJ5dHJlZWltZy5jb21cLzIwMjFcLzA4MTlcL0ZqSEZRQndyMmt5ald2V0JUbzRmZ3o5OGFQUlNfYi5qcGcifQ==" alt="journal_insert_pic_1617789676" _fcksavedurl="http://pic08.babytreeimg.com/2021/0819/FjHFQBwr2kyjWvWBTo4fgz98aPRS_m.jpg" _moz_resizing="true"></p>`

	///^<babytree_meta\s+data\=\"(.*?)\"\s*\><\/babytree_meta\>/
	compile, err := regexp.Compile(`^<babytree_meta\s+data\=\"(.*?)\"\s*\><\/babytree_meta\>`)
	if err != nil {
		panic(err)
	}

	s := compile.FindAllStringSubmatch(demo, -1)
	log.Println(s)
	if len(s) < 1 || len(s[0]) < 1 {
		return
	}

	//log.Println(compile.String())
}
