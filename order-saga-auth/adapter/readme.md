### Why: Adapter?
- Adapter is an impl of port, which stands for external logic including: database calls, thrid-party apis, ...
- Driving port (input port): http-api/web adapter, ...
- Driven port (output port): repository db impl

### Testing
- Adapter testing must be an integration test.
- So that we should spin up the real database connection and then test the adapter.