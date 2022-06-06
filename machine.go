package activation_code

import "github.com/super-l/machine-code/machine"

func GetMachine() *machine.MachineData {
	data := machine.GetMachineData()
	return &data
}
