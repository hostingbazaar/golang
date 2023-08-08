package main

import (
	"encoding/json"
	"net/http"

	"errors"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// hand type for handler//
func CreateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

func GetEmpid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employeeName := mux.Vars(r)["name"]
	var employeeid Employee
	result := Database.First(&employeeid, "emp_name = ?", employeeName)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Employee not found
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})
		} else {
			// Other error
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		}
		return
	}
	// Encode the employee as JSON and send it in the response
	json.NewEncoder(w).Encode(employeeid)
}

func UpdEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employeeUpdate := mux.Vars(r)["name"]
	var employeename Employee
	results := Database.First(&employeename, "emp_name = ?", employeeUpdate)
	if results.Error != nil {
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			// Employee not found
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})
		} else {
			// Other error
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		}
		return
	}
	json.NewDecoder(r.Body).Decode(&employeename)
	Database.Save(&employeename)
	json.NewEncoder(w).Encode(employeename)
}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	Database.Delete(&emp, mux.Vars(r)["name"])
	msg := "delete data successfully"
	json.NewEncoder(w).Encode(msg)

}
