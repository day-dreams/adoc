package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/alecthomas/participle/lexer"
	"os"
	"regexp"
	"strings"
)

// A custom lexer for INI files. This illustrates a relatively complex Regexp lexer, as well
// as use of the Unquote filter, which unquotes string tokens.
var iniLexer = lexer.Must(lexer.Regexp(
	`(?m)` +
		`(\s+)` +
		`|(//api)` +
	//`|(?P<Keyword>return)` +
		`|:` +
		`|(?P<Words>(.*)$)` +
	//`|(?P<Words>[\p{Han}]+)` +
		`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)`))

type INI struct {
	Comments []*Comment `@@*`
}

type Comment struct {
	//PrimaryPrefix string `@"//api"`
	TString string `"//api" @return":"`
	Content string `@Words`
}

func main() {
	filename := os.Args[len(os.Args)-1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	data := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		data = append(data, strings.TrimSuffix(line, "\n"))

	}

	api, err := ParseText(data)
	if err != nil {
		panic(err)
	}

	text := api.ToReadme()
	fmt.Println(text)

}

type TextType = uint

const (
	NotApiComment = iota
	ApiCommentName
	ApiCommentPath
	ApiCommentParam
	ApiCommentMethod
	ApiCommentReturn
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

type Api struct {
	name   string
	path   string
	params []ApiParam
	method string
	rdata  string
}

// IsApiComment 以//api开头的都是ApiComment
func IsApiComment(text string) bool {

	if strings.Index(text, "//api") != 0 {
		return false
	}
	return true
}

func (api *Api) ToReadme() string {

	rv := ""

	rv += fmt.Sprintf("## %s\n", api.name)
	rv += fmt.Sprintf("* 请求URL: %s\n", api.path)
	rv += fmt.Sprintf("* 请求方法: %s\n", api.method)
	rv += fmt.Sprintf("* 请求参数:\n\n")

	rv += fmt.Sprintf("|参数名|是否必选|参数类型|取值范围|说明|\n")
	rv += fmt.Sprintf("|:-:|-|-|-|-|\n")

	for _, param := range api.params {
		//rv += fmt.Sprintf("		* %s\n", param.String())
		rv += fmt.Sprintf("|%s|%s|%s|%s|%s|\n",
			param.name, param.required, param.t, param.scope, param.desc)
	}
	rv += fmt.Sprintf("\n")

	rv += fmt.Sprintf("* 返回数据示例: \n\n")
	rv += fmt.Sprintf("```json\n")
	rv += fmt.Sprintf("%s\n", api.rdata)
	rv += fmt.Sprintf("```\n")
	return rv
}

func ParseText(text []string) (Api, error) {

	api := Api{}

	for _, t := range text {

		if !IsApiComment(t) {
			continue
		}

		if block, err := splitByReg("^//api name:", t); err == nil {
			api.name = block
			//fmt.Printf("block:%v\n", block)
		}

		if block, err := splitByReg("^//api path:", t); err == nil {
			api.path = block
		}
		if block, err := splitByReg("^//api param:", t); err == nil {
			api.params = append(api.params, NewApiParam(block))
		}
		if block, err := splitByReg("^//api method:", t); err == nil {
			api.method = block
		}
		if block, err := splitByReg("^//api return:", t); err == nil {
			api.rdata = block
		}

	}
	return api, nil
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
