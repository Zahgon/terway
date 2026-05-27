package link

// VethNameForPod return host-side veth name for pod
// max veth length is 15
func VethNameForPod(name, namespace, ifName, prefix string) (string, error) {
	_ = "STUB: not implemented"
	// A SHA1 is always 20 bytes long, and so is sufficient for generating the
	// veth name and mac addr.
	return "", nil
}
