# Dashy - The ultimate dashboard for your office tv!
<br>

## Requirements
- Node == 16 (version)
- Go >= 19 (version)

## Get started
The server expects a file named `firebase-config.json` in the server folder, download from firebase or ask someone to give it to you.

This project uses a `makefile` to run most commands.

Run the following:
- `make install-dev-requirements` (downloads nodemon globally to run server in hot-reload)
- `make dev` (runs both server and frontend in watch mode)
- `make build` (builds a binary with the frontend included into the binary)

Other make commands:
The following two command can be run in seperate terminals to make it do the same as `make dev` does.
- `make watch-frontend` (runs the frontend in watch mode)
- `make watch-server` (runs the server in watch mode)