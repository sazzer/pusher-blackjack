import React from 'react';
import { Grid, Container, List } from 'semantic-ui-react';
import Cards from './Cards';
import axios from 'axios';
import Pusher from 'pusher-js';

var pusher = new Pusher('PUSHER_KEY', {
    cluster: 'PUSHER_CLUSTER',
    forceTLS: true
  });

export default class Game extends React.Component {
    state = {
        croupier: [],
        seats: [],
        gameMode: "",
        turn: 0
    };

    componentDidMount() {
        var channel = pusher.subscribe('table-' + this.props.activeTable);
        channel.bind('update', () => {
            this._fetchGame();
        });
        this._fetchGame();
    }

    render() {
        const userSatAtTable = this.state.seats.find(player => player.name === this.props.currentUser.id);

        const players = this.state.seats.map((player, index) => {
            if (player.name) {
                let leave;
                let bet;
                let actions;
                if (player.name === this.props.currentUser.id) {
                    if (this.state.gameMode === 'betting' || player.bet === 0) {
                        leave = <List.Item><button onClick={() => this._leaveTable()}>Leave Table</button></List.Item>
                    }

                    if (player.bet === 0 && this.state.gameMode === 'betting') {
                        bet = (
                            <List.Item>
                                <button onClick={() => this._placeBet()}>Place Bet</button>
                            </List.Item>
                        );
                    } else {
                        bet = (
                            <List.Item>
                                Bet: {player.bet}
                            </List.Item>
                        );
                        if (player.bet > 0 && this.state.turn === index) {
                            actions = [
                                <List.Item><button onClick={() => this._action('hit')}>Hit</button></List.Item>,
                                <List.Item><button onClick={() => this._action('stick')}>Stick</button></List.Item>,
                            ];
                        }
                    }
                }
                return (
                    <Grid.Column textAlign="center" key={`seat-${index}`}>
                        <List>
                            <List.Item>{player.name}</List.Item>
                            <List.Item><Cards cards={player.cards || []} /></List.Item>
                            <List.Item>Current Stack: {player.stack}</List.Item>
                            { bet }
                            { actions }
                            { leave }
                        </List>
                    </Grid.Column>
                );
            } else {
                return (
                    <Grid.Column textAlign="center" key={`seat-${index}`}>
                        <i>Empty Seat</i>
                        { !userSatAtTable && <button onClick={() => this._joinTable(index)}>Sit At Table</button> }
                    </Grid.Column>
                );
            }
        });


        return (
            <Container>
                <Grid>
                    <Grid.Row columns="1">
                        <Grid.Column textAlign="center">
                            Croupier
                            <Cards cards={this.state.croupier} horizontal/>
                        </Grid.Column>
                    </Grid.Row>
                    <Grid.Row columns={this.state.seats.length}>
                        {players}
                    </Grid.Row>
                </Grid>
            </Container>
        );
    }

    _fetchGame() {
        axios({
            baseURL: 'http://localhost:8080',
            url: `/games/${this.props.activeTable}`,
        }).then((response) => {
            this.setState({
                gameMode: response.data.state,
                croupier: response.data.croupier || [],
                turn: response.data.turn,
                seats: response.data.players.map(player => {
                    if (player.id) {
                        return {
                            name: player.id,
                            bet: player.bet,
                            stack: player.stack,
                            cards: player.cards
                        };
                    } else {
                        return {};
                    }
                })
            });
        });
    }

    _joinTable(seat) {
        const formData = new FormData();
        formData.set('seat', seat);

        axios({
            baseURL: 'http://localhost:8080',
            url: `/games/${this.props.activeTable}/${this.props.currentUser.id}`,
            method: 'PUT',
            data: formData,
        });
    }

    _leaveTable() {
        axios({
            baseURL: 'http://localhost:8080',
            url: `/games/${this.props.activeTable}/${this.props.currentUser.id}`,
            method: 'DELETE',
        });
    }

    _placeBet() {
        const bet = prompt('How much to bet?');

        if (bet) {
            const formData = new FormData();
            formData.set('amount', bet);

            axios({
                baseURL: 'http://localhost:8080',
                url: `/games/${this.props.activeTable}/${this.props.currentUser.id}/bet`,
                method: 'PUT',
                data: formData,
            });
        }
    }

    _action(action) {
        const formData = new FormData();
        formData.set('action', action);

        axios({
            baseURL: 'http://localhost:8080',
            url: `/games/${this.props.activeTable}/${this.props.currentUser.id}/action`,
            method: 'POST',
            data: formData,
        });
    }
}
