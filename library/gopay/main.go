package main

import (
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func main() {
	xlog.Info("GoPay Version: ", gopay.Version)
	fmt.Println("GoPay Version: ", gopay.Version)
}
