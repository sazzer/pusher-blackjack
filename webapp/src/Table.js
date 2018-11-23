import React from 'react';
import { Grid, List, Header, Label, Form, Input } from 'semantic-ui-react';
import Game from './Game';

export default class Table extends React.Component {
    state = {
        users: [],
        messages: [],
        newMessage: "",
    };

    constructor(props) {
        super(props);

        props.currentUser.subscribeToRoom({
            roomId: props.activeTable,
            messageLimit: 100,
            hooks: {
                onUserJoined: () => {
                    this._updateUsers();
                },
                onUserLeft: () => {
                    this._updateUsers();
                },
                onNewMessage: (message) => {
                    const messages = this.state.messages;
                    console.log(message);
                    messages.unshift({
                        id: message.id,
                        user: message.senderId,
                        message: message.text
                    });
                    this.setState({
                        messages: messages
                    });
                }
            }
        }).then(() => this._updateUsers());
    }
    _updateUsers() {
        const currentRoom = this.props.currentUser.rooms.find(room => room.id === this.props.activeTable);
        this.setState({
            users: currentRoom.users
        });
    }
    _handleNewMessageChange(e) {
        this.setState({
            newMessage: e.target.value
        });
    }
    _handleSubmit() {
        const { newMessage } = this.state;
        const { currentUser, activeTable } = this.props;
        currentUser.sendMessage({
            text: newMessage,
            roomId: activeTable
        });
        this.setState({
            newMessage: ''
        });
    }

    render() {
        const users = this.state.users.map((user) => (
            <List.Item key={user.id}>
                <List.Content>
                    {user.name}
                </List.Content>
            </List.Item>
        ));

        const messages = this.state.messages
            .map((message) => {
                const user = this.state.users.find(user => user.id === message.user) || {};

                return (
                    <List.Item key={message.id}>
                        <List.Content>
                            <Label ribbon>{ user.name || message.user }</Label>
                            { message.message }
                        </List.Content>
                    </List.Item>
                );
            });

        return (
            <Grid>
                <Grid.Row>
                    <Grid.Column width="12">
                        <Game activeTable={this.props.activeTable} currentUser={this.props.currentUser} />
                    </Grid.Column>
                    <Grid.Column width="4">
                        <Header>Players</Header>
                        <List>{users}</List>
                    </Grid.Column>
                </Grid.Row>
                <Grid.Row>
                    <Grid.Column width="16">
                        <List style={{height: '20em', overflow: 'auto'}}>
                            { messages }
                        </List>
                    </Grid.Column>
                </Grid.Row>
                <Grid.Row>
                    <Grid.Column width={16}>
                        <Form onSubmit={this._handleSubmit.bind(this)}>
                            <Input action='Post'
                                   placeholder='New Message...'
                                   value={this.state.newMessage}
                                   fluid
                                   autoFocus
                                   onChange={this._handleNewMessageChange.bind(this)} />
                        </Form>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        );
    }
}
