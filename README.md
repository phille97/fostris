# FOSTRIS [![Build and Test](https://github.com/phille97/fostris/actions/workflows/qa.yml/badge.svg)](https://github.com/phille97/fostris/actions/workflows/qa.yml)
Fostrande

Demo: https://go.dev/play/p/MBJvtEikqRZ

supports the base64 set of chars and spaces ( because why not ¯\\_(ツ)_/¯ )
```
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/
```

### How to Encode
```
fostrad := fostris.Encode("My name is Fostris")
```

### How to Decode
```
unfostrad := fostris.Decode("fosTRisfoStrIS fOSTriSfoStRIsfoSTriSfoSTRIs foStriSfosTRiS fOsTrisfostRiSfosTRiSfOsTRiSfOStRiSfoStriSfosTRiS")
```
