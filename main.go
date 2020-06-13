package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("Date:  %s\n", now.Format("Jan 2, 2006"))
}
