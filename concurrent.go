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
  url := baseURL + "courses"
  resp, _ := http.Get(url)
  body, _ := ioutil.ReadAll(resp.Body)

  //call the PrintCourseData method for each course
}

func PrintCourseData(body []byte) {
  course := course{}
  json.Unmarshal(body, &course)

  fmt.Println("\nCS "+course.CourseNumber+": "+course.Title)
  fmt.Println("Units: "+ course.Units)
  fmt.Println(course.Description+"\n")
}

func MakeConccurentRequests() {
  urls := [3]string{"http://localhost:8080/", "http://localhost:8080/courses", "http://localhost:8080/course/cs4080"}
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

  //add menu options here 

  //if single class, make single request
  //if all courses, get all courses
  //if concurrent, make concurrent requests
  MakeSingleRequest()
  MakeConccurentRequests()
}
