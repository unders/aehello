package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var initTime = time.Now()

func main() {
	err := errors.New("err just a test")
	fmt.Println(err)

	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Serving the front page.")

	tmpl.Execute(w, time.Since(initTime))
}

var tmpl = template.Must(template.New("front").Parse(`
<html><body>

<p>
Hello, World! 세상아 안녕!
</p>

<p>
This instance has been running for <em>{{.}}</em>.
</p>

</body></html>
`))
