# Pragmateam code challenge server (GO)

## Framework & languages
This project uses
* Golang 1.17

### Run project

- `go run main.go` - Start the server (http://localhost:8081)

### Run tests

- `go test` - Run unit tests on terminal

### Notes

- About the main card topic "Improve test coverage", I would honestly say that I refactored more than I increased test coverage, but that was because in my opinion, more important than increasing the coverage of useless code was isolating valuable business code and then cover tests on it.
- Probably more tests could be created on structure in its current state, but most part are irrelevant at this moment. New tests has to be created as soon as new features that has real business impact has created.

### Things to be done

- Aimed at the business, probably would need to be created more types of alerts (e-mail, push notifications, etc.) or even some default mechanisms to be triggered when the system report temperature below or above the limits, after all, this system could be used for a sort of things (eg.: monitoring temperature of COVID vaccines being transported).
- We are assuming that the setup is correct, but if the setup was incorrect (eg.: min 4 and max 3) probably the application would need to have more validations and should break or alert in some way (Either the setup itself would need to be made by an interface with some control and validations).
- Inside `providers.Sensor`, being using the same type `models.Sensor` to act as a model entity and as a placeholder to desserialize json is not a good ideia. I would probably create a separated type to desserialize json in, and so map this type to the real `models.Sensor` model entity.
- All variable settings (eg.: Sensor provider API endpoint) would need to be put into a config file.
- A service layer could be created to run the flow that, today, is on the controller, making it cleaner.
- All error handling would need to be enhanced.
