package fshare

import "net/http"

type Config struct {
	Username      *string
	Password      *string
	HttpClient    *http.Client
	loginUrl      *string
	loginAppKey   *string
	downloadUrl   *string
	folderListUrl *string
}

func NewConfig(username string, password string) *Config {
	return &Config{
		Username:   PointerString(username),
		Password:   PointerString(password),
		HttpClient: http.DefaultClient,
	}
}

func (cfg *Config) GetLoginURL() string {
	if cfg.loginUrl != nil {
		return *cfg.loginUrl
	}
	return LoginURL
}

func (cfg *Config) GetLoginAppKey() string {
	if cfg.loginAppKey != nil {
		return *cfg.loginAppKey
	}
	return LoginAppKey
}

func (cfg *Config) GetDownloadURL() string {
	if cfg.downloadUrl != nil {
		return *cfg.downloadUrl
	}
	return DownloadURL
}

func (cfg *Config) GetFolderListURL() string {
	if cfg.folderListUrl != nil {
		return *cfg.folderListUrl
	}
	return FolderListURL
}
