package main

// import -----------------------
import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/drone/routes"
	"github.com/kr/pretty"
	"net/http"
)

// Variable Set -----------------
var m map[string]string = make(map[string]string)

// Output Header ----------------
func outputHeader(w http.ResponseWriter) {
	//w.Header().Set("Content-Type", "text/plain")       // DEBUG
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Analyze uri ------------------
func retParam(r *http.Request) string {
	params := r.URL.Query()
	param := params.Get(":param")

	return param
}

// json check ------------------
func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// View Data -------------------
func viewHandler(w http.ResponseWriter, r *http.Request) {

	param := retParam(r)
	outputHeader(w)

	// DEBUG ----------------------
	if param == "debug" {
		debugPrint()
		return
	}

	// Key Check ------------------
	val, ok := m[param]
	if ok == false {
		w.Write([]byte(val))			// Dummy
		w.Write([]byte("{\"Message\":\"Key => "))
		w.Write([]byte(param))
		w.Write([]byte(" not Found!!\"}"))
		w.Write([]byte("\n"))

		return
	}

	w.Write([]byte(m[param]))
	w.Write([]byte("\n"))
}

// Debug Print -----------------
func debugPrint() {
	pretty.Printf("--- m:\n%# v\n\n", m)
}

// Post Data -------------------
func postHandler(w http.ResponseWriter, r *http.Request) {

	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	param := retParam(r)

	// Json Check
	if isJSON(body) == false {
		w.Write([]byte("{\"Message\":\"Invalid Json Data!! Post Aborted...\"}\n"))
		return
	}

	m[param] = body

	outputHeader(w)

	w.Write([]byte("{\"Message\":\"Data Insert => Key:"))
	w.Write([]byte(param))
	w.Write([]byte("\"}\n"))
}

// Delete Data -----------------
func delHandler(w http.ResponseWriter, r *http.Request) {

	param := retParam(r)

	delete(m, param)

	outputHeader(w)
	w.Write([]byte("{\"Message\":\"Delete Key => "))
	w.Write([]byte(param))
	w.Write([]byte("\"}\n"))
}

func main() {

	// Parameter Setting ------------------
	var port_int int
	flag.IntVar(&port_int, "port", 8080, "HTTP-KVS Server Port")
	flag.Parse()
	port_str := ":" + fmt.Sprint(port_int)

	// Start Message ----------------------
	fmt.Print("Starting HTTP-KVS...\n")
	fmt.Print("Server Port => ")
	fmt.Print(port_int)
	fmt.Print("\n")

	// Set Handler ------------------------
	mux := routes.New()

	mux.Get("/:param", viewHandler)  // 取得
	mux.Post("/:param", postHandler) // 更新
	mux.Del("/:param", delHandler)   // 削除

	http.Handle("/", mux)

	// Start HTTP Server ------------------
	http.ListenAndServe(port_str, nil)
}
