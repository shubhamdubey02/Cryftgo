# Metal Blockchain gRPC

Now Serving: **Protocol Version 37**

Protobuf files are hosted at
[https://buf.build/shubhamdubey02/cryftgo](https://buf.build/shubhamdubey02/cryftgo) and
can be used as dependencies in other projects.

Protobuf linting and generation for this project is managed by
[buf](https://github.com/bufbuild/buf).

Please find installation instructions on
[https://docs.buf.build/installation/](https://docs.buf.build/installation/).

Any changes made to proto definition can be updated by running
`protobuf_codegen.sh` located in the `scripts/` directory of AvalancheGo.

Introduction to `buf`
[https://docs.buf.build/tour/introduction](https://docs.buf.build/tour/introduction)

## Protocol Version Compatibility

The protobuf definitions and generated code are versioned based on the
[RPCChainVMProtocol](../version/version.go#L13) defined for the RPCChainVM.
Many versions of an Avalanche client can use the same
[RPCChainVMProtocol](../version/version.go#L13). But each Avalanche client and
subnet vm must use the same protocol version to be compatible.

## Publishing to Buf Schema Registry

- Checkout appropriate tag in cryftgo `git checkout v1.11.3`
- Change to proto/ directory `cd proto`.
- Publish new tag to buf registry. `buf push -t v26`

Note: Publishing requires auth to the shubhamdubey02 org in buf
https://buf.build/shubhamdubey02/repositories