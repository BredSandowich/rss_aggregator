# Gator rss-feed-aggregator

A command-line interface (CLI) RSS feed aggregator built for the boot.dev curriculum on backend development. 
Written in Go to register users, subscribe and track subscriptions, and fetch the latest posts.

## Prerequisites

- Go (version 1.25.3 or higher)
- PostgreSQL (configured and running)

## Installation 
   1. **Install the `gator` CLI**:
   Use the following command to download and install the binary
   ```sh
   go install github.com/BredSandowich/rss_aggregator@latest
   ```
   This will download and install the binary directly to the Go bin path.

## Configuration

Create a .gatorconfig.json file in your home directory with the following structure:

```json
{
  "db_url": "your_database_connection_string",
  "current_user_name": "default_user"
}
```
## Usage
Once installed and configured, you can use the `gator` CLI to interact with the application. Below are some of the available commands:
### Available Commands

- `gator login <username>`: Log in as an existing user.
- `gator register <username>`: Register a new user.
- `gator reset`: Delete all records and reset tables.
- `gator users`: List all registered users.
- `gator agg <time_between_reqs>`: Continuously fetch RSS feeds (e.g., `1m`, `30s`).
- `gator addfeed <name> <url>`: Add a new feed and follow it.
- `gator feeds`: List all feeds in the database.
- `gator follow <url>`: Follow an existing feed.
- `gator following`: List feeds followed by current user.
- `gator unfollow <url>`: Unfollow a feed.
- `gator browse [limit]`: View recent posts (defaults to 2).



License
This project is part of the Boot.dev curriculum and is for educational purposes.