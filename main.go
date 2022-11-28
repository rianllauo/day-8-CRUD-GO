package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public", http.FileServer((http.Dir("./public")))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/form-project", formProject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project-detail/{id}", projectDetail).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")
	route.HandleFunc("/delete-project/{index}", deleteProject).Methods("GET")
	route.HandleFunc("/edit-project/{index}", formEditProject).Methods("GET")
	route.HandleFunc("/edit-project/{index}", editProject).Methods("POST")

	fmt.Println(("server berjalan di port 5000"))
	http.ListenAndServe("localhost:5000", route)
}

type Project struct {
	Id                int
	Title             string
	DateStart         string
	DateEnd           string
	Content           string
	NodeJs            string
	NextJs            string
	ReactJs           string
	Javascript        string
	NodeJsChecked     string
	NextJsChecked     string
	ReactJsChecked    string
	JavascriptChecked string
}

var projects = []Project{
	{
		Title:     "Aplikasi web dumbways",
		DateStart: "11 november 2022",
		DateEnd:   "12 desember 2022",
		Content:   "lorem ipsum dolor si amet",
		// NodeJs:        "public/img/nodejs.svg",
		NextJs:        "public/img/nextjs.svg",
		ReactJs:       "public/img/react.svg",
		Javascript:    "public/img/javascript.svg",
		NodeJsChecked: "none",
	},
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	dateStart := r.PostForm.Get("date-start")
	dateEnd := r.PostForm.Get("date-end")

	nodeJs := r.PostForm.Get("nodeJs")
	nextJs := r.PostForm.Get("nextJs")
	reactJs := r.PostForm.Get("reactJs")
	javascript := r.PostForm.Get("javascript")
	// nodeJs := r.Form["nodeJs"][0] == "true"

	var nodeJsPath = ""
	var nodeJsChecked = ""
	var nextJsPath = ""
	var nextJsChecked = ""
	var reactJsPath = ""
	var reactJsChecked = ""
	var javascriptPath = ""
	var javascriptChecked = ""

	if nodeJs == "true" {
		nodeJsPath = "public/img/nodejs.svg"
		nodeJsChecked = "checked"
	} else {
		nodeJsPath = "d-none"
	}

	if nextJs == "true" {
		nextJsPath = "public/img/nextjs.svg"
		nextJsChecked = "checked"
	} else {
		nextJsPath = "d-none"
	}

	if reactJs == "true" {
		reactJsPath = "public/img/react.svg"
		reactJsChecked = "checked"
	} else {
		reactJsPath = "d-none"
	}

	if javascript == "true" {
		javascriptPath = "public/img/javascript.svg"
		javascriptChecked = "checked"
	} else {
		javascriptPath = "d-none"
	}

	var newProject = Project{
		Title:             title,
		Content:           content,
		DateStart:         dateStart,
		DateEnd:           dateEnd,
		NodeJs:            nodeJsPath,
		NodeJsChecked:     nodeJsChecked,
		NextJs:            nextJsPath,
		NextJsChecked:     nextJsChecked,
		ReactJs:           reactJsPath,
		ReactJsChecked:    reactJsChecked,
		Javascript:        javascriptPath,
		JavascriptChecked: javascriptChecked,
	}
	projects = append(projects, newProject)

	fmt.Println(nodeJs)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func formEditProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/editProject.html")

	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	var ProjectEdit = Project{}

	for i, data := range projects {
		if i == index {
			ProjectEdit = Project{
				Id:                i,
				Title:             data.Title,
				Content:           data.Content,
				DateStart:         data.DateStart,
				DateEnd:           data.DateEnd,
				NodeJs:            data.NodeJs,
				NodeJsChecked:     data.NodeJsChecked,
				NextJsChecked:     data.NextJsChecked,
				ReactJsChecked:    data.ReactJsChecked,
				JavascriptChecked: data.JavascriptChecked,
			}
		}
	}

	// fmt.Println(projects)

	dataEdit := map[string]interface{}{
		"Project": ProjectEdit,
	}

	tmpt.Execute(w, dataEdit)
}

func editProject(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	dateStart := r.PostForm.Get("date-start")
	dateEnd := r.PostForm.Get("date-end")

	nodeJs := r.PostForm.Get("nodeJs")
	nextJs := r.PostForm.Get("nextJs")
	reactJs := r.PostForm.Get("reactJs")
	javascript := r.PostForm.Get("javascript")

	var nodeJsPath = ""
	var nodeJsChecked = ""
	var nextJsPath = ""
	var nextJsChecked = ""
	var reactJsPath = ""
	var reactJsChecked = ""
	var javascriptPath = ""
	var javascriptChecked = ""

	if nodeJs == "true" {
		nodeJsPath = "public/img/nodejs.svg"
		nodeJsChecked = "checked"
	} else {
		nodeJsPath = "d-none"
	}

	if nextJs == "true" {
		nextJsPath = "public/img/nextjs.svg"
		nextJsChecked = "checked"
	} else {
		nextJsPath = "d-none"
	}

	if reactJs == "true" {
		reactJsPath = "public/img/react.svg"
		reactJsChecked = "checked"
	} else {
		reactJsPath = "d-none"
	}

	if javascript == "true" {
		javascriptPath = "public/img/javascript.svg"
		javascriptChecked = "checked"
	} else {
		javascriptPath = "d-none"
	}

	var newProject = Project{
		Title:             title,
		Content:           content,
		DateStart:         dateStart,
		DateEnd:           dateEnd,
		NodeJs:            nodeJsPath,
		NodeJsChecked:     nodeJsChecked,
		NextJs:            nextJsPath,
		NextJsChecked:     nextJsChecked,
		ReactJs:           reactJsPath,
		ReactJsChecked:    reactJsChecked,
		Javascript:        javascriptPath,
		JavascriptChecked: javascriptChecked,
	}

	// projects = append(projects, newProject)
	projects[index] = newProject

	fmt.Println(index)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	dataProject := map[string]interface{}{
		"Projects": projects,
	}

	tmpt.Execute(w, dataProject)
}

func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/addProject.html")

	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/projectDetail.html")

	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var ProjectDetail = Project{}

	for index, data := range projects {
		if index == id {
			ProjectDetail = Project{
				Title:     data.Title,
				Content:   data.Content,
				DateStart: data.DateStart,
				DateEnd:   data.DateEnd,
				NodeJs:    data.NodeJs,
			}
		}
	}

	dataDetail := map[string]interface{}{
		"Project": ProjectDetail,
	}

	tmpt.Execute(w, dataDetail)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	projects = append(projects[:index], projects[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}
