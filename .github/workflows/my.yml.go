name: CI

on: [push]

jobs:
  build:

    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-go@v1
      with:
        go-version: '1.13.8'
    - name: build
      run: go build
    - name: donwload
      uses: actions/upload-artifact@v1
      with:
        name: TODO
        path: TODO
