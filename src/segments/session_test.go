package segments

import (
	"fmt"
	"testing"

	"github.com/jandedobbeleer/oh-my-posh/src/cache"
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime/mock"
	"github.com/jandedobbeleer/oh-my-posh/src/template"

	"github.com/stretchr/testify/assert"
)

func TestSessionSegmentTemplate(t *testing.T) {
	cases := []struct {
		Case           string
		ExpectedString string
		UserName       string
		ComputerName   string
		Template       string
		WhoAmI         string
		Platform       string
		SSHSession     bool
		Root           bool
	}{
		{
			Case:           "user and computer",
			ExpectedString: "john@company-laptop",
			ComputerName:   "company-laptop",
			UserName:       "john",
			Template:       "{{.UserName}}@{{.HostName}}",
		},
		{
			Case:           "user only",
			ExpectedString: "john",
			UserName:       "john",
			Template:       "{{.UserName}}",
		},
		{
			Case:           "user with ssh",
			ExpectedString: "john on remote",
			UserName:       "john",
			SSHSession:     true,
			ComputerName:   "remote",
			Template:       "{{.UserName}}{{if .SSHSession}} on {{.HostName}}{{end}}",
		},
		{
			Case:           "user without ssh",
			ExpectedString: "john",
			UserName:       "john",
			SSHSession:     false,
			ComputerName:   "remote",
			Template:       "{{.UserName}}{{if .SSHSession}} on {{.HostName}}{{end}}",
		},
		{
			Case:           "user with root and ssh",
			ExpectedString: "super john on remote",
			UserName:       "john",
			SSHSession:     true,
			ComputerName:   "remote",
			Root:           true,
			Template:       "{{if .Root}}super {{end}}{{.UserName}}{{if .SSHSession}} on {{.HostName}}{{end}}",
		},
		{
			Case:           "no template",
			ExpectedString: "",
			UserName:       "john",
			SSHSession:     true,
			ComputerName:   "remote",
			Root:           true,
		},
		{
			Case:           "user with ssh using who am i",
			ExpectedString: "john on remote",
			UserName:       "john",
			SSHSession:     false,
			WhoAmI:         "sascha   pts/1        2023-11-08 22:56 (89.246.1.1)",
			ComputerName:   "remote",
			Template:       "{{.UserName}}{{if .SSHSession}} on {{.HostName}}{{end}}",
		},
		{
			Case:           "user with ssh using who am i (windows)",
			ExpectedString: "john",
			UserName:       "john",
			SSHSession:     false,
			WhoAmI:         "sascha   pts/1        2023-11-08 22:56 (89.246.1.1)",
			Platform:       runtime.WINDOWS,
			ComputerName:   "remote",
			Template:       "{{.UserName}}{{if .SSHSession}} on {{.HostName}}{{end}}",
		},
	}

	for _, tc := range cases {
		env := new(mock.Environment)
		env.On("User").Return(tc.UserName)
		env.On("GOOS").Return("burp")
		env.On("Host").Return(tc.ComputerName, nil)

		var SSHSession string
		if tc.SSHSession {
			SSHSession = "zezzion"
		}

		env.On("Getenv", "SSH_CONNECTION").Return(SSHSession)
		env.On("Getenv", "SSH_CLIENT").Return(SSHSession)

		env.On("Platform").Return(tc.Platform)

		var whoAmIErr error
		if tc.WhoAmI == "" {
			whoAmIErr = fmt.Errorf("who am i error")
		}

		env.On("RunCommand", "who", []string{"am", "i"}).Return(tc.WhoAmI, whoAmIErr)

		session := &Session{}
		session.Init(properties.Map{}, env)

		template.Cache = &cache.Template{
			UserName: tc.UserName,
			HostName: tc.ComputerName,
			Root:     tc.Root,
		}

		_ = session.Enabled()
		assert.Equal(t, tc.ExpectedString, renderTemplate(env, tc.Template, session), tc.Case)
	}
}
