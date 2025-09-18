I've analyzed the OCM Support CLI codebase and created a comprehensive CLAUDE.md file that includes:

## Key Information Captured:

1. **Build Commands**: Make targets for building (`make build`, `make install`, `make clean`)
2. **Prerequisites**: Go 1.17+, OCM CLI authentication, proper permissions
3. **Architecture Overview**: 
   - Hierarchical Cobra command structure
   - Package organization by resource type
   - OCM SDK integration patterns
4. **Development Patterns**: How to add new commands and manage resources

## Architecture Understanding:

- **Command Structure**: Root → verb (get/create/delete/patch) → resource type → flags
- **Package Design**: Separate packages for each resource type (account, organization, subscription, etc.)
- **SDK Integration**: Uses OpenShift OCM SDK for API interactions
- **Safety Patterns**: Dry-run mode by default, flexible search criteria

The CLAUDE.md focuses on the essential information future Claude instances need to be productive quickly, avoiding obvious details and emphasizing the big-picture architecture that requires reading multiple files to understand.
