package startup

import (
	"html/template"
	"os"
	"os/user"
	"path/filepath"
)

func GetLaunchAgentsPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, "Library", "LaunchAgents")
}

func CreatePlist(exePath, plistPath string) {
	const plistTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>{{.Label}}</string>
	<key>ProgramArguments</key>
	<array>
		<string>{{.Path}}</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
</dict>
</plist>`
	tmpl, err := template.New("plist").Parse(plistTemplate)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(plistPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := struct {
		Label string
		Path  string
	}{
		Label: filepath.Base(exePath),
		Path:  exePath,
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
