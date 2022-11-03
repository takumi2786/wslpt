package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/golang/glog"
	"golang.org/x/xerrors"
)

/*
標準入力文字列に対するスキャナー(*bufio.Scanner)を取得する。
isDebug=Trueの場合、INPUT_PATHのファイルに書かれたテキスト読み込む
*/
func getScanner(isDebug bool) *bufio.Scanner {
	// NOTE: golangでは、ローカル変数のアドレスを戻り値としても問題ない
	//       https://zenn.dev/rookxx/articles/golang-stack-and-heap
	if DEBUG {
		f, e := os.Open(INPUT_PATH)
		if e != nil {
			log.Fatal(e)
		}
		return bufio.NewScanner(f)
	} else {
		return bufio.NewScanner(os.Stdin)
	}
}

// 標準入力の文字列を読み込む
func initialize() string {
	scanStrVal := func(s *bufio.Scanner) string {
		s.Scan()
		text := s.Text()
		return text
	}

	// スキャナーを取得
	s := getScanner(DEBUG)
	text := scanStrVal(s)
	return text
}

const DEBUG bool = true
const INPUT_PATH = "./input.txt"

func main() {
	var path string
	var replaced string
	fmt.Println("Enter windows path: ")
	path = initialize()
	replaced = strings.Replace(path, "\\", "/", -1)
	re := regexp.MustCompile(`[A-Z]:`)
	matched := re.FindAllStringSubmatch(replaced, -1)
	if len(matched) != 1 {
		err := xerrors.New("failed to get drive symbol. path: " + path)
		glog.Error(err)
	}
	symbol := strings.ToLower(strings.Replace(matched[0][0], ":", "", 1))
	replaced = re.ReplaceAllString(replaced, "/mnt/"+symbol)
	fmt.Println(replaced)
}
