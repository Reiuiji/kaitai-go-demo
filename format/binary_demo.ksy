meta:
  id: binary_demo
  title: Binary Demo Type
  file-extension: rawr
  license: MIT
  endian: be
doc: |
  Binary demo type for testing complex types
seq:
  - id: header
    type: header
  - id: body
    type: body
  - id: message
    type: message
  - id: entries
    type: block_entries
    size: 16

types:
  header:
    seq:
    - id: magic
      type: u2
      valid:
        eq: 0xF0F0
    - id: classification
      type: u1
      enum: classification
  body:
    seq:
    - id: id
      type: u2
    - id: float4
      type: f4
    - id: float8
      type: f8
    - id: count
      type: u1
      valid:
        min: 0
        max: 200
  message:
    seq:
    - id: name
      type: str
      size: 8
      encoding: UTF-8
  block_entries:
    seq:
    - id: block
      type: block
      repeat: eos
  block:
    seq:
      - id: value
        type: u4

enums:
  classification:
    0: unknown
    1: ps4
    2: ps5
    3: xbox
    4: steam