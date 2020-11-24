# protocol-buffer

## Final Objective:
membuat modeling data inventory komputer

## Poin :
1. Mendefinisikan Skema
2. Field : Field rule, Field Types, Field Number
3. Advance field type : Enum, Nested
4. Multifile .proto
   
## install compiler
```
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go
```

## compile 
```
//syntax compile
protoc -I=<source folder proto> --go_out=<source destination > --go_opt=paths=source_relative file.proto

//Example
cd /proto
protoc -I=. --go_out=../pb --go_opt=paths=source_relative todo.proto
```

1. ## **mendefinisikan sebuah message (skema)**
```
    // FORMAT 
    message <Name> {
        /**
         * field tag harus unik dan merupakan integer 
         * field tag yang dapat digunakan : 1 - 2^29-1
         * field tag yang tidak dapat digunakan : 19000 - 19999 
        */
        <Field-rule> <Field-Types> <name> = <field_tag>;
        ...
    }

    //Example 
    message CPU {
        string brand = 1;
        string name = 2;
        unit32 number_cores = 5;
        uint32 numbers_threads = 5;
        double min_ghz = 5;
        double max_ghz = 6;
    }
     
``` 