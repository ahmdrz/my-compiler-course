### Under Construction
---

## Lexer

Example input: 

```
a := 1;
b := 300;

# hello world
if (a > b) {
    print("hello");
}
```

Tokenizer output: 

```
In line 001 token a               type name
In line 001 token :=              type dcln
In line 001 token 1               type dgts
In line 001 token ;               type dlmt
In line 002 token b               type name
In line 002 token :=              type dcln
In line 002 token 300             type dgts
In line 002 token ;               type dlmt
In line 005 token if              type stmt
In line 005 token (               type dlmt
In line 005 token a               type name
In line 005 token >               type oprt
In line 005 token b               type name
In line 005 token )               type dlmt
In line 005 token {               type dlmt
In line 006 token print           type name
In line 006 token (               type dlmt
In line 006 token hello           type dcln
In line 006 token )               type dlmt
In line 006 token ;               type dlmt
In line 007 token }               type dlmt
```