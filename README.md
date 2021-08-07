# calculator

This is a toy project for my go practice.

I'm new to golang. 

This project help me familiar with features of go. Such as
1. how to import local packages
2. difference between pointer of struct and interface
3. error code & enum presentation
4. buildin testing module

Comparing with my experience with cpp, it feels simple just like coding C.

The calculator is a libaray which can calculate a expression, now includes +-*/ but easy to add new binary operator

It calculates expression like (1+3.14)*5, accept blank in expreesion and offer error when not work.

Most confusing problem is I want to accept minus number expression such like 1+-2, Which make things complicate that I can't do it easily while scanning it.

At last, I give up once scanning stratagies and first analyze word then make grammer tree to solve it, so that it can solve 1+(-2).

Tests make you more confident. test first or you would find test be more complicate.

Coding is tour full of adventures and I enjoy it.
