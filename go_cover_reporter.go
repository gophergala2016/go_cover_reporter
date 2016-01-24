package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

type message struct {
	Name string
	Body string
}

const (
	filename = "persist.txt"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler function
	http.HandleFunc("/receiver", receiver)
	http.HandleFunc("/demo_badge", func(w http.ResponseWriter, r *http.Request) {
		coverBadge(w, 0)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	percentString := string(buffer)
	percent, _ := strconv.ParseFloat(strings.TrimSpace(percentString), 64)
	coverage := struct{ Percent float64 }{percent}
	err = pageTemplate.Execute(w, coverage)
	if err != nil {
		log.Print(err)
	}

}

var pageTemplate = template.Must(template.New("pageTemplate").Parse(`
<!DOCTYPE html>
<html>
 <head>
	  <script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
	  <script src="http://d3js.org/d3.v3.min.js"></script>
	  <style>
	    body {
	      font-family: Helvetica, Arial, sans-serif;
	      margin: auto;
	      width: 960px;
	      padding-top: 20px;
	      background-color: #012647;
	    }

	    text {
	      font-family: Helvetica, Arial, sans-serif;
	      font-size: 7rem;
	      font-weight: 400;
	      line-height: 16rem;
	      fill: #1072b8;
	    }

	    h1 {
	    	color: #1072b8;
				text-align: center;
	    }

	    #donut {
	      width: 29rem;
	      height: 29rem;
	      margin: 0 auto;
	    }

	    path.color0 {
	      fill: #1072b8;
	    }

	    path.color1 {
	      fill: #35526b;
	    }
	  </style>
</head>

<body>
	    <h1>Most recent reported code coverage is {{.Percent}}% </h1>

	    <div id="donut" data-donut="{{.Percent}}"></div>

  <script>
    var duration   = 500,
    transition = 200;

    drawDonutChart(
      '#donut',
      $('#donut').data('donut'),
      490,
      490,
      ".35em"
    );

    function drawDonutChart(element, percent, width, height, text_y) {
      width = typeof width !== 'undefined' ? width : 490;
      height = typeof height !== 'undefined' ? height : 490;
      text_y = typeof text_y !== 'undefined' ? text_y : "-.10em";

      var dataset = {
            lower: calcPercent(0),
            upper: calcPercent(percent)
          },
          radius = Math.min(width, height) / 2,
          pie = d3.layout.pie().sort(null),
          format = d3.format("^.2%");

      var arc = d3.svg.arc()
            .innerRadius(radius - 20)
            .outerRadius(radius);

      var svg = d3.select(element).append("svg")
            .attr("width", width)
            .attr("height", height)
            .append("g")
            .attr("transform", "translate(" + width / 2 + "," + height / 2 + ")");

      var path = svg.selectAll("path")
        .data(pie(dataset.lower))
        .enter().append("path")
        .attr("class", function(d, i) { return "color" + i })
        .attr("d", arc)
        .each(function(d) { this._current = d; });

      var text = svg.append("text")
            .attr("text-anchor", "middle")
            .attr("dy", text_y);

      if (typeof(percent) === "string") {
        text.text(percent);
      }  else {
          var progress = 0;
          var timeout = setTimeout(function () {
            clearTimeout(timeout);
            path = path.data(pie(dataset.upper));
            path.transition().duration(duration).attrTween("d", function (a) {
              var i  = d3.interpolate(this._current, a);
              var i2 = d3.interpolate(progress, percent)
              this._current = i(0);
              return function(t) {
                text.text( format(i2(t) / 100) );
                return arc(i(t));
              };
            });
          }, 200);
     }
   };

    function calcPercent(percent) {
      return [percent, 100-percent];
     };
   </script>
  </body>
</html>
`))

func receiver(rw http.ResponseWriter, req *http.Request) {

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var t message
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`\d+.*\d*%`)

	numericalValue := re.FindString(string(t.Body))

	_, err = io.WriteString(file, numericalValue[:len(numericalValue)])
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()

	log.Println(t.Body)
}

func dummyFunction(i int, j int) int {
	return i + j
}
