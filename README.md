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
        fmt.Println(goid.GenerateNext(dictionary, i));
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
...
```

## Roadmap

  - Method to generate IDs with a fixed size
  - Randomized dictionary
