package main

import "fmt"

type ParkingSystem struct {
	big, medium, small int
}

func ParkingSystemConstructor(big int, medium int, small int) ParkingSystem {
	ps := &ParkingSystem{big, medium, small}
	return *ps
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	var target *int

	switch carType {
	case 1:
		target = &ps.big
	case 2:
		target = &ps.medium
	case 3:
		target = &ps.small
	}

	if *target > 0 {
		*target--
		return true
	}

	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func driver__AddCar() {
	obj := ParkingSystemConstructor(1, 1, 0)
	fmt.Println(obj.AddCar(1))
	fmt.Println(obj.AddCar(2))
	fmt.Println(obj.AddCar(3))
	fmt.Println(obj.AddCar(1))
}
