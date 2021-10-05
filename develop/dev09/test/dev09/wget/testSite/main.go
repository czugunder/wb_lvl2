package testSite

import (
	"fmt"
	"net/http"
)

func sendPage(w http.ResponseWriter, req *http.Request) {
	page := `<html>
			<body>
			<a href="http://localhost:8888/one"></a>
			<a href="http://localhost:8888/one.php"></a>
			<a href="http://localhost:8888/one.php?=someargs"></a>
			<a href="http://localhost:8888/one/two"></a>
			<a href="https://yandex.ru"></a>
			</body>
			</html>`
	fmt.Fprintf(w, page)
}

func Server() {
	http.HandleFunc("/", sendPage)
	http.HandleFunc("/one", sendPage)
	http.HandleFunc("/one.php", sendPage)
	http.HandleFunc("/one.php?=someargs", sendPage)
	http.HandleFunc("/one/two", sendPage)
	http.ListenAndServe(":8888", nil)
}
