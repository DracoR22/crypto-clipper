package startup

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func GetStartupPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
}

func CreateShortcut(targetPath, shortcutPath string) {
	script := `
	Set oWS = WScript.CreateObject("WScript.Shell")
	sLinkFile = "` + shortcutPath + `"
	Set oLink = oWS.CreateShortcut(sLinkFile)
	oLink.TargetPath = "` + targetPath + `"
	oLink.Save
	`
	vbsPath := filepath.Join(os.TempDir(), "createShortcut.vbs")
	err := os.WriteFile(vbsPath, []byte(script), 0644)
	if err != nil {
		panic(err)
	}
	defer os.Remove(vbsPath)
	cmd := exec.Command("cscript", vbsPath)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
