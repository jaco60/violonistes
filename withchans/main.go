package main

import (
	"fmt"
	"math/rand"
	"time"
)

const(
	NbViolons = 3
	NbArchets = 2
)

// Channels de communication : on utilise des jetons struct{} car ils occupent 0 octet...
var (
	violonsDispos = make(chan struct{}, NbViolons)
	archetsDispos = make(chan struct{}, NbArchets)
)

func musicien(num int) {
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		// Le musicien tente de prendre un violon
		<- violonsDispos
		fmt.Printf("Le musicien %d a pris un violon\n", num)

		// Le musicien tente de prendre un archet
		<- archetsDispos
		fmt.Printf("Le musicien %d a pris un archet\n", num)

		//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		fmt.Printf("....... Le musicien %d joue du violon....\n", num)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

		// Le musicien repose son violon, puis son archet...
		violonsDispos <- struct{}{}
		archetsDispos <- struct{}{}
		fmt.Printf("....... Le musicien %d a fini de jouer....\n", num)
	}
}


func main() {

	// Remplissage des canaux de communication
	for i := 1; i <= NbViolons; i++ {
		violonsDispos <- struct{}{}
	}
	for i := 1; i <= NbArchets; i++ {
		archetsDispos <- struct{}{}
	}

	// Lancement des musiciens
	for i := 1; i <= 4; i++ {
		go musicien(i)
	}

	// Pour éviter que le prog se termine...
	input := ""
	fmt.Scanln(&input)
}
