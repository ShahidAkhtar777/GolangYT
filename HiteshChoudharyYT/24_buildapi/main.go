package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for cource - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses []Course

// middleware / helpers

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// Controllers - seperate file

// serve Home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API by LCO. </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all cources")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab the request id
	params := mux.Vars(r)
	fmt.Println("Params are", params)
	fmt.Printf("Type of Params: %T", params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this Id.")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// Body is empty ?
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
		return
	}

	// what about {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please add coursename!")
		return
	}

	// Check same title course present or not
	for _, v := range courses {
		if v.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Same title already exixts!")
			return
		}
	}

	randomNum := rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(randomNum.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab request Id
	params := mux.Vars(r)
	fmt.Println("Params are", params)
	fmt.Printf("Type of Params: %T", params)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Given courseId not found!")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab Id
	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode("Provided course Id removed")
			return
		}
	}

	json.NewEncoder(w).Encode("Given courseId not found")
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all course")
	w.Header().Set("Content-Type", "application/json")

	courses = nil
	json.NewEncoder(w).Encode("All courses removed!")
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	// Seeding of fake db
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs",
		CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.in"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack",
		CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")

	// Listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}
