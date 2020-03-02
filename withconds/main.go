package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Variables partagées
var (
	violons         = sync.NewCond(&sync.Mutex{})
	archets         = sync.NewCond(&sync.Mutex{})
	nbViolonsDispos = 3
	nbArchetsDispos = 2
)

func musicien(num int) {
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		violons.L.Lock()
		for nbViolonsDispos == 0 {
			violons.Wait()
		}
		nbViolonsDispos--
		fmt.Printf("Le musicien %d a pris un violon\n", num)
		violons.L.Unlock()

		archets.L.Lock()
		for nbArchetsDispos == 0 {
			archets.Wait()
		}
		nbArchetsDispos--
		fmt.Printf("Le musicien %d a pris un archet\n", num)
		fmt.Printf("Il reste %d violon(s) et %d archet\n", nbViolonsDispos, nbArchetsDispos)
		archets.L.Unlock()

		//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		fmt.Printf("....... Le musicien %d joue du violon....\n", num)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Printf("....... Le musicien %d a fini de jouer....\n", num)

		violons.L.Lock()
		nbViolonsDispos++
		if nbViolonsDispos >= 1 {
			violons.Signal()
		}
		violons.L.Unlock()

		archets.L.Lock()
		nbArchetsDispos++
		if nbArchetsDispos >= 1 {
			archets.Signal()
		}
		archets.L.Unlock()
	}
}

func main() {

	for i := 1; i <= 4; i++ {
		go musicien(i)
	}

	input := ""
	fmt.Scanln(&input)
}
