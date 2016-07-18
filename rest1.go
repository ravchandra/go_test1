package main

import (
    //"fmt"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "strconv"
)

type Employee struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

var emps = []Employee{
    Employee{Id: 1, Name: "Ravi"},
    Employee{Id: 2, Name: "Chandra"},
}

func GetHandler1(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)

    id, _  := strconv.Atoi(vars["id"])

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    for i:=0; i<len(emps); i++{
        if emps[i].Id == id{
            json.NewEncoder(w).Encode(emps[i].Name)
            break
        }
    }

    /*
    if id > len(emps){
        json.NewEncoder(w).Encode("Employee doesn't exist")
    }else{
        json.NewEncoder(w).Encode(emps[id-1].Name)
    }*/

    //w.Write([]byte("GetTest\n"))
}

func GetHandler2(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    json.NewEncoder(w).Encode(emps)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)

    var e Employee
    err := decoder.Decode(&e)

    if err != nil {
        panic(err)
    }

    newe := true
    for i:=0; i<len(emps); i++{
        if emps[i].Id == e.Id{
            newe = false
            break
        }
    }
    if newe == true{
        ne := new(Employee)
        *ne = e    
        emps = append(emps, *ne)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    json.NewEncoder(w).Encode(emps)

    //json.NewEncoder(w).Encode("Recieved POST with employee name "+e.Name)
    /*
    for i := 0; i < len(emps); i++{
        u:=emps[i]
        json.NewEncoder(w).Encode(u)
    }*/
    //w.Write([]byte("PostTest\n"))
}

func DelHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    id, _  := strconv.Atoi(vars["id"])

    for i:=0; i<len(emps); i++{
        if emps[i].Id == id{
            emps = append(emps[:i],emps[i+1:]...)
            break
        }
    }
   
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    json.NewEncoder(w).Encode(emps)
}

func PutHandler(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)

    var e Employee
    err := decoder.Decode(&e)

    if err != nil {
        panic(err)
    }
    for i:=0; i<len(emps); i++{
        if emps[i].Id == e.Id{
            emps[i].Name = e.Name
            break
        }
    }

    //emps[]

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    json.NewEncoder(w).Encode(emps)
}

func main() {
    r := mux.NewRouter()

    //r.HandleFunc("/test/{id}/",GetHandler1).Methods("GET")
    
    r.HandleFunc("/test/",GetHandler2).Methods("GET")
    
    r.HandleFunc("/test/",PostHandler).Methods("POST")
    
    r.HandleFunc("/test/{id}/",DelHandler).Methods("DELETE")
    
    r.HandleFunc("/test/",PutHandler).Methods("PUT")

    http.ListenAndServe(":8081", r)
}
