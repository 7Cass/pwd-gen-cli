package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name: "rnd",
    Usage: "Generate a random 9 characters string",
    Args: true,
    ArgsUsage: "Lenght of the random string",
    Action: func(ctx *cli.Context) error {
      bytes := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@.?");
      length, err := strconv.Atoi(ctx.Args().Get(1))

      if length > 100 {
        log.Fatal("Lenght must be lower than 100")
      }

      if err != nil {
        length = 9
      }

      pwd := make([]byte, length)
      for i := range pwd {
        pwd[i] = bytes[rand.Intn(len(bytes)-1)]
      }
      fmt.Println(string(pwd))
      return nil
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
