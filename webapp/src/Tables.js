import React from 'react';
import { Container, Grid, Header, Segment} from 'semantic-ui-react';
import { TokenProvider, ChatManager } from '@pusher/chatkit';
import TableList from './TableList';
import Table from './Table';

export default class Tables extends React.Component {

  state = {
    joined: [],
    joinable: [],
  };

  _onJoinTable = this._handleJoinTable.bind(this);

  constructor(props) {
    super(props);
    this.chatManager = new ChatManager({
        instanceLocator: 'CHATKIT_INSTANCE_LOCATOR',
        tokenProvider: new TokenProvider({
            url: "http://localhost:8080/chatkit/auth",
        }),
        userId: props.username
    });
    this.chatManager.connect().then(currentUser => {
        this.setState({
            currentUser: currentUser
        });
        setInterval(this._pollRooms.bind(this), 5000);
        this._pollRooms();
    }).catch((e) => {
        console.log('Failed to connect to Chatkit');
        console.log(e);
    });
  }

  _pollRooms() {
    const { currentUser } = this.state;
    return currentUser.getJoinableRooms()
        .then((rooms) => {
            this.setState({
                joined: currentUser.rooms,
                joinable: rooms
            })
        });
  }

  _handleJoinTable(id) {
    const { currentUser } = this.state;
    currentUser.joinRoom({roomId: id})
        .then(() => this._pollRooms())
        .then(() => {
            this.setState({
                activeTable: id
            });
        });
  }

  render() {
    const { currentUser, activeTable } = this.state;

    return (
        <Container>
            <Segment padded>
            <Grid divided>
                <Grid.Row>
                <Grid.Column width="4">
                    <Header>Tables</Header>
                    <TableList joined={this.state.joined}
                               joinable={this.state.joinable}
                               currentTable={activeTable}
                               onJoinTable={this._onJoinTable} />
                </Grid.Column>
                <Grid.Column width="12">
                    { activeTable && <Table key={activeTable} currentUser={currentUser} activeTable={activeTable} /> }
                </Grid.Column>
                </Grid.Row>
            </Grid>
            </Segment>
        </Container>
    );
  }
}
