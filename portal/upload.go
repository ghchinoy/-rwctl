package portal

import (
	"github.com/ghchinoy/rwctl/control"
	"strings"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"mime/multipart"
	"path/filepath"
	"io"
	"net/http"
)

// UploadFiles uploads a file or files to the Portal CMS
func UploadFiles(files []string, config control.Configuration, path string, debug bool) {

	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return
	}

	uploadURI := config.URL + path

	extraParams := map[string]string{
		"none": "really",
	}

	for _, v := range files {
		if debug {
			log.Printf("Uploading %s ...\n", v)
		}
		if strings.HasSuffix(v, ".zip") {
			uploadURI += "?unpack=true"
		}
		//log.Println(uploadURI)
		statusCode, err := uploadFile(client, v, extraParams, uploadURI, debug)
		if err != nil {
			log.Fatalf("Issues. %v : %s", statusCode, err)
		}
		if statusCode != 200 || debug {
			log.Printf("Upload status %v", statusCode)
		}
	}
}

func uploadFile(client *http.Client, uploadFilePath string, extras map[string]string, uploadURI string, debug bool) (int, error) {
	var uploadStatus int

	//var request *http.Request
	request, err := newFileUploadRequest(uploadURI, extras, "File", uploadFilePath)
	if err != nil {
		log.Fatalln(err)
		return uploadStatus, err
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	control.AddCsrfHeader(request, client)

	// debug
	if debug {
		log.Println("* URL", uploadURI)
		log.Println("* Upload Path", uploadFilePath)
		for k, v := range request.Header {
			log.Printf("* %s: %s", k, v)
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return uploadStatus, err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close()

	uploadStatus = resp.StatusCode

	if uploadStatus != 200 {
		b, _ := ioutil.ReadAll(body)
		log.Println("* uploadFile", string(b))
	}

	//log.Printf("Upload status %v", resp.StatusCode)

	return uploadStatus, nil
}

func newFileUploadRequest(uri string, params map[string]string, paramName string, path string) (*http.Request, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	/*
		// for any extra params, map of string keys and vals
			for key, val := range params {
				_ = writer.WriteField(key, val)
			}
	*/

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

