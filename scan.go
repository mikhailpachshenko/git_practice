package main

import (
	"bufio"
	"fmt"
	"os"
)

func strScan(input string) (outStr string) {
	scn := bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		outStr = scn.Text()
		fmt.Println("Out:", outStr)
	}

	return
}

// asdfasfdaffaasdf
