package main
 
import "fmt"
import "time"
import "os"
import "strconv"
 
func main() {
    dt := time.Now()
    sleep := time.Duration(1)
    argsWithoutProg := os.Args[1:]
    counter, err := strconv.Atoi(argsWithoutProg[0])
    if err != nil {
        panic(err)
    }
    increment := counter
    fmt.Println(dt.String())
    time.Sleep(sleep * time.Second)
    for {
        td := time.Duration(counter)
        timein := dt.Add(-td*time.Hour)
        fmt.Println(timein)
        time.Sleep(sleep * time.Second)
        counter = counter + increment
    }
}