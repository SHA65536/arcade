# The Arcade
This is a small collections of terminal games served over SSH. Mostly a learning project for me as I want to get better at making bigger programs.

## How it works (Work in progress)
Users connect to the server using SSH can either play as guests, or register and login with a username and password.

Playing as a guest is the same as being registered except you won't be able to put your scores on the leaderboards and have saved games.

After logging in, the user will see a menu listing all the available games where he can navigate to a game of his choice.

## Packages
The `server` package handles connections, authentication, and the main menu. It uses [Wish](https://github.com/charmbracelet/wish) to handle SSH connections and display using [BubbleTea](https://github.com/charmbracelet/bubbletea).

The `database` package handles interfacing with the database for saving leaderboard scores / user data / saved games.

The `games` package handles all the games, it has a game interface that each game implements.

Server setup and command line interface is in `main.go`.

## License
[GPL-3.0](https://raw.githubusercontent.com/SHA65536/Arcade/main/LICENSE)