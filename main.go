/*
* CursusDB
* Shell
* ******************************************************************
* Copyright (C) 2023 CursusDB
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/peterh/liner"
	"golang.org/x/term"
	"net"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// Curush is the main CursusDB Shell type
type Curush struct {
	TLS         bool
	ClusterHost string
	ClusterPort int
}

var (
	history_fn = filepath.Join(os.TempDir(), ".query_history")
)

func main() {
	var curush Curush

	err := curush.RunShell()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

// RunShell starts CursusDB Shell functionality.  Reading cluster host, port and tls, authenticating, then listing to the user and transmitting the received on shell to the cluster.
func (curush *Curush) RunShell() error {
	flag.BoolVar(&curush.TLS, "tls", false, "Use secure connection.")
	flag.StringVar(&curush.ClusterHost, "host", "", "Cluster host.")
	flag.IntVar(&curush.ClusterPort, "port", 7681, "Cluster host port.")
	flag.Parse()

	if curush.ClusterHost == "" {
		errMsg := "CursusDB cluster host required."
		return errors.New(errMsg)
	}

	if curush.ClusterPort == 0 {
		errMsg := "CursusDB cluster host port required."
		return errors.New(errMsg)
	}

	if !curush.TLS {

		fmt.Print("Username>")
		username, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}
		fmt.Print(strings.Repeat("*", utf8.RuneCountInString(string(username))))

		fmt.Println("")
		fmt.Print("Password>")
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}
		fmt.Print(strings.Repeat("*", utf8.RuneCountInString(string(password))))

		tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", curush.ClusterHost, curush.ClusterPort))
		if err != nil {
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		defer conn.Close()

		text := textproto.NewConn(conn)
		// Authenticate
		err = text.PrintfLine(fmt.Sprintf("Authentication: %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\\0%s", username, password)))))
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		read, err := text.ReadLine()
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		if strings.HasPrefix(read, fmt.Sprintf("%d ", 0)) {

			query := ""

			line := liner.NewLiner()
			defer line.Close()

			line.SetCtrlCAborts(true)

			if f, err := os.Open(history_fn); err == nil {
				line.ReadHistory(f)
				f.Close()
			}
			fmt.Println("")
			for {
				if in, err := line.Prompt("curush>"); err == nil {
					query += in

					query = strings.Join(strings.Split(query, " "), " ")

					if strings.HasSuffix(query, ";") {
						line.AppendHistory(query)
						_, err = conn.Write([]byte(strings.TrimSpace(query) + "\r\n"))
						if err != nil {
							fmt.Println("")
							errMsg := err.Error()
							return errors.New(errMsg)
						}

						read, err := text.ReadLine()
						if err != nil {
							fmt.Println("")
							errMsg := err.Error()
							return errors.New(errMsg)
						}
						fmt.Println(read)
						query = ""
					}

				} else if err == liner.ErrPromptAborted {
					fmt.Println("")
					fmt.Println("Aborted")
					break
				} else {
					fmt.Println("")
					fmt.Println("Error reading line: ", err)
					break
				}
			}
		} else {
			fmt.Println("")
			fmt.Println("Invalid credentials.")
		}

	} else {

		fmt.Print("Username>")
		username, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}
		fmt.Print(strings.Repeat("*", utf8.RuneCountInString(string(username))))
		fmt.Println("")
		fmt.Print("Password>")
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}
		fmt.Print(strings.Repeat("*", utf8.RuneCountInString(string(password))))

		config := tls.Config{ServerName: curush.ClusterHost}

		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", curush.ClusterHost, curush.ClusterPort), &config)
		if err != nil {
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		defer conn.Close()

		text := textproto.NewConn(conn)
		// Authenticate
		err = text.PrintfLine(fmt.Sprintf("Authentication: %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\\0%s", username, password)))))
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		read, err := text.ReadLine()
		if err != nil {
			fmt.Println("")
			errMsg := err.Error()
			return errors.New(errMsg)
		}

		if strings.HasPrefix(read, fmt.Sprintf("%d ", 0)) {

			query := ""

			line := liner.NewLiner()
			defer line.Close()

			line.SetCtrlCAborts(true)

			if f, err := os.Open(history_fn); err == nil {
				line.ReadHistory(f)
				f.Close()
			}
			fmt.Println("")
			for {
				if in, err := line.Prompt("curush>"); err == nil {
					query += in

					query = strings.Join(strings.Split(query, " "), " ")

					if strings.HasSuffix(query, ";") {
						line.AppendHistory(query)
						_, err = conn.Write([]byte(strings.TrimSpace(query) + "\r\n"))
						//err = text.PrintfLine(query) // Because of % we should not use printf
						if err != nil {
							fmt.Println("")
							errMsg := err.Error()
							return errors.New(errMsg)
						}

						read, err := text.ReadLine()
						if err != nil {
							fmt.Println("")
							errMsg := err.Error()
							return errors.New(errMsg)
						}
						fmt.Println(read)
						query = ""
					}

				} else if err == liner.ErrPromptAborted {
					fmt.Println("")
					fmt.Println("Aborted")
					break
				} else {
					fmt.Println("")
					fmt.Println("Error reading line: ", err)
					break
				}
			}
		} else {
			fmt.Println("")
			fmt.Println("Invalid credentials.")
		}

	}

	return nil
}
