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
	http.HandleFunc("/stats/", stats)
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

	ages := make([]int, 67)
	for i := 0; i <= 66; i++ {
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
		log.Printf("Error marshaling the JSON data: %v", err)
		http.Error(w, "Internal marshaling error", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(j))
}

func stats(w http.ResponseWriter, r *http.Request) {
	logreq(r)

	maxcoords := make([][]float64, 100)
	mincoords := make([][]float64, 100)
	for i := 0; i < 86; i++ {
		z := i + 14
		min, max := ages(float64(z))
		maxcoords[i] = []float64{float64(z), max}
		mincoords[i] = []float64{float64(z), min}
	}

	j, err := json.Marshal(struct {
		Min [][]float64
		Max [][]float64
	}{mincoords, maxcoords})
	if err != nil {
		log.Printf("Error marshalling the JSON data: %v", err)
		http.Error(w, "Internal marshalling error", http.StatusInternalServerError)
	}
	fmt.Fprint(w, string(j))
}

var indexTpl = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
  <script language="javascript" type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
  <script language="javascript" type="text/javascript" src="/assets/jquery-ui-1.8.22.custom.min.js"></script>
  <script language="javascript" type="text/javascript" src="/assets/isitcreepy.js"></script>
  <script language="javascript" type="text/javascript" src="/assets/flot/jquery.flot.js"></script>
  <script language="javascript" type="text/javascript" src="/assets/jquery.flot.axislabels.js"></script>
  <style type="text/css">
    body { text-align: center; background-color: white; font-family: Helvetica, Arial, sans-serif; height:100%; margin: 0; padding: 0; }
    div#results { text-align: left; display: inline-block; align: "center" }
    div#content { width: 740px; margin-left: auto; margin-right: auto; }
    div#selector { text-align: left; font-size: 20px}
    div#graphexplain { text-align: left; }
    div#placeholder { width: 740px; height: 400px; text-align: left }
    div#footer {width: 740px; position:absolute; bottom:0; width:100%; height:60px; font-size: 14px }
  </style>
<title>Is it creepy?</title>
<body>
<div id="content">
  <h1>Is it creepy? (To date that person)</h1>
  <a href="http://xkcd.com/314/"><img src="http://imgs.xkcd.com/comics/dating_pools.png"></a>
  <div id="selector">
    <p>Select your age: <select id="age_selector">
    {{ range .Ages }}
    <option>{{ . }}</option>
    {{ end }}
    </select>
    </p>
  </div>
  <div id="results">
  <p><em>Notice that it stops at 80. Not because I think that anyone must stop dating at 80 or any age, it's just for the sake of better data visualization.</em></p>
  </div>
  <div id="placeholder">
  </div>
  <div id="graphexplain">
  </div>
<a href="https://github.com/coredump/"><img style="position: absolute; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_orange_ff7600.png" alt="Fork me on GitHub"></a>
</div>
<div id="footer">There's a blog post about this page <a href="http://coredump.io/blog/2012/08/13/learning-go-lang/">here</a></div>
</body>
</html>
`))
