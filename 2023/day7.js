// const {true
//   Worker, isMainThread, parentPort, workerData,
// } = require('node:worker_threads');
// const Progress = require('progress');
const fs = require('fs');

function readInput(dummy) {
  const input = fs.readFileSync(dummy ? './dummy.day7' : './input.day7', 'utf8').split('\n').filter(line => line != '');
  return input;
}

class Hand {
  constructor(cards, bid, joker = false) {
    this.rank = 0;
    this.bid = bid;
    this.cards = cards;
    this.strength = cards.split('').map(card => joker ? cardWithJokerStrengthMap.get(card) : cardStrengthMap.get(card));
    this.type = joker ? Hand.classifyHandWithJoker(this) : Hand.classifyHand(this);
  }

  static classifyHand(hand) {
    const cards = hand.cards.split('');
    const cardMap = new Map();

    for (let i = 0; i < cards.length; i++) {
      const card = cards[i];
      if (!cardMap.has(card)) {
        cardMap.set(card, 1);
      } else {
        cardMap.set(card, cardMap.get(card) + 1);
      }
    }

    const cardMapValues = [...cardMap.values()];

    const valuesMap = new Map();
    for (let i = 0; i < cardMapValues.length; i++) {
      const val = cardMapValues[i];
      if (!valuesMap.has(val)) {
        valuesMap.set(val, 1);
      } else {
        valuesMap.set(val, valuesMap.get(val) + 1);
      }
    }

    if (valuesMap.has(5)) {
      return 'five-of-kind';
    }

    if (valuesMap.has(4)) {
      return 'four-of-kind';
    }

    if (valuesMap.has(3) && valuesMap.has(2)) {
      return 'full-house';
    }

    if (valuesMap.has(3)) {
      return 'three-of-kind';
    }

    if (valuesMap.has(2)) {
      if (valuesMap.get(2) == 2) {
        return 'two-pair';
      }

      return 'one-pair';
    }

    return 'high-card';
  }

  static classifyHandWithJoker(hand) {
    const cards = hand.cards.split('');
    const cardMap = new Map();

    for (let i = 0; i < cards.length; i++) {
      const card = cards[i];
      if (card == 'J') continue;
      if (!cardMap.has(card)) {
        cardMap.set(card, 1);
      } else {
        cardMap.set(card, cardMap.get(card) + 1);
      }
    }

    const cardMapValues = [...cardMap.values()].sort((a, b) => b - a);
    const jokers = cards.filter(card => card == 'J').length;
    if (jokers > 0) {
      cardMapValues[0] ?
        cardMapValues[0] += jokers :
        cardMapValues[0] = jokers;
    }

    const valuesMap = new Map();
    for (let i = 0; i < cardMapValues.length; i++) {
      const val = cardMapValues[i];
      if (!valuesMap.has(val)) {
        valuesMap.set(val, 1);
      } else {
        valuesMap.set(val, valuesMap.get(val) + 1);
      }
    }

    if (valuesMap.has(5)) {
      return 'five-of-kind';
    }

    if (valuesMap.has(4)) {
      return 'four-of-kind';
    }

    if (valuesMap.has(3) && valuesMap.has(2)) {
      return 'full-house';
    }

    if (valuesMap.has(3)) {
      return 'three-of-kind';
    }

    if (valuesMap.has(2)) {
      if (valuesMap.get(2) == 2) {
        return 'two-pair';
      }

      return 'one-pair';
    }

    return 'high-card';
  }
}

const cardStrengthMap = new Map();
cardStrengthMap.set('A', 14);
cardStrengthMap.set('K', 13);
cardStrengthMap.set('Q', 12);
cardStrengthMap.set('J', 11);
cardStrengthMap.set('T', 10);
cardStrengthMap.set('9', 9);
cardStrengthMap.set('8', 8);
cardStrengthMap.set('7', 7);
cardStrengthMap.set('6', 6);
cardStrengthMap.set('5', 5);
cardStrengthMap.set('4', 4);
cardStrengthMap.set('3', 3);
cardStrengthMap.set('2', 2);

const cardWithJokerStrengthMap = new Map();
cardWithJokerStrengthMap.set('A', 13);
cardWithJokerStrengthMap.set('K', 12);
cardWithJokerStrengthMap.set('Q', 11);
cardWithJokerStrengthMap.set('T', 10);
cardWithJokerStrengthMap.set('9', 9);
cardWithJokerStrengthMap.set('8', 8);
cardWithJokerStrengthMap.set('7', 7);
cardWithJokerStrengthMap.set('6', 6);
cardWithJokerStrengthMap.set('5', 5);
cardWithJokerStrengthMap.set('4', 4);
cardWithJokerStrengthMap.set('3', 3);
cardWithJokerStrengthMap.set('2', 2);
cardWithJokerStrengthMap.set('J', 1);

function part1(dummy = false) {
  const input = readInput(dummy);

  const hands = {
    'high-card': [],
    'one-pair': [],
    'two-pair': [],
    'three-of-kind': [],
    'full-house': [],
    'four-of-kind': [],
    'five-of-kind': [],
  }

  let ans = 0;
  for (let i = 0; i < input.length; i++) {
    const [cards, bid] = input[i].split(' ').filter(line => line != '');
    const hand = new Hand(cards, parseInt(bid), false);
    hands[hand.type].push(hand);
  }

  for (let i = 0; i < Object.keys(hands).length; i++) {
    const hand = hands[Object.keys(hands)[i]];

    if (hand.length == 0) continue;

    hands[Object.keys(hands)[i]] = hand.sort((a, b) => {
      for (let j = 0; j < a.strength.length; j++) {
        if (a.strength[j] > b.strength[j]) {
          return 1;
        }
        if (a.strength[j] < b.strength[j]) {
          return -1;
        }
      }
      return 0;
    });
  }

  let rank = 1;
  for (let i = 0; i < Object.keys(hands).length; i++) {
    const hand = hands[Object.keys(hands)[i]];

    if (hand.length == 0) continue;

    for (let j = 0; j < hand.length; j++) {
      ans += hand[j].bid * rank++;
    }
  }

  return ans;
}

function part2(dummy = false) {
  const input = readInput(dummy);

  const hands = {
    'high-card': [],
    'one-pair': [],
    'two-pair': [],
    'three-of-kind': [],
    'full-house': [],
    'four-of-kind': [],
    'five-of-kind': [],
  }

  let ans = 0;
  for (let i = 0; i < input.length; i++) {
    const [cards, bid] = input[i].split(' ').filter(line => line != '');
    const hand = new Hand(cards, parseInt(bid), true);
    hands[hand.type].push(hand);
  }

  for (let i = 0; i < Object.keys(hands).length; i++) {
    const hand = hands[Object.keys(hands)[i]];

    if (hand.length == 0) continue;

    hands[Object.keys(hands)[i]] = hand.sort((a, b) => {
      for (let j = 0; j < a.strength.length; j++) {
        if (a.strength[j] > b.strength[j]) {
          return 1;
        }
        if (a.strength[j] < b.strength[j]) {
          return -1;
        }
      }
    });
  }

  let rank = 1;
  for (let i = 0; i < Object.keys(hands).length; i++) {
    const hand = hands[Object.keys(hands)[i]];

    if (hand.length == 0) continue;

    for (let j = 0; j < hand.length; j++) {
      ans += hand[j].bid * rank++;
    }
  }

  return ans;
}

//console.log('part1: ', part1());
console.log('part2: ', part2());
