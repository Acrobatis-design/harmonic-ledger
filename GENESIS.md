# Genesis Allocation Constitution (v1.0)

This document defines the genesis allocation, founder allocation constraints, and non-negotiable safeguards
for the Harmonic Ledger project (the “Protocol”). It is written to be publicly auditable and to provide
credible commitments that reduce governance capture risk, market overhang risk, and operational single-point-of-failure risk.

## Status

- This document is binding for the genesis configuration of the first public network (“Genesis Network”).
- Until a production network exists, this is a research specification and may be revised.
- Once a genesis file is published for a production network, any deviation from this constitution must be
  explicitly declared, versioned, and justified in writing.

## Definitions

- **Total Supply (S_total):** The total token supply at genesis for the Genesis Network.
- **Founder Allocation Cap:** The maximum share of S_total allocated to founder-controlled buckets at genesis.
- **Founder-Controlled Buckets:** Buckets A, B, C defined below, subject to strict constraints.
- **Circulating Supply:** Tokens not subject to lockups/vesting restrictions and not held in restricted system reserves.

---

## I. Total Supply Policy

1. **Fixed supply at genesis:** The Genesis Network launches with a fixed total supply **S_total**.
2. **No implied inflation:** No inflation schedule is assumed by default. If a future inflation mechanism is proposed,
   it must be explicitly specified in a separate document and must not retroactively change genesis allocations.

---

## II. Founder Allocation (Hard Cap = 15%)

The Protocol enforces a strict founder allocation cap:

- **Founder Allocation Cap:** **15% of S_total**
- This cap is partitioned into three distinct buckets with different purposes and constraints.
- Funds are separated at genesis into distinct addresses/modules. Inter-bucket transfers are prohibited by default.

### Allocation Table (Genesis)

| Bucket | Share of S_total | Purpose | Key Constraints |
|---|---:|---|---|
| **A. Locked Founder Reserve** | **7%** | Long-term founder upside | Long vesting, non-governing, non-validator, non-liquidity-dominant |
| **B. Research & Stewardship Fund** | **5%** | Protocol R&D, audits, core maintenance | Programmatic treasury rules, transparent spending, non-governing |
| **C. Emergency Stability Buffer** | **3%** | System shock absorption & incident response | Strict triggers, time-locks, capped deployment, no trading/price support |
| **TOTAL** | **15%** | — | — |

---

## III. Bucket A — Locked Founder Reserve (7%)

### Purpose

To provide long-term alignment and upside for the founder(s) without creating early market overhang,
governance capture vectors, or validator centralization.

### Hard Constraints (Non-Negotiable)

A1. **Vesting:** 7% of S_total vests over **6 years** with a **24-month cliff**.  
A2. **No early unlock:** Early unlock is prohibited under all circumstances.  
A3. **Non-governing:** Tokens in Bucket A **must carry zero governance voting power** (directly or indirectly).  
A4. **Non-validator:** Tokens in Bucket A **must not be eligible for validator bonding** or securing committee roles.  
A5. **No early liquidity dominance:** Tokens in Bucket A must not be used for liquidity provisioning during the cliff
    and should not be used to dominate liquidity venues during vesting.

### Vesting Schedule (Normative)

Let **t** be time since genesis, measured in months.

- Cliff period: t < 24 → vested = 0
- Linear vesting thereafter for 48 months: 24 ≤ t ≤ 72
- Fully vested at t > 72

Vested fraction of Bucket A:

- If t < 24:
  - vested_A(t) = 0
- If 24 ≤ t ≤ 72:
  - vested_A(t) = 0.07 * (t - 24) / 48
- If t > 72:
  - vested_A(t) = 0.07

---

## IV. Bucket B — Research & Stewardship Fund (5%)

### Purpose

To fund work required to make the Protocol real and safe, including:
- core development
- formal specification work
- security audits
- reference client maintenance
- testnet operations and tooling
- documentation

### Structure

B1. Bucket B is controlled by a **programmatic on-chain treasury** (or an equivalent transparently auditable mechanism
in early phases), with all outflows publicly visible.  
B2. Bucket B is **not** a personal founder wallet.  
B3. Bucket B tokens must carry **zero direct governance voting power** (unless explicitly and narrowly authorized
for treasury administration only, with strict limits and transparency).

### Release Policy (Normative)

B4. **Maximum release rate:** No more than **1% of S_total per year** may be disbursed from Bucket B.  
B5. **Minimum granularity:** Disbursements must occur at most monthly (no large lump withdrawals).  
B6. **Unspent funds remain locked:** Unspent tokens remain in the treasury.

### Anti-Capture Constraints

B7. Bucket B must not be used to:
- purchase governance outcomes
- bribe validators/builders/relays
- dominate liquidity provisioning
- create hidden market support operations

---

## V. Bucket C — Emergency Stability Buffer (3%)

### Purpose

To reduce systemic risk by providing a bounded, rule-based shock absorber for severe incidents (e.g., protocol exploit,
oracle failure cascades, catastrophic market dislocation affecting protocol stability).

This bucket exists to prevent “social forks” and ad-hoc emergency interventions.

### Hard Constraints (Non-Negotiable)

C1. Bucket C is **not a trading fund** and must not be used for discretionary market operations.  
C2. Bucket C must be protected by:
- threshold control (e.g., m-of-n signers)
- mandatory time-locks
- explicit trigger conditions
- strict per-event caps

C3. Bucket C must carry **zero governance voting power**.

### Trigger Conditions (Normative)

A Bucket C deployment is permitted only if ALL of the following are satisfied:

- **T1: Emergency condition** is met (as defined in an incident policy), such as:
  - confirmed protocol exploit with material loss
  - oracle manipulation causing liquidation cascades beyond defined thresholds
  - failure of a critical system component with measurable systemic risk

- **T2: Multi-party authorization**:
  - approval by at least two independent authorization groups (e.g., a Security Council and a Validator Council),
    each with quorum requirements

- **T3: Time-lock delay**:
  - a minimum delay (e.g., 24–72 hours) between authorization and execution, unless the emergency condition
    explicitly requires faster response and that exception is publicly justified

### Per-Event Cap (Normative)

C4. Each incident response must be capped to a maximum percentage of Bucket C (e.g., ≤ 25% of Bucket C per event),
to prevent depletion and misuse.

### Transparency & Post-Mortem

C5. Every deployment must be accompanied by:
- an on-chain record of the action
- a public post-mortem describing cause, impact, mitigation, and prevention

---

## VI. Governance Separation (Constitutional Principle)

G1. **Founder tokens must not control governance.**  
Specifically, tokens in Buckets A, B, and C must have **zero direct governance voting power**.

G2. Protocol invariants (safety, censorship resistance, MEV-bounded ordering requirements, and treasury restrictions)
must not be changeable by simple token-majority voting.

G3. Any future governance mechanism must include:
- slow activation (time delays)
- bounded parameter ranges
- multiple independent veto/quorum points

---

## VII. Custody Requirements for Founder-Controlled Buckets

To reduce single-point-of-failure risk:

K1. Buckets A, B, and C must be secured using **threshold custody** (e.g., 3-of-5 or 4-of-7) with:
- geographically separated key shares
- hardware-backed signing devices
- documented recovery procedures
- periodic key-rotation procedures

K2. No bucket may be secured by a single hot wallet.

---

## VIII. Public Disclosure Commitments

D1. The project commits to publish:
- the genesis allocations
- vesting contracts/modules (or equivalent)
- treasury outflow reports
- emergency buffer deployments and post-mortems

D2. Any deviation from this constitution must be declared in writing and versioned.

---

## IX. Non-Goals / Explicit Disclaimers

- This document does not constitute financial advice or an offer of securities.
- No promise of price appreciation, yield, or listing is made.
- This is a research-first protocol effort; deployment occurs only after security review and testing.

---

## Appendix A — Minimal Genesis Implementation Notes (Non-Normative)

At genesis, implementers should ensure:

- Separate addresses/modules for A, B, C
- Enforced vesting for A with cliff + linear schedule
- Treasury rules for B (rate-limited disbursement)
- Emergency controls for C (threshold + time-lock + caps)
- On-chain tagging/metadata to ensure public auditability
