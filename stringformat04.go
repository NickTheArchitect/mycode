/* Alta3 Research | RZFeeser
   Using fmt.Printf - String formatting structs and PirateTreasure*/

package main

import (
    "fmt"
)

func main() {

    var uri string = "https://example.org:6001/v2/snacks"
    var k1 string = "req"
    var v1 string = "snickers"
    var k2 string = "quantity"
    var v2 string = "1"
    var k3 string = "size"
    var v3 string = "king"
    var url string = "%[1]s?%[2]s=%[3]s&%[4]s=%[5]s&%[6]s=%[7]s\n"

    fmt.Printf(url, uri, k1, v1, k2, v2, k3, v3)
}

