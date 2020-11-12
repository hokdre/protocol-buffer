# protocol-buffer

## install kompiler
```
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go
```

1. ## mendefinisikan sebuah message (skema)
    ```
        // FORMAT 
        message <Name> {
            <Field-rule> <Field-Types> <name> = <Field-Number>; // field number is must be unique
        }

        //EXAMPLE
        message Person {
            string name = 1;
            int32 id = 2;
            string email = 3;
            
            //Enum : deklarasi bisa didalam ataupun diluar
            enum PhoneType {
                MOBILE = 0; // enum harus dimulai dari 0
                HOME = 1;
                WORK = 2;
            }
            
            //Nested : deklarasi bisa didalam ataupun diluar
            message PhoneNumber {
                string number = 1;
                PhoneType type = 2; //menggunakan PhoneType(enum)
            }
            //menggunakan PhoneNumber type (nested)
            //menggunakan repeated rule, bersifat seperti sebuah array
            repeated PhoneNumber phones = 4;
        }

        // mendefinisikan lebih dari satu message dalam 1 proto file
        message AddressBook {
            repeated Person person = 1;
        }
    ``` 
    ### Field Konfigurasi
   1. field types
      1. Scalar 
      2. Composite (nested message)
      3. Enum (must start from 0 index)
   2. field number (Must be unique start from 1)
   3. field rule
      1. Singular
      2. Repeated (Array)
   
2. Kompile proto file
   ```
   //Format 
   protoc --proto_path=proto <path protofile> --go_out=plugins=grpc:<path source code>
  

   //Example
   protoc --proto_path=proto /proto/*.proto --go_out=plugins=grpc:pb
   ```

3. Implementasi Code
   1. Membuat message
      ```
      // encode
        book := &addressbook.AddressBook{}
        book.Person = []*addressbook.Person{
            &addressbook.Person{
                Name:  "andre",
                Id:    1,
                Email: "hokdre@gmail.com",
            },
        }
        fmt.Println(book)
        out, err := proto.Marshal(book)
        if err != nil {
            fmt.Printf("err marshaling : %s", err)
        }
      ``` 
   2. Membaca message
      ```
      // decode
        decoded := &addressbook.AddressBook{}
        err = proto.Unmarshal(out, decoded)
        if err != nil {
            fmt.Printf("err unmarshaling : %s", err)
        }
        fmt.Println(decoded)
      ``` 