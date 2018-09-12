package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ore  = "ore"
	rock = "ROCK"

	minedOre   = "minedOre"
	smeltedOre = "smeltedOre"
)

type Mine struct {
	Found   chan string
	Mined   chan string
	Smelted []string
}

var mine Mine

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	now := time.Now()
	var emptyCh = make(chan bool)
	mine = Mine{
		Found: make(chan string),
		Mined: make(chan string),
	}

	theMine := generateMine()
	go finder(theMine, emptyCh)

	// Add load balancer
	go miner(emptyCh)
	//go miner(emptyCh)

	smelter(emptyCh)

	fmt.Println(time.Since(now))
}

func finder(materials []string, emptyCh chan bool) {
	defer func() {
		emptyCh <- true
	}()

	for _, material := range materials {
		if material == ore {
			mine.Found <- material
		}
	}

}

func miner(emptyCh chan bool) {
	defer func() {
		emptyCh <- true
	}()

	for {
		select {
		case <-emptyCh:
			return
		case found := <-mine.Found:
			mine.Mined <- upgrader(found)
		}
	}
}

func smelter(emptyCh chan bool) {
	for {
		select {
		case <-emptyCh:
			return
		case mined := <-mine.Mined:
			mine.Smelted = append(mine.Smelted, upgrader(mined))
		}
	}
}

func upgrader(material string) string {
	switch material {
	case ore:
		return minedOre
	case minedOre:
		return smeltedOre
	}
	return ""
}

func generateMine() []string {
	var resultSet []string
	mineSize := randInt(1000000, 3000000)
	fmt.Printf("Size of the mine: %d\n", mineSize)
	for i := 0; i < mineSize; i++ {
		if randInt(0, 20)%2 == 0 {
			resultSet = append(resultSet, ore)
		} else {
			resultSet = append(resultSet, rock)
		}
	}

	return resultSet
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
