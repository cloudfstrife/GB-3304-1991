package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jszwec/csvutil"
	log "github.com/sirupsen/logrus"
)

//Nation 民族
type Nation struct {
	//Name 民族名称
	Name string `csv:"NAME"`
	//RomeCodeInside 罗马代码（对内）
	RomeCodeInside string `csv:"ROME_CODE_INSIDE"`
	//RomeCodeOutSide 罗马代码（对外）
	RomeCodeOutSide string `csv:"ROME_CODE_OUTSIDE"`
	//Code 字母代码
	Code string `csv:"CODE"`
	//Number 数字代码
	Number int64 `csv:"NUMBER"`
}

func main() {
	var ips []byte
	var err error
	var file *os.File
	if file, err = os.Open("GB_3304_91.csv"); err != nil {
		log.Error(err)
	}
	defer file.Close()
	if ips, err = ioutil.ReadAll(file); err != nil {
		log.Error(err)
	}
	var nation []Nation
	if err := csvutil.Unmarshal(ips, &nation); err != nil {
		fmt.Println("error:", err)
	}

	for _, u := range nation {
		fmt.Printf("%+v\n", u)
	}
}
