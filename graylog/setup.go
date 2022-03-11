package graylog

import (
	"bytes"
	"errors"
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

func checkInput(graylogHost string) error {
	req, err := http.NewRequest(
		"GET",
		graylogHost,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("X-Requested-By", "go_tradebits")
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	inputs, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	inpstr := string(inputs)
	if !strings.Contains(inpstr, "Standard GELF UDP input") {
		return errors.New("input not found")
	}
	return nil
}

func setInput(graylogHost string) error {
	err := checkInput(graylogHost)
	if err == nil {
		return nil
	}
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
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 201 {
		return nil
	}
	return errors.New("bad response code")
}

func Setup(graylogHost string) error {
	for i := 0; i < 90; i++ {
		err := setInput(graylogHost)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		return nil
	}
	return errors.New("number of retrys exceeded")
}
