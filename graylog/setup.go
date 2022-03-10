package graylog

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var graylogreq string = `{
    "title": "Standard GELF UDP input",
    "type": "org.graylog2.inputs.gelf.udp.GELFUDPInput",
    "global": true,
    "configuration": {
        "recv_buffer_size": 1048576,
        "tcp_keepalive": false,
        "use_null_delimiter": true,
        "number_worker_threads": 2,
        "tls_client_auth_cert_file": "",
        "bind_address": "0.0.0.0",
        "tls_cert_file": "",
        "decompress_size_limit": 8388608,
        "port": 12201,
        "tls_key_file": "",
        "tls_enable": false,
        "tls_key_password": "",
        "max_message_size": 2097152,
        "tls_client_auth": "disabled",
        "override_source": null
    },
    "node": null
}`

func Setup(graylogHost string, retry int) error {
	req, err := http.NewRequest(
		"POST",
		graylogHost,
		bytes.NewBuffer([]byte(graylogreq)),
	)
	if err != nil {
		log.Fatal("Unable to create request for graylog")
	}
	req.Header.Set("X-Requested-By", "")
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	for retry != 0 {
		retry -= 1
		resp, _ := client.Do(req)
		if resp.StatusCode == 201 {
			return nil
		}
		time.Sleep(time.Second)
		fmt.Println("Graylog connection failed: ", retry)
	}
	return errors.New("graylog setup error")
}
