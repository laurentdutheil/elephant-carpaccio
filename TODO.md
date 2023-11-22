# TODO (aka Roadmap)

## First Version

Purpose: help the facilitator to attribute points to the teams.

- [x] display a qrcode to permit the facilitator to use their smartphone
- [x] use sse to update the board from server for score
- [x] propose some test data to the facilitator during demos
- [x] use sse to update the board from server for team registration

## Second Version

Purpose: each team have to build an api to respond to the requests of the game server.

- [ ] create an API for team registration (name + ip)
- [ ] create a scheduler to ping teams api
- [ ] handle response to give feedback to team (ok + error)
- [ ] add toggles to allow the facilitator to activate new test datasets
- [ ] add the notion of profit with a new graph