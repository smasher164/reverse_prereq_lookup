package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var pre = regexp.MustCompile(`Prerequisite:(.*)`)
var course = regexp.MustCompile(`[[:upper:]]{2,4} \d{4}`)

func populate(code, v string) {
	if !strings.Contains(v, "Prerequisite") {
		return
	}
	arr := pre.FindAllStringSubmatch(v, -1)
	if len(arr) == 0 {
		return
	}
	prereqStr := arr[0][1]
	res := course.FindAllString(prereqStr, -1)
	for _, cstr := range res {
		// courseStr := strings.Fields(r)
		// ccode := a[0]
		// cnum := a[1]
		if m[cstr] == nil {
			m[cstr] = make(map[string]bool)
		}
		m[cstr][code] = true
	}
}

// course code [course number] []prereq
var m map[string]map[string]bool

// var home = template.New(`
// 	<h1>Reverse Prerequisite</h1>
// `)

func main() {
	m = make(map[string]map[string]bool)
	file, err := os.Open("CS1188Data.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	cr := csv.NewReader(file)
	for {
		record, err := cr.Read()
		if err != nil || err == io.EOF {
			break
		}
		last := record[len(record)-1]
		populate(record[1]+" "+record[2], last)
	}
	// rw.Header().Set("Access-Control-Allow-Origin", "*")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write("This is whe")
	})
	http.HandleFunc("/prereq", func(rw http.ResponseWriter, r *http.Request) {
		course := r.FormValue("course")
		rw.Write([]byte(m[course]))
	})
	http.ListenAndServe(":8080", http.HandlerFunc())
}
