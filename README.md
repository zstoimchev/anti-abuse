# Anti-Abuse

Anti-Abuse is an experimental platform for detecting likely account and ban evasion using relationships between accounts, devices, events, and previous abuse decisions.

The system does not attempt to prove that two accounts belong to the same person. Instead, it evaluates available evidence and produces an explainable risk assessment that an integrating application can use when deciding whether to allow, challenge, review, or block an action.

## Problem

A simple account ban is often easy to evade. If `Account A` is blocked, the user can create another `Account B`.

The new account may use a different email address, phone number, network, or other identifiers.

Anti-Abuse models relationships between entities instead of relying on a single identifier such as an IP address.

For example, if `Account A` was banned on `Device X`, and another `Account B` is created on the same `Device X`, the relationship between the new account and a device previously associated with a banned account can be used as a risk signal.

## Goal

The goal of the platform is to provide applications with an API for evaluating abuse risk. The customer application remains responsible for the final decision.

## Technology

The initial implementation uses:

* Go for the API;
* PostgreSQL for persistent storage;
* REST/JSON for communication;
* Docker for local deployment.

The project starts as a modular monolith. Additional infrastructure will only be introduced when required.

## License

$\copyright$ Zhivko Stoimchev 2026. All rights reserved
