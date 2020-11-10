package main

import (
	"flag"

	game "example.com/user/queens/src"
)

//
func main() {
	queensPtr := flag.Int("num", 4, "Game with given queens. MAX=25. num=4")
	shortDisplayPtr := flag.Bool("shortDisplay", false, "Option for short Display list")
	flag.Parse()

	size := *queensPtr
	if size > 25 {
		size = 25
	}

	game.Game(size, *shortDisplayPtr)
}

/*
+---------------------------------------------------------------+
|                            Results                            |
+-----------+-------------------------+-------------------------+
|  Computer | Mac                     | PC                      |
|           +------------+------------+------------+------------+
|           | Go         | Python     | Go         | Python     |
+-----------+------------+------------+------------+------------+
|   Queens  | 15         | 15         | 15         | 15         |
+-----------+------------+------------+------------+------------+
| Solutions |  2,279,184 |  2,279,184 |  2,279,184 |  2,279,184 |
+-----------+------------+------------+------------+------------+
| Time      | 1m37sec    | 116m33sec  |            |            |
+-----------+------------+------------+------------+------------+
// changes from cloned repo at
	g:\My Drive\Programing\Git\Git_youTube\_importGit2\queens.go
*/
