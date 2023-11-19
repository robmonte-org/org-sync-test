package main

import (
	"fmt"

	"github.com/robmonte-org/org-printer-allowed/allowedprinter"
	"github.com/robmonte-org/org-printer-restricted/restrictedprinter"
)

const INCREMENT_ME int = 1

func Print() {
	allowedprinter.PrivatePrint(fmt.Sprintf("dependabot test with access %d", INCREMENT_ME))
	restrictedprinter.PrivatePrint(fmt.Sprintf("dependabot test with access %d", INCREMENT_ME))
}
