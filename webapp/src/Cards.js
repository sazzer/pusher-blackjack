import React from 'react';
import { List } from 'semantic-ui-react';

const cardStyles = {
    clubs: {
        color: 'green'
    },
    diamonds: {
        color: 'blue'
    },
    hearts: {
        color: 'red'
    },
    spades: {
        color: 'black'
    }
};

const cardNames = {
    clubs: '♣',
    diamonds: '♦',
    hearts: '♥',
    spades: '♠'
};

export default function Cards({cards, horizontal}) {
    const cardElements = cards.map((card, index) => (
        <List.Item style={cardStyles[card.suit]} key={`card-${index}`}>
            {card.face} {cardNames[card.suit]}
        </List.Item>
    ));
    return (
        <div>
            <List horizontal={horizontal}>{cardElements}</List>
        </div>
    )
}
