package Non_Blocking_cache
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)
func HttpGetBody(url string) (interface{},error){
	resp,err := http.Get(url)
	if err !=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

}
func incomingURL()<-chan string{
	ch:=make(chan string)
	go func() {
		//Loop statement
		for _, url := range []string{
			"https://en.wikipedia.org/wiki/Juju_(software)",
			"https://en.wikipedia.org/wiki/Agent-based_model",
			"http://www.google.com",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://en.wikipedia.org/wiki/Juju_(software)",
			"https://en.wikipedia.org/wiki/Agent-based_model",
			"http://www.google.com",
		}{
			//Loop body
			ch <- url
		}
		//Closes channel before ending loop
		close(ch)
	}()
		return ch
	}
func ConTest(t *testing.T, nbc *Nbc) {
	var n sync.WaitGroup
	for url := range incomingURL() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := nbc.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s , %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))

		}(url)
	}
	n.Wait()
}
func TestCon(t *testing.T){
	nbc := New(HttpGetBody)
	ConTest(t,nbc)

}

