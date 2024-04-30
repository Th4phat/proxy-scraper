package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s [output file]\n", os.Args[0])
	}
	uri_list := []string{
		"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt",
		"https://free-proxy-list.net",
		"https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all",
		"https://free-proxy-list.net/",
		"https://raw.githubusercontent.com/rdavydov/proxy-list/main/proxies_anonymous/http.txt",
		"https://raw.githubusercontent.com/MuRongPIG/Proxy-Master/main/http.txt",
		"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/http.txt",
		"https://raw.githubusercontent.com/MuRongPIG/Proxy-Master/main/http.txt",
		"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies/http.txt",
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-http.txt",
		"https://raw.githubusercontent.com/roosterkid/openproxylist/main/HTTPS_RAW.txt",
		"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/http.txt",
		"https://raw.githubusercontent.com/rdavydov/proxy-list/main/proxies_anonymous/http.txt",
		"https://raw.githubusercontent.com/officialputuid/KangProxy/KangProxy/http/http.txt",
		"https://raw.githubusercontent.com/yuceltoluyag/GoodProxy/main/raw.txt",
		"https://raw.githubusercontent.com/enseitankado/proxine/main/proxy/http.txt",
		"https://raw.githubusercontent.com/enseitankado/proxine/main/proxy/https.txt",
		"https://raw.githubusercontent.com/Anonym0usWork1221/Free-Proxies/main/http.txt",
		"https://proxyspace.pro/https.txt",
		"https://proxyspace.pro/http.txt",
	}
	outraw := getProxiesa(uri_list)
	nodupe := RemoveDuplicate(outraw)
	err := os.WriteFile(os.Args[1], []byte(strings.Join(nodupe, "\n")), 0644)
	if err != nil {
		fmt.Println(err)
	}

}
func getProxiesa(urls []string) []string {
	var proxies []string //
	for _, url := range urls {
		switch url {
		case "https://free-proxy-list.net/":
			response, err := http.Get(url)
			if err != nil {
				return []string{}
			}
			defer response.Body.Close()
			html, err := io.ReadAll(response.Body)
			if err != nil {
				return []string{}
			}
			re := regexp.MustCompile(`<tr><td>(\d+.\d+.\d+.\d+)</td><td>(\d+)</td>`)
			matches := re.FindAllStringSubmatch(string(html), -1)
			for _, match := range matches {
				proxies = append(proxies, fmt.Sprintf("%s:%s", match[1], match[2]))
			}
		default:
			response, err := http.Get(url)
			if err != nil {
				return []string{}
			}
			defer response.Body.Close()
			html, err := io.ReadAll(response.Body)
			if err != nil {
				return []string{}
			}
			lines := strings.Split(string(html), "\n")
			for _, line := range lines {
				proxies = append(proxies, strings.Trim(line, "\n"))
			}
		}
	}

	return proxies
}

func RemoveDuplicate(proxies []string) []string {
	uniqueProxies := make(map[string]struct{})
	for _, proxy := range proxies {
		uniqueProxies[proxy] = struct{}{}
	}
	var result []string
	for proxy := range uniqueProxies {
		result = append(result, proxy)
	}
	return result
}
