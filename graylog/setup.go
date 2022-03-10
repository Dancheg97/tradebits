package graylog

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func CheckInput(graylogHost string) (bool, error) {
	req, err := http.NewRequest(
		"GET",
		graylogHost,
		nil,
	)
	if err != nil {
		return false, err
	}
	req.Header.Set("X-Requested-By", "go_tradebits")
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	inputs, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	inpstr := string(inputs)
	if !strings.Contains(inpstr, "Standard GELF UDP input") {
		return false, nil
	}
	return true, nil
}

func Setup(graylogHost string, retry int) error {
	req, err := http.NewRequest(
		"POST",
		graylogHost,
		bytes.NewBuffer([]byte(graylogreq)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("X-Requested-By", "go_tradebits")
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	for retry != 0 {
		retry -= 1
		exist, _ := CheckInput(graylogHost)
		if exist {
			return nil
		}
		resp, _ := client.Do(req)
		if resp.StatusCode == 201 {
			return nil
		}
		time.Sleep(time.Second)
		fmt.Println("Graylog connection failed: ", retry)
	}
	return errors.New("graylog setup error")
}
