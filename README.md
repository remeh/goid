# goid

Basic package to generate pseudo-random unique IDs.

## Example

By providing the dictionary:

```
	values := make([]byte, 4)
	values[0] = '0'
	values[1] = 'A'
	values[2] = '2'
	values[3] = 'z'
```

You're able to generate as many IDs as you want by providing the number
of already generated keys :

```
	for i := 0; i < 16; i++ {
        fmt.Println(goid.GenerateNext(i));
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
```

## Roadmap

  - Method to generate IDs with a fixed size
  - Randomized dictionary
