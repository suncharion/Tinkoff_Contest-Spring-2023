package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type IntWrapper struct {
  Value int
}

func (iw *IntWrapper) ChangeIfLesser(newValue int) {
  if iw.Value > newValue {
    iw.Value = newValue
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  length, _ := reader.ReadString('\n')
  text, _ := reader.ReadString('\n')
  length = strings.Trim(length, "\r\n")
  text = strings.Trim(text, "\r\n")
  fmt.Print(FindGoodSubstring(length, text))
}

func FindGoodSubstring(length string, input string) int {
  letters := map[byte]int{
    'a': 0,
    'b': 0,
    'c': 0,
    'd': 0,
  }

  var left, have, need, currentLength int = 0, 0, 4, 0
  var result *IntWrapper = nil

  for right := range input {
    letters[input[right]]++
    if letters[input[right]] == 1 {
      have++
    }
    for have == need {
      if letters[input[left]] == 1 {
        have--
      }
      currentLength = right - left + 1
      if result == nil {
        result = &IntWrapper{
          Value: currentLength,
        }
      } else {
        result.ChangeIfLesser(currentLength)
      }
      letters[input[left]]--
      left++
    }
  }
  if result == nil {
    return -1
  } else {
    return result.Value
  }
}