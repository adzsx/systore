package store

import (
	"os"

	"github.com/adzsx/etwas/internal/utils"
)

type Config struct {
}

func Setup() {
	_, err := os.Open(utils.ConfigPath + "systore.toml")

	if err != nil {
		utils.Verbose(1, "No config file detected. Creating "+utils.ConfigPath+"systore.toml")
		os.Mkdir(utils.ConfigPath, 0774)
		_, err := os.Create(utils.ConfigPath + "systore.toml")
		utils.Err(err)
		utils.Verbose(1, "Created file")
	}

}
