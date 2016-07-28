# grapho
A copycat of gravizo.com by Go

Only support *Graphviz* and *Plant UML* now.

```html
<img src='http://g.gongshw.com/g?
 digraph G {
   main -> parse -> execute;
   main -> init;
   main -> cleanup;
   execute -> make_string;
   execute -> printf
   init -> make_string;
   main -> printf;
 }
'>
```

<img src='http://g.gongshw.com/g?%20%20digraph%20G%20{%20%20%20%20main%20-%3E%20parse%20-%3E%20execute;%20%20%20%20main%20-%3E%20init;%20%20%20%20main%20-%3E%20cleanup;%20%20%20%20execute%20-%3E%20make_string;%20%20%20%20execute%20-%3E%20printf%20%20%20%20init%20-%3E%20make_string;%20%20%20%20main%20-%3E%20printf;%20%20}'>
