## L. Фибоначчи по модулю

| <!-- -->      |             <!-- -->             |
|:-------------:|:--------------------------------:|
| Ограничение времени	|           0.7 секунд         |
|Ограничение памяти	|               64Mb               |
|Ввод | 	стандартный ввод или input.txt  |
|Вывод | стандартный вывод или output.txt |

У Тимофея было n (0 ≤ n ≤ 32) стажёров. Каждый стажёр хотел быть лучше своих предшественников, поэтому i-й стажёр делал столько коммитов, сколько делали два предыдущих стажёра в сумме. Два первых стажёра были менее инициативными — они сделали по одному коммиту. Пусть F(i) — число коммитов, сделанных i -м стажёром (стажёры нумеруются с нуля). Тогда выполняется следующее:F(0) = F(1) = 1 . Для всех i ≥ 2 выполнено F(i) = F(i−1) + F(i−2) . Определите, сколько кода напишет следующий стажёр – найдите F(n) . Решение должно быть реализовано рекурсивно.

### Формат ввода
На вход подаётся n — целое число в диапазоне от 0 до 32.
### Формат вывода
Нужно вывести F(n) .