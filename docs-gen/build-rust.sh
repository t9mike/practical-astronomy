#!/usr/bin/bash

rm -rf ../docs/rust
cd ../rust-implementation
cargo doc --document-private-items --target-dir ../docs/rust
