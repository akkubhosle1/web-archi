package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Read the input URL from standard input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	url, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}
	url = strings.TrimSpace(url)

	// Call the WaybackURL command-line tool to get a list of URLs that match the input URL pattern.
	fmt.Printf("Searching for URLs matching %s...\n", url)
	cmd := exec.Command("waybackurls", url)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching URLs:", err)
		return
	}
	allURLs := strings.Split(string(output), "\n")

	// Categorize the URLs based on their file extension.
	jpgURLs := make(map[string]bool)
	pngURLs := make(map[string]bool)
	jsonURLs := make(map[string]bool)
	svgURLs := make(map[string]bool)
	xmlURLs := make(map[string]bool)
	zipURLs := make(map[string]bool)
	tarURLs := make(map[string]bool)
	htmlURLs := make(map[string]bool)
	phpURLs := make(map[string]bool)
	jspURLs := make(map[string]bool)
	aspxURLs := make(map[string]bool)
	otherURLs := make(map[string]bool)
    cfmURLs := make(map[string]bool)
    batURLs := make(map[string]bool)
	cURLs := make(map[string]bool)
	cgiURLs := make(map[string]bool)
	cssURLs := make(map[string]bool)
	dllURLs := make(map[string]bool)
	exeURLs := make(map[string]bool)
	htmURLs := make(map[string]bool)
	incURLs := make(map[string]bool)
	jhtmlURLs := make(map[string]bool)
	jsURLs := make(map[string]bool)
	jsaURLs := make(map[string]bool)
	logURLs := make(map[string]bool)
	mdbURLs := make(map[string]bool)
	nsfURLs := make(map[string]bool)
	pcapURLs := make(map[string]bool)
	php2URLs := make(map[string]bool)
	php3URLs := make(map[string]bool)
	php4URLs := make(map[string]bool)
	php5URLs := make(map[string]bool)
	php6URLs := make(map[string]bool)
	php7URLs := make(map[string]bool)
	phpsURLs := make(map[string]bool)
	phtURLs := make(map[string]bool)
	phtmlURLs := make(map[string]bool)
	plURLs := make(map[string]bool)
	regURLs := make(map[string]bool)
	shURLs := make(map[string]bool)
	shtmlURLs := make(map[string]bool)
	sqlURLs := make(map[string]bool)
	swfURLs := make(map[string]bool)
	txtURLs := make(map[string]bool)
	woff2URLs := make(map[string]bool)
	woffURLs := make(map[string]bool)
	ttfURLs := make(map[string]bool)

	for _, url := range allURLs {
		ext := getFileExtension(url)
		switch ext {
		case ".jpg":
			jpgURLs[url] = true
		case ".png":
			pngURLs[url] = true
		case ".json":
			jsonURLs[url] = true
		case ".svg":
			svgURLs[url] = true
		case ".xml":
			xmlURLs[url] = true
		case ".zip":
			zipURLs[url] = true
		case ".jsp":
			jspURLs[url] = true
		case ".aspx":
			aspxURLs[url] = true
		case ".cfm":
			cfmURLs[url] = true
		case ".bat":
			batURLs[url] = true
		case ".c":
			cURLs[url] = true
		case ".cgi":
			cgiURLs[url] = true
		case ".css":
			cssURLs[url] = true
		case ".dll":
			dllURLs[url] = true
		case ".exe":
			exeURLs[url] = true
		case ".htm":
			htmURLs[url] = true
		case ".inc":
			incURLs[url] = true
		case ".jhtml":
			jhtmlURLs[url] = true
		case ".js":
			jsURLs[url] = true
		case ".jsa":
			jsaURLs[url] = true
		case ".log":
			logURLs[url] = true
		case ".mdb":
			mdbURLs[url] = true
		case ".nsf":
			nsfURLs[url] = true
		case ".pcap":
			pcapURLs[url] = true
		case ".php2":
			php2URLs[url] = true
		case ".php3":
			php3URLs[url] = true
		case ".php4":
			php4URLs[url] = true
		case ".php5":
			php5URLs[url] = true
		case ".php6":
			php6URLs[url] = true
		case ".php7":
			php7URLs[url] = true
		case ".phps":
			phpsURLs[url] = true
		case ".pht":
			phtURLs[url] = true
		case ".phtml":
			phtmlURLs[url] = true
		case ".pl":
			plURLs[url] = true
		case ".reg":
			regURLs[url] = true
		case ".sh":
			shURLs[url] = true
		case ".shtml":
			shtmlURLs[url] = true
		case ".sql":
			sqlURLs[url] = true
		case ".swf":
			swfURLs[url] = true
		case ".txt":
			txtURLs[url] = true
		case ".woff2":
			woff2URLs[url] = true
		case ".woff":
			woffURLs[url] = true
		case ".ttf":
			ttfURLs[url] = true
		default:
			otherURLs[url] = true
		}
	}

	// Save the URLs to files.
	saveURLs("jpg_urls.txt", jpgURLs)
	saveURLs("png_urls.txt", pngURLs)
	saveURLs("json_urls.txt", jsonURLs)
	saveURLs("svg_urls.txt", svgURLs)
	saveURLs("xml_urls.txt", xmlURLs)
	saveURLs("zip_urls.txt", zipURLs)
	saveURLs("tar_urls.txt", tarURLs)
	saveURLs("html_urls.txt", htmlURLs)
	saveURLs("php_urls.txt", phpURLs)
	saveURLs("jsp_urls.txt", jspURLs)
	saveURLs("aspx_urls.txt", aspxURLs)
	saveURLs("cfm_urls.txt", cfmURLs)
	saveURLs("bat_urls.txt", batURLs)
	saveURLs("c_urls.txt", cURLs)
	saveURLs("cgi_urls.txt", cgiURLs)
	saveURLs("css_urls.txt", cssURLs)
	saveURLs("dll_urls.txt", dllURLs)
	saveURLs("htm_urls.txt", htmURLs)
	saveURLs("inc_urls.txt", incURLs)
	saveURLs("jhtml_urls.txt", jhtmlURLs)
	saveURLs("js_urls.txt", jsURLs)
	saveURLs("jsa_urls.txt", jsaURLs)
	saveURLs("log_urls.txt", logURLs)
	saveURLs("mdb_urls.txt", mdbURLs)
	saveURLs("nsf_urls.txt", nsfURLs)
	saveURLs("pcap_urls.txt", pcapURLs)
	saveURLs("php2_urls.txt", php2URLs)
	saveURLs("php3_urls.txt", php3URLs)
	saveURLs("php4_urls.txt", php4URLs)
	saveURLs("php5_urls.txt", php5URLs)
	saveURLs("php6_urls.txt", php6URLs)
	saveURLs("php7_urls.txt", php7URLs)
	saveURLs("phps_urls.txt", phpsURLs)
	saveURLs("pht_urls.txt", phtURLs)
	saveURLs("phtml_urls.txt", phtmlURLs)
	saveURLs("pl_urls.txt", plURLs)
	saveURLs("reg_urls.txt", regURLs)
	saveURLs("sh_urls.txt", shURLs)
	saveURLs("shtml_urls.txt", shtmlURLs)
	saveURLs("sql_urls.txt", sqlURLs)
	saveURLs("swf_urls.txt", swfURLs)
	saveURLs("txt_urls.txt", txtURLs)
	saveURLs("xml_urls.txt", xmlURLs)
	saveURLs("other_urls.txt", otherURLs)

}

// getFileExtension returns the file extension of a URL.
func getFileExtension(url string) string {
	parts := strings.Split(url, ".")
	if len(parts) > 1 {
		return "." + parts[len(parts)-1]
	}
	return ""
}

// saveURLs saves the URLs to a file.
func saveURLs(filename string, urls map[string]bool) {
if len(urls) == 0 {
return
}
file, err := os.Create(filename)
if err != nil {
fmt.Fprintln(os.Stderr, "Error creating file:", err)
return
}
defer file.Close()
for url := range urls {
fmt.Fprintln(file, url)
}
fmt.Printf("Saved %d URLs to %s\n", len(urls), filename)
}
