> Stormkit Command Line Interface

## Installation

```bash
# Using npm
npm install @stormkit/cli

# Using yarn
yarn add @stormkit/cli

# Using pnpm
pnpm install @stormkit/cli
```

## Options

### env pull

```
stormkit env pull

Pull environment variables from the environment.

Options:
      --help     Show help                                             [boolean]
      --version  Show version number                                   [boolean]
  -d, --debug    Prints additional logs.
```

Usage:

1. Create an API Key from **Environments** > **Your environment** > **API Keys**. 
2. Create a `.env` file in your project root.
3. Add `SK_API_KEY=<your_api_key>` to the .env file

```bash
npx @stormkit/cli env pull
```

### api

```
stormkit api

Starts an API development server

Options:
      --help     Show help                                             [boolean]
      --version  Show version number                                   [boolean]
  -p, --port     Specify the port on which the API server should listen.
                                                               [default: "9090"]
  -d, --dir      Specify the directory in which the API routes are located. This
                 path is relative to project root.              [default: "api"]
```

Usage:


```json
{
  "scripts": {
    "dev:api": "stormkit api"
  }
}
```

```bash
npm run dev:api -- --port 9090 --dir src/api
```

## Testing locally

Run the `bin` command in the scripts. For example to test the e2e folder:

```
npm run bin -- api -d e2e
```

## License

MIT