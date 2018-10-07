package client

type DnsRecord struct {
	Address  string `json:"address,omitempty"`
	HostName string `json:"hostname,omitempty"`
}

// Create creates a dns record
func (c *Client) Create(address, hostname string) (*DnsRecord, error) {
	dns := &DnsRecord{
		Address:  address,
		HostName: hostname,
	}
	if err := c.doRequest(*dns); err != nil {
		return nil, err
	}
	return dns, nil
}

// Delete deletes a dns record
func (c *Client) Delete(address string) (*DnsRecord, error) {
	dns := &DnsRecord{
		Address:  address,
		HostName: "delete",
	}
	if err := c.doRequest(*dns); err != nil {
		return nil, err
	}
	return dns, nil
}
