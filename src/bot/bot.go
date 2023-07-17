package bot

import (
    "discord-bot/config"
    "log"
    "github.com/bwmarrin/discordgo"
)

var BotIDstring
var goBot *discordgo.Session

funcRun() {
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
funcmessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// Ignore all messages created by the bot itself
    if m.Author.ID == BotID {
        return
    }
// If the message is "Hi" reply with "Hi Back!!"
    ifm.Content == "Hi" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Hi Back")
    }
}
    }
}