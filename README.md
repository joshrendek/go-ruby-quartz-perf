### Example Output

```
Calculating -------------------------------------
         ruby-bcrypt     1.000  i/100ms
           go-bcrypt     1.000  i/100ms
            ruby-add   140.469k i/100ms
              go-add     1.675k i/100ms
-------------------------------------------------
         ruby-bcrypt     15.430  (± 0.0%) i/s -    309.000
           go-bcrypt     11.415  (± 0.0%) i/s -    229.000
            ruby-add      8.786M (± 5.9%) i/s -    175.024M
              go-add     16.191k (±15.1%) i/s -    314.900k

Comparison:
            ruby-add:  8785574.8 i/s
              go-add:    16191.5 i/s - 542.60x slower
         ruby-bcrypt:       15.4 i/s - 569370.87x slower
           go-bcrypt:       11.4 i/s - 769657.63x slower
```
