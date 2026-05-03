# utf8

UTF-8 encoder and decoder implemented from scratch in Go.

No standard library encoding functions used: every bit operation is explicit. Built to understand what happens at the byte level when text moves across a wire.

## Usage and internals

Article available at [https://x.com/mxdvf](x.com/mxdvf) (coming soon!), explained step-by-step using first principles.

## Private use area demo

Unicode reserves U+E000-U+F8FF for private use → valid codepoints with no assigned meaning. The demo encodes U+E001 to bytes, ships it over the wire and decodes it through a custom glyph registry that maps it to a private symbol.

Open `demo/index.html` to see it live.

## Benchmarks

```
BenchmarkEncodeRune    75837970    15.17 ns/op    3 B/op    1 allocs/op
BenchmarkDecodeRune    348052320    3.45 ns/op    0 B/op    0 allocs/op
```

Decode is ~4x faster than encode. The gap comes from allocation: encode returns a new `[]byte` each call, decode reads from an existing slice and returns an integer.
