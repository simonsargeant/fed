package load

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Items(args ...string) map[string]int {
	items := make(map[string]int)
	for _, arg := range args {
		str := strings.Split(arg, ":")

		i, err := strconv.Atoi(str[1])
		if err != nil {
			log.Fatalf("Error parsing item args: invalid quantity from %s: %s", arg, err)
		}

		items[str[0]] = i
	}

	return items
}
