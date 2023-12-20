package api

import (
	"log"
	"github.com/UtkarshM-hub/Hex/internal/application/core"
)

// struct which consist of Arith interface
// basically api communicating with our core logic
type API struct{
	Arith core.Arithmatic_Interface
}

func New(Arithmatic *core.Arithmatic) *API {
	return &API{Arith: Arithmatic};
}

func (A *API) GetAddition(a, b int32) (int32, error){
	val,err:=A.Arith.Addition(a,b)
	if err!=nil{
		log.Fatal(err)
	}
	return val,nil
}

func (A *API) GetSubtraction(a, b int32) (int32, error){
	val,err:=A.Arith.Subtraction(a,b)
	if err!=nil{
		log.Fatal(err)
	}
	return val,nil
}

func (A *API) GetMultiplication(a, b int32) (int32, error){
	val,err:=A.Arith.Multiplication(a,b)
	if err!=nil{
		log.Fatal(err)
	}
	return val,nil
}

func (A *API) GetDivision(a, b int32) (int32, error){
	val,err:=A.Arith.Division(a,b)
	if err!=nil{
		log.Fatal(err)
	}
	return val,nil
}
