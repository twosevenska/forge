# Forge

![](http://i.imgur.com/p0DJ4uh.jpg)

> Then Brokkr brought forward his gifts: he gave to Odin the ring, saying that eight rings of the same weight would drop from it every ninth night; to Freyr he gave the boar, saying that it could run through air and water better than any horse, and it could never become so dark with night or gloom of the Murky Regions that there should not be sufficient light where be went, such was the glow from its mane and bristles. Then he gave the hammer to Thor, and said that Thor might smite as hard as he desired, whatsoever might be before him, and the hammer would not fail; and if he threw it at anything, it would never miss, and never fly so far as not to return to his hand; and if be desired, he might keep it in his sark, it was so small; but indeed it was a flaw in the hammer that the fore-haft was somewhat short.

## Quick description

The Forge is basically a module responsible for creating and managing resources used in games for [Aesir](https://github.com/twosevenska/Aesir). Examples: Create an item, critter, npc, player char possibly even adventures or campaigns. So in a way think of it like the middleman API between the other modules and the DB.

## How it works

Basically Forge receives an API request, depending on the type it either fetches something from the database or creates/updates it. That's it!