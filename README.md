# SupplyChainGuard Shared Contracts (`scg-contracts`)

This repository contains the **canonical, versioned Protobuf contracts** (types, enums, value objects, DTOs, and domain events) used by all SupplyChainGuard microservices.

## ✨ Why use this library?

- **Single source of truth:** All microservices (Go, gRPC, event bus, REST, JS/TS, Python, etc.) use the same, versioned contract definitions.
- **Business-aligned:** Every field, type, and event is named using SupplyChainGuard's _Ubiquitous Language Glossary_ (e.g., `Unit`, `Consignment`, `Partner`, `Custodian`, `Provenance`, `EventLog`), not just tech jargon.
- **Regulatory ready:** Designed for LkSG, ESG, GDPR, and global supply chain compliance and traceability.
- **Plug-and-play for Go:** `option go_package` is set for all `.proto` files, enabling easy Go codegen and direct microservice import.
- **AI/analytics friendly:** Supports codegen for TypeScript, Python, Java, and easy downstream data lake use.
- **Buf.build integration:** Leverages buf.build for schema validation, linting, and code generation.

---

## 📦 Structure

The repository is organized as follows:

```
proto/
  └── scg/                      # Root namespace for all SupplyChainGuard contracts
      ├── shared/v1/            # Common types, enums, and utilities used across domains
      │   ├── enums.proto       # Shared enumerations (e.g., CurrencyCode)
      │   ├── error.proto       # Standard error types
      │   ├── events.proto      # EventEnvelope and common event structures
      │   └── types.proto       # Reusable value objects (Address, Money, etc.)
      │
      ├── partner/v1/           # Partner domain (suppliers, customers, logistics providers)
      │   ├── enums.proto       # Partner-specific enumerations
      │   ├── events.proto      # Partner domain events
      │   └── models.proto      # Partner domain models
      │
      ├── unit/v1/              # Unit domain (physical items, batches, SKUs)
      │   ├── enums.proto       # Unit-specific enumerations
      │   ├── events.proto      # Unit domain events
      │   └── models.proto      # Unit domain models
      │
      ├── consignment/v1/       # Consignment domain (shipments, orders)
      │   ├── enums.proto       # Consignment-specific enumerations
      │   ├── events.proto      # Consignment domain events
      │   └── models.proto      # Consignment domain models
      │
      ├── compliance/v1/        # Compliance domain (regulations, certifications)
      │   ├── enums.proto       # Compliance-specific enumerations
      │   ├── events.proto      # Compliance domain events
      │   └── models.proto      # Compliance domain models
      │
      ├── custody/v1/           # Custody domain (chain of custody, ownership)
      │   ├── enums.proto       # Custody-specific enumerations
      │   ├── events.proto      # Custody domain events
      │   └── models.proto      # Custody domain models
      │
      ├── digital_twin/v1/      # Digital Twin domain (IoT, sensors, real-time data)
      │   ├── enums.proto       # Digital Twin-specific enumerations
      │   ├── events.proto      # Digital Twin domain events
      │   └── models.proto      # Digital Twin domain models
      │
      ├── event_log/v1/         # Event Log domain (audit, history, traceability)
      │   └── events.proto      # Event Log domain events
      │
      ├── inventory/v1/         # Inventory domain (stock levels, locations)
      │   ├── enums.proto       # Inventory-specific enumerations
      │   ├── events.proto      # Inventory domain events
      │   └── models.proto      # Inventory domain models
      │
      ├── notification/v1/      # Notification domain (alerts, messages)
      │   ├── enums.proto       # Notification-specific enumerations
      │   ├── events.proto      # Notification domain events
      │   └── models.proto      # Notification domain models
      │
      ├── provenance/v1/        # Provenance domain (origin, history)
      │   └── events.proto      # Provenance domain events
      │
      ├── user/v1/              # User domain (accounts, permissions)
      │   ├── enums.proto       # User-specific enumerations
      │   ├── events.proto      # User domain events
      │   └── models.proto      # User domain models
      │
      └── warehouse/v1/         # Warehouse domain (facilities, storage)
          ├── enums.proto       # Warehouse-specific enumerations
          ├── events.proto      # Warehouse domain events
          └── models.proto      # Warehouse domain models
```

### Domain Organization

Each domain follows a consistent structure:

- **enums.proto**: Contains all enumerations specific to the domain
- **events.proto**: Contains all domain events that can be published/subscribed
- **models.proto**: Contains all domain models, DTOs, and value objects

### Naming Conventions

- **UUIDs**: All entity identifiers use the suffix `_uuid` (not `_id`) for consistency and to emphasize their globally unique nature
- **Enums**: All enum types are singular with domain prefix (e.g., `UnitType`, `PartnerStatus`)
- **Events**: All events are named using past tense verbs (e.g., `UnitCreated`, `PartnerRegistered`)
- **Models**: All models are named using singular nouns (e.g., `Unit`, `Partner`)

#### Examples of Proper Field Naming

```
// Good - uses _uuid suffix for identifiers
message Unit {
  string unit_uuid = 1;
  string manufacturer_partner_uuid = 2;
}

// Bad - uses inconsistent naming (_id instead of _uuid)
message Unit {
  string unit_id = 1;
  string manufacturer_partner_id = 2;
}
```

#### Examples of Proper Event Naming

```
// Good - uses specific event names that describe business transitions
message UnitRegistered { ... }
message UnitAllocatedToConsignment { ... }
message UnitQuarantined { ... }

// Bad - uses generic "StatusChanged" events
message UnitStatusChanged { ... }
```

### Documentation Standards

All proto files should include comprehensive documentation comments:

- Use `//` style comments for consistency
- Document all messages, fields, and enum values
- Include purpose, usage, and any constraints for each element
- Document the business meaning, not just the technical details

Example of well-documented message:

```protobuf
// EventLogEntry represents a single event in the audit log
// for traceability and compliance purposes
message EventLogEntry {
  // Metadata about the event (timestamp, actor, etc.)
  proto.scg.shared.v1.EventEnvelope meta = 1;

  // UUID of the entity this event relates to
  string entity_uuid = 2;

  // Type of entity (Unit, Consignment, etc.)
  string entity_type = 3;

  // Type of event that occurred
  string event_type = 4;

  // JSON serialized event data
  string data_json = 5;
}
```

### Versioning Strategy

All contracts are versioned using directory structure (e.g., `v1`, `v2`). Breaking changes require a new version directory.

## 🚀 Usage

### Go

```go
import (
    sharedv1 "github.com/hbttundar/scg-contracts/gen/go/proto/scg/shared/v1"
    consignmentv1 "github.com/hbttundar/scg-contracts/gen/go/proto/scg/consignment/v1"
)

func ProcessConsignment(c *consignmentv1.Consignment) {
    // Use the generated Go structs
}
```

### Code Generation

Run the code generation script to generate Go code:

```bash
./scg proto-gen
```

This will generate Go code in the `gen/go` directory.

### Package Management Commands

The `scg` script includes several commands for managing the codebase:

```bash
# Show all available commands
./scg help

# Build the code
./scg build

# Run tests with race detection and coverage
./scg test

# Run linter on the codebase
./scg lint

# Run linter and fix issues
./scg lint-fix

# Run security checks (govulncheck and gosec)
./scg security

# Generate Protobuf code
./scg proto-gen

# Run all CI checks (build, test, lint, security)
./scg ci

# Install required tools
./scg install-tools
```

These commands help maintain the codebase and ensure code quality.

Note: The `.env` file is already included in `.gitignore` to prevent accidentally committing sensitive information.

## 🛠️ Development Guidelines

### Adding New Proto Files

1. Place new proto files in the appropriate domain directory
2. Follow the naming conventions for messages, fields, and enums
3. Include proper documentation comments for all messages, fields, and enum values
4. Ensure all time fields use `google.protobuf.Timestamp`
5. Set the correct `go_package` option following the pattern: `github.com/hbttundar/scg-contracts/gen/go/scg/{domain}/v1;{domain}v1`
6. Use fully qualified import paths for shared models and enums
7. Run `./scg proto-gen` to verify your changes

### Event Design Guidelines

1. Use specific event names that describe business transitions (e.g., `UnitRegistered`, not `UnitStatusChanged`)
2. Always include the `EventEnvelope` as the first field in every event message
3. Use past tense verbs for event names (e.g., `UnitCreated`, not `CreateUnit`)
4. Include only the necessary data in events (entity UUID and changed fields)
5. Ensure events are immutable and represent facts that have occurred

### Buf Integration

This repository uses [buf.build](https://buf.build) for schema validation, linting, and code generation:

- **buf.yaml**: Defines the module and linting rules
- **buf.gen.yaml**: Configures code generation plugins
- **.buf/work.yaml**: Configures workspace for multi-module repositories

To use buf directly:

```bash
# Lint proto files
buf lint

# Generate code
buf generate

# Check for breaking changes
buf breaking --against '.git#branch=main'
```
### Troubleshooting

#### Common Issues
1. **Code generation errors**: If `./scg proto-gen` fails, check that:
   - All imports are correct
   - All package names are consistent
   - The `buf` tool is installed correctly

## 📚 Glossary

- **Unit**: A physical item, batch, or SKU in the supply chain
- **Consignment**: A shipment or order containing multiple units
- **Partner**: An organization participating in the supply chain (supplier, buyer, carrier, etc.)
- **Custody**: The current possession or control of a unit
- **Provenance**: The origin and history of a unit
- **Digital Twin**: A virtual representation of a physical object with real-time data
- **Compliance**: Adherence to regulations, standards, and certifications
- **Event Log**: A record of all events for audit and traceability purposes

## 🤝 Contributing

### Contribution Guidelines

1. **Fork the repository**: Create your own fork of the repository to work on changes
2. **Create a feature branch**: Make your changes in a new branch
3. **Follow the naming conventions**: Ensure your changes adhere to the naming conventions
4. **Add documentation**: Include comprehensive documentation for any new messages, fields, or enums
5. **Run linting**: Use `buf lint` to check for any linting issues
6. **Check for breaking changes**: Use `buf breaking` to ensure your changes don't break compatibility
7. **Submit a pull request**: Include a clear description of the changes and any relevant context
