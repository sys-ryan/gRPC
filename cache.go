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

func CacheUpdateTime(name string, cdata CachingData) CachingData{ 
        return CachingData{
                        serviceName: cdata.serviceName,
                        lastCalled: time.Now(),
                        c_input: cdata.c_input,
                        c_output: cdata.c_output,
                }
}


func (cm CacheMap) CacheIn(name string, cdata CachingData) {
        cm.cmap[name] = cdata
}

func (cm CacheMap) CacheGet(name string, input string) (CachingData, bool) {
        data, ok := cm.cmap[name]
        if ok && (data.c_input == input) {
                fmt.Println("last called : ")
                cm.cmap[name] = CacheUpdateTime(name, data)
                fmt.Println(cm.cmap[name].lastCalled)
                return data, ok
        } else {
                return CachingData{}, false
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

        fmt.Println("<HIT>\n")
        data, ok := Cache.CacheGet("Service name", "Input data")
        fmt.Println(data)
        fmt.Println(ok)
        fmt.Println()

        fmt.Println("<MISS - Wrong data>\n")
        data, ok = Cache.CacheGet("Service name", "Wrong data")
        fmt.Println(ok)
        fmt.Println()

        fmt.Println("<MISS - Wrong name>\n")
        data, ok = Cache.CacheGet("Wrong name", "123")
        fmt.Println(ok)
        fmt.Println()

        fmt.Println("done")
}
*/
