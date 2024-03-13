# TODO (aka Roadmap)

## First Version

Purpose: help the facilitator to attribute points to the teams.

- [x] display a qrcode to permit the facilitator to use their smartphone
- [x] use sse to update the board from server for score
- [x] propose some test data to the facilitator during demos
- [x] use sse to update the board from server for team registration
- [x] add notion of risk mitigation, cost of delay and business value to the score
- [ ] permit to generate some test data without decimal
- [ ] add a page to present backlog
- [ ] add pages to debrief of Story Points, Business Value, Risk Mitigation and Cost of Delay

## Second Version

Purpose: each team have to build an api to respond to the requests of the game server.

- [x] create an API for team registration (name + ip)
- [ ] create a scheduler to ping teams api
- [ ] handle response to give feedback to team (ok + error)
- [ ] add toggles to allow the facilitator to activate new test datasets
- [ ] add the notion of profit with a new graph