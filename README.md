# Simple Particle Swarm Optimization Sample for Golang Practice

Langage : Golang

## Abstract

### Purpose

Optimize following function :
```math
F(x_1,x2) = {x_1}^2 + {x_2}^2, S_0 = {(x_1, x_2)| |x_1| \le 5, |x_2| \le 5}
```
### Update Rule

```math
v^i_j(n + 1) = wv^i_j(n) + \rho_1(p^i_j - x^i_j(n)) + \rho_2(l^i_j - x^i_j(n)) \\
x^i_j(n + 1) = x^i_j(n) + v^i_j(n)
```

### Graph Structure

Complete Graph

