package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sftp"
	"strings"

	"golang.org/x/crypto/ssh"
)

const (
	user     = "xx"
	ip_port  = "x.x.x.x:22"
	password = "xxx"
)

func main() {

	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd, HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }}
	Client, err := ssh.Dial("tcp", ip_port, &Conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer Client.Close()
	/*
		if session, err := Client.NewSession(); err == nil {

			session.Stdout = os.Stdout
			if err := session.Run("ls /; ls /DATA/applog/main_dev_services_log"); err != nil {
				if err != nil {
					if err.Error() != "Process exited with: 1. Reason was:  ()" {
						fmt.Println(err.Error())
					}
				}
			}
		}
	*/
	c, err := sftp.NewClient(Client, sftp.MaxPacket(6e9))
	if err != nil {
		log.Fatalf("unable to start sftp subsytem: %v", err)
	}
	defer c.Close()

	path := "/DATA/applog/main_dev_services_log/server.log"
	localpath := "/home/bush/Downloads"
	fs, e := c.Open(path)
	if e != nil {
		log.Println(e)
		os.Exit(-1)
	}
	filename := filepath.Base(path)
	info, _ := fs.Stat()
	File, err := os.Create(fmt.Sprintf(`%s/%s`, strings.TrimRight(localpath, `/`), filename))
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println("保存路径:", fmt.Sprintf(`%s/%s`, strings.TrimRight(localpath, `/`), filename))
	defer File.Close()
	io.Copy(File, io.LimitReader(fs, info.Size()))
}
