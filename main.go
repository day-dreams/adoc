package main

import (
	"adoc.zhangnan.xyz/api"
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	apis, err := ParseText(data)
	if err != nil {
		panic(err)
	}

	for _, api := range apis {
		readme := api.ToReadme()
		fmt.Println(readme)
	}

}

func ParseText(text []string) ([]api.Api, error) {

	rv := []api.Api{}

	a := api.Api{}
	for _, t := range text {

		if !api.IsApiComment(t) {
			continue
		}

		if api.IsApiBegin(t) {
			a = api.Api{}
			continue
		}

		if api.IsApiEnd(t) {
			rv = append(rv, a)
			continue
		}

		if ok, block := api.IsApiName(t); ok {
			a.Name = block
		}
		if ok, block := api.IsApiPath(t); ok {
			a.Path = block
		}
		if ok, param := api.IsApiParam(t); ok {
			a.Params = append(a.Params, param)
		}
		if ok, method := api.IsApiMethod(t); ok {
			a.Method = method
		}
		if ok, rdata := api.IsApiReturn(t); ok {
			a.Rdata = rdata
		}

	}

	return rv, nil
}
