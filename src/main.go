package main // must have main for entry

import ( // unused pkg cause error
	"fmt"
	"os"
)

// *functions*
func getNumber() int {
	return 8282
}

func log(message string) { // void

}

// func add(a,b int) int // short
func add(a int, b int) int { // 2 in 1 out
	return 0 // 1 out
}

func power(name string) (int, bool) { // 1 in 2 out
	return 0, false // 2 out
}

// takes a pointer to a User and make it grow
func Grow(usr *User) {
	usr.Age += 100
}

// *structs*
type User struct {
	Name string
	Age int
	Boss *User // field is pointer to another struct
}
// add method Senior to a User
// type User is receiver of Senior method
func (usr *User) Senior() {
	usr.Name += "Senior"
	fmt.Println(usr.Name, " is now a Senior")
}
type Ninja struct { // composition
	*User // enables methods of User
	Attack int
}
func (usr *Ninja) Senior() { // mimic overloading
	usr.Name += "NinjaSenior"
	fmt.Println(usr.Name, " is now a NinjaSenior")
}

// User factory
func NewUser(name string, age int) *User {// using pointer to User struct, (optional, can use `User`` to return whole copy of User)
	return &User{ // returning address of new User, (optional, can use `User``)
		Name: name,
		Age: age,
		Boss: &User { 
			Name: "Boss",
			Age: 100,
			Boss: nil,
		},
	}
}


func main() { // must called main for entry point
	// *basics*
	// Args[0] is path to current bin
	if len(os.Args) > 2{ // size of str, dict, array
		os.Exit(1)
	}
	var number int // declare
	number = 82 // assign
	var num string = "num" // unused var cause error
	n, nfunc := 2, getNumber() // multi-assign
	num, newvar := "samenum", "newvar"// at least one new var to use :=
	if len(os.Args) == 2{
		fmt.Println("Args[1]:", os.Args[1])
	}
	fmt.Printf("number:%d\n", number) // print format
	fmt.Printf("num:%s n:%d \n", num, n)
	fmt.Printf("nfunc:%d newvar:%s\n", nfunc, newvar)
	_, exists := power("wh") // _ not assign, any type
	if exists == false {
		// error
	}

	// *structs*
	wh := User{
		Name: "WuHao",
	}
	wh.Age = 31
	fmt.Println(wh.Name, wh.Age)
	whp := &User{"Hao",18, &wh} // the address of a new User
	fmt.Println("whp", whp.Name, whp.Age, whp.Boss)
	Grow(whp) // make User old
	fmt.Println("whp-grown", whp.Name, whp.Age)
	whp.Senior() // make User Senior
	fmt.Println("whp-senior", whp.Name, whp.Age)

	whnew := new(User) // allocate memory for User struct
	whnew.Name = "whnew"
	whnew2 := &User{ //same as above
		Name: "whnew2",
	}
	fmt.Println("whnew and 2", whnew.Name, whnew2.Name)

	ninja0 := &Ninja{// use composition
		User: &User{"Ninja 0", 14, &wh},
		Attack: 10,
	}
	ninja0.Senior()// method of User
	fmt.Println(ninja0.Name, ninja0.User.Name)// .User can be explicitly used
}