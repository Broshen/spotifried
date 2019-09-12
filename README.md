TODO:

- ~~handle rate limiting for requests~~

- ~~logging all activity~~

- handling errors gracefully

- ~~configure sensitive info as environment variables & exclude from git~~

- prettify song comparison response JSON

	- show added_at fields for both users

	- move array into an object

	- add fields for both user names

- analyze playlists & albums

	- allow users to import:

		- all songs from playlists they created (regardless of whether or not the songs are saved to their library)

		- all songs from playlists they've saved (can be created by others)

		- all songs from all albums they've saved, regardless of whether or not the songs in their albums are also saved to their library

	- add user preference fields to keep track of what their full "library" consists of

	- add a popup? to the reload buttons for the user to SET their import preferences

		- initial "Lets do it!" button will still only fetch library & top data

		- /fetch endpoint will also take parameters on whether or not to fetch playlist/album data

		- getAllUserData function takes in such parameters

			- make sure when getAllUserData saves, it does not override user preferences to defaults


- ~~analyze artists~~

    - ~~get genres of all artists~~

- analyze trends over time

- ~~analyze TOP (not ALL) tracks/artists - /top endpoint~~

	- ~~common TOP track/artists~~

- add suggestions - get songs by artists common to both users that only one user has saved

- ~~caching - add an intermediate caching layer so that we don't always SPAM spotify with requests~~

- add tests

- search functions

- clean up & parallelize go code

- ~~sort common artists by # of songs in common~~

Deploying:
1. Build the backend golang binary `go build`
2. Build the react frontend `cd frontend && npm run build`
3. Commit the changes
4. Push to github, and merge to master
5. Go to heroku, and deploy

Local workflow:
1. setup required API keys & database URLs - obtain from respective sources (e.g. Spotify API keys, etc), and either save as environment variables, or as golang variables - see `initialize.go` for a list of required resources
2. have two commmand line tabs open - 1 to run the backend in the top level dir, and 1 to run the frontend in the /frontend dir
3. in the first tab, run `go build && ./spotifried` to build & start the backend. rerun this every time you make changes to the backend code that you want reflected in your local development environment
4. in the second tab, run `npm start` to start the react frontend. changes made will be automatically reflected
