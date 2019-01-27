package wrapgroup

import (
	"bufio"
	"fmt"
	"github.com/imroc/req"
	"io"
	"os"
	"strings"
	"testing"
)
func TestGenerate(t *testing.T) {
	wg := Generate(50, 2)

	f, err := os.Open("urls.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		wg.Add()
		go func(line string) {
			defer wg.Done()
			download(line)
		}(line)
	}
	wg.Wait()
}

func download(pic string){
	progress := func(current, total int64) {
		fmt.Println(float32(current)/float32(total)*100, "%")
	}
	a := strings.Split(pic, "/")
	al := len(a)

	//req.SetTimeout(30 * time.Second)
	r, e := req.Get(pic, req.DownloadProgress(progress))
	if e == nil {
		filPath := "./pics/" + a[al - 1]
		fmt.Println(filPath)
		r.ToFile(filPath)
		fmt.Println("download complete")
	} else {
		fmt.Println("can not fetch the img")
	}
}
