package main 

import (
    "fmt"
    "os"
    "path/filepath"
    "strconv"
)

type DataFile struct { 
    filepath string  
}

func (d *DataFile) write(key string, val string) {
    file, err := os.OpenFile(d.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    defer file.Close()
   
    _, err = file.WriteString(key + ": " + val + "\n")

    if err != nil { 
        fmt.Println("Unable to write to file")
    }

    fmt.Println("Successfully written to file") 
}

type Cactus struct {
    // key->offset
    keyDir map[string]int
    activeFile DataFile 
} 

func Init(dbdir string) (*Cactus, error) {
    
    // read all files from this path and count the number of files in it 
    files, err := filepath.Glob(fmt.Sprintf("%s/*.db", dbdir))

    if err != nil {
        return nil, err
    }

    var activeFilePath string

    if len(files) != 0 {
        activeFilePath = dbdir + "cactus" + strconv.Itoa(len(files)) + ".db"
    } else {
        activeFilePath = dbdir + "cactus1.db"
    }

    cactus := &Cactus{
        keyDir: make(map[string]int),
        activeFile: DataFile{
            filepath: activeFilePath,
        },
    }

    return cactus, nil 
}

func (d *Cactus) Put(key string, val string) error {
    
    d.activeFile.write(key, val) 
    return nil 
}

func (d *Cactus) Get(key string) (string, error) {
    fmt.Println("reading from cactus") 
    return "hello", nil 
}


