// see type Visited: we indicated that it can be used concurrently. Do not
// assume that a method can be used concurrently unless explicitly documented!
// Also it's considerd good practice to keep Mutex directly above variable it is
// guarding and to include a comment to make that clear

package main

import (
	"fmt"
	"sync"
)

// Visited tracks how often an url has been visited
// its methods can be safely used concurrently
type Visited struct {
	// mu guards the visited map
	mu      sync.Mutex     // unlocked by default
	visited map[string]int // [url]count
}

// VisitLink checks map for url and adjusts counter
// adds new url to map if not already there
func (v *Visited) VisitLink(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url] // if url doesn't exist will be 0
	count++
	v.visited[url] = count // if url dodesn't exist will create new entry
}

func main() {

	site_index := Visited{
		visited: map[string]int{
			"someurl.com": 1,
			"another.org": 1,
			"popular.net": 7,
			"new.co.uk":   0,
		},
	}

	site_index.VisitLink("popular.net") // 8
	site_index.VisitLink("someurl.com") // 2
	site_index.VisitLink("unknown.ru")  // will be added, count 1
	site_index.VisitLink("popular.net") // 9
	fmt.Printf("%v\n", site_index.visited)

}
