// we write seomthing similar to 144_mutexes_structs.go but now use goroutines
// and see if the mutex is properly acting as a guard

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitSite(URL string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[URL]
	count++
	v.visited[URL] = count
}

func main() {

	siteDB := Visited{
		visited: map[string]int{
			"yandex.com":     0,
			"golang.org":     0,
			"duckduckgo.com": 0,
		},
	}

	someurls := []string{"yandex.com", "golang.org",
		"duckduckgo.com", "wikipedia.org", "golang.org",
		"yandex.com", "golang.org"}

	// start 100 goroutines visiting sites from []string someurls randomly
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		go siteDB.VisitSite(someurls[rand.Intn(7)])
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("%v\n", siteDB.visited)

}

// you can try not using a mutex to see what happens, mostly you will get
// 	fatal error: concurrent map read and map write
// 	fatal error: concurrent map writes
// the likelihood of a fatal error depends on how many goroutines we have
// w/ i < 20 in the for loop it might still work, with i < 2000 it surely won't
