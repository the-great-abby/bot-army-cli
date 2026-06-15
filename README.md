# Bot Army CLI

A Go-based REPL interface for Bot Army personal OS. Connect to NATS and query bridge responders for briefings, context, and LLM chat.

## Features

- **Morning Briefing** — `brief` command queries `bridge.brief.today`
- **Context Query** — `context` command queries `bridge.context.current`
- **LLM Chat** — `chat <query>` command queries `bridge.chat`
- **NATS Integration** — Connects to NATS server for all operations

## Quick Start

### Local (against production NATS on port 4222)

```bash
make build
make run
```

### Docker (against Docker NATS)

```bash
make docker-run
```

## Environment

- `NATS_URL`: NATS server URL (default: `nats://localhost:4222`)

## Commands

```
brief             - Get today's briefing
context           - Show current context
chat <query>      - Chat with LLM
help              - Show this help
quit              - Exit
```

## Example Session

```
> brief
📋 Today's Briefing...
<briefing content>

> context
🎯 Current Context...
<context JSON>

> chat What are my priorities?
🤖 Chat: What are my priorities?
<LLM response>

> quit
Goodbye!
```

## Development

```bash
make test         # Test NATS connectivity
make clean        # Clean build artifacts
```

## License

MIT
