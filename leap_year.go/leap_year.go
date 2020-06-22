package main

import (
    "fmt"
)

func main() {

  // year := 1996
  // year := 1997
  // year := 1900
  // year := 2000
  leap_year_identifier(year)

}

func leap_year_identifier(year int) {


  if year % 400 == 0 {
    fmt.Println(year, "is a leap year.")
  } else if year % 4 == 0 && year % 100 != 0 {
    fmt.Println(year, "is a leap year.")
  } else {
    fmt.Println(year, "is not a leap year.")
  }
}

// Do this in HTML and JS!
// Write a function that will take any given year and return whether it is a leap year or not. Remember that a leap year is:
//
// Every year that is evenly divisible by 4 Except if it is evenly divisible by 100... Unless it is also divisible by 400
//
// Test Data... Make sure it is working!
//
// 1997 is not a leap year, it should return false
// 1996 is a leap year, it should return true
// 1900 is not a leap year, it should return false
// 2000 is a leap year, it should return true
// How to structure it...
//
// We want a custom function called isLeapYear
// Bonus!
//
// Ask the user what number they want to test
