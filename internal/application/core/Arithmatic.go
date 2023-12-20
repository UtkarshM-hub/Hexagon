package core

import "errors"

// Core Arithmatic structure
type Arithmatic struct {

}

// returns new Arithmatic Struct reference
func New() *Arithmatic {
	return &Arithmatic{};
}

// Core methods
func (Arith Arithmatic) Addition(a,b int32) (int32,error){
	return (a+b),nil
}

func (Arith Arithmatic) Multiplication(a,b int32) (int32,error){
	return (a*b),nil
}
func (Arith Arithmatic) Division(a,b int32) (int32,error){
	if a==0 || b==0{
		return 0,errors.New("One of the provided value is 0")
	}
	return (a/b),nil
}
func (Arith Arithmatic) Subtraction(a,b int32) (int32,error){
	return (a-b),nil
}