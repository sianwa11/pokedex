# Pokedex CLI

A command-line Pokedex application built with Go that interacts with the [PokeAPI](https://pokeapi.co/).

## Features

- Explore location areas in the Pokemon world
- Catch Pokemon with a random chance of success
- View your caught Pokemon
- Navigate through different location areas
- Cache API responses for better performance

## Commands

- `help` - Display a list of available commands
- `map` - Display the names of 20 location areas
- `mapb` - Display the previous 20 location areas
- `explore {location-name}` - List Pokemon in a specific location area
- `catch {pokemon-name}` - Attempt to catch a Pokemon
- `pokedex` - View list of caught Pokemon
- `exit` - Exit the Pokedex

## Installation

1. Clone the repository:
```bash
git clone https://github.com/sianwa11/pokedex.git
```

2. Navigate to the project directory:
```bash
cd pokedex
```

3. Run the application:
```bash
go run .
```

## Usage

After starting the application, you'll be presented with a prompt:
```
Pokedex > 
```

You can enter any of the available commands. For example:
```
Pokedex > explore canalave-city-area
```

## Cache

The application implements caching to reduce API calls and improve performance. Cached data expires after 5 minutes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
