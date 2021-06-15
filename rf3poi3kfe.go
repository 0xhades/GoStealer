package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var homeDir string

func normalizePath(path string) string {
	var err error
	if homeDir == "" {
		homeDir, err = os.UserHomeDir()
	}

	if err != nil {
		return ""
	}

	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(homeDir, path[2:])
	}
	return path
}

func fileCopy(src, dst string) bool {
	source, err := os.Open(src)
	if err != nil {
		return false
	}
	defer source.Close()

	destination, err := os.OpenFile(dst, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err == nil {
		return true
	}
	return false
}

func getAllFiles(dict string, list *[]string) {

	if stat, err := os.Stat(dict); err == nil && stat.IsDir() {
		files, err := ioutil.ReadDir(dict)
		if err != nil {
			return
		}

		for _, f := range files {
			if f.IsDir() {
				getAllFiles(filepath.Join(dict, f.Name()), list)
			} else {
				*list = append(*list, filepath.Join(dict, f.Name()))
			}
		}

	}

}

func walkSearch(filename string, path string) []string {

	var AllFilesWithin []string
	var cookiesFiles []string

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return nil
		}

		for _, f := range files {
			if f.IsDir() {
				getAllFiles(filepath.Join(path, f.Name()), &AllFilesWithin)
			} else {
				AllFilesWithin = append(AllFilesWithin, filepath.Join(path, f.Name()))
			}
		}

	}

	for _, f := range AllFilesWithin {
		_, file := filepath.Split(f)
		if file == filename {
			cookiesFiles = append(cookiesFiles, f)
		}
	}

	return cookiesFiles

}

func appendSliceToFile(filename string, data []string) error {

	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, v := range data {
		if _, err := f.WriteString(v + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func appendToFile(filename string, data string) error {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		return err
	}
	return nil
}

func writeToFile(filename string, data string) error {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

func stringSliceContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func readLinesToSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func readFile(path string) (string, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}
func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func downloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func runFile(filepath string) error {
	command := exec.Command(filepath)
	return command.Start()
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
