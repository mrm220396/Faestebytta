# Faestebytta Bot

Faestebytta Bot is a Telegram bot that provides various features for its users. This bot is developed using the Go programming language and utilizes the Telegram Bot API.

## Features

- **Kick Users**: Kick a user from the group chat by Replying them.
- **Ban Users**: Ban a user from the group chat by Replying them. 
- **Unban Users**: Unban a user from the group chat by Replying them. [To be implemented]
- **Pin Messages**: Pin a message in the group chat. [
- **Unpin Messages**: Unpin a message in the group chat.
- **Roll**: Roll a dice <Number>d<Number>
- **Echo**: Repeat in a quotation what you've said
## Getting Started

To use Faestebytta Bot, you need to have a Telegram account and join a group chat where the bot is added. Once you have joined the group chat, you can start using the available features by mentioning the bot in the chat.

## Installation

To install Faestebytta Bot, you need to have Go installed on your system. Once you have installed Go, you can run the following command to install the bot:

> go get -u https://github.com/faestebytta/faestebytta_bot


## Configuration

The bot requires a configuration file to function properly. You need to create a `config.json` file in the `settings` folder with the following structure:


Replace `YOUR_TELEGRAM_BOT_TOKEN` with your actual bot token obtained from the BotFather on Telegram.

## Usage

To run Faestebytta Bot, navigate to the project directory and run the following command:

```go
go run main.go
```


Once the bot is running, you can use the available features by mentioning the bot in the chat.

## License

This project is licensed under the BSD 3-Clause License. See the `LICENSE` file for details.


