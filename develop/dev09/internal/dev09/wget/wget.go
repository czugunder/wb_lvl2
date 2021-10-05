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

// Wget - основной тип в скарпере
type Wget struct {
	domain    string
	saveDir   string
	transport *http.Transport
	pages     map[string]bool
	mu        *sync.RWMutex
	wg        *sync.WaitGroup
}

// NewWget создает экземпляр Wget
func NewWget() *Wget {
	return &Wget{
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

// SetDomain задает домен для скачивания
func (w *Wget) SetDomain(domain string) {
	w.domain = domain
}

// SetSaveDirectory задает путь директории, куда сохранить сайт
func (w *Wget) SetSaveDirectory(saveDir string) {
	w.saveDir = saveDir
	if err := os.Chdir(saveDir); err != nil {
		log.Fatalln("can't set save directory")
	}
}

// SaveSite запускает процесс сохранения сайта
func (w *Wget) SaveSite() {
	w.AddPage(w.domain)
	w.GetSite()
}

// GetSite последовательно скачивает все страницы сайта, исключая уже скаченные
func (w *Wget) GetSite() {
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

// ProcessPage обрабатывает страницу, сохраняет если все ок
func (w *Wget) ProcessPage(url string) {
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

// AddPage добавляет ссылку в очередь на обработку
func (w *Wget) AddPage(url string) {
	w.mu.Lock()
	w.pages[url] = false
	w.mu.Unlock()
}

// SetInspected поднимает флаг обработки
func (w *Wget) SetInspected(url string) {
	w.mu.Lock()
	w.pages[url] = true
	w.mu.Unlock()
}

// GetSavePathAndName выдает путь и название для сохранения страницы исходя из ссылки
func (w *Wget) GetSavePathAndName(url string) (path, name string) {
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

// IsURLInternal проверяет, относится ли ссылка к данному домену
func (w *Wget) IsURLInternal(url string) (is bool) {
	if strings.Index(url, w.domain) == 0 {
		is = true
	}
	return
}

// Save сохраняет страницу
func (w *Wget) Save(url string, page []byte) (bool, error) {
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

// ContainsExtension проверяет есть ли в ссылке расширение файла
func (w *Wget) ContainsExtension(name string) int {
	return strings.Index(name, ".")
}

// IsJunkPage проверяет ссылку на аргументы, если есть возвращает true
func (w *Wget) IsJunkPage(url string) bool {
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

// ParseLinks собирает все ссылки со страницы
func (w *Wget) ParseLinks(page string) []string {
	r := regexp.MustCompile(`href="(.*?)"`)
	return r.FindAllString(page, -1)
}

// GetPage скачивает страницу по ссылке
func (w *Wget) GetPage(url string) ([]byte, error) {
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
