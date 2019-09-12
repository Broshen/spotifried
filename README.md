TODO:

- ~~handle rate limiting for requests~~

- ~~logging all activity~~

- handling errors gracefully

- ~~configure sensitive info as environment variables & exclude from git~~

- prettify song comparison response JSON

	- show added_at fields for both users

	- move array into an object

	- add fields for both user names

- analyze playlists

- analyze albums

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
