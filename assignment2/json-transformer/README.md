# code-challenge-json-transformer

DevX RUN SEAL JSON Transformer Coding Challenge

## Requirements

- **Must** transform the schema-less input `JSON file` to the desired output following the transformation criteria.
- **Must** print the output to `stdout`.

### Input

<details>
  <summary>Toggle</summary>

```json
{
  "number_1": {
    "N": "1.50"
  },
  "string_1": {
    "S": "784498 "
  },
  "string_2": {
    "S": "2014-07-16T20:55:46Z"
  },
  "map_1": {
    "M": {
      "bool_1": {
        "BOOL": "truthy"
      },
      "null_1": {
        "NULL ": "true"
      },
      "list_1": {
        "L": [
          {
            "S": ""
          },
          {
            "N": "011"
          },
          {
            "N": "5215s"
          },
          {
            "BOOL": "f"
          },
          {
            "NULL": "0"
          }
        ]
      }
    }
  },
  "list_2": {
    "L": "noop"
  },
  "list_3": {
    "L": [
      "noop"
    ]
  },
  "": {
    "S": "noop"
  }
}
```

</details>

### Output

<details>
  <summary>Toggle</summary>

```json
[
  {
    "map_1": {
      "list_1": [
        11,
        false
      ],
      "null_1": null
    },
    "number_1": 1.5,
    "string_1": "784498",
    "string_2": 1405544146
  }
]
```

</details>

### Transformation

- All `Nth` level values in the input are represented as `Strings` with pertinent data type information.
    - In the example [above](#input), `number_1` will be the field `key`, `"1.50"` would be the associated `value`,
      and **N** denotes the value's data type.
    - Other fields with different data types follow a similar pattern.
        - The `Input` has no schema restrictions other than this.
- Implementation should consider the following conventions for transformation.

#### JSON Field Keys

- **Must** sanitize the keys of trailing and leading whitespace before processing.
- **Must** represent keys with `String` data type.
- **Must** omit fields with empty keys.
- **Must** omit all invalid fields.

##### Data Types

- **S** represents the `String` data type.
    - It stores a `String` value.
    - **Transformation criteria**.
        - **Must** transform value to the `String` data type.
        - **Must** sanitize the value of trailing and leading whitespace before processing.
        - **Must** transform `RFC3339` formatted `Strings` to `Unix Epoch` in `Numeric` data type.
        - **Must** omit fields with empty values.


- **N** represents the `Number` data type.
    - It stores any `Numeric` value (positive, negative, int, float, etc.).
    - **Transformation criteria**.
        - **Must** be transformed to the relevant `Numeric` data type.
        - **Must** sanitize the value of trailing and leading whitespace before processing.
        - **Must** strip the leading zeros.
        - **Must** omit fields with invalid `Numeric` values.


- **BOOL** represents the `Boolean` data type.
    - It stores a `Boolean` value.
    - **Transformation criteria**.
        - **Must** be transformed to the `Boolean` data type.
        - **Must** transform `1`, `t`, `T`, `TRUE`, `true`, or `True` to `true`.
        - **Must** transform `0`, `f`, `F`, `FALSE`, `false`, or `False` to `false`.
        - **Must** sanitize the value of trailing and leading whitespace before processing.
        - **Must** omit fields with invalid `Boolean` values.


- **NULL** represents the `Null` data type.
    - It denotes if a value is supposed to be `Null` using a `Boolean` data type.
    - **Transformation criteria**.
        - **Must** represent a `null` literal when the value is `1`, `t`, `T`, `TRUE`, `true`, or `True`.
        - **Must** omit the field when the value is `0`, `f`, `F`, `FALSE`, `false`, or `False`.
        - **Must** sanitize the value of trailing and leading whitespace before processing.
        - **Must** omit fields with invalid `Boolean` values.


- **L** represents the `List` data type.
    - It stores a `List` of heterogeneous data types except for the `Null`, `List`, or `Map` data types.
    - **Transformation criteria**.
        - **Must** be transformed into the apt data types.
        - **Must** not contain empty `Strings`.
        - **Must** maintain the input order.
        - **Must** omit fields with unsupported data types.
        - **Must** omit fields with empty `List`.
        - The `List` can contain duplicates.


- **M** represents the `Map` data type.
    - It stores the unordered collection of heterogeneous data types including the `Map` data type.
    - **Transformation criteria**.
        - **Must** adhere to all the data type criteria defined in this document.
        - **Must** be lexically sorted.
        - **Must** omit fields with empty `Map`.

### Submission

- **Must** select a language available on [Replit](https://replit.com/templates).
- **Must** select a random name for the repo using the [generator](https://mrsharpoblunto.github.io/foswig.js).
- **Must** host the solution on `GitHub` and import it in `Replit` with the apt language/template.
- **Must** not contain any reference or content from this **confidential** document except the [Input](#input), in part
  or entirety.
- **Must** contain a `README.md` file with `local` execution instructions and `Replit` import/setup and execution
  instructions.
    - A sample [README.md](submission/README.md) is provided in the submission dir.
- **Must** report the implementation processing time.
- **Must** provide the `GitHub` repo link post completion.
- **Must** be prepared for live coding on `Replit` with an invitation link.
- Usage of third-party libraries is highly discouraged.
- Please feel free to submit a **partial solution** with apt justifications.
    - A solution with only the missing data type transformations will be deemed a **partial solution**.
        - E.g., you can skip the `List` or `Map` data type transformations with a justification of being complex to implement.

## Version

`v0.0.1`

## Maintainers

**DevX RUN SEAL**