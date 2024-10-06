package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	godotenv.Load()

	b, err := bot.New(os.Getenv("BOT_TELEGRAM"))
	if err != nil {
		log.Println("An error occured while the bot initialization: ", err)
	}

	// miscellanious, remove this if you dont want to see it.
	info, err := b.GetMe(ctx)
	if err != nil {
		log.Println(err)
	}

	infoTotal := fmt.Sprintf("| [ID] => %d\n| [NAME] => %s\n| [USERNAME] => @%s", info.ID, info.FirstName, info.Username)
	fmt.Println(infoTotal)

	b.Start(ctx)
}
