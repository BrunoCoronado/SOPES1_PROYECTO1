// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (

	//websocket
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/gorilla/websocket"

	//ram monitor
	"fmt"
	//"runtime"
	"github.com/shirou/gopsutil/mem"
	
	//cpu monitor
	"strings"
	
	//graficar
	chart "github.com/wcharczuk/go-chart"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 1 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 7 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 1 * time.Second
)

var (
	addr      = flag.String("addr", ":3000", "http service address")
	homeTempl1 = template.Must(template.New("").Parse(ramHTML))
	homeTempl2 = template.Must(template.New("").Parse(cpuHTML))
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func dealwithErr(err error) {
	if err != nil {
			fmt.Println(err)
			//os.Exit(-1)
	}
}

func getInfoMemoria(lastMod time.Time) ([]byte, time.Time, error) {
    // memory
    vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	
	contenido := "Memoria Total: " + strconv.FormatUint(vmStat.Total, 10) + " bytes\n"
	contenido = contenido + "Memoria Libre: " + strconv.FormatUint(vmStat.Free, 10) + " bytes\n"
	contenido = contenido + "Porcentaje de memoria utilizado: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%\n"
	return []byte(contenido), lastMod, err
}

func getCPUSample() (idle, total uint64) {
    contents, err := ioutil.ReadFile("/proc/stat")
    if err != nil {
        return
    }
    lines := strings.Split(string(contents), "\n")
    for _, line := range(lines) {
        fields := strings.Fields(line)
        if fields[0] == "cpu" {
            numFields := len(fields)
            for i := 1; i < numFields; i++ {
                val, err := strconv.ParseUint(fields[i], 10, 64)
                if err != nil {
                    fmt.Println("Error: ", i, fields[i], err)
                }
                total += val // tally up all the numbers to get total ticks
                if i == 4 {  // idle is the 5th field in the cpu line
                    idle = val
                }
            }
            return
        }
    }
    return
}

func getInfoCPU(lastMod2 time.Time) ([]byte, time.Time) {
	idle0, total0 := getCPUSample()
    time.Sleep(3 * time.Second)
    idle1, total1 := getCPUSample()

    idleTicks := float64(idle1 - idle0)
    totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks
	

	contenido2 := "CPU Total: " + fmt.Sprintf("%f",totalTicks) + " bytes\n"
	contenido2 = contenido2 + "CPU Ocupado: " + fmt.Sprintf("%f",totalTicks-idleTicks) + " bytes\n"
	contenido2 = contenido2 + "Porcentaje del CPU utilizado: " + fmt.Sprintf("%f",cpuUsage) + "%\n"
	return []byte(contenido2), lastMod2
}

func graficarRam(){
	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),        //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(100)}.Values(), //generates a []float64 randomly from 0 to 100 with 100 elements.
	}

	polyRegSeries := &chart.PolynomialRegressionSeries{
		Degree:      3,
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Series: []chart.Series{
			mainSeries,
			polyRegSeries,
		},
	}

	f, _ := os.Create("outputram.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writer(ws *websocket.Conn, lastMod time.Time) {
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-fileTicker.C:
			var p []byte
			var err error

			p, lastMod, err = getInfoMemoria(lastMod)

			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if p != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	var lastMod time.Time
	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	go writer(ws, lastMod)
	reader(ws)
}

func serveWs2(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	var lastMod time.Time
	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	go writer(ws, lastMod)
	reader(ws)
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func index(w http.ResponseWriter, r *http.Request) {    
    file, err := os.Open("index.html")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    b, err := ioutil.ReadAll(file)
    fmt.Fprintf(w, BytesToString(b))
}

func ramm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rammonitor" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	graficarRam()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p, lastMod, err := getInfoMemoria(time.Time{})
	if err != nil {
		p = []byte(err.Error())
		lastMod = time.Unix(0, 0)
	}
	var v = struct {
		Host    string
		Data    string
		LastMod string
	}{
		r.Host,
		string(p),
		strconv.FormatInt(lastMod.UnixNano(), 16),
	}
	homeTempl1.Execute(w, &v)
}

func cpum(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cpumonitor" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//graficarRam()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p, lastMod := getInfoCPU(time.Time{})
	var v = struct {
		Host    string
		Data2    string
		LastMod string
	}{
		r.Host,
		string(p),
		strconv.FormatInt(lastMod.UnixNano(), 16),
	}
	homeTempl2.Execute(w, &v)
}


func main() {
	http.HandleFunc("/",index);
	http.HandleFunc("/rammonitor", ramm)
	http.HandleFunc("/cpumonitor", cpum)
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/rs", serveWs2)

	fmt.Printf("Corriendo correctamente el proyecto en el puerto 3000...\n")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}

const ramHTML = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset=”UTF-8”>
		<title>P1 Processes Monitor</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
		<script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
		<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    </head>
	<body>
		<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<a class="navbar-brand" href="#">P1 Processes Monitor</a>
			<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarSupportedContent">
				<ul class="navbar-nav mr-auto">
					<li class="nav-item active">
						<a class="nav-link" href="/">Home <span class="sr-only">(current)</span></a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="/cpumonitor">CPU Monitor</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="/rammonitor">RAM Monitor</a>
					</li>
				</ul>
			</div>
		</nav>
        <pre id="fileData">{{.Data}}</pre>
        <script type="text/javascript">
            (function() {
                var data = document.getElementById("fileData");
                var conn = new WebSocket("ws://{{.Host}}/ws?lastMod={{.LastMod}}");
                conn.onclose = function(evt) {
                    data.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log('file updated');
                    data.textContent = evt.data;
                }
            })();
		</script>
		<img src="outputram.png" alt="Grafica de la Ram no encontrada"> 
    </body>
</html>
`

const cpuHTML = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset=”UTF-8”>
		<title>P1 Processes Monitor</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
		<script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
		<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    </head>
	<body>
		<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<a class="navbar-brand" href="#">P1 Processes Monitor</a>
			<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarSupportedContent">
				<ul class="navbar-nav mr-auto">
					<li class="nav-item active">
						<a class="nav-link" href="/">Home <span class="sr-only">(current)</span></a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="/cpumonitor">CPU Monitor</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="/rammonitor">RAM Monitor</a>
					</li>
				</ul>
			</div>
		</nav>
        <pre id="fileData">{{.Data2}}</pre>
        <script type="text/javascript">
            (function() {
                var data = document.getElementById("fileData");
                var conn2 = new WebSocket("rs://{{.Host}}/rs?lastMod={{.LastMod}}");
                conn2.onclose = function(evt) {
                    data.textContent = 'Connection closed';
                }
                conn2.onmessage = function(evt) {
                    console.log('file updated 2');
                    data.textContent = evt.data;
                }
            })();
		</script>
		<img src="outputram.png" alt="Grafica del CPU no encontrada"> 
    </body>
</html>
`
