package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type Customer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	MobileNumber string `json:"mobilenumber"`
}
var ( 
	db *gorm.DB
)
func Connect(){
	
	
	d, err := gorm.Open("mysql", "root:Shravan@123#/sample?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	db=d
}
func GetDB() *gorm.DB{
	return db
}

var customers []Customer
func main(){
	
	db, err := gorm.Open("mysql", "root:Shravan@123#@/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("unable to connect database",err.Error())
	}
	customers = append(customers, Customer{ID: "1", Name: "shravan",Email:"shravan@gmail.com",MobileNumber:"7702750330"})
	customers = append(customers, Customer{ID: "2", Name: "chintu",Email:"chintu@gmail.com",MobileNumber:"7702750331"})
	//result:=db.Model("customers").Create(&customers)
	//fmt.Println(result)

	fmt.Println(db)
	fmt.Println(err)
	
    
    

    

    
	r := mux.NewRouter()
	
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", createCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	fmt.Printf("starting server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
	
}
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
func  deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range customers {
		if item.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(customers)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range customers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mm Customer
	_ = json.NewDecoder(r.Body).Decode(&mm)
	
	customers = append(customers, mm)
	json.NewEncoder(w).Encode(customers)

}
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range customers {
		if item.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)
			var mm Customer
			_ = json.NewDecoder(r.Body).Decode(&mm)
			mm.ID = params["id"]
			customers = append(customers, mm)

		}
	}
	json.NewEncoder(w).Encode(customers)
}
