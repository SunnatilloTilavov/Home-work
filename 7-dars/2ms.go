package main

import (
  "fmt"
  "strings"
  "unicode"
)

func main() {
    fmt.Println(findIngLyWords("Walking is healthy Swiftly they ran."))
}

func findIngLyWords(text string) []string {
  words := strings.Fields(text)

  result := []string{}

  for _, word := range words {
    if len(words) > 2 && unicode.IsUpper(rune(word[0])) && strings.HasSuffix(word, "ing") || strings.HasSuffix(word, "ly")  {
      result = append(result, word)
    }
  }
  return result
}