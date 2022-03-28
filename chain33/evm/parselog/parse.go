package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func main() {
	// 将ms的统一转换成s
	fmt.Println("将需要解析的时间复制到命令行里")
	scanner := bufio.NewScanner(os.Stdin)
	str := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "#" {
			break
		}
		str = append(str, s)
	}
	fmt.Println("``````````````````````````````````````````")
	for _, v := range str {
		f, err := strconv.ParseFloat(v, 10)
		if err != nil {
			panic(err)
		}
		if f > 10 {
			f = f / 1000
		}
		fmt.Println(f)
	}
}
