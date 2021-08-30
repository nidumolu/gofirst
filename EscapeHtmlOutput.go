package main

import (
   "fmt"
   "html"
)

func main() {
   rawString := `<script>alert("Test...by ..Srini");</script>`
   safeString := html.EscapeString(rawString)

   fmt.Println("Unescaped: " + rawString)
   fmt.Println("Escaped: " + safeString)
}