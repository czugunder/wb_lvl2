package wget

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type wget struct {
	domain    string
	saveDir   string
	transport *http.Transport
	pages     map[string]bool
	mu        *sync.RWMutex
	wg        *sync.WaitGroup
}

func NewWget() *wget {
	return &wget{
		pages: make(map[string]bool),
		mu:    &sync.RWMutex{},
		wg:    &sync.WaitGroup{},
		transport: &http.Transport{
			MaxIdleConns:       5,
			IdleConnTimeout:    10 * time.Second,
			DisableCompression: true,
		},
	}
}

func (w *wget) SetDomain(domain string) {
	w.domain = domain
}

func (w *wget) SetSaveDirectory(saveDir string) {
	w.saveDir = saveDir
	if err := os.Chdir(saveDir); err != nil {
		log.Fatalln("can't set save directory")
	}
}

func (w *wget) SaveSite() {
	w.AddPage(w.domain)
	w.GetSite()
}

func (w *wget) GetSite() {
	var pageLen int
	for {
		pageLen = len(w.pages)
		w.mu.RLock()
		for url, inspected := range w.pages {
			if !inspected {
				w.wg.Add(1)
				go w.ProcessPage(url)
			}
		}
		w.mu.RUnlock()
		w.wg.Wait()
		if pageLen == len(w.pages) {
			break
		}
	}
	w.transport.CloseIdleConnections()
}

func (w *wget) ProcessPage(url string) {
	defer w.wg.Done()
	page, err := w.GetPage(url)
	w.SetInspected(url)
	if err != nil {
		log.Printf("couldn't get page %s because %s\n", url, err.Error())
	}
	isJunk, err := w.Save(url, page)
	if err != nil {
		log.Printf("couldn't save page %s because %s\n", url, err.Error())
	}
	if !isJunk {
		links := w.ParseLinks(string(page))
		var cleanLink string
		for _, v := range links {
			cleanLink = v[6 : len(v)-1]
			if w.IsURLInternal(cleanLink) {
				w.AddPage(cleanLink)
			}
		}
	}
}

func (w *wget) AddPage(url string) {
	w.mu.Lock()
	w.pages[url] = false
	w.mu.Unlock()
}

func (w *wget) SetInspected(url string) {
	w.mu.Lock()
	w.pages[url] = true
	w.mu.Unlock()
}

func (w *wget) GetSavePathAndName(url string) (path, name string) {
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	prep := strings.Replace(url, w.domain, "", 1)
	if prep == "" {
		path = ""
		name = "index"
	} else {
		slash := strings.LastIndex(prep, "/")
		if slash == -1 {
			log.Fatalln(url) // debug
		}
		path = prep[:slash]
		name = prep[slash+1:]
	}
	return
}

func (w *wget) IsURLInternal(url string) (is bool) {
	if strings.Index(url, w.domain) == 0 {
		is = true
	}
	return
}

func (w *wget) Save(url string, page []byte) (bool, error) {
	if w.IsJunkPage(url) {
		return true, nil
	}
	relPath, name := w.GetSavePathAndName(url)
	if relPath != "" {
		if err := os.MkdirAll(w.saveDir+relPath, os.ModePerm); err != nil {
			return false, err
		}
	}
	if ext := w.ContainsExtension(name); ext == -1 {
		name += ".html"
	}
	if file, err := os.Create(w.saveDir + relPath + "/" + name); err != nil {
		return false, err
	} else {
		if _, err = file.Write(page); err != nil {
			return false, err
		}
	}
	return false, nil
}

func (w *wget) ContainsExtension(name string) int {
	return strings.Index(name, ".")
}

func (w *wget) IsJunkPage(url string) bool {
	if strings.Contains(url, "?") {
		return true
	}
	if strings.Contains(url, "=") {
		return true
	}
	if strings.Contains(url, "&") {
		return true
	}
	return false
}

func (w *wget) ParseLinks(page string) []string {
	r := regexp.MustCompile(`href="(.*?)"`)
	return r.FindAllString(page, -1)
}

func (w *wget) GetPage(url string) ([]byte, error) {
	client := &http.Client{Transport: w.transport}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buff, nil
}
