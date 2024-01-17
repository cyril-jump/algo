## C. Соседи

| <!-- -->      |            <!-- -->             |
|:-------------:|:-------------------------------:|
| Ограничение времени	|            1 секунда            |
|Ограничение памяти	|              64Mb               |
|Ввод | стандартный ввод или input.txt  |
|Вывод |стандартный вывод или output.txt |

Дана матрица. Нужно написать функцию, которая для элемента возвращает всех его соседей. Соседним считается элемент, находящийся от текущего на одну ячейку влево, вправо, вверх или вниз. Диагональные элементы соседними не считаются.

Например, в матрице A соседними элементами для (0, 0) будут 2 и 0. А для (2, 1) –— 1, 2, 7, 7.
### Формат ввода
В первой строке задано n — количество строк матрицы. Во второй — количество столбцов m. Числа m и n не превосходят 1000. В следующих n строках задана матрица. Элементы матрицы — целые числа, по модулю не превосходящие 1000. В последних двух строках записаны координаты элемента, соседей которого нужно найти. Индексация начинается с нуля.
### Формат вывода
Напечатайте нужные числа в возрастающем порядке через пробел.