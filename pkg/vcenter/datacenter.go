package vcenter

const (
	listDCpath       = "/datacenter"
	listClustersPath = "/cluster?datacenters=%s"
)

// ListDatacenters returns a list of all datacenters managed by this vcenter
func (vc *Vcenter) ListDatacenters() (DatacenterResult, error) {
	dc := &DatacenterResult{}
	err := vc.get(dc, listDCpath)
	if err != nil {
		vc.log.Error("failed to retrieve datacenters")
	}
	return *dc, err
}

// ListDcClusters lists all clusters for the given datacenter ID
// The datacenter ID can be found in the results of `ListDatacenters()`
func (vc *Vcenter) ListDcClusters(datacenterID string) (ClusterListResult, error) {
	clr := &ClusterListResult{}
	err := vc.get(clr, listClustersPath, datacenterID)
	if err != nil {
		vc.log.Error("failed to list clusters")
	}
	return *clr, err
}
