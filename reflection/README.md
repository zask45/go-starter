# Reflection

_Reflection_?

Apa maksudnya?

Jadi gini. Misalnya kita diminta nulis fungsi `walk(x interface{}, fn func(string))`

Kita masukkin `interface` sama `function`sebagai type-nya kan?

Nah ini yang dimaksud `reflection`.

 > Reflection itu bisa dibilang kemampuan program untuk meriksa strukturnya sendiri, terutama melalui types.


 ## What is interface{}

Kalo kita liat ini,

```
walk(x interface{}, fn func(string))
``` 

ada `interface{}` kan?

Maksudnya ini?

Selama ini kita ngebuat `function` dengan tipe-tipe data kayak `int` `string` atau tipe data sendiri kayak `BankAccount`. Nah kalo tipe data yang dimasukkin kayak gini, `compiler` bisa komplain kalo kita masukkin argumen dengan tipe data yang salah. 

Dari sini, gimana kalo kita mau masukkin nilai dengan tipe data apa aja?

Di sini lah kita pake `interface{}`! Tipe data interface bakal nerima value dengan tipe data apa aja. 

Kok gak pake `any` aja? 
Bisa juga kalo mau pake any. Kenapa? Karena `any` itu sebenernya alias untuk `interface{}`.

## Kenapa gak pake interface dan buat function flexible?

Kita gak bisa pake `interface{}` for everything. Pengecekkan kevalidan tipe data itu salah satu hal yang krusial di pemrograman. Misal kita cuma minta input dalam bentuk `int` masa kita mau accept input-an user dalem bentuk `char`? Gak mungkin kan? Makanya kita gak bisa pake `interface{}` untuk semua kasus.

