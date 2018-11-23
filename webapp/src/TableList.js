// src/Tables.js

import React from 'react';
import { List, Label } from 'semantic-ui-react';

export default class Tables extends React.Component {
    render() {
        const { currentTable } = this.props;

        const joinedRooms = this.props.joined.map(room => ({
            id: room.id,
            name: room.name,
            size: room.userIds.length,
            member: true
        }));

        const joinableRooms = this.props.joinable.map(room => ({
            id: room.id,
            name: room.name,
            size: room.userIds.length,
            member: false
        }));

        const rooms = [].concat(joinedRooms, joinableRooms)
            .sort((a, b) => {
                var nameA = a.name.toUpperCase(); // ignore upper and lowercase
                var nameB = b.name.toUpperCase(); // ignore upper and lowercase
                if (nameA < nameB) {
                  return -1;
                }
                if (nameA > nameB) {
                  return 1;
                }

                // names must be equal
                return 0;
            });

        const tables = rooms.map(table => (
            <List.Item as="a" key={table.id} onClick={() => this.props.onJoinTable(table.id)}>
                <List.Content>
                    <Label ribbon={table.id === currentTable} color={table.member ? 'red' : 'grey'}>
                        {table.name}
                        <Label.Detail>
                            ({table.size})
                        </Label.Detail>
                    </Label>
                </List.Content>
            </List.Item>
        ));

        return (
            <List>
                { tables }
            </List>
        );
    }
}
