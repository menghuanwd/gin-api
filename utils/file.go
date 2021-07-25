package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func WriteLocalFile(url string, path string)  (str string) {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(path)
	io.Copy(out, bytes.NewReader(body))

	sEnc := base64.StdEncoding.EncodeToString(body)

	fmt.Println(GetMD5Hash(sEnc))

	return GetMD5Hash(sEnc)
}
