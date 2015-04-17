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
	"time"
)

// Variable Set -----------------
var m map[string]string = make(map[string]string)

// counter
var post_cnt int
var get_cnt int
var miss_cnt int
var del_cnt int

// time
var t time.Time

// Output Header ----------------
func outputHeader(w http.ResponseWriter) {
	//w.Header().Set("Content-Type", "text/plain")       // DEBUG
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE")
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
	if param == "_debug_" {
		debugPrint()
		return
	}

	// Key Check ------------------
	val, ok := m[param]
	if ok == false {
		w.Write([]byte(val)) // Dummy
		w.Write([]byte("{\"Message\":\"Key => "))
		w.Write([]byte(param))
		w.Write([]byte(" not Found!!\"}"))
		w.Write([]byte("\n"))

		miss_cnt++
		return
	}

	w.Write([]byte(m[param]))
	w.Write([]byte("\n"))

	get_cnt++
}

// Debug Print -----------------
func debugPrint() {

	fmt.Print("\nProess Start at => ")
	fmt.Print(t)
	fmt.Print("\nItems => ")
	fmt.Print(len(m))
	fmt.Print("\nTotal Post => ")
	fmt.Print(post_cnt)
	fmt.Print("\nTotal Get_Hit => ")
	fmt.Print(get_cnt)
	fmt.Print("\nTotal Get_Miss => ")
	fmt.Print(miss_cnt)
	fmt.Print("\nTotal Delet => ")
	fmt.Print(del_cnt)

	pretty.Printf("\n--- m:\n%# v\n\n", m)
}

// Post Data -------------------
func postHandler(w http.ResponseWriter, r *http.Request) {

	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	param := retParam(r)

	/* 所々の理由により中止
	// Json Check
	if isJSON(body) == false {
		w.Write([]byte("{\"Message\":\"Invalid Json Data!! Post Aborted...\"}\n"))
		return
	}
	*/

	m[param] = body

	outputHeader(w)

	w.Write([]byte("{\"Message\":\"Data Insert => Key:"))
	w.Write([]byte(param))
	w.Write([]byte("\"}\n"))

	post_cnt++
}

// Delete Data -----------------
func delHandler(w http.ResponseWriter, r *http.Request) {

	param := retParam(r)

	delete(m, param)

	outputHeader(w)
	w.Write([]byte("{\"Message\":\"Delete Key => "))
	w.Write([]byte(param))
	w.Write([]byte("\"}\n"))

	del_cnt++
}

func main() {

	// Parameter Setting ------------------
	var port_int int
	flag.IntVar(&port_int, "port", 8080, "HTTP-KVS Server Port")
	flag.Parse()
	port_str := ":" + fmt.Sprint(port_int)

	// Timer
	t = time.Now() // 現在時刻を得る

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
