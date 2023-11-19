package main

import (
	"fmt"

	"github.com/robmonte-org/org-private-printer/privateprinter"
)

const INCREMENT_ME int = 1

func Print() {
	privateprinter.PrivatePrint(fmt.Sprintf("dependabot test with access %d", INCREMENT_ME))
}
