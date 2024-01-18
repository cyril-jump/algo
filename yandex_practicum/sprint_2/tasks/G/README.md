## G. Стек - MaxEffective

| <!-- -->      | <!-- -->        |
|:-------------:|:---------------:|
| Ограничение времени	| 0.2 секунды |
|Ограничение памяти	| 64Mb      |
|Ввод |	стандартный ввод или input.txt|
|Вывод | стандартный вывод или output.txt|

Реализуйте класс StackMaxEffective, поддерживающий операцию определения максимума среди элементов в стеке. Сложность операции должна быть O(1). Для пустого стека операция должна возвращать None. При этом push(x) и pop() также должны выполняться за константное время.

### Формат ввода
В первой строке записано одно число — количество команд, оно не превосходит 100000. Далее идут команды по одной в строке. Команды могут быть следующих видов:

- push(x) — добавить число x в стек. Число x не превышает 105;
- pop() — удалить число с вершины стека;
- get_max() — напечатать максимальное число в стеке;
- top() — напечатать число с вершины стека;

Если стек пуст, при вызове команды get_max нужно напечатать «None», для команды pop и top — «error».
### Формат вывода
Для каждой команды get_max() напечатайте результат её выполнения. Если стек пустой, для команды get_max() напечатайте «None». Если происходит удаление из пустого стека — напечатайте «error».