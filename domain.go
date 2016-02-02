package main

import (
	"bufio"
	"net"
	"strings"
)

type Domain struct {
	Name string
}

func NewDomain(name string) *Domain {
	return &Domain{
		Name: name,
	}
}

func (s *Domain) Exists() (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(s.Name + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}
	return true, nil
}
