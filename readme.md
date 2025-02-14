# OpenAPI Client Library Generation Plan

This document outlines our plan to automatically generate client libraries for our API based on the OpenAPI spec. We support multiple languages including **Go**, **Python**, **JavaScript**, and **JVM-based languages**. The goal is to ensure that our client libraries stay in sync with changes to the API and are automatically updated in their respective repositories.

## Overview

- **Source:** The OpenAPI specification is maintained in a central repository.
- **Targets:** Client libraries for Go, Python, JavaScript, and JVM.
- **Approach:** Automated generation with periodic/diff-based checks and CI/CD integration.

## Architecture

1. **Source Repository**
   - The OpenAPI specification is maintained in a central repository.
   - Changes to the spec trigger client generation processes either immediately (via webhook/CI trigger) or on a scheduled basis (e.g., weekly).

2. **Client Repositories**
   - Separate repositories (or branches) for each client library.
   - Generated code is committed to these repositories, with diffs reviewed via pull requests or merge requests.

3. **Generation Tooling**
   - Use [OpenAPI Generator](https://openapi-generator.tech/) (or similar tooling) to generate client code.
   - Maintain configuration files for each language to customize code generation (naming conventions, package management, etc.).

4. **CI/CD Pipeline**
   - Set up CI/CD pipelines (GitHub Actions, Jenkins, GitLab CI, etc.) to:
     - Detect changes in the OpenAPI spec.
     - Trigger the generation of client libraries.
     - Compare generated output with the current repository state.
     - Commit changes or create a pull request if there are differences.

## Implementation Steps

### 1. Setup Environment

- **Tool Installation:** Ensure that `openapi-generator-cli` is installed in the CI/CD environment.
- **Configuration:** Create a configuration file (e.g., YAML or JSON) for each target language with required generation options.
- **Docker (Optional):** Use Docker containers to ensure consistent environments for code generation.

### 2. Automated Detection of Changes

- **Webhook Integration:** Configure a webhook on the OpenAPI spec repository to trigger CI jobs upon commits.
- **Scheduled Checks:** Alternatively, or in addition, set up a scheduled CI job (e.g., weekly) to:
  - Pull the latest OpenAPI spec.
  - Regenerate client libraries.
  - Compare generated code with the existing client repository using a diff tool.

### 3. Client Library Generation Process

- **Fetch Spec:** Pull the latest OpenAPI spec from the central repository.
- **Generate Clients:**
  - Run the `openapi-generator-cli` for each target language using the respective configuration.
  - Example command:
    ```bash
    openapi-generator-cli generate -i path/to/openapi.yaml -g go -o ./generated/go
    ```
- **Diff & Validation:**
  - Compare the generated code with the existing client library.
  - Run tests/builds to ensure the generated code compiles and passes basic sanity checks.

### 4. Update Process

- **Automated Commits/PRs:**
  - If differences are detected, automatically commit the changes or create a pull request in the respective client repository.
  - Include meaningful commit messages and changelogs to document updates.
- **Review Process:**
  - Developers review the diffs to ensure that the changes are expected.
  - After approval, the changes are merged into the main branch.

### 5. Monitoring and Reporting

- **Logging:** Implement logging within the CI/CD pipelines to capture the status of generation and updates.
- **Alerts:** Configure notifications (via email, Slack, etc.) for failures in the generation process.
- **Metrics:** Optionally, track metrics on how often client libraries are updated and any failures in the generation process.

## Issues We Need to Resolve

As we progress with this automation plan, the following issues require further discussion and resolution:

- **Handling Existing Clients and Named Functions:**  
  How do we manage scenarios where existing clients have custom-named functions that might conflict with generated ones?  
  *Plan:* Consider implementing a naming strategy or aliasing mechanism within the generator configurations. We may also need a migration guide to help developers transition smoothly if function names are updated.

- **`next_url` Pagination:**  
  Our API uses `next_url` for pagination. Should we generate a customer function to handle this pattern?  
  *Plan:* Evaluate the possibility of integrating custom templates or hooks in the generation process that can incorporate a standardized pagination function across all clients.

- **Websocket Streaming:**  
  How do we handle API endpoints that involve websocket streaming, given that the generation tool may not fully support streaming out-of-the-box?  
  *Plan:* Investigate creating custom code templates or post-generation scripts that add websocket support. Collaboration with the team to define a common interface for streaming clients might be necessary.

- **Code Examples for Documentation:**  
  How do we generate or maintain code examples for our docs alongside the client libraries?  
  *Plan:* Consider integrating code example generation as part of the client build process, or maintain a separate repository where code examples are automatically updated based on generated clients. These examples can then be referenced in our documentation.

## Future Enhancements

- **Extended Testing:** Integrate more comprehensive testing suites for the generated clients.
- **Customization Hooks:** Allow for custom patches or manual overrides for specific parts of the client libraries.
- **Versioning:** Implement version tagging in client repositories corresponding to API version changes.

## Summary

This plan ensures that our client libraries for Go, Python, JavaScript, and JVM remain in sync with our API. The automated generation, diff-checking, and CI/CD integration minimize manual work and reduce the risk of out-of-date client libraries. We will continue to iterate on this process, addressing the issues outlined above to improve the overall quality and maintainability of our generated clients.

