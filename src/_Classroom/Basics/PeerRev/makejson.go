package main

import (
  "fmt"
  "bufio"
  "os"
  "encoding/json"
)

func main() {
  userMap := make(map[string] string)
  reader := bufio.NewScanner(os.Stdin)

  var userName string
  fmt.Println("Enter a name: ")
  reader.Scan()
  userName = reader.Text()

  var userAddress string
  fmt.Println("Enter an address: ")
  reader.Scan()
  userAddress = reader.Text()

  userMap["name"] = userName
  userMap["address"] = userAddress
  jsonString, _ := json.Marshal(userMap)
  fmt.Println(string(jsonString))
}
