package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadKeyBytes(t *testing.T) {
	createPemFile()
	rez := readKeyBytes("test.pem")
	if len(rez) < 5 {
		t.Error("file have not been read properly")
	}
	os.RemoveAll("test.pem")
}

func TestReadConnectionFile(t *testing.T) {
	createConnectionFile()
	rez := readConnecitonAdress("test.cfg")
	if len(rez) < 5 {
		t.Error("file have not been read properly")
	}
	os.RemoveAll("test.cfg")
}

func createPemFile() {
	contents := `-----BEGIN RSA PUBLIC KEY-----
MIICCgKCAgEAjdgFBLDvmGQvl3o87UZAxUCuGwKqym92Iv4FGujlcr95oPxkIoXt
LYnZUNXBAPdDRqeag8dhTbftCuywd99xJybV4uoqlOSghd2lu5nrtixrcIF/x3KH
+5brTzLjGaTV5WsK9hLXgnEBdPDJ6FadCo1Se8/6D5oAvKQYoqNNgZqiJBYsf9rO
0bPQqwqSaghWcfGVO/tbJY0OUfhOlhqokeGsvFa3c9g5Ca0kHvpY8u/yF33xiCl7
yadza68yxHu5X5/ciqEOldzsyK5SauyVANGHeIio5X9yVHKiEUeNLkRh3Zd1QLZv
LT6iHyBbhFzN+aI7nWUZx8mqTQPcOoJdUgXpwzsbq2ODgqZf6vLG+vonoW8A+Irt
pyP3t/SdgzwtTt9j8ssK08ma4K+mC5vHdyNt2zd0bT05ZUl2X0hGtcGgkusBmfXO
gV0CD6KynI7E25kNSFm5puCxPiZ5tisyGdrfgBM5lK8tWl5AZPBbPmNNnriUKOIr
C72EsXnxV5F3ZX626t1YUKMpfmkr9H4yicSjHYtZDVInF1VeNOW3eMUUgiw0kyEZ
ZZLQigUd1IuLkKEvoW9R1gtXKXy9e3xv8v0uhvdLL7UCEA8bDCAmzJ8fOTQLIfxe
Ad3bM7JcmnzdglKgGTQU16tIzhizQHOKGM3gNf2W1f3TJyGU8lseA50CAwEAAQ==
-----END RSA PUBLIC KEY-----`
	ioutil.WriteFile("test.pem", []byte(contents), 0644)
}

func createConnectionFile() {
	contents := `120.129.90.21`
	ioutil.WriteFile("test.cfg", []byte(contents), 0644)
}
