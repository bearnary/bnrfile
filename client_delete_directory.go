package bnrfile

func (c *defaultClient) ForceDeleteDirectory(name string) error {
	return ForceDeleteDirectory(name)
}
