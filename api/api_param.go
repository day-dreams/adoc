package api

import (
	"errors"
	"fmt"
	"strings"
)

type ApiParam struct {
	name     string //参数名字
	required string //是否必选
	t        string //参数类型
	scope    string //参数取值范围
	desc     string //说明
}

func (param *ApiParam) String() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s",
		param.name, param.required, param.t, param.scope, param.desc)
}

func NewApiParam(body string) ApiParam {
	blocks := strings.Split(body, "`")
	if len(blocks) != 5 {
		fmt.Printf("wrong param comment format:[%s]\n", body)
		panic(errors.New("wrong param comment"))
	}
	return ApiParam{
		name:     blocks[0],
		required: blocks[1],
		t:        blocks[2],
		scope:    blocks[3],
		desc:     blocks[4]}
}
