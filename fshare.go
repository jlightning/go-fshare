package fshare

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
)

type commonApiResult struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
}

type loginResult struct {
	commonApiResult
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
}

type downloadResult struct {
	commonApiResult
	Location string `json:"location"`
}

type folderResult struct {
	LinkCode string `json:"linkcode"`
}

type Client struct {
	cfg       *Config
	token     string
	sessionID string
}

func NewClient(cfg *Config) *Client {
	return &Client{cfg: cfg}
}

func (c *Client) Login() error {
	if c.cfg == nil || c.cfg.Username == nil || c.cfg.Password == nil {
		return errors.New("invalid username / password")
	}

	jsonData, err := json.Marshal(map[string]string{
		"user_email": *c.cfg.Username,
		"password":   *c.cfg.Password,
		"app_key":    c.cfg.GetLoginAppKey(),
	})
	req, err := http.NewRequest("POST", c.cfg.GetLoginURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	resp, err := c.cfg.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var loginResult loginResult
	if err = json.NewDecoder(resp.Body).Decode(&loginResult); err != nil {
		return err
	}
	if loginResult.Code != 200 || resp.StatusCode != 200 {
		return errors.New("cannot login: " + loginResult.Msg)
	}

	c.token = loginResult.Token
	c.sessionID = loginResult.SessionID

	return nil
}

func (c *Client) GetDownloadURL(fshareURL string) (string, error) {
	jsonData, err := json.Marshal(map[string]string{
		"token": c.token,
		"url":   fshareURL,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.cfg.GetDownloadURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "session_id", Value: c.sessionID})

	resp, err := c.cfg.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var downloadResult downloadResult
	if err = json.NewDecoder(resp.Body).Decode(&downloadResult); err != nil {
		return "", err
	}
	if downloadResult.Location == "" || resp.StatusCode != 200 {
		return "", errors.New("cannot get download url: " + downloadResult.Msg)
	}

	return downloadResult.Location, nil
}

func (c *Client) GetFileInfo(url string) ([]string, error) {
	jsonData, err := json.Marshal(map[string]interface{}{
		"token": c.token,
		"url":   url,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", FileInfoURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session_id", Value: c.sessionID})
	resp, err := c.cfg.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	logBody(&resp.Body)

	return nil, nil
}

func (c *Client) IsFolderUrl(url string) bool {
	return strings.HasPrefix(url, FolderBaseURL)
}

func (c *Client) GetAllFolderURLs(fshareFolderURL string) (res []string, err error) {
	resLock := sync.Mutex{}
	wg := sync.WaitGroup{}

	pageChannel := make(chan int, 10)
	stopChannel := make(chan error, 100)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				page, ok := <-pageChannel
				if !ok {
					break
				}
				urls, err := c.GetFolderURLs(fshareFolderURL, page)
				if err != nil || len(urls) == 0 {
					stopChannel <- err
				}
				resLock.Lock()
				res = append(res, urls...)
				resLock.Unlock()
			}
		}()
	}

	page := 1

forLoop:
	for {
		select {
		case pageChannel <- page:
		case <-stopChannel:
			close(pageChannel)
			break forLoop
		}
		page++
	}

	wg.Wait()
	return
}

func (c *Client) GetFolderURLs(fshareFolderURL string, page int) (res []string, err error) {
	jsonData, err := json.Marshal(map[string]interface{}{
		"token":     c.token,
		"url":       fshareFolderURL,
		"dirOnly":   0,
		"pageIndex": page,
		"limit":     1,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.cfg.GetFolderListURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session_id", Value: c.sessionID})
	resp, err := c.cfg.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	logBody(&resp.Body)

	var folderResult []folderResult
	if err = json.NewDecoder(resp.Body).Decode(&folderResult); err != nil {
		return nil, err
	}
	if folderResult == nil || resp.StatusCode != 200 {
		return nil, errors.New("cannot get folder urls")
	}

	for _, item := range folderResult {
		res = append(res, FileBaseURL+item.LinkCode)
	}

	return
}
