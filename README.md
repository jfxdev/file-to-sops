# File to SOPS

## Description

File to SOPS is a tool to ease file content encryption process, making them compatible with Mozilla SOPs

## Usage

```shell
filets -f certificate.pem -o yaml
```

## Output
```yaml
content: |
    Zm9vYmFyYml6
metadata:
  hash: 4a7334879405e4e72e90a3a4e4300d7968bd45fdba924e497a1dc58e4661ad8361d7cd52b87256da39925d6ba9d85a4f4c7e4b2f26938fac97754f1f66e25585
  filename: certificate.pem
```

- Limited to 64kb file