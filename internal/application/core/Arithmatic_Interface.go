package core

// Interface which the arithmatic core will implement
type Arithmatic_Interface interface{
	Addition(int32,int32) (int32,error);
	Subtraction(int32,int32) (int32,error);
	Multiplication(int32,int32) (int32,error);
	Division(int32,int32) (int32,error);
}