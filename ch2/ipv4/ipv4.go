package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var ipv4_matcher = regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)
var subaddr_matcher = regexp.MustCompile(`[0-9]{1,3}`)

type IPv4Addr struct {
	subaddr [4]int
}

func New(ip_str string) (*IPv4Addr, error) {
	if !(ipv4_matcher.MatchString(ip_str)) {
		return nil, errors.New("Invalid format for IP address.")
	}
	in_subaddr := subaddr_matcher.FindAllString(ip_str, 4)
	var addr IPv4Addr
	var err error
	for i, elem := range in_subaddr {
		addr.subaddr[i], err = strconv.Atoi(elem)
		if err != nil {
			return nil, err
		}
	}
	return &addr, nil
}

func (ip IPv4Addr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.subaddr[0], ip.subaddr[1], ip.subaddr[2], ip.subaddr[3])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter IPv4 address: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	ip_address, err := New(text[:len(text)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Your IP address is: ")
	fmt.Println(ip_address)
}
