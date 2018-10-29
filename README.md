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

# start while structure !
while
(a < 5) 
{
    print(a);
    a++;
}

# this is multi-line comment
# my name is ahmadreza zibaei
# how are you ?

float_number := 3.141592;
sc_number := 2e-10;

c := a + b;

for (i := 0;i < 10;i++) {
    print(i);
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
In line 011 token while           type stmt
In line 011 token (a              type name
In line 011 token <               type oprt
In line 011 token 5               type dgts
In line 011 token )               type dlmt
In line 012 token {               type dlmt
In line 013 token print           type name
In line 013 token (               type dlmt
In line 013 token a               type name
In line 013 token )               type dlmt
In line 013 token ;               type dlmt
In line 014 token a               type name
In line 014 token +               type mtsl
In line 014 token +               type mtsl
In line 014 token ;               type dlmt
In line 015 token }               type dlmt
In line 021 token float_number    type name
In line 021 token :=              type dcln
In line 021 token 3.141592        type dgts
In line 021 token ;               type dlmt
In line 022 token sc_number       type name
In line 022 token :=              type dcln
In line 022 token 2e-10           type dgts
In line 022 token ;               type dlmt
In line 024 token c               type name
In line 024 token :=              type dcln
In line 024 token a               type name
In line 024 token +               type mtsl
In line 024 token b               type name
In line 024 token ;               type dlmt
In line 026 token for             type stmt
In line 026 token (               type dlmt
In line 026 token i               type name
In line 026 token :=              type dcln
In line 026 token 0               type dgts
In line 026 token ;               type dlmt
In line 026 token i               type name
In line 026 token <               type oprt
In line 026 token 10              type dgts
In line 026 token ;               type dlmt
In line 026 token i               type name
In line 026 token +               type mtsl
In line 026 token +               type mtsl
In line 026 token )               type dlmt
In line 026 token {               type dlmt
In line 027 token print           type name
In line 027 token (               type dlmt
In line 027 token i               type name
In line 027 token )               type dlmt
In line 027 token ;               type dlmt
In line 028 token }               type dlmt
```