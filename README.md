# gbgolang
golang courses on geekbrains

2
Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?
Унарные операторы имеют наивысший приоритет. Оператор разыменования - унарный.
Вероятно компилятор GO когда встречает оператор * определят, какой тип имеет оперант справа от него.
Если указатель, то выполняется его разыменование. Иначе - пытается перемножить левый и правый операнды.

3
Калькулятор использует структуру Operation, в которой два int'а и string

Размер int'а равен разрядности системы (64 / 8 (бит в байте) = 8 байт)

Вероятно string реализован, как массив байт, а значит в своей основе это указатель на первый байт в массиве.
Указатель - адрес в памяти, а значит его размер равен как и у int'а

Таким образом, как в этой структуре не переставлять местами поля, память будет использоваться 
максимально экономоно


4
Калькулятор реализован в отдельном пакете
