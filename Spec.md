# Harmonic Ledger Protocol Specification (Draft v0.1)

This document specifies the core architecture, invariants, and operational principles
of the Harmonic Ledger protocol.

The protocol is designed as a research-first distributed ledger system that prioritizes:
- structural security over incentives,
- bounded extractable value,
- decentralization through causal separation,
- and stability under adversarial and market stress.

This specification intentionally precedes any production implementation.

---

## 1. Design Goals

The Harmonic Ledger is designed to achieve the following high-level goals:

1. Avoid global linear transaction ordering except where causally necessary.
2. Bound toxic MEV by removing pre-commit orderflow visibility.
3. Prevent long-lived coordination and cartel formation.
4. Scale through parallelism based on causality, not trust.
5. Remain stable (over-damped) under congestion and market shocks.
6. Avoid single points of economic, governance, or operational failure.

---

## 2. Non-Goals

The protocol explicitly does NOT aim to:

- Optimize for maximum short-term throughput at the cost of decentralization.
- Provide guaranteed price stability or returns.
- Replace all existing chains or ecosystems.
- Solve cross-venue (CEX–DEX) arbitrage at the global market level.
- Eliminate all MEV; only structurally bound its most harmful forms.

---

## 3. Core Invariants

The following invariants must hold at all times:

### I1. Safety
Once finalized, valid state transitions must not be reverted except under explicitly defined catastrophic recovery procedures.

### I2. Censorship Resistance
No single actor or small coalition can indefinitely prevent the inclusion of valid transactions.

### I3. MEV Boundedness
No participant should gain extractable value from pre-commit knowledge of transaction contents.

### I4. Decentralization
No persistent role exists whose control implies disproportionate system power.

### I5. Locality of Failure
Failures must remain local and must not propagate system-wide.

---

## 4. High-Level Architecture

The protocol is structured around four conceptual layers:

1. Transaction Intake Layer
2. Ordering & Execution Layer
3. Finality Layer
4. Data Availability Layer

Each layer operates with minimal assumptions about the others.

---

## 5. Transaction Intake Layer (Encrypted Pre-Commit)

### Purpose

To accept user transactions without revealing actionable content prior to commitment.

### Core Properties

- Transactions enter the system as encrypted payloads.
- Transaction validity (format, fees, size) is checked without revealing semantics.
- Admission produces a receipt or commitment that is publicly verifiable.

---

## 6. Causal Domains (CDs)

### Definition

A Causal Domain (CD) is a dynamically determined subset of state objects
that are currently interacting or potentially conflicting.

### Properties

- Transactions operating on disjoint CDs do not require global ordering.
- CDs may split or merge based on observed conflict graphs.
- CDs operate independently for ordering and execution.

---

## 7. Ordering Mechanism

### Principles

- Ordering is deterministic and verifiable.
- Ordering does not rely on auctions.
- Randomness may be used only as a tie-breaker.

### Constraints

- Ordering must not depend on cleartext mempool visibility.
- Ordering authority must rotate on non-stationary schedules.

---

## 8. Execution Model

- Execution is deterministic.
- Parallel execution is permitted for non-conflicting transactions.
- Conflicts are resolved within the relevant CD.

---

## 9. Finality

### Requirements

- Finality must be fast and explicit.
- Finality must not depend on a single proposer.
- Finality failures must be detectable and punishable.

---

## 10. Data Availability

- All finalized state transitions must be publicly available.
- Data availability failures must be detectable independently of execution correctness.
- Data availability must not depend on trusted off-chain parties.

---

## 11. Governance Principles

- Protocol invariants are not subject to simple token-majority voting.
- Governance operates slowly and within bounded parameter ranges.
- Emergency powers are limited, auditable, and revocable.

---

## 12. Security Model (Outline)

The protocol assumes:
- Rational and adversarial participants.
- Partial network asynchrony.
- Attempted cartel formation.
- Attempted key compromise and operational failures.

The protocol is designed to remain safe and usable under these conditions.

---

## 13. Open Questions

The following areas require further specification and research:

- Exact CD split/merge algorithms.
- Threshold encryption committee design.
- Fee control parameters and stability proofs.
- Emergency recovery mechanisms.
- Formal verification scope.

---

## 14. Status

This specification is a draft.
It defines structure and intent but not a finalized implementation.

---

## 15. Authority Separation Model

The protocol separates authority into independent functional roles in order to prevent
persistent power concentration and cartel formation.

The following authority roles are defined:

- **Admission Authority**: accepts encrypted transactions and issues admission receipts.
- **Decryption Authority**: performs threshold decryption of committed batches.
- **Ordering Authority**: determines transaction order within a Causal Domain.
- **Execution Authority**: executes state transitions deterministically.
- **Finality Authority**: finalizes ordered state transitions.

No single entity or committee may simultaneously control more than one authority role
within the same Causal Domain.

Authority roles rotate over time using non-stationary schedules.

---

## 16. Committee Rotation and Time Structure

Authority committees rotate according to non-fixed schedules designed to prevent
predictable phase-locking and long-term coordination.

Time is treated as a logical variable rather than a fixed global clock.

Rotation schedules may include:
- randomized epoch boundaries,
- bounded jitter,
- and non-constant rotation intervals.

Exact rotation mechanisms are implementation-dependent but must satisfy Invariant I4
(Decentralization).

---

## 17. Fee and Congestion Control (Stability Principle)

Fees are designed to regulate load, not to allocate priority.

The protocol aims to remain over-damped under congestion:

- sudden demand spikes must not produce runaway fee escalation,
- local congestion must not propagate globally,
- and recovery from overload must be smooth.

Fee mechanisms must prioritize variance minimization and tail-risk control
over short-term revenue maximization.

---

## 18. Emergency Powers and Failure Containment

The protocol may define limited emergency mechanisms to contain catastrophic failures.

Emergency mechanisms must satisfy:

- explicit activation criteria,
- multi-party authorization,
- time delays where feasible,
- public transparency and post-mortem requirements.

Emergency mechanisms must not allow discretionary control over user funds
or protocol invariants.

---

## 19. Explicit Prohibitions

The following design choices are explicitly prohibited:

- Global cleartext mempools.
- Auction-based transaction ordering.
- Permanent privileged proposer roles.
- Governance mechanisms based solely on token-majority voting.
- Emergency actions without public justification.

---

## 20. Architectural Commitments

The Harmonic Ledger commits to the following architectural principles:

- Causality precedes chronology.
- Information precedes power.
- Stability precedes efficiency.
- Failure must remain local.
- No single variable must control security, governance, and economics simultaneously.

