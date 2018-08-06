package main

import (
	"fmt"

	"github.com/jay-dee7/v4.O/cmd"
	"github.com/jay-dee7/v4.O/core"
	"github.com/jay-dee7/v4.O/wallet"
)

func main() {
	cmd.Execute()

	core.NewCore()
	s := wallet.NewWallet("blockchain-4.0")
	fmt.Println(s)
}
