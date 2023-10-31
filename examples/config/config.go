package config

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

// Params hold examples configuration.
type Params struct {
	IP             net.IP
	Port           uint16
	ServerCertPool *x509.CertPool
	ClientCert     tls.Certificate
}

func (p *Params) ReadEnv() error {
	ipEnv := os.Getenv("IP_ADDRESS")
	ip := net.ParseIP(ipEnv).To4()
	if ip == nil {
		return fmt.Errorf("bad IP_ADDRESS env value %q", ipEnv)
	}

	portEnv := os.Getenv("PORT")
	port, err := strconv.ParseUint(portEnv, 10, 16)
	if err != nil {
		return fmt.Errorf("bad PORT env value %q: %v", portEnv, err)
	}

	serverCertPemPathEnv := os.Getenv("SERVER_CERT_PEM_PATH")
	if len(serverCertPemPathEnv) == 0 {
		return fmt.Errorf(
			"bad SERVER_CERT_PEM_PATH env value %q", serverCertPemPathEnv)
	}

	serverCertPem, err := readfile(serverCertPemPathEnv)
	if err != nil {
		return fmt.Errorf("can't read server certificate pem file: %v", err)
	}
	serverCertPool := x509.NewCertPool()
	ok := serverCertPool.AppendCertsFromPEM(serverCertPem)
	if !ok {
		return fmt.Errorf("can't prepare server cert pool")
	}

	clientCertPemPathEnv := os.Getenv("CLIENT_CERT_PEM_PATH")
	if len(clientCertPemPathEnv) == 0 {
		return fmt.Errorf(
			"bad CLIENT_CERT_PEM_PATH env value %q", clientCertPemPathEnv)
	}
	clientKeyPemPathEnv := os.Getenv("CLIENT_KEY_PEM_PATH")
	if len(clientKeyPemPathEnv) == 0 {
		return fmt.Errorf(
			"bad CLIENT_KEY_PEM_PATH env value %q", clientKeyPemPathEnv)
	}

	clientCert, err := tls.LoadX509KeyPair(
		clientCertPemPathEnv, clientKeyPemPathEnv)
	if err != nil {
		return fmt.Errorf("load client key pair has failed: %v", err)
	}

	p.IP = ip
	p.Port = uint16(port)
	p.ServerCertPool = serverCertPool
	p.ClientCert = clientCert

	return nil
}

func readfile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

// ReadInput performs reading value specified by label,
func ReadInput(label string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	os.Stdout.WriteString(label)
	os.Stdout.WriteString(": ")
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	s = strings.TrimSpace(s)

	return s, nil
}
