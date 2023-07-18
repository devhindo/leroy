package bot

import (
	"fmt"
	"leroy/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string = "1129975760430563358"
var goBot *discordgo.Session

func Run() {
// create bot session
    goBot, err := discordgo.New("Bot " + config.Token)
    if err != nil {
        log.Fatal(err)
        return
    }
    // make the bot a user
    user, err := goBot.User("@me")
    if err != nil {
        log.Fatal(err)
        return
    }
    BotID = user.ID
    goBot.AddHandler(messageHandler)
    err = goBot.Open()
    if err != nil {
        return
    }
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// Ignore all messages created by the bot itself
    if m.Author.ID == BotID {
        return
    }
    // print bot permissions


    // If the message is "Hi" reply with "Hi Back!!"
    if m.Content == "Hi" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Hi Back")

    } else {
        fmt.Println(m.Message.Content)
        _, _ = s.ChannelMessageSend(m.ChannelID, "Hi Back")

    }
}
