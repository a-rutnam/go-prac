package main

import (
    // "strings"
    "fmt"
    "unicode"
)

func main() {
  // statement := "Can I help you?"
  // statement := "YO"
  // statement := "Serge"
  statement := "bla"

  serge_says(statement)
}

func serge_says(statement string) {

  last_character := statement[len(statement)-1:]

  if last_character == "?" {
    fmt.Println("Sure")
  } else if IsUpper(statement) {
    fmt.Println("Woah, chill out!")
  } else if statement == "Serge" {
    fmt.Println("Fine. Be that way!")
  } else {
    fmt.Println("Whatever")
  }

}

func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

// func IsUpper(r rune) bool
// Warmup - Serge Says
// Examples
// Serge answers 'Sure.' if you ask him a question.
//
// He answers 'Woah, chill out!' if you yell at him (ALL CAPS).
//
// He says 'Fine. Be that way!' if you address him without actually saying anything.
//
// He answers 'Whatever.' to anything else.
//
// Create a function that takes an input and returns Serge's response.
