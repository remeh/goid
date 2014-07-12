# goid

Basic package to generate pseudo-random unique IDs.

## Example

By providing the dictionary:

```
	dictionary := make([]byte, 4)
	dictionary[0] = '0'
	dictionary[1] = 'A'
	dictionary[2] = '2'
	dictionary[3] = 'z'
```

You're able to generate as many IDs as you want by providing the number
of already generated keys :

```
	for i := 0; i < 20; i++ {
        fmt.Println(goid.GenerateNext(dictionary, i, 0));
    }
```

Displays :

```
0
A
2
z
0A
AA
2A
zA
02
A2
22
z2
0z
Az
2z
zz
00A
A0A
20A
z0A
0AA
AAA
2AA
zAA
02A
A2A
```

This algorithm is limited to generate `(2^31) - 1` different values (signed int max) due to its design. It's a design choice : it could have been easily extended by using int64 but the performance should have been impacted.

## Roadmap

  - Randomized dictionary
