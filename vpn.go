package main

import (
	"os/exec"
	"strings"
	"time"
	"encoding/json"
    "io/ioutil"
	"syscall"
 )

func main() {

	type Config struct {
		NetextenderPath string
		User string
		Pass string
		Server string
		Domain string
	}

	read_config, _ := ioutil.ReadFile("./vpnconfig.json")

	var config Config
    _ = json.Unmarshal(read_config, &config)

	vpn_config := [5]string{config.NetextenderPath, config.Server, config.Domain, config.User, config.Pass}

	for {
		vpnConnection(vpn_config[0], vpn_config[1], vpn_config[2], vpn_config[3], vpn_config[4])
		time.Sleep(15 * time.Second)
	}
}

func vpnConnection(netExtenderPath string, server string, domain string, user string, pass string) {
	cmd_status := exec.Command(netExtenderPath + "\\NECLI", "showstataus")
	cmd_status.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdout, _ := cmd_status.Output()

	vpn_status := strings.TrimRight(strings.Split(strings.Split(string(stdout), "\n")[0], " ")[1], "\r\n")

	if vpn_status != "Connected" {

		cmd_kill := exec.Command("taskkill", "/IM", "NEGui.exe", "/F")
		cmd_kill.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_, _ = cmd_kill.Output()

		cmd_connect := exec.Command(netExtenderPath + "\\NECLI", "connect", "-s", server, "-d", domain, "-u", user, "-p", pass)
		cmd_connect.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_, _ = cmd_connect.Output()
	}
}