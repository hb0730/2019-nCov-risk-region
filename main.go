package main

import (
	"2019-nCov-risk-region/source"
	"encoding/json"
	"flag"
	"net/http"
)

var (
	serverAddr = flag.String("server", ":80", "Listen Server (ip:port)")
)

func main() {
	flag.Parse()
	handle()
	err := http.ListenAndServe(*serverAddr, nil)
	if err != nil {
		panic(err)
	}
}
func handle() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			write(w, resultFailed("invalid http method"))
			return
		}
		name := request.URL.Query().Get("type")
		if name == "" {
			name = "gov"
		}
		result := source.All(name)
		write(w, resultSuccess(result))
	}))
	http.Handle("/high", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			write(w, resultFailed("invalid http method"))
			return
		}
		name := request.URL.Query().Get("type")
		if name == "" {
			name = "gov"
		}
		write(w, resultSuccess(source.HighRisk(name)))
	}))
	http.HandleFunc("/middle", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			write(w, resultFailed("invalid http method"))
			return
		}
		name := request.URL.Query().Get("type")
		if name == "" {
			name = "gov"
		}
		write(w, resultSuccess(source.MiddleRisk(name)))
	}))
}
func write(w http.ResponseWriter, rt []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rt)
}
func resultSuccess(data interface{}) (bt []byte) {
	return result(true, data, "")
}
func resultFailed(err string) (bt []byte) {
	return result(false, nil, err)
}
func result(success bool, data interface{}, err string) (rt []byte) {
	r := Result{
		Success: success,
		Data:    data,
		Error:   err,
	}
	rt, _ = json.Marshal(&r)
	return
}

type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}
