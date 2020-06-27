package login

import (
	"errors"
	"os/user"
	"path/filepath"

	"github.com/hashload/boss/internal/pkg/configuration"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewLoginCommand(config *configuration.Configuration) *cobra.Command {
	var removeLogin bool

	var cmd = &cobra.Command{
		Use:   "login",
		Short: "Register login to repo",
		Run: func(cmd *cobra.Command, args []string) {
			if removeLogin {
				delete(config.Auth, args[0])
				config.SaveConfiguration()
				return
			}

			var auth *configuration.Auth
			var repo string
			if len(args) > 0 && args[0] != "" {
				repo = args[0]
				auth = config.Auth[args[0]]
			} else {
				repoPrompt := promptui.Prompt{
					Label: "Url to login (ex: github.com)",
					Validate: func(input string) error {
						if input == "" {
							return errors.New("Empty is not valid!!")
						}
						return nil
					},
				}
				repo, err := repoPrompt.Run()
				print(err)
				auth = config.Auth[repo]
			}

			if auth == nil {
				auth = &configuration.Auth{}
			}

			useSSHPrompt := promptui.Prompt{
				Label:     "Use SSH",
				IsConfirm: true,
			}

			option, err := useSSHPrompt.Run()
			print(option)
			auth.UseSsh = err == nil
			if auth.UseSsh {
				sshPathPrompt := promptui.Prompt{
					Label:     "Path of ssh private key",
					Default:   getSSHKeyPath(),
					AllowEdit: true,
				}

				sshPath, _ := sshPathPrompt.Run()
				auth.Path = sshPath
			} else {
				passPrompt := promptui.Prompt{
					Mask:  rune('*'),
					Label: "Password",
				}

				userPrompt := promptui.Prompt{
					Label: "Username",
				}

				user, _ := userPrompt.Run()
				pass, _ := passPrompt.Run()

				auth.SetUser(user)
				auth.SetPassword(pass)
			}
			config.Auth[repo] = auth
			config.SaveConfiguration()

		},
	}
	cmd.Flags().BoolVarP(&removeLogin, "rm", "", false, "remove login")

	return cmd
}

func getSSHKeyPath() string {
	usr, e := user.Current()
	if e != nil {
		log.Fatal(e)
	}
	return filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
}
