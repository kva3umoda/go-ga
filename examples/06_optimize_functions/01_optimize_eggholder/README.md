# Оптимизация функции Eggholder
![](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Eggholder_function.pdf/page1-800px-Eggholder_function.pdf.jpg)
https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Eggholder_function.pdf/page1-800px-Eggholder_function.pdf.jpg
Функция Eggholder (она называется так потому, что ее форма напоминает
подставку для яиц) часто используется для тестирования алгоритмов оптими-
зации. Нахождение единственного глобального минимума этой функции счи-
тается трудной задачей из-за большого количества локальных минимумов.

Математически функция описывается выражением:
```
f(x, y) = -(y + 47) * sin(sqrt((abs(x/2+(y+47))))) - x * sin(sqrt(abs(x - (y + 47))))
```

Известно, что глобального минимума эта функция достигает в точке:
```
x=512, y=404.2319, f(x, y)=-959.6407
```
