package convertapi

import (
	"io"
	"net/http"
	"os"
)

type ResFile struct {
	client   *http.Client
	resp     *http.Response
	FileName string
	FileSize int
	Url      string
}

func (this *ResFile) download() (err error) {
	if this.resp == nil {
		this.resp, err = this.client.Get(this.Url)
	}
	return
}

func (this *ResFile) Read(p []byte) (n int, err error) {
	err = this.download()
	if err == nil {
		n, err = this.resp.Body.Read(p)
	}
	return
}

func (this *ResFile) ToFile(file *os.File) (err error) {
	_, err = io.Copy(file, this)
	return
}

func (this *ResFile) ToFilePath(path string) (err error) {
	file, err := os.Create(path)
	defer file.Close()
	if err == nil {
		err = this.ToFile(file)
	}
	return
}
