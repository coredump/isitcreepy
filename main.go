package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/calc/", calc)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Starting server")
	log.Panic(http.ListenAndServe(":9080", nil))

}

func logreq(r *http.Request) {
	log.Printf("%v - %v - %v", r.RemoteAddr, r.Method, r.URL)
}

func ages(age float64) (float64, float64) {
	return math.Ceil(age/2 + 7), math.Ceil(age*2 - 14)
}

func index(w http.ResponseWriter, r *http.Request) {
	logreq(r)

	ages := make([]int, 66)
	for i := 0; i < 66; i++ {
		ages[i] = i + 14
	}
	tpldata := struct {
		Ages []int
	}{ages}

	if err := indexTpl.Execute(w, tpldata); err != nil {
		log.Printf("%v - %v - %v - Index template execution failed", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "Index template execution failed", http.StatusInternalServerError)
		return
	}

}

func calc(w http.ResponseWriter, r *http.Request) {
	logreq(r)

	value, err := strconv.ParseFloat(strings.TrimLeft(r.URL.Path, "/calc/"), 64)
	if err != nil {
		log.Printf("Wrong data on URL: %v - %v", r.URL, err)
		http.Error(w, "Wrong data on the URL", http.StatusInternalServerError)
		return
	}
	min, max := ages(value)
	result := struct {
		Min float64
		Max float64
	}{min, max}

	j, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshaling the JSON data")
		http.Error(w, "Internal marshaling error", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(j))
}

var indexTpl = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
  <script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
  <script type="text/javascript" src="/assets/isitcreepy.js"></script>
  <script type="text/javascript" src="/assets/jquery-ui-1.8.22.custom.min.js"></script>
  <style type="text/css">
    body { text-align: center; background-color: white; font-family: 'Helvetica, Arial, sans'}
    div#results { text-align: left; display: inline-block; align: "center" }
    div#content { width: 740px; margin-left: auto; margin-right: auto; }
    div#selector { text-align: left; }
  </style>
<title>Is it creepy?</title>
<body>
<div id="content">
  <h1>Is it creepy? (To date that person)</h1>
  <img src="http://imgs.xkcd.com/comics/dating_pools.png">
  <div id="selector">
    <p>Select your age: <select id="age_selector">
    {{ range .Ages }}
    <option>{{ . }}</option>
    {{ end }}
    </select>
    </p>
  </div>
  <div id="results">
  </div>
<a href="https://github.com/coredump/"><img style="position: absolute; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_orange_ff7600.png" alt="Fork me on GitHub"></a>
</div>
</body>
</html>
`))
