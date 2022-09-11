# sops-for-files

## Mozilla SOPS helper to encrypt a file content

This `cli` encode content from an file and create a `secret.yaml` file in same directory.

So you can run `sops` after this operation, and encrypt your file.

This is useful to secure store certificates, keys and confidential documents in a git repository.

- Limited to 64kb file