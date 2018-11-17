### Under Construction
---

## Lexer

Example input: 

```
a := 1;
b := 300;
if a <> b {
    print("hello");
}

/*

hello
world

*/

c := 2
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
In line 003 token if              type stmt
In line 003 token a               type name
In line 003 token <>              type oprt
In line 003 token b               type name
In line 003 token {               type dlmt
In line 004 token print           type name
In line 004 token (               type dlmt
In line 004 token hello           type dcln
In line 004 token )               type dlmt
In line 004 token ;               type dlmt
In line 005 token }               type dlmt
In line 014 token c               type name
In line 014 token :=              type dcln
In line 014 token 2               type dgts
```