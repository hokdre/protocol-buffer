# protocol-buffer

## Final Objective:
membuat modeling data inventory komputer

## Poin :
1. Mendefinisikan Skema
2. Multifile .proto
3. Field : Field rule, Field Types, Field Number
4. Advance field type : Enum, Nested
   
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
    syntax = "proto3";
    package pb;
    option go_package = "github.com/protocol-buffer/pb";

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

    message ... {
        ....
    }
     
```  
2. ## **Multifile .proto** <br />
   ### directory : <br /> 
   /proto <br />
    &nbsp; cpu.proto <br />
    &nbsp; memory.proto <br />
    
    ### from : cpu.Proto
   ```
   import "memory.proto";
   import "...";
   import "...";

   message ...  {
       Memory memory = 1;
   }
   ```

3. ## **Field :**
   1. ### **Field Types ** <br />
      1. **Enum Type :FIELD NUMBER HARUS DIMULAI DARI 0** <br />
        ```
            message ... {
                enum <EnumName> {
                    DEFAULT = 0;
                    ....
                }

                EnumName fieldName = 1;
            }

                //example
                message Memory {
                    enum Unit {
                        UNKNOW = 0;
                        BIT = 1;
                        BYTE = 2;
                        KILOBYTE = 3;
                        MEGABYTE = 4;
                        GIGABYTE = 5;
                        TERABYTE = 6;
                    }

                    Unit unit = 1;
                }
        ```  
      2. **Composite** <br />
        ```
        message M1 {
            ...
        }

        message M2 {
             M2 composite = 1;
        }
        ``` 
      3. **Date** <br />
         ```
         import "google/protobuf/timestamp.proto";

         message ... {
             ...
             google.protobuf.Timestamp updated_at = 1;
         }
         ``` 
    
   2. ### **Field Rule : array ** <br />
      1. **Repeated: Array**
      ```
         message M1 {

         }

         message M2 {
            repeated M2 composite = 1;
         }
      ``` 
      1. **oneof : pilihan**
        ```
            message Weight {
                oneof name {
                    double weight_kg = 1;
                    double weight_lb = 2;
                }
            }
        ``` 

   
