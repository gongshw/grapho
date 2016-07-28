# grapho
A copycat of gravizo.com by Go

Only support *Graphviz* and *Plant UML* now.

## Graphviz Diagram
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

## Plant UML Class Diagram

<img src='http://g.gongshw.com/g?
@startuml;
class Object << general >>;
Object <|--- ArrayList;
note top of Object : In java, every class\nextends this one.;
note "This is a floating note" as N1;
note "This note is connected\nto several objects." as N2;
Object .. N2;
N2 .. ArrayList;
class Foo;
note left: On last defined class;
@enduml;
'>

<img src='http://g.gongshw.com/g?%20@startuml;%20class%20Object%20%3C%3C%20general%20%3E%3E;%20Object%20%3C|---%20ArrayList;%20note%20top%20of%20Object%20:%20In%20java,%20every%20class\nextends%20this%20one.;%20note%20%22This%20is%20a%20floating%20note%22%20as%20N1;%20note%20%22This%20note%20is%20connected\nto%20several%20objects.%22%20as%20N2;%20Object%20..%20N2;%20N2%20..%20ArrayList;%20class%20Foo;%20note%20left:%20On%20last%20defined%20class;%20@enduml;'>


## Plant UML Sequnce Diagram

<img src='http://g.gongshw.com/g?
@startuml;
Alice -> Bob: Authentication Request;
alt successful case;
    Bob -> Alice: Authentication Accepted;
else some kind of failure;
    Bob -> Alice: Authentication Failure;
    group My own label;
    	Alice -> Log : Log attack start;
        loop 1000 times;
            Alice -> Bob: DNS Attack;
        end;
    	Alice -> Log : Log attack end;
    end;
else Another type of failure;
   Bob -> Alice: Please repeat;
end;
@enduml;
'>

<img src='http://g.gongshw.com/g?%20@startuml;%20Alice%20-%3E%20Bob:%20Authentication%20Request;%20alt%20successful%20case;%20Bob%20-%3E%20Alice:%20Authentication%20Accepted;%20else%20some%20kind%20of%20failure;%20Bob%20-%3E%20Alice:%20Authentication%20Failure;%20group%20My%20own%20label;%20Alice%20-%3E%20Log%20:%20Log%20attack%20start;%20loop%201000%20times;%20Alice%20-%3E%20Bob:%20DNS%20Attack;%20end;%20Alice%20-%3E%20Log%20:%20Log%20attack%20end;%20end;%20else%20Another%20type%20of%20failure;%20Bob%20-%3E%20Alice:%20Please%20repeat;%20end;%20@enduml;}'>


## Plant UML Use Case Diagram

<img src='http://g.gongshw.com/g?
@startuml;
:Main Admin: as Admin;
(Use the application) as (Use);
User -> (Start);
User --> (Use);
Admin ---> (Use);
note right of Admin : This is an example.;
note right of (Use);
  A note can also;
  be on several lines;
end note;
note "This note is connected\nto several objects." as N2;
(Start) .. N2;
N2 .. (Use);
@enduml;
'>

<img src='http://g.gongshw.com/g?%20@startuml;%20class%20Object%20%3C%3C%20general%20%3E%3E;%20Object%20%3C|---%20ArrayList;%20note%20top%20of%20Object%20:%20In%20java,%20every%20class\nextends%20this%20one.;%20note%20%22This%20is%20a%20floating%20note%22%20as%20N1;%20note%20%22This%20note%20is%20connected\nto%20several%20objects.%22%20as%20N2;%20Object%20..%20N2;%20N2%20..%20ArrayList;%20class%20Foo;%20note%20left:%20On%20last%20defined%20class;%20@enduml;;'>
