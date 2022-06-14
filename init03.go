/* Alta3 Research | RZFeeser
   init - order of initialization  */

package main 

import "fmt"

func init() {
    WhatIsThe = 0
}

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
    return 42
}

func main() {
    if WhatIsThe == 0 {
        fmt.Println("The cake is a lie.")
    }
}

