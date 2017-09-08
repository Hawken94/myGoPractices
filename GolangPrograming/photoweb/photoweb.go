package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		/* io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
		" enctype=\"mutipart /form-data\">"+
		"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
		"<input type=\"submit\" value=\"Upload\" />"+
		"</form>") */

		/* if err := renderHTML(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} */
		renderHTML(w, "upload", nil)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageID
	if ok := isExists(imagePath); !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images

	renderHTML(w, "list", locals)
}

// 渲染模板
func renderHTML(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	tmpl += ".html"
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				log.Println("WARN: panic in %v.panic - %v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			/* if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			} */
			fi, err := os.Stat(file)
			if err != nil || fi.IsDir() {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
