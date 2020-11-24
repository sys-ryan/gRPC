package main
  
import (
        "time"
        "fmt"
)

type CachingData struct{
        serviceName string
        lastCalled time.Time
        c_input string
        c_output string
}

type CacheMap struct{
        cmap map[string]CachingData
}

func NewCache() CacheMap {
        cm := CacheMap{cmap: map[string]CachingData{} }
        return cm
}


func (cm CacheMap) CacheIn(name string, cdata CachingData) {
        cm.cmap[name] = cdata
}

func (cm CacheMap) CacheGet(name string) (CachingData, bool) {
        data, ok := cm.cmap[name]
        if ok {
                return data, ok
        } else {
                fmt.Println("[RYAN] Cache MISS")
                return CachingData{}, ok
        }
}

/*
func main(){
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
*/
