# IPFS / IPNS peer ID generator

This tool generates IPFS public and private keypair until it finds public key
which contains required substring. Keys are stored in local directory. If you
like one of them, you can move it to ~/.ipfs/keystore/ to use it with IPFS.

## Installation

```
go get github.com/meehow/peer-id-generator
```

## Usage


➜  ~ peer-id-generator **dupa**

12D3KooWDpLVwxMacesha42m**dupa**3H8jU7bbJdAmsMD7iPjzrkfK

12D3KooWLhGVro7Ks**DUPa**4FX3LpWC61xACoj7AEZg5wmz6VqDGJ9

12D3KooWMJm**DUPA**tjcaHo8LARo5VHQ5uSxrt8Mdi5fytiRnnfD2f

^C

➜  ~ cp 12D3KooWMJm**DUPA**tjcaHo8LARo5VHQ5uSxrt8Mdi5fytiRnnfD2f ~/.ipfs/keystore/dupa

➜  ~ ipfs name publish --key=dupa {ipfs-path}

Published to 12D3KooWMJm**DUPA**tjcaHo8LARo5VHQ5uSxrt8Mdi5fytiRnnfD2f: {ipfs-path}



## Examples

In [examples](examples) you can find key pairs which contain word "funny"
at the end of Peer ID. They are just for demonstration purposes.
Please consider these private keys as already leaked.
