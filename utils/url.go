package utils

import "strings"

func FormatM3u8Suffix(url string, siteKey string) string {
	url = strings.Replace(url, "/data/org", "", -1)
	if strings.Contains(url, "index.m3u8") {
		url = strings.Replace(url, "index.m3u8", siteKey+".m3u8", -1)
	}

	return url
}
