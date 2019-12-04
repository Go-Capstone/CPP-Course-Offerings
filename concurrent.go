package main

import (
     "fmt"
     "os"
     "strings"
     "bufio"
     "net/http"
     "encoding/json"
     "time"
     "io/ioutil"
)

const baseURL string = "http://localhost:8080/"

type course struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Department string `json:"Department"`
	CourseNumber string `json:"CourseNumber"`
	Units string `json:"Units"`
	Description string `json:"Description"`
}

func MakeRequest(url string, ch chan<-string) {
  start := time.Now()
  resp, _ := http.Get(url)
  secs := time.Since(start).Seconds()
  body, _ := ioutil.ReadAll(resp.Body)
  ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func MakeSingleRequest(){
  scanner := bufio.NewReader(os.Stdin)
  var courseName string
  fmt.Print("Enter Class : ")
  courseName, _ = scanner.ReadString('\n')
  courseName = strings.TrimSuffix(courseName, "\n")
  url := baseURL + "courses/" + courseName
  resp, _ := http.Get(url)
  body, _ := ioutil.ReadAll(resp.Body)

  PrintCourseData(body)
}

func GetAllCourses(){
 
  classes := [18]string{"cs1300", "cs1400", "cs2400", "cs2450", "cs2520", "cs2560", "cs2600", "cs2640", "cs2990", "cs3010", "cs3110", "cs3310", "cs3520", "cs3560", "cs3650", "cs4080", "cs4310", "cs4800"}
  
  for i := 0; i < 18; i++{
	url := baseURL + "courses/" + classes[i]
    resp, _ := http.Get(url)
    body, _ := ioutil.ReadAll(resp.Body) 

  PrintCourseData(body)
  }
}

func PrintCourseData(body []byte) {
  course := course{}
  json.Unmarshal(body, &course)

  fmt.Println("\nCS "+course.CourseNumber+": "+course.Title)
  fmt.Println("Units: "+ course.Units)
  fmt.Println(course.Description+"\n")
}

func MakeConccurentRequests() {
  urls := [3]string{"http://localhost:8080/courses/cs1300", "http://localhost:8080/courses/cs3560", "http://localhost:8080/course/cs4080"}
  start := time.Now()
  ch := make(chan string)
  for _,url := range urls{
      go MakeRequest(url, ch)
  }
  for range urls{
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
    
	fmt.Println()
    var choice int
    for ok := true; ok; ok = (choice != 4) {
     fmt.Println("1. Get single course\n2. Get all courses\n3. Make concurrent request\n4. Quit")
     fmt.Scan(&choice) 
	 
	 if(choice == 1){
	    fmt.Println()
	    MakeSingleRequest()
	    fmt.Println()
	 }
	 if(choice == 2){
	    fmt.Println()
	    GetAllCourses()
	    fmt.Println()
	 }
	 if(choice == 3){
	    fmt.Println()
		MakeConccurentRequests()
		fmt.Println()
	 }
   }
}