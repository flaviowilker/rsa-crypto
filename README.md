# rsa-crypto
Simple RSA encryption algorithm implementation

## Features
Create Key Pair:
```sh
$ rsa-crypto -keypair
```
Encrypt text:
```sh
$ rsa-crypto -encrypt
```
Decrypt text:
```sh
$ rsa-crypto -decrypt
```
Verbose commands:
```sh
$ rsa-crypto -v
```
## Execution
All flags are optional, you can execute without flags, with all flags, with one flag, etc. Without flags you is choosing the default execution. Follow the instructions shown in the program.

## Examples

### Example 1 (all flags)
```sh
$ rsa-crypto -keypair -encrypt -decrypt -v
RSA Crypto
----------
Key numbers
n:  731
e:  79
d:  9391
Prime numbers
p:  43
q:  17
Other numbers
z:  672

Key pair
Public Key (n, e): (731, 79)
Private Key (n, d): (731, 9391)

Public Key
Write the N number
731
Write the E number
79
Enter text: hi   

Encrypt text
char 1: h
char ascii:  104
char encrypted:  77
char 2: i
char ascii:  105
char encrypted:  635

Encrypted: [77 635]

Private Key
Write the N number
731
Write the D number
9391

Write separated by space " "
Enter the encrypted text: 77 635

Decrypt text
char 1: h
char encrypted:  77
char ascii:  104
char 2: i
char encrypted:  635
char ascii:  105

Decrypted: hi
```

### Example 2 (without flags)
```sh
$ rsa-crypto
RSA Crypto
----------
Enter text: hi

Key pair
Public Key (n, e): (5251, 41)
Private Key (n, d): (5251, 5353)

Encrypted: [4768 3783]

Decrypted: hi
```