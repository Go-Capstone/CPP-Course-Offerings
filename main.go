package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type course struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Department string `json:"Department"`
	CourseNumber string `json:"CourseNumber"`
	Units string `json:"Units"`
	Description string `json:"Description"`
}

type allCourses []course

var courses = allCourses {
	{
		ID: "cs1300",
		Title: "Discrete Structures",
		Department: "Computer Science",
		CourseNumber: "1300",
		Units: "4",
		Description: "Logic, proof techniques, sets, basic counting rules, relations, functions and recursion. Modular arithmetic, prime numbers. Graphs and trees. Boolean algebra, computer logic, and combinational circuits.",
	},
	{
		ID: "cs1400",
		Title: "Introduction to Programming and Problem Solving",
		Department: "Computer Science",
		CourseNumber: "1400",
		Units: "4",
		Description: "Problem-solving methods. Object-oriented concepts including classes, inheritance, and polymorphism. Basic control structures and data types. File I/O and exception handling. Program documentation and testing.",
	},
	{
		ID: "cs2400",
		Title: "Data Structures and Advanced Programming",
		Department: "Computer Science",
		CourseNumber: "2400",
		Units: "4",
		Description: "Abstract data types and their implementation using linear and non-linear data structures. Interfaces and generics. Advanced file access. Recursive structures and operations. Big-O notation and introduction to algorithm analysis.",
	},
	{
		ID: "cs2450",
		Title: "Programming Graphical User Interfaces",
		Department: "Computer Science",
		CourseNumber: "2450",
		Units: "3",
		Description: "Computer interfaces. Usability of interactive systems. GUI development processes. GUI components. Input and viewing devices. Event-handling. Animation use in GUIs. Problem-solving techniques. Introduction to Human-Computer Interface.",
	},
	{
		ID: "cs2520",
		Title: "Python for Programmers ",
		Department: "Computer Science",
		CourseNumber: "2520",
		Units: "3",
		Description: "Python basics. Data structures and control structures. Functions and functional forms. Objects, classes, inheritance, and polymorphism. I/O and exceptions. Advanced features.",
	},
	{
		ID: "cs2560",
		Title: "C++ Programming ",
		Department: "Computer Science",
		CourseNumber: "2560",
		Units: "3",
		Description: "C++ fundamentals. Recursion, structures, class encapsulation, inheritance, polymorphism, and exception handling. Memory management using pointers. Template functions and template classes. Software reuse and object-oriented programming.",
	},
	{
		ID: "cs2600",
		Title: "Unix and Scripting ",
		Department: "Computer Science",
		CourseNumber: "2600",
		Units: "3",
		Description: "Fundamentals of the UNIX operating system and scripting. Tools for file management, communication, process control, and program development. Introduction to various UNIX shells and programming (scripting) in the shells.",
	},
	{
		ID: "cs2640",
		Title: "Programming Graphical User Interfaces",
		Department: "Computer Science",
		CourseNumber: "2640",
		Units: "3",
		Description: "Number representation, integer arithmetic. Von Neumann machine. Instruction set architecture. Addressing modes. Assembly programming. Arrays and records. Subroutines and macros. Interrupts. I/O interfacing and communication.",
	},
	{
		ID: "cs2990",
		Title: "Special Topics for Lower Division Students ",
		Department: "Computer Science",
		CourseNumber: "2990",
		Units: "3",
		Description: "Group study of a selected well defined topic or area not covered by a regularly offered course. Total credit limited to 6 units applicable to a degree, with a maximum of 2 sections per semester.",
	},
	{
		ID: "cs3010",
		Title: "Numerical Methods",
		Department: "Computer Science",
		CourseNumber: "3010",
		Units: "3",
		Description: "Error analysis, floating-point representations, roots of nonlinear equations, systems of linear equations, interpolation, Chebyshev approximation, least squares approximation, numerical integration and differentiation.",
	},
	{
		ID: "cs3110",
		Title: "Formal Languages and Automata",
		Department: "Computer Science",
		CourseNumber: "3110",
		Units: "3",
		Description: "Finite automata. Non-determinism. Regular expressions and languages. Context-free grammars and push down automata. Ambiguity. Closure properties. Normal forms. Context-free and non-context-free languages.",
	},
	{
		ID: "cs3310",
		Title: "Design and Analysis of Algorithms",
		Department: "Computer Science",
		CourseNumber: "3310",
		Units: "3",
		Description: "Algorithms for fundamental problems. Efficiency analysis using asymptotic notation. Principal algorithm design techniques and their tradeoffs. NP-theory and approaches used to address intractability.",
	},
	{
		ID: "cs3520",
		Title: "Symbolic Programming",
		Department: "Computer Science",
		CourseNumber: "3110",
		Units: "3",
		Description: "Finite automata. Non-determinism. Regular expressions and languages. Context-free grammars and push down automata. Ambiguity. Closure properties. Normal forms. Context-free and non-context-free languages.",
	},
	{
		ID: "cs3560",
		Title: "Object-Oriented Design and Programming ",
		Department: "Computer Science",
		CourseNumber: "3560",
		Units: "3",
		Description: "Elements of the object model. Abstraction, encapsulation, modularity, hierarchy and polymorphism. Object-oriented design principles. Design patterns. Implementation and programming of system design. Object and portable data. Comprehensive examples using a case study approach.",
	},
	{
		ID: "cs3650",
		Title: "Symbolic Programming",
		Department: "Computer Science",
		CourseNumber: "3650",
		Units: "3",
		Description: "Flip flops, sequential machines. Computer arithmetic. Data path and control unit design. Pipelining. Memory hierarchy and storage technology. Multiprocessing and alternative architectures.",
	},
	{
		ID: "cs4080",
		Title: "Concepts of Programming Languages",
		Department: "Computer Science",
		CourseNumber: "4080",
		Units: "3",
		Description: "Concepts in programming languages. Virtual machines and abstraction. Syntax and semantics. Declarations and types. Scoping and binding. Data abstraction. Control and abstraction. Subprograms and implementations. Exception handling. Programming paradigms.",
	},
	{
		ID: "cs4310",
		Title: "Operating Systems",
		Department: "Computer Science",
		CourseNumber: "4310",
		Units: "3",
		Description: "Overview of operating systems. Operating system structures. Process management. Concurrency and synchronization. Deadlock. Processor management. Scheduling and dispatch. Memory management. Virtual memory. Device management. File systems. Security, privacy and acceptable use.",
	},
	{
		ID: "cs4800",
		Title: "Software Engineering",
		Department: "Computer Science",
		CourseNumber: "4800",
		Units: "3",
		Description: "Models of the software development process and metrics. Software requirements and specifications. Methodologies, tools and environments. Human-computer interaction. Software architecture, design and implementation techniques. Project management. Cost estimation. Testing and validation. Automated build, deployment and continuous integration. Maintenance and evolution.",
	},


}
func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to your Go home page!")
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	courseID := mux.Vars(r)["id"]

	for _, singleCourse := range courses{
		if singleCourse.ID == courseID {
			json.NewEncoder(w).Encode(singleCourse)
		}
	}
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
