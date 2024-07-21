package main 

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    _ "net/http/pprof"
    "time"
)

type Handler struct {
    store Store 
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {

    fmt.Println("Get!") 
    key := r.URL.Query().Get("key") 
    
    fmt.Println("Looking for key: ", key) 
    value, err := h.store.Get(key)

    if err != nil {
        fmt.Println("Get call failed: ", err) 
    }

    w.Write([]byte(value))
}

func (h *Handler) put(w http.ResponseWriter, r *http.Request) {
   
    body, err := ioutil.ReadAll(r.Body) 
    
    if err != nil {
        fmt.Println("Couldn't read body", err) 
        http.Error(w, "Couldn't read body", 
            http.StatusInternalServerError)
        return  
    } 

    values, err := url.ParseQuery(string(body)) 
    
    if err != nil {
        http.Error(w, "Couldn't parse query", http.StatusInternalServerError)
        return 
    }

    key := values.Get("key")
    value := values.Get("value") 

    err = h.store.Put(key, value)

    if err != nil {
        http.Error(w, "Unable to put key-value pair into db", 
            http.StatusInternalServerError)
        return 
    }
}
    

func main() {
    // start pprof HTTP server in separate goroutine for debugging 

    go func() {
        fmt.Println("Starting pprof server on :6060")
		http.ListenAndServe(":6060", nil)
    }()


    /*
    store := Store{
        data: make(map[string]string, 10000),
    }

    BenchmarkWrites(store, 1000000)
    BenchmarkReads(store, 1000000)
    */

    h := Handler{
        store: Store{
            data: make(map[string]string, 40000),
        },
    }
    http.HandleFunc("/get", h.get)
    http.HandleFunc("/put", h.put)

    fmt.Println("Cactus Server Running..")
    go http.ListenAndServe(":3001", nil) 
    time.Sleep(1 * time.Second)

    BenchHTTPut(10000)
    // BenchHTTPGet(10000)
    fmt.Scanln()
}
