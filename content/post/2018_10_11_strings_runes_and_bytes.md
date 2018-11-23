---
title: "Demystifying Go's Strings, Runes, and Bytes"
date: 2018-10-11T11:40:19-07:00
draft: true
---

TL;DR Runes are a renameing of Unicode's 'code points', and are a potentially multibyte representation of a UTF-8 character. String indexing produces bytes, not characters.

I've found Go intuitive. 

While looking up how to do something, after reading the documentation I usually think "oh, of course" rather than "oh. I suppose that works." Therefore I was surprised when trying to manipulate strings; it wasn't clear and simple. Plus, I found Rob Pike's blogpost ['Strings, bytes, runes and characters in Go'](https://blog.golang.org/strings) strangly obtuse on a first read.

Up to this point in my dabbling, I have avoided the need to do low level string manuipulations. However while doing programming practice problems, that changed. To check character uniquiness in a string, or to check palindromes etc, I just wanted to have the string as an array of characters. Since I didn't know immediately what to do, evidently I misunderstood something.

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

Range over the string, one character at a time. Then compare that character against every remaining character. O(n^2) time and O(1) space. Not pretty but it should work.

However line 5 throws an error: **`'Invalid operation: char == s[j] (mismatched types rune and byte)'`**

Wait a second, what is a 'rune'? Shouldn't that error say 'character'?

Evidently in one code block I've managed to iterate across the string in two different ways, to get both a byte and a rune.

In otherwords, how are **`'for i, char := range s'`** and **`'s[j]'`** different?

---

## Character Encodings

Before we can talk about Runes in Go, we have to talk character encodings and why simple ASCII mappings are no long sufficent. 

As Pike suggests in his blogpost, I would recommend reading Joel Spolsky's excellent blogpost on [Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/). I found it greatly clarified the need for some intermediate representation between bytes and strings, plus the historical context of how characters can be varying length.

In short, 

* ASCII is dead as it doesn't consistently handle characters beyond the first 128.
* Multiple bytes are needed to handle the globally connected internet's multitude of languages.
* There are many standards, of which [UTF-8](https://en.wikipedia.org/wiki/UTF-8) is the most universally used and recognised format. [1]
    * There are between one and four bytes per character. [2]
    * Each sequence refers to a "code point", the Unicode Consortium's way of referring to a 'complete' UTF-8 value.
* All Go source code is UTF-8 encoded.

Three quick examples:

### The letter 'f' 

<table>
<tr><td><b>Latin Character</b></td><td>f</td></tr>
<tr><td><b>ASCII Character</b></td><td>0x66</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0066</td></tr>
</table>

Length: 1 Byte

### The English word 'Forest'

<table>
<tr>
<td><b>Latin Character</b></td><td>F</td><td>o</td><td>r</td><td>e</td><td>s</td><td>t</td>
<tr><td><b>ASCII Character</b></td><td>0x46</td><td>0x6F</td><td>0x72</td><td>0x65</td><td>0x73</td><td>0x74</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0046</td><td>U+006F</td><td>U+0072</td><td>U+0065</td><td>U+0073</td><td>U+0074</td></tr>
</table>

Length: 6 Bytes

### The French word 'For^et'

<table>
<tr><td><b>Latin Character</b></td><td>F</td><td>o</td><td>r</td><td>^e</td><td>t</td></tr>
<tr><td><b>ASCII Character</b></td><td>0x46</td><td>0x6F</td><td>0x72</td><td>0x65</td><td>0x74</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0046</td><td>U+006F</td><td>U+0072</td><td>U+0065 U+005E</td><td>U+0074</td></tr>
</table>

Length: 7 Bytes

<table>
<tr><td><b>Latin Character</b></td><td>F</td><td>o</td><td>r</td><td>^e</td><td>t</td></tr>
<tr><td><b>ASCII Character</b></td><td>0x46</td><td>0x6F</td><td>0x72</td><td>0x65</td><td>0x74</td></tr>
<tr><td><b>Unicode Code Point</b></td><td>U+0046</td><td>U+006F</td><td>U+0072</td><td>U+00EA</td><td>U+0074</td></tr>
</table>

Length: 6 Bytes

These two are equivalent. 

Note the second byte in the first example on the ^e character. UTF-8 considers both a 'e' followed by a '^', and their hex sum to be valid. 

Hence UTF-8 encoded strings are variably lengthed, and often have multiple valid encodings.

Clearly UTF-8 is very important. But what are these rune things?

---

## Runes

Simply put, the designers of Go found 'code point' to be an unwieldly phrase. Therefore they added the word 'rune' to the Go lexicon as a synonyum for 'code point'; it means the same but is one syllable less. Whenever you see 'rune', read 'code point'.

---


## Sources & Further Reading

['Strings, bytes, runes and characters in Go'](https://blog.golang.org/strings)
[Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

http://standards.iso.org/ittf/PubliclyAvailableStandards/c069119_ISO_IEC_10646_2017.zip

https://tools.ietf.org/html/rfc3629

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
