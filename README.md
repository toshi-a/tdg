# TDG (Test Data Generator)

## Overview

TDG is test data generator for CLI.  
To use it, specify the definition file in which items and output methods are entered, and the number of output lines, and call it.

### Features currently available
* Use it like this:
* Sequential number
* Random number
* Random enumeration item
* Static text
* Per item quoting on/off
* Output in sql insert statement

## Usage

Use it like this:

```bash
$ tdg PATH_TO_DEFINITION_JSON_FILE NUMBER_OF_OUTPUT_ROWS
```

for example. path of ./sample.json definition json file exists and 120 test data are output based on that definition:

```bash
$ tdg ./sample.json 120
INSERT INTO table1 (id, uuid, seq_int_digit5, seq_int_initial100, enum_list, name, address, random_int, num, created_at, modified_at) VALUES (0, 'a2053e2f-075f-5697-a579-434abb08d13f', '00000', 100, true, 'names of', 'list list of', 36, 50, NOW(), NOW());
INSERT INTO table1 (id, uuid, seq_int_digit5, seq_int_initial100, enum_list, name, address, random_int, num, created_at, modified_at) VALUES (1, 'f9886ca4-5dac-5b0f-b3bb-937dbfb61f84', '00001', 101, true, 'firstnames names', 'list states countries', 80, 756, NOW(), NOW());
INSERT INTO table1 (id, uuid, seq_int_digit5, seq_int_initial100, enum_list, name, address, random_int, num, created_at, modified_at) VALUES (2, 'f43930c8-4375-5d4a-bc2e-2df87955f878', '00002', 102, true, 'of names', 'of states of', 69, 210, NOW(), NOW());
...
```

## Definition JSON file

The contents of the definition JSON file are as follows.

```json
{
  "items": [
    {
      "name": "id",
      "generator": "sequential",
      "params": {}
    },
    {
      "name": "uuid",
      "quote": true,
      "generator": "uuid",
      "params": {}
    },
    {
      "name": "seq_int_digit5",
      "generator": "sequential",
      "quote": true,
      "params": {
        "digit": 5
      }
    },
    {
      "name": "seq_int_initial100",
      "generator": "sequential",
      "params": {
        "initial": 100
      }
    },
    {
      "name": "enum_list",
      "generator": "enum",
      "params": {
        "list": ["true", "false"]
      }
    },
    {
      "name": "name",
      "generator": "enum",
      "quote": true,
      "params": {
        "lists": [
          ["names", "of", "firstnames"],
          ["names", "of", "lastnames"]
        ]
      }
    },
    {
      "name": "address",
      "generator": "enum",
      "quote": true,
      "params": {
        "lists":[
          ["list", "of", "streets"],
          ["list", "of", "states"],
          ["list", "of", "countries"]
        ]
      }
    },
    {
      "name": "random_int",
      "generator": "integer",
      "params": {
        "min": 10,
        "max": 100
      }
    },
    {
      "name": "created_at",
      "generator": "static",
      "params": {
        "value": "NOW()"
      }
    },
    {
      "name": "modified_at",
      "generator": "static",
      "params": {
        "value": "NOW()"
      }
    }
  ],
  "output": {
    "writer": "sql",
    "params": {
      "table_name": "table1"
    }
  }
}
```

#### Description of each item

* **items**  
  An array of objects containing settings for the item
  *  **name**  
    Item name
  *  **generator**  
    Specifies the name of the generator that generates the item's data. See below for available generator names.
  *  **quote (Optional)**  
    A boolean whether to quote the item. Referred to the output method.
  *  **params**  
    Specify options for the generator specified by "generator". The available options are different for each generator.
* **output**  
  Output settings. Contains the following settings
  *  **writer**  
     Specifies the name of the writer that writes each record to the destination. See below for available writers.
  *  **params**  
     Specify options for the writer specified by "writer". The available options are different for each writer.


## Now available generators

### static
  A generator that uses the specified value as is.
#### params
* **value**: Value to output as is

### uuid
  A generator that generates UUIDs
#### params 
do not have

### integer
  A generator that generates random numbers.
#### params
* **min**  
  Lower bound for generated values.
* **max**  
  Maximum value to generate.

### sequential
  A generator that generates sequential numbers.
#### params
* **digit (optional)**  
  The number of digits in the generated value. Fill empty sentences with 0.
* **initial (optional)**  
  Initial value of the sequence number to generate.

### enum
  A generator that returns random items of the given array.
#### params
* **list (optional)**  
  The number of digits in the generated value. Fill empty sentences with 0.
* **lists (optional)**  
  Initial value of the sequence number to generate.

### date
A generator that generates random date.
#### params
* **min**  
  Lower bound for generated values.  
  Accepted formats are: [2006-01-02T15:04:05 MST](https://pkg.go.dev/time#Time.Format)
* **max**  
  Maximum value to generate.  
  Accepted formats are: [2006-01-02T15:04:05 MST](https://pkg.go.dev/time#Time.Format)
* **format (Optional, default: "2006-01-02 15:04:05")**  
  Specifies the format of the output date.  
  See the [golang documentation](https://pkg.go.dev/time#Time.Format) for the format.

## Now available writers

### sql
  A writer that writes out the values generated by each generator as an insert statement.
#### params
* **table_name**
  Table name to insert into.

## Develop environment (tool)

* [golang](https://go.dev) 1.19

## Build
If you have "go" installed, you should be able to build with the following command in the top directory of the project.

```bash
go build -o tdg ./cmd/tdg/main.go
```

## Scripts
The following commands are created and used for development on UNIX-based OS.

Create executable files for several environments under the build directory.
```bash
./build_all.sh
```

After building the executable file for the execution environment, refer to sample.json and execute
```bash
./build_run.sh
```
I don't write many tests, but I run all the tests in the project.
```bash
./test_all.sh
```

## Caution
This is a program I made for my own use as well as learning the go language.  
We will continue to add functions in the future, but please note that currently only the minimum error handling is done.  
Please take a look at the code and advise if there is a better way to implement it.
