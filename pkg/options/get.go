package options

// GetOptions is the standard query options for getting a resource through the Calico API.
type GetOptions struct {
	// When specified:
	// - if unset, then the result is returned from remote storage based on quorum-read flag;
	// - if set to non zero, then the result is at least as fresh as given rv.
	// +optional
	ResourceVersion string
}