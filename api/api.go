package api

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Api struct {
	Name   string
	Path   string
	Params []ApiParam
	Method string
	Rdata  string
}

// IsApiComment 以//api开头的都是ApiComment
func IsApiComment(text string) bool {

	if strings.Index(text, "//api") != 0 &&
		strings.Index(text, "//@api") != 0 {
		return false
	}
	return true
}

// IsApiBegin 判断是否遇到了api注释的开头标记
func IsApiBegin(text string) bool {
	if text == "//@api" {
		return true
	}
	return false
}

// IsApiEnd 判断是否遇到了api注释的结尾标记
func IsApiEnd(text string) bool {
	if text == "//@api end" {
		return true
	}
	return false
}

// IsApiName 判断text是否是api注释的名字，如果是，返回该name
func IsApiName(text string) (ok bool, name string) {

	block, err := splitByReg("^//api name:", text)
	if err != nil {
		return false, ""
	}
	return true, block
}

// IsApiPath 判断text是否是api注释的请求路径，如果是，返回该path
func IsApiPath(text string) (ok bool, path string) {

	block, err := splitByReg("^//api path:", text)
	if err != nil {
		return false, ""
	}
	return true, block
}

// IsApiParam 判断text是否是api注释的请求参数，如果是，返回该param
func IsApiParam(text string) (ok bool, param ApiParam) {

	block, err := splitByReg("^//api param:", text)
	if err != nil {
		return false, ApiParam{}
	}
	return true, NewApiParam(block)
}

// IsApiMethod 判断text是否是api注释的请求方法，如果是，返回该method
func IsApiMethod(text string) (ok bool, method string) {

	block, err := splitByReg("^//api method:", text)
	if err != nil {
		return false, ""
	}
	return true, block
}

// IsApiDesc 判断text是否是api注释的返回实例数据，如果是，返回该示例数据
func IsApiReturn(text string) (ok bool, rdata string) {

	block, err := splitByReg("^//api return:", text)
	if err != nil {
		return false, ""
	}
	return true, block
}

func splitByReg(pattern string, text string) (string, error) {

	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("err:%v\n", reg)
		return "", err
	}

	block := reg.FindString(text)
	if len(block) == 0 {
		return "", errors.New("no such pattern found")
	}

	return text[len(block):], nil
}

func (api *Api) ToReadme() string {

	rv := ""

	rv += fmt.Sprintf("## %s\n", api.Name)
	rv += fmt.Sprintf("* 请求URL: %s\n", api.Path)
	rv += fmt.Sprintf("* 请求方法: %s\n", api.Method)
	rv += fmt.Sprintf("* 请求参数:\n\n")

	rv += fmt.Sprintf("|参数名|是否必选|参数类型|取值范围|说明|\n")
	rv += fmt.Sprintf("|:-:|-|-|-|-|\n")

	for _, param := range api.Params {
		//rv += fmt.Sprintf("		* %s\n", param.String())
		rv += fmt.Sprintf("|%s|%s|%s|%s|%s|\n",
			param.name, param.required, param.t, param.scope, param.desc)
	}
	rv += fmt.Sprintf("\n")

	rv += fmt.Sprintf("* 返回数据示例: \n\n")
	rv += fmt.Sprintf("```json\n")
	rv += fmt.Sprintf("%s\n", api.Rdata)
	rv += fmt.Sprintf("```\n")
	return rv
}
