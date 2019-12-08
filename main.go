package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func pow(num int, s int) int64 {
	sum := 1
	for i := 0; i < s; i++ {
		sum *= num
	}
	return int64(sum)
}

var (
	sort = flag.String("sort", "size", "sort output")
	d    = flag.String("d", ".", "Directory to process")
	a    = flag.Bool("a", false, "Print all info")
)

func hrSize(fsize int64) string {
	if fsize < 1024 {
		return strconv.Itoa(int(fsize)) + "B"
	}
	if fsize >= 1024 && fsize < pow(1024, 2) {
		fsize = fsize / 1024
		return strconv.Itoa(int(fsize)) + "KB"
	}
	if fsize >= pow(1024, 2) && fsize < pow(1024, 3) {
		fsize = fsize / pow(1024, 2)
		return strconv.Itoa(int(fsize)) + "MB"
	}
	if fsize >= pow(1024, 3) && fsize < pow(1024, 4) {
		fsize = fsize / pow(1024, 3)
		return strconv.Itoa(int(fsize)) + "GB"
	}
	return "0"
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize, _ := strconv.Atoi(strconv.Itoa(int(file.Size())))
	fsize := hrSize(int64(fSize))
	fmt.Printf("%s %s %s \n", fsize, time, file.Name())
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			if *sort == "size" {
				fmt.Println(file.Name())
			} else {
				fmt.Println("Cannot sort without date output")
			}
		}
	}
}
