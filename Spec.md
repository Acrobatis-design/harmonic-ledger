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

---
---

## 21. Threat Model

This section defines the adversarial assumptions under which the Harmonic Ledger
protocol is expected to remain safe, usable, and decentralized.

The threat model is explicit and conservative. Any security claim made by the protocol
is valid only within these assumptions.

---

## 21.1 Adversary Classes

The protocol considers the following classes of adversaries:

### A1. Economic Adversaries
Actors seeking to extract profit through:
- MEV strategies (front-running, sandwiching, reordering),
- market manipulation,
- liquidation cascades,
- censorship for economic advantage.

These adversaries may control substantial capital and infrastructure.

### A2. Cartel Adversaries
Coalitions of validators, operators, or infrastructure providers attempting to:
- coordinate transaction ordering,
- control block inclusion,
- dominate liquidity or governance,
- establish long-lived privileged positions.

### A3. Network Adversaries
Actors capable of:
- delaying or dropping messages,
- causing partial network partitions,
- exploiting network-level asymmetries,
but not breaking cryptographic primitives.

### A4. Key Compromise Adversaries
Attackers who may:
- compromise individual user keys,
- compromise some validator or operator keys,
- exploit operational security failures.

Mass simultaneous compromise of all threshold participants is considered unlikely
but not impossible.

### A5. Governance Adversaries
Actors attempting to:
- capture governance mechanisms,
- exploit voter apathy,
- introduce protocol changes that weaken invariants.

---

## 21.2 Explicit Assumptions

The protocol assumes:

A. Standard cryptographic primitives (hash functions, signatures, encryption)
   are computationally secure.

B. Adversaries do not control a supermajority of all independent authority roles
   across the entire system simultaneously.

C. Network asynchrony is partial, not total (messages are eventually delivered).

D. Not all users, validators, or operators are malicious simultaneously.

E. Some degree of rational behavior exists, but worst-case coordination is assumed
   where economically viable.

---

## 21.3 Out-of-Scope Threats

The protocol does NOT attempt to protect against:

- Total global internet shutdown.
- Complete cryptographic break (e.g. hash/signature primitives).
- Simultaneous compromise of all independent threshold participants.
- Adversaries with unlimited computational and physical resources.

Such threats are considered outside the realistic operating domain.

---

## 21.4 Security Goals Under Adversarial Conditions

Under the above threat model, the protocol aims to guarantee:

G1. **Safety:** finalized state transitions are not reverted.
G2. **Bounded MEV:** no adversary can extract value from pre-commit transaction knowledge.
G3. **Censorship Resistance:** no single actor or cartel can indefinitely censor transactions.
G4. **Decentralization:** no persistent privileged role can be established.
G5. **Failure Containment:** failures remain local and do not cascade system-wide.
G6. **Recoverability:** limited key compromise does not lead to catastrophic loss.

---

## 21.5 Design Responses to Threats

The protocol responds to identified threats as follows:

- Encrypted pre-commit transaction intake limits MEV from orderflow visibility.
- Authority separation prevents single-role dominance.
- Non-stationary committee rotation limits cartel persistence.
- Causal Domains localize contention and failure.
- Threshold-native custody reduces single-key catastrophic risk.
- Governance separation prevents token-majority capture of invariants.

---

## 21.6 Residual Risks

The following risks are acknowledged but not fully eliminated:

- Cross-venue arbitrage (CEX–DEX latency differences).
- Post-reveal MEV strategies.
- Social coordination attacks outside the protocol.
- Human error in implementation and operations.

These risks are considered manageable and bounded within the protocol design.
---

## 22. Execution Model

This section defines how transactions are interpreted, how state is represented,
how conflicts are detected, and how execution proceeds within the Harmonic Ledger.

The execution model is intentionally explicit to avoid hidden global ordering
assumptions and to enable safe parallelism.

---

## 22.1 State Model

The global protocol state is composed of discrete **state objects**.

A state object is any uniquely identifiable unit of mutable state, such as:
- an account balance,
- a smart contract instance,
- a liquidity position,
- a governance or system parameter object.

Each state object has:
- a unique identifier,
- a current version,
- a deterministic transition function.

State objects are the minimal units of conflict.

---

## 22.2 Transaction Structure

A transaction consists of:

- **Inputs:** references to one or more state objects.
- **Intents:** declared operations to be performed on those objects.
- **Read Set:** state objects whose current values are read.
- **Write Set:** state objects whose values are modified.
- **Fee Commitment:** a bounded fee envelope.
- **Encrypted Payload:** operation details hidden prior to commitment.

Transactions must explicitly declare their read and write sets.

---

## 22.3 Conflict Definition

Two transactions are said to **conflict** if and only if:

- their write sets intersect, or
- one transaction writes to a state object that the other reads.

Transactions whose read and write sets are disjoint are non-conflicting
and may be executed in parallel without ordering constraints.

---

## 22.4 Causal Domains (CDs)

A **Causal Domain (CD)** is defined as a dynamically formed set of state objects
that are currently involved in conflicting transactions.

Properties:

- Transactions are assigned to a CD based on their declared read/write sets.
- Each CD maintains its own ordering and execution context.
- CDs may merge when conflicts span previously separate domains.
- CDs may split when conflict graphs become disconnected.

No global ordering exists across independent CDs.

---

## 22.5 Transaction Lifecycle

The lifecycle of a transaction proceeds as follows:

1. **Submission:** the user submits an encrypted transaction with declared read/write sets.
2. **Admission:** the transaction is admitted if format, fees, and declarations are valid.
3. **CD Assignment:** the transaction is assigned to a Causal Domain.
4. **Ordering:** transactions within a CD are ordered deterministically.
5. **Execution:** ordered transactions are executed against state objects.
6. **Finality:** executed transitions are finalized and become irreversible.

---

## 22.6 Deterministic Execution

Execution within a CD must satisfy:

- Determinism: identical ordered inputs produce identical state transitions.
- No hidden side effects: execution affects only declared write-set objects.
- Explicit failure: failed execution must be detectable and attributable.

Execution engines must not rely on implicit global state.

---

## 22.7 Parallelism Guarantees

The protocol guarantees:

- Parallel execution of non-conflicting transactions.
- Independent progress of distinct CDs.
- Localized rollback or failure handling.

Global stalls caused by local contention are explicitly prohibited.

---

## 22.8 Execution Failure Handling

If a transaction fails during execution:

- its effects are discarded,
- it does not block unrelated transactions,
- failure is recorded for auditability.

Repeated failure patterns may be subject to policy controls
(e.g. rate limits or fee penalties).

---

## 22.9 Implementation Notes (Non-Normative)

Possible implementation approaches include:

- Optimistic parallel execution with conflict checks.
- Pre-declared access lists enforced at admission.
- DAG-based execution within CDs.

The specific execution engine is left to implementation choice,
provided all invariants defined in this specification are preserved.
---

## 23. Cryptographic Plumbing and Information Flow

This section defines how cryptography is used to control information flow,
prevent pre-commit orderflow leakage, and bound extractable value.

The design principle is explicit:
information must not be visible before it becomes non-actionable.

---

## 23.1 Cryptographic Objectives

The cryptographic subsystem must ensure:

C1. Transaction contents are not visible prior to commitment.
C2. No single party can decrypt transactions unilaterally.
C3. Decryption occurs only after ordering becomes binding.
C4. Ordering can be verified independently of transaction secrecy.
C5. Randomness used in ordering cannot be biased by participants.

---

## 23.2 Encrypted Transaction Submission

Each transaction submitted to the protocol consists of:

- A **cleartext header**, containing:
  - declared read set,
  - declared write set,
  - fee commitment,
  - transaction size metadata.

- An **encrypted payload**, containing:
  - operation details,
  - parameters,
  - signatures and authorization data.

The encrypted payload is encrypted using a **threshold public key**
associated with the relevant authority committee.

No authority may access plaintext transaction contents
before the reveal phase.

---

## 23.3 Admission Receipts (Pre-Commit Acknowledgement)

Upon successful admission, the protocol issues a **public admission receipt**
containing:

- a commitment hash of the encrypted payload,
- a timestamp or logical sequence marker,
- the assigned Causal Domain identifier.

Admission receipts provide users with a verifiable guarantee
that the transaction has entered the system
without revealing its contents.

Admission receipts do not imply ordering or finality.

---

## 23.4 Commit–Reveal Lifecycle

The protocol enforces a strict commit–reveal separation:

1. **Commit Phase**
   - Encrypted transactions are accepted and committed.
   - Ordering authorities see only commitments and metadata.
   - No plaintext transaction data is accessible.

2. **Ordering Phase**
   - Transactions within a Causal Domain are ordered
     based on commitments, declared sets, and deterministic rules.
   - Randomness may be used only as a tie-breaker.

3. **Reveal Phase**
   - After ordering is fixed, threshold decryption is performed.
   - Plaintext transaction contents are revealed to execution authorities.

This ordering-before-reveal property is mandatory.

---

## 23.5 Threshold Decryption Model

Decryption is performed via a threshold scheme:

- A decryption key is split among multiple independent participants.
- A quorum of shares is required to reconstruct plaintext.
- No single participant can decrypt alone.

Key rotation must occur periodically to limit long-term compromise risk.

Failure of individual participants must not block progress,
provided quorum assumptions hold.

---

## 23.6 Randomness Generation

Randomness used for ordering tie-breaks and committee rotation must satisfy:

- Unpredictability prior to commitment.
- Verifiability after use.
- Resistance to single-party manipulation.

Acceptable sources include:
- threshold-generated randomness,
- verifiable random functions (VRFs),
- commit–reveal randomness schemes with penalties for non-reveal.

Randomness sources must not depend on transaction plaintext.

---

## 23.7 Information Visibility Matrix

At no point may any single actor observe all of the following simultaneously:

- plaintext transaction contents,
- transaction ordering discretion,
- final execution authority.

The protocol enforces separation such that:

- Admission authorities see ciphertext only.
- Ordering authorities see commitments and metadata only.
- Decryption authorities act only after ordering is fixed.
- Execution authorities receive plaintext only post-reveal.

This separation is fundamental to MEV boundedness.

---

## 23.8 Failure Modes and Safeguards

If threshold decryption fails due to insufficient participation:

- fallback procedures may be triggered,
- penalties or slashing may apply to non-participating authorities,
- alternative decryption committees may be activated.

Fallbacks must not reintroduce cleartext mempools
or discretionary plaintext access.

---

## 23.9 Explicit Prohibitions

The following are explicitly prohibited:

- Cleartext transaction gossip prior to commitment.
- Single-key decryption of user transactions.
- Ordering based on plaintext content.
- Randomness derived from transaction data.

---

## 23.10 Security Rationale

By enforcing encryption before ordering and reveal after ordering,
the protocol eliminates pre-commit information asymmetry.

This structurally bounds:
- front-running,
- sandwich attacks,
- private orderflow exploitation,
- builder and relay cartels.

Residual MEV after reveal is acknowledged but bounded.


