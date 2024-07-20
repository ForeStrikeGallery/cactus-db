package main 

import (
   "fmt"
   "errors"
)

var KeyNotFoundErr = errors.New("Key Not Found")

type Store struct {
   data map[string]string
}

func (s *Store) Put(key string, value string) error {
   s.data[key] = value  
   fmt.Println("successfully put value in data")
   return nil 
}

func (s *Store) Get(key string) (string, error) {

   if val, ok := s.data[key]; ok {
      fmt.Println("Value found : ", key, val) 
      return val, nil
   } else {
      return "", KeyNotFoundErr  
   }

   return "hi", nil 
}

