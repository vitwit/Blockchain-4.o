package main

import (
	"fmt"

	"github.com/vitwit/Blockchain-4.o/core"
	"github.com/vitwit/Blockchain-4.o/wallet"
)

func main() {
	core.NewCore()
	s := wallet.NewWallet("blockchain-4.0")
	fmt.Println(s)
}
