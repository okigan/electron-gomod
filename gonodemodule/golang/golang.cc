#include <node.h>
#include <v8.h>

#include "markdown.h"

void HelloMethod (const v8::FunctionCallbackInfo<v8::Value> &args) {


  v8::Isolate *isolate = args.GetIsolate();
  // Call exported Go function, which returns a C string
  char *c = Hello();
  // return the value
  args.GetReturnValue().Set(v8::String::NewFromUtf8(isolate, c).ToLocalChecked());
  free(c);
}

// add method to exports
void Init (v8::Local<v8::Object> exports) {
  NODE_SET_METHOD(exports, "hello", HelloMethod);
}

// create module
NODE_MODULE(NODE_GYP_MODULE_NAME, Init)