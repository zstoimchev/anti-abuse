# Anti-Abuse Risk Platform

Anti-Abuse is an experimental platform for detecting likely account and ban evasion using relationships between accounts, devices, events, and previous abuse decisions.

The platform does not attempt to prove that two accounts belong to the same person. Instead, it evaluates available evidence and returns an explainable risk assessment that a customer application can use to allow, challenge, review, or block an action.

## Problem

Simple account bans are often insufficient.

A user can be banned and return using a new account, email address, phone number, network, or device. Relying on a single identifier such as an IP address creates both easy evasion and false positives.

Anti-Abuse instead models relationships between entities and evaluates multiple signals.

```text
Account A
   │
   └── Device X
          │
          └── Account B

Account A → banned
Account B → new account

Result:
Account B has elevated risk because Device X was previously associated with a banned account.
```

## Current Scope

The first version focuses on a customer-local anti-abuse system.

It provides:

* account registration
* device registration
* account-device relationships
* security and abuse events
* account bans
* rule-based risk evaluation
* explainable risk signals

The initial system intentionally does not include blockchain, distributed consensus, machine-learning models, endpoint agents, or cross-customer reputation sharing.

## Example

A customer application can ask the platform to evaluate a newly created account:

```http
POST /v1/risk/evaluate
```

Example response:

```json
{
  "risk": 0.92,
  "decision": "review",
  "signals": [
    "device_previously_linked_to_banned_account"
  ]
}
```

The customer remains responsible for the final action.

## Architecture

```text
Customer Application
        │
        │ REST / JSON
        ▼
┌─────────────────────┐
│     Anti-Abuse API  │
│                     │
│ Accounts            │
│ Devices             │
│ Events              │
│ Bans                │
│ Relationships       │
│ Risk Engine         │
└──────────┬──────────┘
           │
           ▼
      PostgreSQL
```

The first implementation is intentionally a modular monolith.

## Technology

* Go
* PostgreSQL
* REST / JSON
* Docker
* Docker Compose

## Future Work

Later versions may add:

* configurable risk rules
* additional device and credential signals
* customer dashboard
* registry versioning
* Sparse Merkle Tree commitments and proofs
* independently verifiable registry history
* optional blockchain anchoring
* endpoint/device agents where justified

Cryptographic verification is intended to support auditing and registry integrity rather than perform identity detection itself.

## Status

Early development / proof of concept.

The first milestone is a complete local demonstration where:

```text
Account A registers on Device X
        ↓
Account A is banned
        ↓
Account B registers on Device X
        ↓
Risk evaluation detects the relationship
        ↓
HIGH RISK / REVIEW
```
