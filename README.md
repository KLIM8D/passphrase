# password_generator
Random password generator made in Go

I wasn't satisfied with how most password generators (mainly in password
managers) work. Since passwords consisting of words have a higher bitwise value,
and are therefore inherently more secure, it would make more sense to build a
password of arbitrary words.

# How-to
Call the binary with the complete path to your dictionary file, and the
number of words you want your password to consist of. E.g. `./password_generator
/path/to/file 5` creates a password consisting of 5 words.

# TODO
* Since most services require numbers/symbols in the password, it would
make some sense to add these somewhere too...
