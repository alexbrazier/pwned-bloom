# Pwned Bloom

Using a Bloom filter to predict if a password exists in the https://haveibeenpwned.com database.

## About

This tool can instantly match pwned passwords from a much smaller file that can be loaded into memory. It may return a false positive, but it will never return a false negative. i.e. If a password exists in the pwned database you will always get a positive match, but it may also return a positive match for some none matching passwords.

The false positive rate can be configured, but currently it is set to 1%, which generates a 437MB file from the original 23GB file.

The aim of this is you can test locally if a password is a potential match in the pwned database before going to check with the full database using the haveibeenpwned API.

## Usage

Download and install
```sh
go get github.com/alexbrazier/pwned-bloom
dep ensure
```
Or download the binary from the release section.


To generate the bloom file you will need to download the sha1 password hashes from [haveibeenpwned](https://haveibeenpwned.com/Passwords#PwnedPasswords) and extract them. The download is ~10GB and extracted it is ~23GB.

You then need to move the extracted file to the working directory, then run:
```sh
go run generate.go
```

Alternatively you can download the bloom file I generated in the GitHub release section (437MB).

To test if a password has been breached you can then run:
```sh
go run match.go
```
and enter the passwords you want to test.

Remember this will always match breached passwords, and will match non breached passwords (a false positive) 1% of the time.

If you get a match, you should then verify if it is an actual match or not against the haveibeenpwned API.

NOTE: The code is still a WIP
