package main

import (
	"bufio"
	"fmt"
	"io"
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

func loopInt(input int) (out int) {
	for {
		_, err := fmt.Scan(&out)
		if err == io.EOF {
			return
		}
		fmt.Println("out:", out)
	}
	return
}

<<<<<<< HEAD
// asdfasfdaffaasdf
=======
// some text about included
>>>>>>> parent of a75e055 (some changes minor)
