# PokedexCLI

A command-line interface (CLI) application to explore the Pokémon world. Built with Go and powered by the PokeAPI, PokedexCLI allows users to navigate Pokémon locations, catch Pokémon, and manage their own Pokedex, all from the terminal.

## Features

- Explore Pokémon locations and discover unique areas.
- Catch Pokémon and add them to your Pokedex.
- Inspect detailed stats for caught Pokémon.
- View a list of all Pokémon in your Pokedex.
- Simple, intuitive CLI commands for seamless interaction.

## Getting Started

### Prerequisites

- Go (version 1.23.1 or later)
- Internet connection (to fetch data from PokeAPI)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Fepozopo/pokedexcli.git
   ```
2. Navigate to the project directory:
   ```bash
   cd pokedexcli
   ```
3. Build the application:
   ```bash
   go build -o pokedexcli
   ```
4. Run the application:
   ```bash
   ./pokedexcli
   ```

## Usage

Upon starting, the CLI presents a prompt for entering commands:
```text
Pokedex >
```

### Available Commands

| Command              | Description                                                         |
|----------------------|---------------------------------------------------------------------|
| `help`              | Displays all available commands with a brief description.          |
| `map`               | Displays 20 location areas in the Pokémon world.                   |
| `mapb`              | Displays the previous 20 locations.                                |
| `explore <area>`    | Lists all Pokémon found in the specified area.                     |
| `catch <pokemon>`   | Attempts to catch the specified Pokémon.                           |
| `inspect <pokemon>` | Displays details of a Pokémon already in your Pokedex.             |
| `pokedex`           | Lists all caught Pokémon.                                          |
| `exit`              | Exits the application.                                             |

## Project Structure

- `main.go`: Entry point of the application.
- `commands.go`: Implements the logic for CLI commands.
- `internal/`: Contains additional modules and helpers for API interaction and caching.

---

Enjoy catching them all with PokedexCLI!
