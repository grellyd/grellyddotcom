---
title: "Demystifying Go's Strings, Runes, and Bytes"
date: 2018-10-11T11:40:19-07:00
draft: true
---

---

**TL;DR Runes are a renameing of Unicode's 'code points', and are a potentially multibyte representation of a UTF-8 character. String indexing produces bytes, not characters. Either use a `range` statement or [unicode/utf8](https://golang.org/pkg/unicode/utf8/) to access full runes.**

I've found Go intuitive. 

While looking up how to do something, after reading the documentation I usually think "oh, of course" rather than "oh. I suppose that works." Therefore I was surprised when trying to manipulate strings, it wasn't clear and simple. Plus, I found Rob Pike's blogpost ['Strings, bytes, runes and characters in Go'](https://blog.golang.org/strings) strangly obtuse on a first read.

Up to this point in my dabbling, I have avoided the need to do low level string manuipulations. However while doing programming practice problems, that changed. Generally, I just wanted to have the string as an array of characters. Since I didn't know immediately what to do, evidently I misunderstood something.

---

## The Problem

Lets take a simple problem from _Cracking the Coding Interview_ as an example: 

    1.1 - *IsUnique*: Implement an algorithm to determine if a string has all unique characters. 

We will be using the variant where we use no external data structures; an in-place check.

My intuition said to try:

{{< highlight go "linenos=table" >}}
// IsUnique checks if a string has all unique characters
func IsUnique(s string) bool {
	for i, char := range s {
		for j := i+1; j < len(s); j++ {
			if char == s[j] {
				return false
			}
		}
	}
	return true
}
{{< / highlight >}}

Range over the string, one character at a time. Then compare that character against every remaining character. O(n<sup>2</sup>) time and O(1) space. Not pretty but it should work.

However line 5 throws an error: **`'Invalid operation: char == s[j] (mismatched types rune and byte)'`**

Wait a second, what is a 'rune'? Shouldn't that error say 'character'?

Evidently in one code block I've managed to iterate across the string in two different ways, to get both a byte and a rune.

In otherwords, how are **`'for i, char := range s'`** and **`'s[j]'`** different?

---

## Character Encodings

Before we can talk about Runes in Go, we have to discuss character encodings and why simple ASCII mappings are no long sufficent. 

As Pike suggests in his blogpost, I would recommend reading Joel Spolsky's excellent blogpost on [Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/). I found it greatly clarified the need for some intermediate representation between bytes and strings, plus the historical context of how characters can be varying length.

In short, 

* ASCII is dead as it doesn't consistently handle characters beyond the first 128.
* Multiple bytes are needed to handle the globally connected internet's multitude of languages.
* There are many standards, of which [UTF-8](https://en.wikipedia.org/wiki/UTF-8) is the most universally used and recognised format.<sup>[1]</sup>
* All Go source code is UTF-8 encoded.

where UTF-8 is:

* Unicode Transformation Format (UTF) of the Universal Character Set ([UCS](https://www.iso.org/standard/69119.html)) using 8 bit (one byte/octet) sequence components..
* There are between one and four bytes per character.<sup>[2]</sup>
* Each character sequence refers to a "code point", the Unicode Consortium's way of referring to a 'complete' UTF-8 value.

Three quick examples:

### The letter 'f' 

<table>
<tr><td><b>Latin Character</b></td><td>f</td></tr>
<tr><td><b>ASCII Character</b></td><td>0x66</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0066</td></tr>
</table>

**Length:** 1 Byte

### The English word 'Forest'

<table>
<tr>
<td><b>Latin Character</b></td><td>F</td><td>o</td><td>r</td><td>e</td><td>s</td><td>t</td>
<tr><td><b>ASCII Character</b></td><td>0x46</td><td>0x6F</td><td>0x72</td><td>0x65</td><td>0x73</td><td>0x74</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0046</td><td>U+006F</td><td>U+0072</td><td>U+0065</td><td>U+0073</td><td>U+0074</td></tr>
</table>

**Length:** 6 Bytes

### The French word 'For&#234;t'

<table>
<tr><td><b>Latin Character</b></td><td>F</td><td>o</td><td>r</td><td>&#234;</td><td>t</td></tr>
<tr><td><b>ASCII Character</b></td><td>0x46</td><td>0x6F</td><td>0x72</td><td>0x65</td><td>0x74</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0046</td><td>U+006F</td><td>U+0072</td><td>U+0065 U+005E</td><td>U+0074</td></tr>
</table>

**Length:** 7 Bytes

<table>
<tr><td><b>Latin Character</b></td><td>F</td><td>o</td><td>r</td><td>&#234;</td><td>t</td></tr>
<tr><td><b>ASCII Character</b></td><td>0x46</td><td>0x6F</td><td>0x72</td><td>0x65</td><td>0x74</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0046</td><td>U+006F</td><td>U+0072</td><td>U+00EA</td><td>U+0074</td></tr>
</table>

**Length:** 6 Bytes

These two are equivalent. 

Note the second byte in the first example on the &#234; character. UTF-8 considers both a 'e' followed by a '^', and their hex sum to be valid. 

---

## Runes

But how do 'runes' fit into this? 

Simply put, the designers of Go found 'code point' to be an unwieldly phrase. Therefore they added the word 'rune' to the Go lexicon as a synonyum for 'code point'; it means the same but is one syllable less. Whenever you see 'rune', read 'code point'.

Source?

---

## The Solution

Armed with this new understanding, lets solve the original problem:

From [Effective Go](https://golang.org/doc/effective_go.html#for), the range statement "does more work for you, breaking out individual Unicode code points by parsing the UTF-8", hence how we got the runes above. Alright lets try nested `range` statements.

Also time to change `char` to `r`, as `char` is misleading.

<b>TODO: change '!' to '_'</b>

{{< highlight go "linenos=table" >}}
// IsUnique checks if a string has all unique characters
func IsUnique(s string) bool {
	for i, r1 := range s {
		for !, r2 := range s[i+1:] {
            if r1 == r2 {
                return false
            }
        }
	}
	return true
}
{{< / highlight >}}

Alright. That works. What if we don't want to use the range statement for the second loop? In that case, we have to use the [unicode/utf8](https://golang.org/pkg/unicode/utf8/) standard package. Specifically we are going to use [#DecodeRune](https://golang.org/pkg/unicode/utf8/#DecodeRune).

{{< highlight go "linenos=table" >}}
import "unicode/utf8"

// IsUnique checks if a string has all unique characters
func IsUnique(s string) bool {
	for i, r1 := range s {
		b := []byte(s[i+1:])
        for len(b) > 0 {
            r2, size := utf8.DecodeRune(b)
            if r1 == r2 {
                return false
            }
            b = b[size:]
        }
	}
	return true
}
{{< / highlight >}}

What if we wanted to compare against array indicies, per our original idea? Lets also clean up that byte array with a helper:

{{< highlight go "linenos=table" >}}
import "unicode/utf8"

// IsUniqueDecoded checks if a string has all unique characters
func IsUniqueDecoded(s string) bool {
    runes := runeArray(s)
	for i, r := range s {
		for j := i+1; j < len(runes); j++ {
			if r == runes[j] {
				return false
			}
		}
	}
	return true
}

func runeArray(s string) (runes []rune) {
    b := []byte(s)
    for len(b) > 0 {
        r, size := utf8.DecodeRune(b)
        runes = append(runes, r)
        b = b[size:]
    }
    return runes
}
{{< / highlight >}}


---


## Conclusion

Armed with our newfound knowledge, the Go blog ['Strings, bytes, runes and characters in Go'](https://blog.golang.org/strings) has a nice summary:

<blockquote>
Go source code is always UTF-8.
A string holds arbitrary bytes.
A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
Those sequences represent Unicode code points, called runes.
No guarantee is made in Go that characters in strings are normalized.
</blockquote>



## Sources & Further Reading

['Strings, bytes, runes and characters in Go'](https://blog.golang.org/strings)
[Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

http://standards.iso.org/ittf/PubliclyAvailableStandards/c069119_ISO_IEC_10646_2017.zip

https://tools.ietf.org/html/rfc3629


Go Language Specification: Rune Literals [https://golang.org/ref/spec#Rune_literals](https://golang.org/ref/spec#Rune_literals)

---

## Footnotes:

1.

One cautionary note: 

Joel Spolsky's post was written in 2003. I came away from reading his post thinking that UCS-2, a precursor to UTF-16, was the best encoding to use. He even notes that was how he chose to use encodings within his business. 

However looking at the adoption chart off wikipedia, it is clear he wrote that at a time when UTF-8 adption was much lower than it is now. 

{{< figure src="https://upload.wikimedia.org/wikipedia/commons/c/c4/Utf8webgrowth.svg" alt="Webpage Encoding Adoption Chart" >}}

Without delving into the encoding history since 2003, I would suggest that larger multibyte encodings such as UTF-16 have since fallen out of favour due to UTF-8's smaller size for the English speaking world, and lack of backwards-compatability with ASCII. 

2.

The RFC document refers to "In UTF-8, characters from the U+0000..U+10FFFF range (the UTF-16 accessible range) are encoded using sequences of 1 to 4 octets." An octet is defined as a set of 8 bits. The RFC document uses 'octet' instead of 'byte' as it is more exact, for historically a 'byte' could be something other than eight bits. In this post, I use the two interchangably.


<b>
TODO:

- Fix references
- CSS on blockquote for shading
- int32 of every rune
