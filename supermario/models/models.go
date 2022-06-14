/* Alta3 Research | RZFeeser@alta3.com
   Understanding Go Receiver Functions (i.e. methods)

   func(receiver_name Type) method_name(parameter_list)(return_type){
   // Code
   }

*/

package supermario

type Player struct {
    Lives int
    Stage int
    Inventory []string
}

// recv function to add a life
func (p *Player) Greenmushroom() {
    p.Lives++
}

// recv function to add an inventory item
func (p *Player) Pickup(powerup string) {
    p.Inventory = append(p.Inventory, powerup)
}

// rec function to check on the current stasge
func (p Player) CanWhistle() bool {
    return p.Stage >= 5
    }

