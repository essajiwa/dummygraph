package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	log.Println("Start generating RDF.")
	o, err := os.OpenFile("feed.rdf.gz", os.O_WRONLY|os.O_CREATE, 0755)
	check(err)

	w := gzip.NewWriter(o)
	var str string
	runningAct := 1

	// Generate Shop with randomize total activity

	for i := 1; i <= 50000; i++ {
		str = fmt.Sprintf("<_:shop%d> <idshop> \"%d\"^^<xs:int> .\n", i, i)
		w.Write([]byte(str))
		// Each shop have random between 100 and 200 total activity
		totalActivity := random(100, 200)
		for j := 1; j <= totalActivity; j++ {
			// Activity
			str = fmt.Sprintf("<_:act%d.%d> <idactivity> \"%d\"^^<xs:int> .\n", i, j, runningAct)
			w.Write([]byte(str))
			// Create Time
			str = fmt.Sprintf("<_:act%d.%d> <create_time> \"%d\"^^<xs:int> .\n", i, j, runningAct)
			w.Write([]byte(str))
			// Sambungin ke shop pake doing
			str = fmt.Sprintf("<_:shop%d> <doing> _:act%d.%d .\n", i, i, j)
			w.Write([]byte(str))

			runningAct++
		}
	}

	totalFollow := 0
	// Generate 1 million user
	for i := 1; i <= 1000000; i++ {
		str = fmt.Sprintf("<_:user%d> <iduser> \"%d\"^^<xs:int> .\n", i, i)
		w.Write([]byte(str))

		// Set user fav toko random tiap user follow minimal 10 sampai 500 toko
		totalShopFollowed := random(10, 500)
		for s := totalShopFollowed; s <= 500; s++ {
			str = fmt.Sprintf("<_:user%d> <favorite> _:shop%d .\n", i, s)
			w.Write([]byte(str))
			totalFollow++
		}
	}

	log.Println("Finished generating RDF.")
	fmt.Printf("Total Activity : %d\n", runningAct)
	fmt.Printf("Total Favorite : %d\n", totalFollow)

	err = w.Flush()
	check(err)

	err = w.Close()
	check(err)

	err = o.Close()
	check(err)
}
