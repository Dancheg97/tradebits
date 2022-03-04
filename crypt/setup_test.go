package crypt

import (
	"testing"
)

func getTestKey() string {
	return `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAL+lgvRKvBFhN8ZsfXJ9L8GOsnh/k1mhRZU0GtMBN0ODZfrZtawo
9oU/H2NHEhFTZUCIeMAfDMUnW+Mozy7cM2sCAwEAAQJAbhGZTMU/jpvPdN8KjZ7o
trpGNn7PGkNNu4vZfCbOnq7yc44lg9LH6ME+XJ5mil1BfULyf4wN+z6BUSkgc5Ed
AQIhAOPEk3p5N/4qNfLgxo7rAMyaWXf7G8XzGpnsRbcFYN4RAiEA12a/8IjoSyTV
3v1h252QYUv+5/EKXeUr9UaG6ZwBLbsCIBIHybT0S55vMP4dFRrTN6j6vwJkenpd
GEn/DjXC0zxBAiEAv6DpoVQrVK6wlnFVXcwPZn/1huMqFd6L/OmBnNLtOYECIDV9
sr1ru20w//oaYzsK98D7AdYBMadvOn9sIiSa6Zrh
-----END RSA PRIVATE KEY-----`
}

func TestSetup(t *testing.T) {
	rez := Setup(getTestKey())
	if rez != nil {
		t.Error(rez)
	}
}
