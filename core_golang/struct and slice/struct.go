package main

import "fmt"

type contactInfo struct{
	email  string
	zipCode int
	mobileNo int
}

// type employee struct{
// 	empId int
// 	firstName string
// 	lastName string
// 	contact contactInfo
// }

type employee struct{
	empId int
	firstName string
	lastName string
	 contactInfo
}
func main(){
	//1st type to declare struct
	// pratiksha := employee{empId: 1,firstName: "Pratiksha",lastName: "Nagoshe"}
	// fmt.Println(pratiksha)

	//2nd one
	//if value not assign it will take zero value or empty string
	// var josh employee
	// fmt.Println(josh)
	// fmt.Printf("%+v", josh)

	//3rd one
	// var josh employee
	// josh.empId = 1
	// josh.firstName ="Josh"
	// josh.lastName = "alexzender"
	// fmt.Println(josh)
	// fmt.Printf("%+v", josh)


	viru := employee{
		empId: 1,
		firstName: "Viru",
		lastName: "Ram",
		contactInfo: contactInfo{
			email: "viraujay@yahoo.com",
			zipCode: 441904,
			mobileNo: 7788665544,
		},
	}

	viru.updateFirstName("Jay")
	viru.print()
}

func (e employee) print(){
	fmt.Printf("%v",e)
}

func (employeePointer *employee) updateFirstName(newfirstName string){
	(*employeePointer).firstName = newfirstName
}