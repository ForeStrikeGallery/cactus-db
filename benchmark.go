package main

import (
    "bytes"
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

func BenchmarkWrites(s Store, count int) {
    start := time.Now()

    for i := 0; i < count; i++ {
        start_i := time.Now()
        s.Put(string(i), string(i)) 
        fmt.Printf("Write time for %d: %.5f ms", i, time.Since(start_i).Seconds() * 1000)
    }

    fmt.Printf("Write time for set: %.5f s", time.Since(start).Seconds()) 
}


func BenchmarkReads(s Store, count int) {
    start := time.Now()

    for i := 0; i < count; i++ {
        start_i := time.Now()

        s.Get(string(i))

        fmt.Printf("Read time for %d: %.5f ms\n", i, time.Since(start_i).Seconds() * 1000)
    }

    fmt.Printf("Read time for set of %d: %.5f s", count, time.Since(start).Seconds()) 
}

func BenchHTTPGet(count int) {
    start := time.Now()

    for i := 0; i < count; i++ {
        start_i := time.Now()

        url := "http://localhost:3001/get?key=" + string(i)
        http.Get(url) 

        fmt.Printf("Read time for %d: %.5f ms\n", i, time.Since(start_i).Seconds() * 1000)
    }

    fmt.Printf("Read time for set of %d: %.5f s", count, time.Since(start).Seconds()) 
}

func BenchHTTPut(count int) {
    start := time.Now()

    for i := 0; i < count; i++ {
        start_i := time.Now()

        data := map[string]string{
            "key": string(i),
            "value": string(i), 
        }

        jsonData, _:= json.Marshal(data)

        url := "http://localhost:3001/put"
        http.Post(url, "application/json", bytes.NewBuffer(jsonData))

        fmt.Printf("Write time for %d: %.5f ms\n", i, time.Since(start_i).Seconds() * 1000)
    }

    fmt.Printf("Write time for set: %.5f s", time.Since(start).Seconds()) 
}

