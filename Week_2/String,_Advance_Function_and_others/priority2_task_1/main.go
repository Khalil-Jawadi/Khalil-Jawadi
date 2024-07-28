package main

import (
	"fmt"
)

type Vehicle struct {
	Name      string
	Brand     string
	Type      string
	RentPrice float64
}

type Customer struct {
	Name           string
	Address        string
	RentedVehicles []*Vehicle
}

type RentalSystem struct {
	Customers []Customer
	Vehicles  []Vehicle
}

func (rs *RentalSystem) AddCustomer(name, address string) {
	rs.Customers = append(rs.Customers, Customer{Name: name, Address: address})
}

func (rs *RentalSystem) AddVehicle(name, brand, vType string, rentPrice float64) {
	rs.Vehicles = append(rs.Vehicles, Vehicle{Name: name, Brand: brand, Type: vType, RentPrice: rentPrice})
}

func (rs *RentalSystem) RentVehicle(customerName, vehicleName string) {
	for i := range rs.Customers {
		if rs.Customers[i].Name == customerName {
			for j := range rs.Vehicles {
				if rs.Vehicles[j].Name == vehicleName {
					rs.Customers[i].RentedVehicles = append(rs.Customers[i].RentedVehicles, &rs.Vehicles[j])
					rs.Vehicles = append(rs.Vehicles[:j], rs.Vehicles[j+1:]...)
					return
				}
			}
		}
	}
}

func (rs *RentalSystem) ReturnVehicle(customerName, vehicleName string) {
	for i := range rs.Customers {
		if rs.Customers[i].Name == customerName {
			for j, v := range rs.Customers[i].RentedVehicles {
				if v.Name == vehicleName {
					rs.Vehicles = append(rs.Vehicles, *v)
					rs.Customers[i].RentedVehicles = append(rs.Customers[i].RentedVehicles[:j], rs.Customers[i].RentedVehicles[j+1:]...)
					return
				}
			}
		}
	}
}

func (rs *RentalSystem) ListCustomers() {
	fmt.Println("Customers:")
	for _, customer := range rs.Customers {
		fmt.Printf("Name: %s, Address: %s, Rented Vehicles: ", customer.Name, customer.Address)
		for _, vehicle := range customer.RentedVehicles {
			fmt.Printf("%s ", vehicle.Name)
		}
		fmt.Println()
	}
}

func (rs *RentalSystem) ListRentedVehicles() {
	fmt.Println("Rented Vehicles:")
	for _, customer := range rs.Customers {
		for _, vehicle := range customer.RentedVehicles {
			fmt.Printf("Vehicle: %s, Rented by: %s\n", vehicle.Name, customer.Name)
		}
	}
}

func main() {
	rentalSystem := RentalSystem{}

	rentalSystem.AddCustomer("Alice", "123 Main St")
	rentalSystem.AddCustomer("Bob", "456 Oak St")

	rentalSystem.AddVehicle("Corolla", "Toyota", "Sedan", 100)
	rentalSystem.AddVehicle("Civic", "Honda", "Sedan", 120)
	rentalSystem.AddVehicle("F-150", "Ford", "Truck", 150)

	rentalSystem.RentVehicle("Alice", "Corolla")
	rentalSystem.RentVehicle("Bob", "Civic")

	rentalSystem.ListCustomers()
	rentalSystem.ListRentedVehicles()

	rentalSystem.ReturnVehicle("Alice", "Corolla")
	rentalSystem.ListCustomers()
	rentalSystem.ListRentedVehicles()
}
