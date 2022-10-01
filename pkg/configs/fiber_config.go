// ./pkg/configs/fiber_config.go
package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	readTimeOutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeOutSecondsCount),
	}

}
