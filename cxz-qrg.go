package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/skip2/go-qrcode"
)

/*
# 传参变量定义
*/
var output string
var isFile bool
var isHelp bool

func initFlag() {
	flag.StringVar(&output, "m", "", "input string")
	flag.BoolVar(&isFile, "f", false, "input file.")
	flag.BoolVar(&isHelp, "h", false, "Help")
}

func main() {
	initFlag()
	flag.Parse()
	if isHelp {
		flag.Usage()
		return
	}
	if output != "" {
		if isFile {
			f, err := os.Open(output)
			if err != nil {
				fmt.Println("[ERR] " + err.Error())
				return
			}
			fmt.Println("[INFO] Opened file: " + f.Name())
			defer f.Close()
			if os.IsNotExist(err) {
				fmt.Println("[ERR] The file you specified is not exist.")
				return
			}
			var buffer = make([]byte, 1024)
			n, err := f.Read(buffer)
			if err != nil {
				fmt.Println("[ERR] Failed to read file.")
			}
			bts := strconv.Itoa(n)
			fmt.Println("[INFO] Read " + bts + " Bytes data from the file.")
			if ToPic("./fileoutput.png", string(buffer)) {
				fmt.Println("[INFO] Finished.")
			}
			return
		}
		if ToPic("./output.png", output) {
			fmt.Println("[INFO] Finished.")
		}
	} else {
		flag.Usage()
	}

}

// ToPic 将文本转化为二维码
func ToPic(filename, msg string) bool {
	err := qrcode.WriteFile(msg, qrcode.Highest, 256, filename)
	if err != nil {
		return false
	}
	return true
}

// 00001010
// 11110110
