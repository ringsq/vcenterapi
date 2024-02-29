package vcenter

const (
	vmListPath = "/vm?clusters=%s"
)

// ListClusterVMs returns all VMs for the given clusterID.  The clusterID
// can be found in the results of ListDcClusters
func (vc *Vcenter) ListClusterVMs(clusterID string) (VMlistResult, error) {
	vms := &VMlistResult{}
	err := vc.get(vms, vmListPath, clusterID)
	if err != nil {
		vc.log.Error("could not list VMs")
	}
	return *vms, err
}
