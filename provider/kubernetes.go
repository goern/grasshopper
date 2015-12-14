/*
 Copyright 2015 Red Hat, Inc.

 This file is part of Grasshopper.

 Grasshopper is free software: you can redistribute it and/or modify
 it under the terms of the GNU Lesser General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 Grasshopper is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Lesser General Public License for more details.

 You should have received a copy of the GNU Lesser General Public License
 along with Grasshopper. If not, see <http://www.gnu.org/licenses/>.
*/

package provider

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	certFile    string = "/Users/goern/Gocode/src/github.com/goern/grasshopper/openshift-master.crt"
	keyFile     string = "/Users/goern/Gocode/src/github.com/goern/grasshopper/openshift-master.key"
	caFile      string = "/Users/goern/Gocode/src/github.com/goern/grasshopper/ca.crt"
	apiEndpoint string = "https://10.0.2.15:8443/api/v1/"
)

var ClientCertificate tls.Certificate
var caCertificate *tls.Certificate
var caCertificatePool *x509.CertPool
var tlsConfig *tls.Config

func Ping() {
	// Load client certifiacte
	ClientCertificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA certificate
	caCertificate, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertificatePool := x509.NewCertPool()
	caCertificatePool.AppendCertsFromPEM(caCertificate)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{ClientCertificate},
		RootCAs:      caCertificatePool,
	}
	tlsConfig.BuildNameToCertificate()

	// Setup HTTPS client
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))

}
