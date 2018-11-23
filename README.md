# Creating a realtime Blackjack game

This is an example of using Pusher Channels and Pusher ChatKit to create a real-time BlackJack game.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

In order to run this project, ensure that the following software is all installed correctly:

* [Go](https://golang.org/) - at least version 1.11.
* [Node.js](https://nodejs.org/en/)
* [Create React App](https://github.com/facebook/create-react-app)

Also ensure that this project is checked out in an appropriate place under the $GOPATH.

## Running the examples

Before the example can be run, [Pusher Credentials](https://dashboard.pusher.com/) and [Pusher ChatKit Credentials](https://dash.pusher.com/) will need to be obtained by registering a new Application, and populating the credentials into the appropriate places.

### Running the Backend Service

Ensure that Go are installed and set up on your machine. Download the necessary dependencies by executing `go get`, and then run the backend by running `go run blackjack.go`.

### Running the Pundit UI

Ensure that Node.js is installed on your machine. From the `webapp` directory execute `npm install` to download the depdendencies and then `npm start` to run the application.

## Built With

* [Pusher](https://pusher.com/) - APIs to enable devs building realtime features
* [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Create React App](https://github.com/facebook/create-react-app) - Create React apps with no build configuration.
* [Semantic UI](https://react.semantic-ui.com/introduction) - User Interface is the language of the web


