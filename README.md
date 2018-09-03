# **Go (Golang): Explorando a Linguagem do Google**
## *Professor: Leonardo Leitão*

*Concorrência vs Paralelismo*

Go é uma linguagem concorrente (**Round-robin**)

**Paralelismo** - excutar código simultaneamente em processadores físicos diferentes.

**Concorrência** - intercalar (administrar) vários processos ao mesmo tempo e isso pode ocorrer em um único processador físico.

**Concorrência viabiliza paralemismo**

* É possivel que a concorrência seja melhor que o paralelismo!
* Paralelismo exige muita mais do SO e do hardware.

**Concorrência** -  é a forma de estruturar o seu programa, é a composição de processos (tipicamente funções) que executam de forma independente.