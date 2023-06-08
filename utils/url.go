package utils

import "strings"

const (
	ImgUrlPrefix  = "/data/org"
	DownUrlPrefix = "/data/down"
)

func FormatUrl(url string) string {
	if url == "" {
		return ""
	}

	if strings.Contains(url, ImgUrlPrefix) {
		url = strings.Replace(url, ImgUrlPrefix, "", -1)
	} else if strings.Contains(url, DownUrlPrefix) {
		url = strings.Replace(url, DownUrlPrefix, "", -1)
	}

	return url
}

func FormatM3u8Suffix(url string, siteKey string) string {
	url = strings.Replace(url, "/data/org", "", -1)
	if strings.Contains(url, "index.m3u8") {
		url = strings.Replace(url, "index.m3u8", siteKey+".m3u8", -1)
	}

	return url
}
