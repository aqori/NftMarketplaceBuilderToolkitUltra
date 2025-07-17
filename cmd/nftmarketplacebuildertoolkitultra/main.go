// cmd/nftmarketplacebuildertoolkitultra/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacebuildertoolkitultra/internal/nftmarketplacebuildertoolkitultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacebuildertoolkitultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
