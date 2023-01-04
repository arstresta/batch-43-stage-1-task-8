package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title": "Personal Web",
	"Text":  "Mau Makan",
}

// Array of object
// nama = []string{"Dira", "Dwi"}

// type persegi interface {
// 	panjang() float64
// 	lebar() float64
// }

type Blog struct {
	Title     string
	Post_date string
	Author    string
	Content   string
}

var Blogs = []Blog{
	{
		Title:     "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		Post_date: "12 Jul 2021 | 22:30 WIB",
		Author:    "Abel Dustin",
		Content:   "Test",
	},
}

type Project struct {
	ProjectName string
	StartDate   string
	EndDate     string
	Durasi      string
	Description string
	Node        string
	React       string
	TypeScript  string
	Next        string
}

var Projects = []Project{
	{
		ProjectName: "Dumbways Mobile App",
		StartDate:   "10 Jan 2021",
		EndDate:     "20 Jan 2021",
		Durasi:      "12 Jul 2021 | 22:30 WIB",
		Node:        `<i class="fa-brands fa-node-js"></i>`,
		React:       `<i class="fa-brands fa-react"></i>`,
		TypeScript:  `<i class="fa-brands fa-playstation"></i>`,
		Next:        `<i class="fa-brands fa-google-play"></i>`,
		Description: "Test",
	},
}

func main() {
	route := mux.NewRouter()
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", helloWorld).Methods("GET")
	route.HandleFunc("/home", home).Methods("GET").Name("home")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/project/{id}", projectDetail).Methods("GET")
	route.HandleFunc("/project", addProject).Methods("POST")
	route.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/coba/{id}", coba).Methods("GET")

	fmt.Println("Server Is Running On Port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message: " + err.Error()))
		return
	}

	respData := map[string]interface{}{
		"Data":     Data,
		"Projects": Projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var tmpl, err = template.ParseFiles("views/detail-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message" + err.Error()))
		return
	}

	resp := map[string]interface{}{
		"Data": Data,
		"Id":   id,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Title : " + r.PostForm.Get("name"))
	projectName := r.PostForm.Get("name")
	startDate := time.Now().String()
	endDate := time.Now().String()
	durasi := "10 Bulan"
	description := r.PostForm.Get("description")
	Node := r.PostForm.Get("node")
	React := r.PostForm.Get("react")
	Next := r.PostForm.Get("next")
	TypeScript := r.PostForm.Get("typescript")

	var newProject = Project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Durasi:      durasi,
		Description: description,
		Node:        Node,
		React:       React,
		TypeScript:  TypeScript,
		Next:        Next,
	}
	Projects = append(Projects, newProject)
	fmt.Println(Projects)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content/type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println(id)
	Projects = append(Projects[:id], Projects[id+1:]...)
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func coba(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content/type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	Projects[id].ProjectName = "oke"

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact-form.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}
