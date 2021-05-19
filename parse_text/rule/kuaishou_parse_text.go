package rule

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type RuleKuaiPhotoGetService struct {
}

type RuleKuaiPhotoGetResult struct {
	ErrList []interface{}
	List    []*RuleKuaiPhotoGetResultOne
}
type RuleKuaiPhotoGetResultOne struct {
	PhotodId      string
	PhotoTitle    string
	ShareUrl      string
	PlayCnt       string
	PoiName       string
	PoiProvince   string
	PoiCity       string
	CreateTime    string
	CoverUrl      string
	AuthorName    string
	AuthorIconUrl string
	FansCount     string
	CommentCount  string
	ShareCount    string
	LikeCnt       string
	PoiAddress    string
	PhotoTag      string
	PhotoDuration string
	WebpUrl       string
	PhotoWidth    string
	PhotoHeight   string
	CoverWidth    string
	CoverHeight   string
	PublishTime   string
	AuthorId      string
	Verified      string
}

func (result *RuleKuaiPhotoGetResult) GetError() []interface{} {
	return result.ErrList
}
func (result *RuleKuaiPhotoGetResult) GetList() interface{} {
	return result.List
}

func (service *RuleKuaiPhotoGetService) Parse(f *os.File) (IResult, error) {
	reader := bufio.NewReader(f)
	result := &RuleKuaiPhotoGetResult{ErrList: make([]interface{}, 0), List: []*RuleKuaiPhotoGetResultOne{}}
	for {
		line, isPrefix, err := reader.ReadLine()
		_ = isPrefix
		if err != nil {
			if err == io.EOF {
				return result, nil
			}
			result.ErrList = append(result.ErrList, err)
			continue
		}

		//todo
		if len(result.List) > 3 {
			break
		}

		//photoId	photoTitle	shareUrl	playCnt	poiName	poiProvince	poiCity	photoDuration	createTime	coverUrl
		splitList := bytes.Split(line, []byte("\t"))

		size := 26
		if len(splitList) != size {
			err := fmt.Errorf("数据\t分隔后长度不符合， len=%d, rawLine=%s", len(splitList), line)
			result.ErrList = append(result.ErrList, err)
			continue
		}

		i := 0
		getIndex := func() int {
			i++
			return i
		}
		one := &RuleKuaiPhotoGetResultOne{
			PhotodId:      string(splitList[i]),
			PhotoTitle:    string((splitList[getIndex()])),
			ShareUrl:      string((splitList[getIndex()])),
			PlayCnt:       string((splitList[getIndex()])),
			PoiName:       string((splitList[getIndex()])),
			PoiProvince:   string((splitList[getIndex()])),
			PoiCity:       string((splitList[getIndex()])),
			CreateTime:    string((splitList[getIndex()])),
			CoverUrl:      string((splitList[getIndex()])),
			AuthorName:    string((splitList[getIndex()])),
			AuthorIconUrl: string((splitList[getIndex()])),
			FansCount:     string((splitList[getIndex()])),
			CommentCount:  string((splitList[getIndex()])),
			ShareCount:    string((splitList[getIndex()])),
			LikeCnt:       string((splitList[getIndex()])),
			PoiAddress:    string((splitList[getIndex()])),
			PhotoTag:      string((splitList[getIndex()])),
			PhotoDuration: string((splitList[getIndex()])),
			WebpUrl:       string((splitList[getIndex()])),
			PhotoWidth:    string((splitList[getIndex()])),
			PhotoHeight:   string((splitList[getIndex()])),
			CoverWidth:    string((splitList[getIndex()])),
			CoverHeight:   string((splitList[getIndex()])),
			PublishTime:   string((splitList[getIndex()])),
			AuthorId:      string((splitList[getIndex()])),
			Verified:      string((splitList[getIndex()])),
		}

		if one, err := service.check(one); err != nil {
			result.ErrList = append(result.ErrList, err)
		} else {
			result.List = append(result.List, one)
		}

	}
	return nil, nil
}

func (service *RuleKuaiPhotoGetService) check(one *RuleKuaiPhotoGetResultOne) (*RuleKuaiPhotoGetResultOne, error) {
	if one == nil {
		return nil, errors.New("one==nil")
	}

	if one.PhotodId == "" {
		return nil, fmt.Errorf("photoId==空, one=%#v", *one)
	}
	if one.CoverUrl == "" {
		return nil, fmt.Errorf("CoverUrl==空, one=%#v", *one)
	}

	return one, nil
}
