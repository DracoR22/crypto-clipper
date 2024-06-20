package startup

import (
	"html/template"
	"os"
	"os/user"
	"path/filepath"
)

func GetAutostartPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".config", "autostart")
}

func CreateDesktopFile(exePath, desktopFilePath string) {
	const desktopFileTemplate = `[Desktop Entry]
Type=Application
Exec={{.Path}}
Hidden=false
NoDisplay=false
X-GNOME-Autostart-enabled=true
Name={{.Name}}
Comment=Start {{.Name}} at login
`
	tmpl, err := template.New("desktopFile").Parse(desktopFileTemplate)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(desktopFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := struct {
		Name string
		Path string
	}{
		Name: filepath.Base(exePath),
		Path: exePath,
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
