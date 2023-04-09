# O'zgaruvchilar

#### Golangda har xil turdagi o'zgaruvchilar mavjud, masalan:

- `int` - 123 yoki -123 kabi butun sonlarni (butun raqamlarni) saqlaydi
- `flaot32` - 19.99 yoki -19.99 kabi o'nlik kasr sonlarni saqlaydi
- `string` - "Salom Dunyo" kabi matnni saqlaydi. Satr qo'shtirnoq bilan yoziladi
- `bool` - mantiq qiymatlarini saqlaydi: `true` yoki  `false`

## O'zgaruvchilarni e'lon qilish

Go da o'zgaruvchini e'lon qilishning ikkita yo'li mavjud:

#### 1. `var` kalit so'zi yordamida:

`var` kalit so'zidan keyin o'zgaruvchi nomi va turidan foydalaning:

> #### Sintaksis
>
> `var` *variablename type = value*

#### 2. `:=` belgisi yordamida:

`:=` belgisidan keyin o'zgaruvchi qiymatini kiriting:

> #### Sintaksis
>
> *variablename type := value*

**Eslatma:** Bu holda, qiymatdan o'zgaruvchining turi aniqlanadi **( kompilyator qiymatdan kelib chiqqan holda o'zgaruvchining turini belgilaydi degan ma'noni anglatadi).**

**Eslatma:** `:=` o'zgaruvchini unga qiymat bermasdan e'lon qilish mumkin emas .

## Boshlang'ich qiymatga ega o'zgaruvchan deklaratsiya

Agar o'zgaruvchining qiymati boshidan ma'lum bo'lsa, siz o'zgaruvchini e'lon qilishingiz va unga bitta satrda qiymat belgilashingiz mumkin:

```go
package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}
```

## Boshlang'ich qiymatsiz o'zgaruvchan deklaratsiya

Go'da barcha o'zgaruvchilar ishga tushiriladi. Shunday qilib, agar siz o'zgaruvchini boshlang'ich qiymatisiz e'lon qilsangiz, uning qiymati uning turining standart qiymatiga o'rnatiladi:

```go
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

## Deklaratsiyadan keyin qiymatni belgilash

O'zgaruvchi e'lon qilingandan keyin unga qiymat berish mumkin. Bu qiymat dastlab ma'lum bo'lmagan holatlar uchun foydalidir.

```go
package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}
```

> __Eslatma: `:=` yordamida o'zgaruvchiga qiymat bermasdan e'lon qilib bo'lmaydi.

## `var` va `:=` o'rtasidagi farq

