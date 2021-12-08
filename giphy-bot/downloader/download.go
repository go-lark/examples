package downloader

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

// DownloadFileWithProxy ...
func DownloadFileWithProxy(path, proxy, imageURL string) error {
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	proxyURL, err := url.Parse(proxy)
	tr := &http.Transport{
		Proxy:        http.ProxyURL(proxyURL),
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// DownloadParallelly ...
func DownloadParallelly(path, imageURL string) error {
	res, err := http.Head(imageURL)
	if err != nil {
		return err
	}

	if lengthHeader, ok := res.Header["Content-Length"]; ok && len(lengthHeader) > 0 {
		contentLength, err := strconv.Atoi(lengthHeader[0])
		if err != nil {
			return err
		}

		numOfRoutines := 50
		lengthOfSub := contentLength / numOfRoutines
		diff := contentLength % numOfRoutines
		body := make([]string, numOfRoutines+1)
		var wg sync.WaitGroup
		for routineIndex := 0; routineIndex < numOfRoutines; routineIndex++ {
			min := lengthOfSub * routineIndex
			max := lengthOfSub * (routineIndex + 1)
			if routineIndex == numOfRoutines-1 {
				max += diff // Add the remaining bytes in the last request
			}
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				client := &http.Client{}
				req, _ := http.NewRequest("GET", imageURL, nil)
				rangeHeader := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1) // Add the data for the Range header of the form "bytes=0-100"
				req.Header.Add("Range", rangeHeader)
				resp, err := client.Do(req)
				if err != nil {
					return
				}
				defer resp.Body.Close()
				reader, _ := ioutil.ReadAll(resp.Body)
				body[index] = string(reader)
			}(routineIndex)
		}
		wg.Wait()

		var bufs []byte
		for routineIndex := 0; routineIndex < numOfRoutines; routineIndex++ {
			buf := body[routineIndex]
			bufs = append(bufs, buf...)
		}
		ioutil.WriteFile(path, bufs, 0o666)
	}
	return nil
}
