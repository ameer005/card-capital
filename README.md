# Card Capital 🃏💸

_A trading card economy simulator built by someone learning markets the hard way_

---

## What is Card Capital?

Card Capital is a simulation game where players trade cards in a dynamic market driven by supply, demand, and algorithmic pricing.

You start with limited currency. You invest in cards. Prices move. Chaos happens.

Your goal is simple:

> Become the richest player without breaking the economy.

This project is a mix of game design, system design, and economic experimentation. It exists to explore how markets behave when you build them from scratch.

---

## Why this project exists

This project is an deliberate attempt to:

- understand how markets and pricing systems work
- design a stable in-game economy
- practice building complex systems without tutorials
- get comfortable implementing ideas from scratch

It is not a real trading platform. It is a sandbox for experimentation.

---

## Core Concepts

### Card

Tradable assets with:

- price
- supply
- ownership distribution
- creator rules (limited or unlimited issuance)

Cards can be created by players once they reach certain wealth thresholds.

### Player

Participants in the market who:

- hold currency
- own cards
- make buy/sell decisions
- try to optimize profit

### Market

The engine that:

- processes trades
- updates prices
- enforces rules
- prevents infinite money exploits

### Pricing Engine

An algorithm that adjusts prices based on:

- supply and demand
- trade volume
- randomness
- anti-whale balancing

### Simulation Loop

A system that runs repeated trading cycles to stress-test the economy.

If the economy collapses, the algorithm gets redesigned.

---

## Goals

Primary goals:

- build a functioning simulated economy
- design fair pricing mechanics
- prevent infinite money glitches
- learn system architecture

Secondary goal:

- get better at implementing complex ideas independently

---

## Known Risks

The simulated economy may:

- inflate uncontrollably
- crash instantly
- reward broken strategies
- expose flaws in the pricing model

These failures are intentional learning signals.

---

## Future Ideas

- Multiplayer mode
- UI visualization
- Advanced pricing models
- Player-created assets
- Anti-whale mechanics
- Less terrible math

---

## Philosophy

Build → test → break → fix → repeat.

Complex systems are understood by constructing them, not by watching tutorials.

---

## License

Open for experimentation. Use, modify, and extend freely.
