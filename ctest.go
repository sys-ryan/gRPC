// Cache test 
// go run cache.go ctest.go 

package main
  
import (
        "fmt"
        "time"
)


func main() {

        Cache := NewCache()

        cdata := CachingData{
                serviceName: "Service name",
                lastCalled: time.Now(),
                c_input: "Input data",
                c_output: "Output data",
        }

        fmt.Println(Cache)
        fmt.Println(cdata)

        fmt.Println()
        Cache.CacheIn("Service name", cdata)
        Cache.CacheGet("Service name")
        Cache.CacheGet("Wrong name")

        fmt.Println()
        fmt.Println("done")
}
