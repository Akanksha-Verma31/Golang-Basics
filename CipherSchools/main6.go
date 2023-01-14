/*
Determine if a triangle is equilateral, isosceles, or scalene.

An equilateral triangle has all three sides the same length.

An isosceles triangle has at least two sides the same length. (It is sometimes specified as having exactly two sides the same length, but for the purposes of this exercise we'll say at least two.)

A scalene triangle has all sides of different lengths.
*/
package main

import "fmt"
import "errors"

func main(){
	var side1 int
	var side2 int
	var side3 int
	fmt.Println("Enter sides of the triangle : ")
	fmt.Scan(&side1)
	fmt.Scan(&side2)
	fmt.Scan(&side3)
	s, err := check(side1, side2, side3)
	if err != nil{
		fmt.Println("Okay")
	}
	fmt.Println(s)
	// s1:=1
	// s2:=1
	// s3:=1
	//checkTriangle(s1,s2,s3)
}

func check(s1, s2, s3 int)(string, error){
	if (s1+s2<=s3 || s1+s3<=s2 || s2+s3<=s1) || (s1<=0 || s2<=0 || s3<=0){
		 fmt.Println("Not a triangle")
		 return " ", errors.New("Not is not a triangle")
	}else if s1==s2 && s2==s3 && s3==s1{
		fmt.Println("Equilateral")
		return " ", nil
	}else if s1==s2 || s2==s3 || s3==s1{
		fmt.Println("Isosceles")
		return " ", nil
	}else{
		fmt.Println("Scalene")
		return " ", nil
	}
	
}

// func checkTriangle(s1 int, s2 int, s3 int){
//     if (s1 == s2) && (s2==s3) && (s3 == s1){
// 				fmt.Println("Equilateral Triangle")
// 	}
// 	else if ((s1 == s2)&&(s1==s3) || (s2==s1)&&(s2==s3) || (s3==s1)&&(s3==s2)){
// 	     fmt.Println("Isosceles")
// 	}
// 	else if (s1 != s2) && (s2 != s3) && (s3 != s1){
// 		fmt.Println("Scalene")
// 	}
// }