package main

import (
    "fmt"
    "math/rand"
"reflect"
)

func main() {

  students := make(map[string]int)

  students["Lara"] = 1
  students["Samuel"] = 2
  students["Andie"] = 17
  students["Kelly"] = 7
  students["Mouse"] = 2

  // fmt.Println(students)
  candidates :=

  for _, element := range students {
    if element > 1 {
      // fmt.Println("Key:", key, "=>", "Element:", element)
      fmt.Println("Random Key: ", MapRandomKeyGet(students).(string))


    }
  }

}

func MapRandomKeyGet(mapI interface{}) interface{} {x
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}
//
// The names of those who are witnessed clapping Badger more than twice in the first four days of any one week will go into a draw. On Friday morning, a random name will be drawn from that list and the 'winner' will be forced to deliver the solution to Friday's warmup.
//
// Create a program that Badger can use for this task - ideally, you should create an object (revengeOfBadger) to do this, but if you can get it working using standalone functions, that's cool.
//
// Your program should:
//
// Track how many times each student in the class has clapped this week (just make up the numbers);
// Include a collection of candidates for the Friday draw (ie, a list of names of people who have clapped Badger more than twice that week);
// Pick a random student to deliver the solution to Friday's warmup.
// If no one has clapped enough that week, the program should indicate that Badger has to do his own damned warmup.
