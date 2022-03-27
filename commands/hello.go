package commands

import (
	"bytes"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"os/exec"
	"path/filepath"
)

func GetCommand() components.Command {
	return components.Command{
		Name:        "my-command",
		Description: "This is my-command's description",
		Aliases:     []string{"mc"},
		Arguments:   getArguments(),
		Flags:       getFlags(),
		EnvVars:     getEnvVar(),
		Action: func(c *components.Context) error {
			return helloCmd(c)
		},
	}
}

func getArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "shout",
			Description:  "Makes output uppercase.",
			DefaultValue: false,
		},
	}
}

func getEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

func helloCmd(c *components.Context) error {
	if len(c.Arguments) == 0 {
		message := "Hello :)"
		// You log messages using the following log levels.
		log.Output(message)
		log.Debug(message)
		log.Info(message)
		log.Warn(message)
		log.Error(message)
	}

	// We will add an API to get the resources directory. For now,
	resourcesPath := "resources"
	commandName := "test.sh"
	commandPath := filepath.Join(resourcesPath, commandName)

	// You can also send additional arguments to the `exec.Command` function
	cmd := exec.Command(commandPath)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	log.Output("The command output:")
	log.Info(string(stdout.Bytes()))
	log.Error(string(stderr.Bytes()))

	return err
}
