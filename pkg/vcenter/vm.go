package vcenter

const (
	getVMpath       = "/vm/%s"
	getVMInterfaces = "/vm/%s/guest/networking/interfaces"
)

// GetVM returns the details of the given vmID, which can be found in the
// ListClusterVMs() results
func (vc *Vcenter) GetVM(vmID string) (VMDetail, error) {
	vm := &VMDetail{}
	err := vc.get(vm, getVMpath, vmID)
	if err != nil {
		vc.log.Error("could not get VM")
	}
	vm.VMinterfaces, _ = vc.GetInterfaces(vmID)
	return *vm, err
}

func (vc *Vcenter) GetInterfaces(vmID string) ([]VMinterface, error) {
	interfaces := make([]VMinterface, 0)
	err := vc.get(&interfaces, getVMInterfaces, vmID)
	return interfaces, err
}
