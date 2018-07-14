// Author Boris HUISGEN <bhuisgen@hbis.fr>

package vaultapi

type Raw interface {
	RawGet(path string, data interface{}) error
	RawList(path string, data interface{}) error
	RawPost(path string, body string, data interface{}) error
	RawPut(path string, body string) error
	RawDelete(path string) error
}

func (c *client) RawGet(path string, data interface{}) error {
	return c.get(path, data)
}

func (c *client) RawList(path string, data interface{}) error {
	return c.list(path, data)
}

func (c *client) RawPost(path string, body string, data interface{}) error {
	return c.post(path, body, data)
}

func (c *client) RawPut(path string, body string) error {
	return c.put(path, body)
}

func (c *client) RawDelete(path string) error {
	return c.delete(path)
}
