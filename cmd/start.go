/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Abhishek-Srivastava/reverse-proxy/internal/app"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the reverse proxy",
	Long: `Starts the reverse proxy with the arguments. For example:

reverse-proxy start --ipaddress 172.200.18.22 --port 443 --protocol https --proxyport 8080 --httptimeout 10 --certfile cert.crt --keyfile key.key --insecure false
reverse-proxy start --ipaddress 172.200.18.22`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var ip, serverport, proxyserverport, serverprotocol, cert, key string
		var timeout int64
		var insec bool
		ip, err = cmd.Flags().GetString(ipaddress)
		if err != nil {
			log.Fatal(err)
		}
		serverport, err = cmd.Flags().GetString(port)
		if err != nil {
			log.Fatal(err)
		}
		proxyserverport, err = cmd.Flags().GetString(proxyport)
		if err != nil {
			log.Fatal(err)
		}
		serverprotocol, err = cmd.Flags().GetString(protocol)
		if err != nil {
			log.Fatal(err)
		}
		cert, err = cmd.Flags().GetString(certfile)
		if err != nil {
			log.Fatal(err)
		}
		key, err = cmd.Flags().GetString(keyfile)
		if err != nil {
			log.Fatal(err)
		}
		timeout, err = cmd.Flags().GetInt64(httptimeout)
		if err != nil {
			log.Fatal(err)
		}
		insec, err = cmd.Flags().GetBool(insecure)
		if err != nil {
			log.Fatal(err)
		}
		rvProxy := app.New(ip, serverport, serverprotocol,
			proxyserverport, cert, key, timeout, insec)
		rvProxy.RunProxy()
	},
}

const (
	ipaddress   = "ipaddress"
	port        = "port"
	proxyport   = "proxyport"
	protocol    = "protocol"
	certfile    = "certfile"
	keyfile     = "keyfile"
	httptimeout = "httptimeout"
	insecure    = "insecure"
)

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().String(ipaddress, "", "ip of the remote server")
	if err := startCmd.MarkFlagRequired(ipaddress); err != nil {
		log.Fatal(err)
	}
	startCmd.Flags().String(port, "443", "port of the remote server")
	startCmd.Flags().String(protocol, "https", "protocol to connect http/https")
	startCmd.Flags().String(proxyport, "8080", "proxyport on which the revereproxy would be served")
	startCmd.Flags().String(certfile, "./certs/revpro.crt", "certificate file with path for tls")
	startCmd.Flags().String(keyfile, "./certs/revpro.key", "certificate key file with path for tls")
	startCmd.Flags().Int64(httptimeout, 10, "http timeout value in seconds")
	startCmd.Flags().Bool(insecure, false, "set true to disables https")

}
