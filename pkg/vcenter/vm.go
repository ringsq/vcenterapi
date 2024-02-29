package vcenter

const (
	getVMpath = "/vm/%s"
)

// GetVM returns the details of the given vmID, which can be found in the
// ListClusterVMs() results
func (vc *Vcenter) GetVM(vmID string) (VMDetail, error) {
	vm := &VMDetail{}
	err := vc.get(vm, getVMpath, vmID)
	if err != nil {
		vc.log.Error("could not get VM")
	}
	return *vm, err
}
