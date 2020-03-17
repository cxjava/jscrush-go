# jscrush-go
golang version of run JSCrush


## Javascript crusher by @aivopaas.

JSCrush was originally created to pack my first js1k entry, a Tetris clone.

I have made some optimizations and bug fixes based on @tpdown's crusher:

- Size is reported in actual bytes, not characters as it was in previous versions.
- Use single or double quotes, output is optimal for both.
- Closure-compile or minify by hand before using the crusher.
- Single line comments, leading whitespace and linebreaks are stripped so you can easily crush commented hand-minified code.
- Recommended style for functions: A=function(a,b,c){return a+b},B=function(a,b,c){x=A(a,2)}.
- Use same function signatures - exactly same parameters, to allow more compression.
- Try to use as few different variable names as possible.

# Licensed under MIT license.
