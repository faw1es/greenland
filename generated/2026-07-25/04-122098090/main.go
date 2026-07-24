package main

import "fmt"

const generatedID = "30118198758-04"

func calculateTip(tagihan float64) (float64, float64) {
    var tipRate float64

    if tagihan >= 50 && tagihan <= 300 {
        tipRate = 0.15
    } else {
        tipRate = 0.20
    }

    tip := tagihan * tipRate
    total := tagihan + tip

    return tip, total
}

func main() {
    tagihan := 200.0

    tip, total := calculateTip(tagihan)

    fmt.Printf(
        "ID: %s\nTagihan: Rp. %.2f\nTip: Rp. %.2f\nTotal: Rp. %.2f\n",
        generatedID,
        tagihan,
        tip,
        total,
    )
}
