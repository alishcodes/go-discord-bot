### Alice-Bot

Alice is an interactive Discord bot created with DiscordGo and the Ollama dolphin3 model. Alice is the mischievous alter-ego of her creator.

#### Features

- Typing simulation: Displays "Typing..." indicator on discord while processing responses.
- AI powered: Integrates a backend API to generate interactive, evil, witty and sometimes playful replies.
- Personalized interaction: Alice can reference usernames in her responses, making interactions feel more engaging and lively.
- Configurable: API endpoint, bot token, bot trigger word and AI model settings stored in an `.env` file.

#### Configuration

| Variable | Description |
| -------  |   -------   |
| API_ENDPOINT | the API endpoint for your AI model. |
| DISCORD_BOT_TOKEN | the bot token provided by Discord. |
| AI_MODEL | the name of your AI model (ex: "alice:latest") |
| BOT_TRIGGER | the phrase to trigger your bot when typed (ex: "alice")|


