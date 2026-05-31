package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
	"log"
	"io"
	"strconv"
	
	"github.com/fatih/color"
	"golang.org/x/term"
)

const (
	toolPath       = "/data/data/com.termux/files/home/go/bin/nsscan"
	expiryDate     = "2026-12-30"
)

func generateDateKey(t time.Time) int {
	year, month, day := t.Date()
	return year*10000 + int(month)*100 + day
}

func getExpiryKey() int {
	expiryTime, _ := time.Parse("2006-01-02", expiryDate)
	return generateDateKey(expiryTime)
}

func isExpired() bool {
	currentKey := generateDateKey(time.Now())
	expiryKey := getExpiryKey()
	return currentKey > expiryKey
}

func checkExpiration() bool {
	if isExpired() {
		lockFile := "/data/data/com.termux/files/home/go/pkg/sumdb/sum.golang.org/goo"
		os.WriteFile(lockFile, []byte("go"), 0644)
		return true
	}
	
	if _, err := os.Stat("/data/data/com.termux/files/home/go/pkg/sumdb/sum.golang.org/goo"); err == nil {
		return true
	}
	
	return false
}

var (
	colorR1 = color.New(color.FgHiRed)
	colorR2 = color.New(color.FgRed)
	colorG1 = color.New(color.FgHiGreen)
	colorG2 = color.New(color.FgGreen)
	colorY1 = color.New(color.FgHiYellow)
	colorY2 = color.New(color.FgYellow)
	colorB1 = color.New(color.FgHiBlue)
	colorB2 = color.New(color.FgBlue)
	colorM1 = color.New(color.FgHiMagenta)
	colorM2 = color.New(color.FgMagenta)
	colorC1 = color.New(color.FgHiCyan)
	colorC2 = color.New(color.FgCyan)
	colorW1 = color.New(color.FgHiWhite)
)

var (
	globalFlagThreads         = 100
	scanDirectFlagFilename    string
	scanDirectFlagOutput      = "results.txt"
	scanDirectFlagHideLocations = []string{
		"https://jio.com/BalanceExhaust",
		"http://choof.ooredoo.dz",
		"http://internet.djezzy.dz", 
		"https://safaricom.zerod.live/?c=77",
	}
	scanDirectFlagMethod      = "HEAD"
	scanDirectFlagTimeout     = 10
	scanDirectFlagInputType   string
	scanDirectFlagRawInput    string
	scanDirectFlagSkipCodes   []int
	scanDirectFlagCustomPort  string
	fileMutex                 sync.Mutex 
)

var serverColors = map[string]*color.Color{
	"cloudflare":    colorG1,
	"cf-ray":       colorG1,
	"akamai":       colorY1,
	"akamaigslb":   colorY1,
	"cloudfront":   colorC1,
	"awselb":       colorC1,
	"amazons3":     colorC1,
	"fastly":       colorM1,
	"varnish":      colorM1,
	"microsoft":    colorC2,
	"azure":        colorC2,
	"cachefly":     colorY2,
	"alibaba":      colorY2,
	"alicdn":       colorY2,
	"tencent":      colorM2,
	"nginx":        colorY1,
	"bunny":        colorM1,
	"bunnycdn":     colorM1,
	"BunnyCDN-CCU1-1124": colorM1,
	"bunnycdn-ccu1-1124": colorM1,
	"apache":       colorY1,
	"wit":          colorC1,
	"wit application server": colorC1,
	"volt-adc":     colorM1,
	"volt":         colorM1,
	"sffe":         colorM1,
	"ESF":          colorM1,
	"esf":          colorM1,
	"google frontend": colorM1,
	"Google Frontend": colorM1,
	"Golfe2":       colorM1,
	"golfe2":       colorM1,
	"gind1":        colorM1,
	"gind2":        colorM1,
	"cdn77":        colorC2,
	"stackpath":    colorC2,
	"google":       colorM1,
	"akamai.net":   colorY1,
	"edgekey":      colorY1,
	"fastly-ssl":   colorM1,
	"google cloud": colorM1,
	"gws":          colorM1,
	"msedge":       colorC2,
	"windows-azure": colorC2,
	"qcloud":       colorM2,
	"limelight":    colorY1,
	"llnw":         colorY1,
	"incapsula":    colorR1,
	"imperva":      colorR1,
	"incapdns":     colorR1,
	"sucuri":       colorR1,
	"sucuri cloudproxy": colorR1,
	"ovh":          colorM2,
	"keycdn":       colorC1,
	"leaseweb":     colorY1,
	"beluga":       colorC2,
	"belugacdn":    colorC2,
	"highwinds":    colorY1,
	"netdna":       colorY1,
	"quantil":      colorM1,
	"chinacache":   colorM2,
	"wangsu":       colorM2,
	"baidu":        colorM2,
	"baiduyun":     colorM2,
	"openresty":    colorB2,
	"cowboy":       colorY1,
	"lighttpd":     colorY1,
	"iis":          colorC2,
	"microsoft-iis": colorC2,
	"tomcat":       colorY1,
	"jetty":        colorY1,
	"litespeed":    colorY1,
	"caddy":        colorY1,
	"cloudflare-waf": colorG1,
	"barracuda":     colorR1,
	"f5":            colorR1,
	"big-ip":        colorR1,
	"fortinet":      colorR1,
	"fortiweb":      colorR1,
	"palo alto":     colorR1,
	"radware":       colorR1,
	"citrix":        colorR1,
	"netscaler":     colorR1,
	"kona":          colorY1,
	"prolexic":      colorY1,
	"reblaze":       colorR1,
	"denyall":       colorR1,
	"wallarm":       colorR1,
	"signal sciences": colorR1,
	"haproxy":      colorY1,
	"envoy":        colorY1,
	"traefik":      colorY1,
	"alb":          colorC1,
	"nlb":          colorC1,
	"cloud":        colorG1,
	"google lb":    colorM1,
	"azure lb":     colorC2,
	"wordpress":    colorY1,
	"joomla":       colorY1,
	"drupal":       colorY1,
	"shopify":      colorM1,
	"magento":      colorM1,
	"express":      colorY1,
	"nodejs":       colorY1,
	"gunicorn":     colorY1,
	"php":          colorY1,
	"ruby":         colorY1,
	"python":       colorY1,
}

func getServerColor(server string) *color.Color {
	if strings.TrimSpace(server) == "" || strings.ToLower(server) == "unknown" {
		return colorW1
	}
	serverLower := strings.ToLower(server)
	for k, v := range serverColors {
		if strings.Contains(serverLower, k) {
			return v
		}
	}
	return colorB1
}

var globalCtx *Ctx 

func main() {
	if checkExpiration() {
		colorR1.Print("\n  ⛔ Tool expired on 30 December 2026\n")
		fmt.Println()
		colorG1.Print("  💬 Contact @NS_SCAN to renew\n")
		fmt.Println()
		colorC1.Print("  🔗 Telegram channel: @NS_SCAN_CH\n\n")
		os.Exit(1)
	}

	log.SetOutput(io.Discard)
	ClearScreen()
	PrintBanner()

	if len(os.Args) > 1 {
		processInput(os.Args[1])
	} else {
		AskScanDirectOptions()
	}

	dir := filepath.Dir(scanDirectFlagOutput)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	os.WriteFile(scanDirectFlagOutput, []byte(""), 0644)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		showCursor()
		fmt.Println()
		colorR1.Println("\n\n 🛑 Scan interrupted by user!")
		if globalCtx != nil {
			globalCtx.mx.Lock()
			totalComplete := globalCtx.ScanComplete
			totalHits := len(globalCtx.ScanSuccessList)
			globalCtx.mx.Unlock()
			
			colorY1.Printf(" 📊 Progress till now -> Total Scanned: %d | Total Live Found (Saved): %d\n", totalComplete, totalHits)
		}
		colorG1.Printf(" 📦 Full results with server/codes are already saved inside: %s\n", scanDirectFlagOutput)
		
		colorR1.Print("\n ⛔ Press Enter to exit...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		os.Exit(0)
	}()

	printHeaders()
	RunDirectScan()
	
	colorR1.Print("\n ⛔ Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')

	colorY1.Print("\n 👋 BYE BYE") 
	colorW1.Print(" | ")
	colorC1.Print("NSSCAN") 
	colorW1.Print(" | ") 
	colorG1.Print("Your")
	colorW1.Print(":")
	colorM1.Print("—͟͞͞𝙉𝙎 𝙃𝙖𝙘𝙠𝙚𝙧")
	fmt.Println("\n")
}

func ClearScreen() {
	fmt.Print("\033c")
}

func PrintBanner() {
	fmt.Print("\n")
	banner := []struct {
		text  string
		color *color.Color
	}{
		{"  _   _  ____   ____    _    _   _ ", color.New(color.FgHiBlue)},
		{" | \\ | |/ ___| / ___|  / \\  | \\ | |", color.New(color.FgBlue)},
		{" |  \\| |\\___ \\| |     / _ \\ |  \\| |", color.New(color.FgHiCyan)},
		{" | |\\  | ___) | |___ / ___ \\| |\\  |", color.New(color.FgCyan)},
		{" |_| \\_|____/  \\____/_/   \\_\\_| \\_|", color.New(color.FgHiGreen)},
	}

	for _, line := range banner {
		line.color.Println(line.text)
		time.Sleep(15 * time.Millisecond)
	}

	footer := []struct {
		text  string
		color *color.Color
	}{
		{"\n Made By NS Hacker ", colorB2},
		{"Simple, Fast, Accurate ", colorC1},
		{"🔗 t.me/NS_SCAN", colorG1},
	}

	for _, part := range footer {
		for _, c := range part.text {
			part.color.Printf("%c", c)
			time.Sleep(1 * time.Millisecond)
		}
	}
	fmt.Println("\n")
}

func printHeaders() {
	colorW1.Printf("%-17s%-5s%-5s%-14s %s\n", "---", "----", "----", "------", "----")
	colorC1.Printf("%-17s", "IP")
	colorB1.Printf("%-5s", "PORT")
	colorY1.Printf("%-5s", "CODE")
	colorM1.Printf("%-14s ","SERVER")
	colorG1.Printf("%s\n","HOST")

	colorW1.Printf("%-17s%-5s%-5s%-14s %s\n", "---", "----", "----", "------", "----")  
	fmt.Println()
}

func processInput(userInput string) {
	userInput = strings.TrimSpace(userInput)

	if userInput == "" {
		colorR1.Println("    ⛔ Input cannot be empty !")
		return
	}

	if _, _, err := net.ParseCIDR(userInput); err == nil {
		scanDirectFlagInputType = "cidr"
		scanDirectFlagRawInput = userInput
		return
	}

	pathsToTry := []string{
		userInput,
		userInput + ".txt",
		filepath.Join(os.Getenv("HOME"), userInput),
		filepath.Join(os.Getenv("HOME"), userInput+".txt"),
		filepath.Join("/sdcard", userInput),
		filepath.Join("/sdcard", userInput+".txt"),
		filepath.Join("/sdcard/Download", userInput),
		filepath.Join("/sdcard/Download", userInput+".txt"),
		filepath.Join("/sdcard/Download/Telegram", userInput),
		filepath.Join("/sdcard/Download/Telegram", userInput+".txt"),
		filepath.Join("/sdcard/Telegram/Telegram Files", userInput),
		filepath.Join("/sdcard/Telegram/Telegram Files", userInput+".txt"),
	}

	for _, path := range pathsToTry {
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			scanDirectFlagInputType = "file"
			scanDirectFlagFilename = path
			
			file, err := os.Open(path)
			if err != nil {
				return
			}
			defer file.Close()
			
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if _, _, err := net.ParseCIDR(line); err == nil {
					scanDirectFlagInputType = "file-with-cidr"
					break
				}
			}
			return
		}
	}

	if strings.HasSuffix(userInput, ".txt") || !strings.Contains(userInput, ".") {
		colorR1.Printf("    ❌ File '%s' not found in common storage paths !\n", userInput)
		fmt.Println()
		return
	}

	scanDirectFlagInputType = "single"
	scanDirectFlagRawInput = userInput
}

func AskScanDirectOptions() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()

	for {
		colorC1.Print(" 📥 Enter ( file / example.com / CIDR ) = ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)		

		if userInput == "" {
			colorR1.Println("    ❌ Input cannot be empty !\n")
			continue
		}

		processInput(userInput)
		if scanDirectFlagInputType == "file" {
			colorY1.Printf("    📂 Found file : %s ✅\n", scanDirectFlagFilename)
		} else if scanDirectFlagInputType == "" {
			continue
		}
		break
	}
	fmt.Println()

	colorG1.Print(" ⏳ Threads to use")
	colorW1.Print(" default-")
	colorY1.Print("(100) = ")
	threadStr, _ := reader.ReadString('\n')
	threadStr = strings.TrimSpace(threadStr)
	if threadStr != "" {
		fmt.Sscanf(threadStr, "%d", &globalFlagThreads)
	} else {
		globalFlagThreads = 100
	}
	fmt.Println()

	colorB1.Print(" ⏰ Timeout in seconds")
	colorW1.Print(" default-")
	colorY1.Print("(10) = ")
	timeoutStr, _ := reader.ReadString('\n')
	timeoutStr = strings.TrimSpace(timeoutStr)
	if timeoutStr != "" {
		fmt.Sscanf(timeoutStr, "%d", &scanDirectFlagTimeout)
	} else {
		scanDirectFlagTimeout = 10
	}
	fmt.Println()

	colorY1.Print(" 💾 Output file name")
	colorW1.Print(" default-")
	colorY1.Print("(results.txt) = ")
	outStr, _ := reader.ReadString('\n')
	outStr = strings.TrimSpace(outStr)
	if outStr != "" {
		scanDirectFlagOutput = outStr
	} else {
		scanDirectFlagOutput = "results.txt"
	}
	fmt.Println()

	if !strings.Contains(scanDirectFlagOutput, "/") && !strings.Contains(scanDirectFlagOutput, "\\") {
		colorM1.Print(" 📁 Output save to ")
		colorG2.Print("termux")
		colorY1.Print("-(t)")
		colorW1.Print(" default ")
		colorG2.Print("storage-")  
		colorY1.Print("(s) = ")  
		location, _ := reader.ReadString('\n')
		location = strings.ToLower(strings.TrimSpace(location))
		
		var baseDir string
		if location == "t" {
			baseDir = os.Getenv("HOME")
		} else {
			baseDir = "/sdcard"
		}
		
		scanDirectFlagOutput = filepath.Join(baseDir, scanDirectFlagOutput)
		fmt.Println()
	}

	for {
		colorC1.Print(" 🌐 Custom ports, or ")
		colorY1.Print("leave blank = ")
		portStr, _ := reader.ReadString('\n')
		portStr = strings.TrimSpace(portStr)

		if portStr == "" {
			scanDirectFlagCustomPort = ""
			break
		}

		if portStr == "80" || portStr == "443" {
			colorR1.Println("    ❌ 80 and 443 are default ports, no need to specify!")
			fmt.Println()
			continue
		}

		ports := strings.Split(portStr, ",")
		validPorts := make([]string, 0)
		hasInvalid := false
		
		for _, p := range ports {
			p = strings.TrimSpace(p)
			portNum, err := strconv.Atoi(p)
			if err != nil || portNum < 1 || portNum > 65535 {
				colorR1.Printf("    ❌ Invalid port number: %s (must be 1-65535)\n", p)
				hasInvalid = true
				continue
			}
			if portNum == 80 || portNum == 443 {
				colorR1.Printf("    ❌ %d is default and should not be specified\n", portNum)
				hasInvalid = true
				continue
			}
			validPorts = append(validPorts, p)
		}
		
		if hasInvalid {
			fmt.Println()
			continue
		}

		if len(validPorts) > 0 {
			scanDirectFlagCustomPort = strings.Join(validPorts, ",")
			break
		}
	}
	fmt.Println()

	colorR1.Print(" 📬 HTTP Method GET/POST/TRACE")
	colorW1.Print(" default-")
	colorY1.Print("(HEAD) = ")
	method, _ := reader.ReadString('\n')
	method = strings.ToUpper(strings.TrimSpace(method))

	validMethods := map[string]bool{
		"GET":     true,
		"POST":    true,
		"HEAD":    true,
		"PUT":     true,
		"DELETE":  true,
		"CONNECT": true,
		"OPTIONS": true,
		"TRACE":   true,
		"PATCH":   true,
	}

	if method == "" {
		method = "HEAD"
	} else if !validMethods[method] {
		colorR1.Println("    ❌ Invalid HTTP method! Using default (HEAD)")
		method = "HEAD"
	}

	scanDirectFlagMethod = method
	fmt.Println()

	colorB2.Print(" 🚦 Skip status codes?")
	colorY1.Print(" (302,307,etc)")
	colorB2.Print(" or ")
	colorY1.Print("leave blank = ")
	skipCodesStr, _ := reader.ReadString('\n')
	skipCodesStr = strings.TrimSpace(skipCodesStr)

	scanDirectFlagSkipCodes = []int{}

	if strings.ToLower(skipCodesStr) != "n" && skipCodesStr != "" {
		codes := strings.Split(skipCodesStr, ",")
		for _, codeStr := range codes {
			codeStr = strings.TrimSpace(codeStr)
			if code, err := strconv.Atoi(codeStr); err == nil {
				if code >= 100 && code <= 599 {
					scanDirectFlagSkipCodes = append(scanDirectFlagSkipCodes, code)
				} else {
					colorR1.Printf("    ❌ Invalid HTTP status code: %d (must be 100-599)\n", code)
				}
			} else {
				colorR1.Printf("    ❌ Invalid number: %s\n", codeStr)
			}
		}
	}
	fmt.Println()
}

func trimString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-1] + ".."
	}
	return s
}

type scanDirectRequest struct {
	Domain string
}

type scanDirectResponse struct {
	Request    *scanDirectRequest
	NetIPList  []net.IP
	StatusCode int
	Server     string
	Location   string
}

type Ctx struct {
	ScanSuccessList []interface{}
	ScanComplete    int64
	dataList        []*QueueScannerScanParams
	mx              sync.Mutex
	startTime       time.Time
}

func (c *Ctx) Log(a ...interface{}) {
	fmt.Printf("\r\033[2K%s\n", fmt.Sprint(a...))
}

func (c *Ctx) Logf(f string, a ...interface{}) {
	c.Log(fmt.Sprintf(f, a...))
}

func (c *Ctx) LogReplace(a ...string) {
	scanSuccess := len(c.ScanSuccessList)
	scanComplete := func() int64 { c.mx.Lock(); defer c.mx.Unlock(); return c.ScanComplete }()
	scanCompletePercentage := float64(scanComplete) / float64(len(c.dataList)) * 100
	elapsed := time.Since(c.startTime).Round(time.Second)
	
	s := fmt.Sprintf(
		"%s - %.2f%% - %d/%d - Hit:%d - Sec%v ≈ %s", 
		scanDirectFlagMethod,
		scanCompletePercentage, 
		scanComplete, 
		len(c.dataList), 
		scanSuccess,
		elapsed.Seconds(),
		strings.Join(a, " "),
	)

	if termWidth, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		w := termWidth - 3
		if len(s) >= w {
			s = s[:w] + "..."
		}
	}

	fmt.Print("\r\033[2K", s, "\r")
}

func (c *Ctx) LogReplacef(f string, a ...interface{}) {
	c.LogReplace(fmt.Sprintf(f, a...))
}

func (c *Ctx) ScanSuccess(a interface{}, fn func()) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if fn != nil {
		fn()
	}

	c.ScanSuccessList = append(c.ScanSuccessList, a)
}

type QueueScannerScanParams struct {
	Name string
	Data interface{}
}

type QueueScannerScanFunc func(c *Ctx, a *QueueScannerScanParams)
type QueueScannerDoneFunc func(c *Ctx)

type QueueScanner struct {
	threads  int
	scanFunc QueueScannerScanFunc
	queue    chan *QueueScannerScanParams
	wg       sync.WaitGroup
	ctx      *Ctx
}

func NewQueueScanner(threads int, scanFunc QueueScannerScanFunc) *QueueScanner {
	t := &QueueScanner{
		threads:  threads,
		scanFunc: scanFunc,
		queue:    make(chan *QueueScannerScanParams, threads*1),
		ctx:      &Ctx{startTime: time.Now()},
	}
	globalCtx = t.ctx 

	for i := 0; i < t.threads; i++ {
		go t.run()
	}

	return t
}

func (s *QueueScanner) run() {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		a, ok := <-s.queue
		if !ok {
			break
		}
		s.scanFunc(s.ctx, a)

		s.ctx.mx.Lock()
		s.ctx.ScanComplete++
		s.ctx.mx.Unlock()

		s.ctx.LogReplace(a.Name)
	}
}

func (s *QueueScanner) Add(dataList ...*QueueScannerScanParams) {
	s.ctx.dataList = append(s.ctx.dataList, dataList...)
}

func (s *QueueScanner) Start(doneFunc QueueScannerDoneFunc) {
	hideCursor()
	defer showCursor()

	for _, data := range s.ctx.dataList {
		s.queue <- data
	}
	close(s.queue)

	s.wg.Wait()

	if doneFunc != nil {
		doneFunc(s.ctx)
	}
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func newHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(scanDirectFlagTimeout) * time.Second,
				KeepAlive: 0,
			}).DialContext,
			TLSHandshakeTimeout:   time.Duration(scanDirectFlagTimeout) * time.Second,
			ResponseHeaderTimeout: time.Duration(scanDirectFlagTimeout) * time.Second,
		},
		Timeout: time.Duration(scanDirectFlagTimeout) * time.Second,
	}
}

func performRequest(domain, scheme, method string) *http.Response {
	url := fmt.Sprintf("%s://%s", scheme, domain)
	
	if scanDirectFlagCustomPort != "" {
		url = fmt.Sprintf("%s://%s:%s", scheme, domain, scanDirectFlagCustomPort)
	}

	httpReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}

	client := newHTTPClient()
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil
	}
	return httpRes
}

func scanDirect(c *Ctx, p *QueueScannerScanParams) {
	req := p.Data.(*scanDirectRequest)
	method := scanDirectFlagMethod
	if method == "" {
		method = "HEAD"
	}

	var httpRes *http.Response
	var port string
	
	if scanDirectFlagCustomPort != "" {
		port = scanDirectFlagCustomPort
		httpRes = performRequest(req.Domain, "http", method)
		if httpRes == nil {
			httpRes = performRequest(req.Domain, "https", method)
		}
	} else {
		httpRes = performRequest(req.Domain, "http", method)
		if httpRes != nil {
			port = "80"
		} else {
			httpRes = performRequest(req.Domain, "https", method)
			if httpRes != nil {
				port = "443"
			}
		}
	}

	if httpRes == nil {
		return
	}
	defer httpRes.Body.Close()

	for _, skipCode := range scanDirectFlagSkipCodes {
		if httpRes.StatusCode == skipCode {
			return
		}
	}

	hServer := httpRes.Header.Get("Server")
	hLocation := httpRes.Header.Get("Location")

	for _, hideLoc := range scanDirectFlagHideLocations {
		if hLocation == hideLoc {
			return
		}
	}

	netIPs, _ := net.LookupIP(req.Domain)
	ip := "unknown"
	if len(netIPs) > 0 {
		ip = netIPs[0].String()
	}

	if hServer == "" || strings.ToLower(hServer) == "unknown" {
		hServer = "Unknown"
	}
	
	serverColor := getServerColor(hServer)

	res := &scanDirectResponse{
		Request:    req,
		NetIPList:  netIPs,
		StatusCode: httpRes.StatusCode,
		Server:     hServer,
		Location:   hLocation,
	}
	c.ScanSuccess(res, nil)

	// 🔥 [UPDATED RUNTIME FILE APPENDING WITH STATUS & SERVER]
	fileMutex.Lock()
	f, err := os.OpenFile(scanDirectFlagOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		// Clean aur customized details file me turant append hongi
		dataLine := fmt.Sprintf("%s | PORT: %s | CODE: %d | SERVER: %s\n", req.Domain, port, httpRes.StatusCode, hServer)
		f.WriteString(dataLine)
		f.Close()
	}
	fileMutex.Unlock()

	s := fmt.Sprintf(  
		"%-17s%-5s%-4d %-14s %s",  
		trimString(ip, 15),  
		port,  
		httpRes.StatusCode,  
		trimString(hServer, 12),  
		req.Domain,  
	)  
	c.Log(serverColor.Sprint(s))
}

func RunDirectScan() {
	var domains []string
	var domainInput []string
	startTime := time.Now()
	maxCIDRIPsPerRange := 1000
	currentProcessingCIDR := ""
	var allSuccessResponses []*scanDirectResponse
	var allDomainInput []string

	switch scanDirectFlagInputType {
	case "file":
		file, err := os.Open(scanDirectFlagFilename)
		if err != nil {
			colorR1.Printf("    ⛔ Error opening file : %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domain := strings.TrimSpace(scanner.Text())
			if domain != "" {
				domains = append(domains, domain)
				domainInput = append(domainInput, domain)
				allDomainInput = append(allDomainInput, domain)
			}
		}

	case "file-with-cidr":
		file, err := os.Open(scanDirectFlagFilename)
		if err != nil {
			colorR1.Printf("    ⛔ Error opening file : %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			
			if _, ipnet, err := net.ParseCIDR(line); err == nil {
				currentProcessingCIDR = line
				colorY1.Printf("\n🔍 Processing CIDR: %s\n", currentProcessingCIDR)
				
				domains = domains[:0]
				domainInput = domainInput[:0]
				
				ip := ipnet.IP
				ipCount := 0
				for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
					domains = append(domains, ip.String())
					domainInput = append(domainInput, ip.String())
					allDomainInput = append(allDomainInput, ip.String())
					ipCount++
					
					if ipCount >= maxCIDRIPsPerRange {
						colorY1.Printf("    ⚠️ Limited to first %d IPs from this CIDR\n", maxCIDRIPsPerRange)
						break
					}
				}
				
				queueScanner := NewQueueScanner(globalFlagThreads, scanDirect)

				for _, domain := range domains {
					queueScanner.Add(&QueueScannerScanParams{
						Name: domain,
						Data: &scanDirectRequest{
							Domain: domain,
						},
					})
				}
				
				var batchResponses []*scanDirectResponse
				var mx sync.Mutex
				
				queueScanner.Start(func(c *Ctx) {
					mx.Lock()
					defer mx.Unlock()
					
					for _, data := range c.ScanSuccessList {
						if res, ok := data.(*scanDirectResponse); ok {
							batchResponses = append(batchResponses, res)
						}
					}
				})
				
				allSuccessResponses = append(allSuccessResponses, batchResponses...)
				
			} else {
				if currentProcessingCIDR != "" {
					currentProcessingCIDR = ""
					colorY1.Println("\n 🏁 Finished CIDR batch")
				}
				
				domains = []string{line}
				domainInput = []string{line}
				allDomainInput = append(allDomainInput, line)
				
				queueScanner := NewQueueScanner(globalFlagThreads, scanDirect)

				queueScanner.Add(&QueueScannerScanParams{
					Name: line,
					Data: &scanDirectRequest{
						Domain: line,
					},
				})
				
				var batchResponses []*scanDirectResponse
				var mx sync.Mutex
				
				queueScanner.Start(func(c *Ctx) {
					mx.Lock()
					defer mx.Unlock()
					
					for _, data := range c.ScanSuccessList {
						if res, ok := data.(*scanDirectResponse); ok {
							batchResponses = append(batchResponses, res)
						}
					}
				})
				
				allSuccessResponses = append(allSuccessResponses, batchResponses...)
			}
		}
		
		if currentProcessingCIDR != "" {
			colorY1.Println("\n 🏁 Finished CIDR batch")
		}
		printFinalSummary(allSuccessResponses, startTime)
		return

	case "cidr":
		ip, ipnet, err := net.ParseCIDR(scanDirectFlagRawInput)
		if err != nil {
			colorR1.Printf("    ⛔ Error parsing CIDR : %v\n", err)
			return
		}

		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
			domains = append(domains, ip.String())
			domainInput = append(domainInput, ip.String())
			allDomainInput = append(allDomainInput, ip.String())
		}

	case "single":
		domains = append(domains, scanDirectFlagRawInput)
		domainInput = append(domainInput, scanDirectFlagRawInput)
		allDomainInput = append(allDomainInput, scanDirectFlagRawInput)
	}

	if len(domains) == 0 {
		colorR1.Println(" ⛔ No valid targets found !")
		return
	}

	queueScanner := NewQueueScanner(globalFlagThreads, scanDirect)

	for _, domain := range domains {
		queueScanner.Add(&QueueScannerScanParams{
			Name: domain,
			Data: &scanDirectRequest{
				Domain: domain,
			},
		})
	}
	
	queueScanner.Start(func(c *Ctx) {
		var mx sync.Mutex
		mx.Lock()
		defer mx.Unlock()
		
		for _, data := range c.ScanSuccessList {
			if res, ok := data.(*scanDirectResponse); ok {
				allSuccessResponses = append(allSuccessResponses, res)
			}
		}
	})
	
	printFinalSummary(allSuccessResponses, startTime)
}

func printFinalSummary(successResponses []*scanDirectResponse, startTime time.Time) {
	if len(successResponses) == 0 {
		colorR1.Println("\n ⛔ No live hosts found !")
		return
	}
	fmt.Println()
	colorG1.Printf("\n 📦 All detailed results are successfully saved inside: %s\n", scanDirectFlagOutput)
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
