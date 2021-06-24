{
  "targets": [
    {
      "target_name": "addon",
      "sources": [ "nodegomodule.cc" ],
      "libraries": [ "<!(pwd)/../gomodule/build/gomodule.so" ]
    }
  ]
}