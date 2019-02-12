# IPFS / IPNS peer ID generator

This tool generates IPFS public and private keypair until it finds public key
which contains required substring. Keys are stored in local directory. If you
like one of them, you can move it to ~/.ipfs/keystore/ to use it with IPFS.

## Installation

```
go get github.com/meehow/peer-id-generator
```

## Usage

<pre>
➜  ~ peer-id-generator dupa
12D3KooWDpLVwxMacesha42m<span style="color:green">dupa</span>3H8jU7bbJdAmsMD7iPjzrkfK
12D3KooWLhGVro7Ks<span style="color:green">DUPa</span>4FX3LpWC61xACoj7AEZg5wmz6VqDGJ9
12D3KooWMJm<span style="color:green">DUPA</span>tjcaHo8LARo5VHQ5uSxrt8Mdi5fytiRnnfD2f
^C
➜  ~ cp 12D3KooWMJmDUPAtjcaHo8LARo5VHQ5uSxrt8Mdi5fytiRnnfD2f ~/.ipfs/keystore/dupa
➜  ~ ipfs name publish --key=dupa QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn
Published to 12D3KooWMJmDUPAtjcaHo8LARo5VHQ5uSxrt8Mdi5fytiRnnfD2f: /ipfs/QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn
</pre>
